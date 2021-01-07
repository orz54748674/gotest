package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"git.jiaxianghudong.com/go/utils"
)

var (
	from           int64
	end            int64
	appid          int32 = 432
	channelid      int32 = 200
	clientversion  int64 = 1
	devicecode           = "1111111111111111111111111111111111111111"
	password             = "E10ADC3949BA59ABBE56E057F20F883E"
	userNamePrefix       = "robotddz3d"
	urlPrefix            = "https://gw.weiletest.com/"
)

//注册
type ReqRegister struct {
	Type          int32             `json:"type"`          // 注册类型（1-普通账号注册，2-手机号注册）：手机号注册必须填写验证码
	Username      string            `json:"username"`      // 用户名
	Nickname      string            `json:"nickname"`      // 昵称
	Password      string            `json:"password"`      // 密码
	DeviceCode    string            `json:"deviceCode"`    // 设备码
	CaptchaToken  string            `json:"captchaToken"`  // 验证码标识
	Captcha       string            `json:"captcha"`       // 验证码
	Sex           int32             `json:"sex"`           // 性别,1:男,0:女
	Avatarid      int32             `json:"avatarid"`      // 头像id(-1用于自定义头像，此时地址存于avatarimg)
	Avatarimg     string            `json:"avatarimg"`     // 头像地址
	Region        string            `json:"region"`        // 区域码
	Refereeid     int64             `json:"refereeid"`     // 推荐人id
	Appid         int32             `json:"appid"`         // 应用id
	Channelid     int32             `json:"channelid"`     // 渠道id
	ClientVersion int64             `json:"clientVersion"` // 客户端版本
	Ext           map[string]string `json:"ext"`           // 扩展字段
}

// ReqCreateRole 创建角色
type ReqCreateRole struct {
	Sex      int32  `json:"sex"`      //性别 1:男 0:女
	Nickname string `json:"nickname"` //昵称
	Face     int32  `json:"face"`     //脸型
	Head     int32  `json:"head"`     //发型
	HColor   int32  `json:"hcolor"`   //发色
}

func main() {
	flag.Int64Var(&from, "s", 1041, "register start id, include")
	flag.Int64Var(&end, "e", 1042, "register end id, not include")
	flag.Parse()

	regMsg := &ReqRegister{}
	roleMsg := &ReqCreateRole{}

	rand.Seed(time.Now().Unix())
	for i := from; i < end; i++ {
		regMsg.Type = 1
		regMsg.Username = fmt.Sprintf("%v%v", userNamePrefix, i)
		regMsg.Nickname = ""
		regMsg.Password = password
		regMsg.DeviceCode = devicecode
		regMsg.Appid = appid
		regMsg.Channelid = channelid
		regMsg.ClientVersion = clientversion
		regMsg.Ext = make(map[string]string, 0)
		regMsg.Ext["type"] = "2"

		body, _ := json.Marshal(regMsg)

		url := urlPrefix + "logonsvr/register?format=json"
		header := map[string]string{"targetip": "192.168.13.96"}
		if _, err := utils.Post(url, body, header, true); err != nil {
			fmt.Printf("register account [%s] Err:%v\n", regMsg.Username, err.Error())
			continue
		}

		//登陆
		url = urlPrefix + "logonsvr/loginbyname?format=json"
		param := map[string]interface{}{
			"appid":         appid,
			"channelid":     channelid,
			"clientVersion": clientversion,
			"deviceCode":    devicecode,
			"password":      password,
			"region":        "0",
			"username":      regMsg.Username,
		}
		var uid int64
		var token string
		Post(url, param, func(data map[string]interface{}) {
			code := data["code"].(float64)
			if code != 0 {
				fmt.Println(data["msg"])
				return
			}
			token = data["token"].(string)
			uid = int64(data["userid"].(float64))
		}, nil)

		//创建角色
		roleMsg.Sex = rand.Int31n(1)
		roleMsg.Nickname = fmt.Sprintf("r%v", i)
		roleMsg.Face = 1
		roleMsg.Head = 101
		roleMsg.HColor = 201
		url = urlPrefix + "ddzhallapi/createRole?format=json"
		body, _ = json.Marshal(roleMsg)
		header = GetAuthorizeHeader(uid, token)
		if _, err := utils.Post(url, body, header, true); err != nil {
			fmt.Printf("create role [%s] Err:%v\n", regMsg.Username, err.Error())
			continue
		}

		fmt.Printf("Register %v\n", regMsg.Username)
	}

	fmt.Println("Done")
}

func Post(url string, data map[string]interface{}, cb func(map[string]interface{}), header map[string]string) {
	param := utils.JsonEncode(data)

	req, err := http.NewRequest("POST", url, strings.NewReader(param))
	if err != nil {
		fmt.Printf("Post error: %v url: %s\n", err, url)
		return
	}
	if header != nil {
		for k := range header {
			req.Header.Add(k, header[k])
		}
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Post error: %v url: %s\n", err, url)
		return
	}
	defer resp.Body.Close()

	var buffer [1024]byte
	n, err := resp.Body.Read(buffer[0:])

	if err != nil && n == 0 {
		return
	}

	str := string(buffer[0:n])
	ret := make(map[string]interface{})
	//fmt.Println(str)

	utils.JsonDecode(str, &ret)
	cb(ret)
}

func GetAuthorizeHeader(uid int64, token string) map[string]string {
	if token == "" {
		fmt.Printf("player[%d] haven't login yet.", uid)
		return nil
	}

	nonce := MakeNonce()
	timenow := fmt.Sprintf("%d", time.Now().Unix())
	struid := fmt.Sprintf("%d", uid)
	sign := Sign(nonce + timenow + struid + token)

	header := make(map[string]string)

	header["X-UID"] = struid
	header["X-NONCE"] = nonce
	header["X-TS"] = timenow
	header["X-SIGN"] = sign

	header["X-APPID"] = strconv.Itoa(int(appid))
	header["X-CHANNELID"] = strconv.Itoa(int(channelid))
	header["X-CVER"] = ""

	return header
}

func MakeNonce() string {
	format := "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx"
	var result string

	for _, c := range format {
		if c == 120 || c == 121 {
			r := (byte)(rand.Intn(16))
			if c == 120 {
				result += fmt.Sprintf("%1x", r)
			} else {
				result += fmt.Sprintf("%1x", r&0x3|0x8)
			}
		} else {
			result += fmt.Sprintf("%c", c)
		}
	}
	return result
}

func Sign(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write(String2Bytes(str))
	cipherStr := md5Ctx.Sum(nil)
	md5str := hex.EncodeToString(cipherStr)
	return strings.ToUpper(md5str)
}

func String2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
