package main

import (
	"fmt"
	"github.com/jan-bar/golibs"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const Length = 4 /* 表示2048长宽均4个格子 */

var (
	NumData [Length][Length]int /* 存储格子数据的二维数组 */
	TmpData = struct {
		data [Length * Length]struct{ i, j int }
		cnt  int
	}{} /* 缓存数据所在i,j值,cnt表示个数 */
	ValData = struct {
		data [Length]int
		cnt  int
	}{} /* 缓存每次处理有效数据,cnt表示个数 */
	Score     [2]int         /* 0存当前获得分数,1表示历史最高分 */
	ScoreFile string         /* 缓存分数的文件 */
	ColorMap  = map[int]int{ /* 每种数字的颜色值 */
		0:    golibs.ForegroundRed | golibs.ForegroundGreen | golibs.ForegroundBlue,
		2:    golibs.ForegroundRed,
		4:    golibs.ForegroundBlue | golibs.ForegroundRed,
		8:    golibs.ForegroundGreen,
		16:   golibs.ForegroundRed | golibs.ForegroundGreen,
		32:   golibs.ForegroundGreen | golibs.ForegroundBlue,
		64:   golibs.ForegroundRed,
		128:  golibs.ForegroundBlue | golibs.ForegroundGreen | golibs.ForegroundIntensity,
		256:  golibs.ForegroundGreen,
		512:  golibs.ForegroundRed | golibs.ForegroundGreen | golibs.ForegroundIntensity,
		1024: golibs.ForegroundGreen | golibs.ForegroundBlue,
		2048: golibs.ForegroundRed | golibs.ForegroundIntensity,
	}
	api = golibs.NewWin32Api()
)

func init() {
	api.Clear()                        /* 清屏 */
	api.CenterWindowOnScreen(500, 330) /* 居中显示,并设置窗口大小 */
	api.SetWindowText("2048游戏")
	api.ShowHideCursor(false)            /* 不显示光标 */
	api.ShowScrollBar(golibs.SB_BOTH, 0) /* 不显示滚动条 */
	fmt.Println(` ---------------------------
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |
 ---------------------------
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |
 ---------------------------
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |
 ---------------------------
|      |      |      |      |
|      |      |      |      |
|      |      |      |      |
 ---------------------------`) /* 画外框 */
	api.GotoXY(30, 1)
	fmt.Print("按上下左右按键控制!")

	ScoreFile = filepath.Join(os.TempDir(), "game2048.score")
	byt, err := ioutil.ReadFile(ScoreFile)
	if err == nil { /* 读文件正常 */
		Score[1], _ = strconv.Atoi(string(byt)) /* 转换异常时分数为0 */
	}
	rand.Seed(time.Now().Unix())
}

func main() {
	var (
		i, j   int /* 循环变量 */
		record = func(i, j int) {
			if 0 != NumData[i][j] { /* 记录不为空的位置 */
				ValData.data[ValData.cnt] = NumData[i][j]
				ValData.cnt++
			}
			TmpData.data[TmpData.cnt].i = i
			TmpData.data[TmpData.cnt].j = j
			TmpData.cnt++ /* 缓存本次4个点的位置 */
		}
	)
	RandNum() /* 开局随机出现1个位置 */
	for {
		RandNum()      /* 空白位置新增一个2或4 */
		PrintNumData() /* 打印游戏界面 */
		UpdateScore()  /* 打印并,更新分数 */
		if i = SucceedOrFail(); i >= 0 {
			api.GotoXY(30, 5)
			fmt.Print([]string{"你输了比赛!", "你赢了比赛!"}[i])
			break
		}

		switch golibs.WaitKeyBoard() {
		case golibs.KeyUp:
			for j = 0; j < Length; j++ {
				for i, ValData.cnt, TmpData.cnt = 0, 0, 0; i < Length; i++ {
					record(i, j)
				}
				ProcessingData() /* 处理上面得到的数据 */
			}
		case golibs.KeyDown:
			for j = 0; j < Length; j++ {
				for i, ValData.cnt, TmpData.cnt = Length-1, 0, 0; i >= 0; i-- {
					record(i, j)
				}
				ProcessingData() /* 处理上面得到的数据 */
			}
		case golibs.KeyLeft:
			for i = 0; i < Length; i++ {
				for j, ValData.cnt, TmpData.cnt = 0, 0, 0; j < Length; j++ {
					record(i, j)
				}
				ProcessingData() /* 处理上面得到的数据 */
			}
		case golibs.KeyRight:
			for i = 0; i < Length; i++ {
				for j, ValData.cnt, TmpData.cnt = Length-1, 0, 0; j >= 0; j-- {
					record(i, j)
				}
				ProcessingData() /* 处理上面得到的数据 */
			}
		} // end switch
	}
	_, _ = fmt.Scanln() /* 避免一闪而逝 */
}

/**
* 处理数据
* 本次需要处理的数据
* 均已经缓存到ValData 和 TmpData中了
**/
func ProcessingData() {
	var i, tmpScore = 0, 0
	switch ValData.cnt {
	case 1: /* 1个有效值,直接放到最底下 */
		NumData[TmpData.data[i].i][TmpData.data[i].j] = ValData.data[0]
	case 2:
		if ValData.data[0] == ValData.data[1] { /* 2个有效值相等 */
			tmpScore = ValData.data[0] * 2 /* 加分数 */
			NumData[TmpData.data[i].i][TmpData.data[i].j] = tmpScore
		} else { /* 2个有效值不等 */
			NumData[TmpData.data[i].i][TmpData.data[i].j] = ValData.data[0]
			i++
			NumData[TmpData.data[i].i][TmpData.data[i].j] = ValData.data[1]
		}
	case 3:
		if ValData.data[0] == ValData.data[1] {
			tmpScore = ValData.data[0] * 2
			NumData[TmpData.data[i].i][TmpData.data[i].j] = tmpScore /* 加分数 */
			i++
			NumData[TmpData.data[i].i][TmpData.data[i].j] = ValData.data[2]
		} else if ValData.data[1] == ValData.data[2] {
			NumData[TmpData.data[i].i][TmpData.data[i].j] = ValData.data[0]
			i++
			tmpScore = ValData.data[1] * 2
			NumData[TmpData.data[i].i][TmpData.data[i].j] = tmpScore /* 加分数 */
		} else {
			NumData[TmpData.data[i].i][TmpData.data[i].j] = ValData.data[0]
			i++
			NumData[TmpData.data[i].i][TmpData.data[i].j] = ValData.data[1]
			i++
			NumData[TmpData.data[i].i][TmpData.data[i].j] = ValData.data[2]
		}
	case 4:
		if ValData.data[0] == ValData.data[1] {
			tmpScore = ValData.data[0] * 2
			NumData[TmpData.data[i].i][TmpData.data[i].j] = tmpScore /* 加分数 */
			i++
			if ValData.data[2] == ValData.data[3] {
				tmpScore += ValData.data[2] * 2
				NumData[TmpData.data[i].i][TmpData.data[i].j] = ValData.data[2] * 2
			} else {
				NumData[TmpData.data[i].i][TmpData.data[i].j] = ValData.data[2]
				i++
				NumData[TmpData.data[i].i][TmpData.data[i].j] = ValData.data[3]
			}
		} else if ValData.data[1] == ValData.data[2] {
			NumData[TmpData.data[i].i][TmpData.data[i].j] = ValData.data[0]
			i++
			tmpScore = ValData.data[1] * 2
			NumData[TmpData.data[i].i][TmpData.data[i].j] = tmpScore /* 加分数 */
			i++
			NumData[TmpData.data[i].i][TmpData.data[i].j] = ValData.data[3]
		} else if ValData.data[2] == ValData.data[3] {
			NumData[TmpData.data[i].i][TmpData.data[i].j] = ValData.data[0]
			i++
			NumData[TmpData.data[i].i][TmpData.data[i].j] = ValData.data[1]
			i++
			tmpScore = ValData.data[2] * 2
			NumData[TmpData.data[i].i][TmpData.data[i].j] = tmpScore /* 加分数 */
		} else {
			return /* 4个都不相等,不用继续了 */
		}
	default:
		return /* 没有可用数据,可以直接退出 */
	}

	Score[0] += tmpScore       /* 本次得分加到当前游戏得分中 */
	for i++; i < Length; i++ { /* 清零 */
		NumData[TmpData.data[i].i][TmpData.data[i].j] = 0
	}
}

/**
* 在空位随机出现一个2或4
**/
func RandNum() {
	var (
		i, j int /* 循环变量 */
	)
	for i, TmpData.cnt = 0, 0; i < Length; i++ { /* 从上往下 */
		for j = 0; j < Length; j++ { /* 从左往右 */
			if 0 == NumData[i][j] {
				TmpData.data[TmpData.cnt].i = i
				TmpData.data[TmpData.cnt].j = j
				TmpData.cnt++ /* 记录空白位置i,j值 */
			}
		}
	}

	if TmpData.cnt > 0 { /* 当有空白位置才填值 */
		i = rand.Intn(TmpData.cnt) /* 随机选一个空白位置 */
		j = rand.Intn(2)           /* j=0,写2,j=1,写4 */
		NumData[TmpData.data[i].i][TmpData.data[i].j] = j*2 + 2
	}
}

/**
* 判断当前成功还是失败
* 返回-1表示还能玩
* 返回0表示输了
* 返回1表示赢了
**/
func SucceedOrFail() int {
	var i, j int                 /* 循环变量 */
	for i = 0; i < Length; i++ { /* 从上往下 */
		for j = 0; j < Length; j++ { /* 从左往右 */
			if 2048 <= NumData[i][j] {
				return 1 /* 赢了 */
			}
			if 0 == NumData[i][j] {
				return -1 /* 有空位,能继续 */
			}

			if i+1 < Length && NumData[i][j] == NumData[i+1][j] {
				return -1 /* 和 下 相同 */
			}
			if j+1 < Length && NumData[i][j] == NumData[i][j+1] {
				return -1 /* 和 右 相同 */
			}
		}
	}
	return 0 /* 不能继续,也没赢,就是输了 */
}

/**
* 更新最高分
**/
func UpdateScore() {
	api.GotoXY(30, 3)
	fmt.Printf("当前分数:%d,历史最高:%d", Score[0], Score[1])
	if Score[0] > Score[1] {
		Score[1] = Score[0] /* 不光临时变量要更新,分数文件也要更新 */
		_ = ioutil.WriteFile(ScoreFile, []byte(strconv.Itoa(Score[0])), os.ModePerm)
	}
}

/**
* 画界面
* 每个数字都不同颜色
**/
func PrintNumData() {
	var i, j, x, tmp int
	for i = 0; i < Length; i++ {
		for x, j = 4*i+2, 0; j < Length; j++ {
			api.GotoXY(7*j+1, x)
			if tmp = NumData[i][j]; tmp > 0 {
				api.TextBackground(ColorMap[tmp]) /* 前置文字上色 */
				fmt.Printf(" %4d ", tmp)
			} else {
				fmt.Print("      ")
			}
		} /* end for j */
	} /* end for i */
	api.TextBackground(ColorMap[0]) /* 重置为白色 */
}

/** 下面是2048游戏规则
 ---------------------------
|      |      |      |      |
| 4096 | 4096 | 4096 | 4096 |
|      |      |      |      |
 ---------------------------
|      |      |      |      |
| 4096 | 4096 | 4096 | 4096 |
|      |      |      |      |
 ---------------------------
|      |      |      |      |
| 4096 | 4096 | 4096 | 4096 |
|      |      |      |      |
 ---------------------------
|      |      |      |      |
| 4096 | 4096 | 4096 | 4096 |
|      |      |      |      |
 ---------------------------
开始时棋盘内随机出现两个数字，出现的数字仅可能为2或4。

玩家可以选择上下左右四个方向（电脑用户请使用上下左右键，手机和平板用户直接在棋盘内向四个方向拖动），若棋盘内的数字出现位移或合并，视为有效移动。

玩家选择的方向上若有相同的数字则合并，每次有效移动可以同时合并，但不可以连续合并。

合并所得的所有新生成数字相加即为该步的有效得分。

玩家选择的方向行或列前方有空格则出现位移。

每有效移动一步，棋盘的空位（无数字处）随机出现一个数字（依然可能为2或4）。

棋盘被数字填满，无法进行有效移动，判负，游戏结束。

棋盘上出现2048，判胜，游戏结束。
*/
