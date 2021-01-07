package main

import "fmt"

const ROOM_TYPE_MASK = 0xff00
const ROOM_LEVEL_MASK = 0x00ff
const GAME_TYPE_MASK = 0xff0000

//房间类型
const (
	ROOM_TYPE_NORMAL          = 1 + iota //金币场
	ROOM_TYPE_FRIEND                     //朋友场
	ROOM_TYPE_DASHISAI_DAILY             //大师赛-日循环赛
	ROOM_TYPE_DASHISAI_WEEKLY            //大师赛-周循环赛
	ROOM_TYPE_CHUANGGUAN                 //闯关

	ROOM_TYPE_MAX = 0xff
)

//场次级别
const (
	ROOM_LEVEL_XINSHOU = 1 + iota //新手场
	ROOM_LEVEL_PUTONG             //普通场
	ROOM_LEVEL_JINGYING
	ROOM_LEVEL_TUHAO

	ROOM_LEVEL_Max = 0xff
)

//游戏类型
const (
	GAME_TYPE_UNKNOWN = iota //未知类型
	GAME_TYPE_POKER          //扑克类
	GAME_TYPE_MAHJONG        //麻将类

	GAME_TYPE_MAX = 0xff
)

//00000000 00000000 00000000 00000000
//00000000 00000001 00000001 00000001 扑克-金币场-新手场
//00000000 00000001 00000010 00000000 扑克-朋友场
//00000000 00000010 00000001 00000001 麻将-金币场-新手场
//00000000 00000001 00000011 00000001 扑克-大师赛-日循环赛-初级场
//00000000 00000001 00000100 00000000 扑克-大师赛-周循环赛
func main() {
	//fmt.Println(ROOM_TYPE_MASK)
	//fmt.Println(ROOM_LEVEL_MASK)
	var roomType int32 = ROOM_TYPE_NORMAL<<8 | ROOM_LEVEL_XINSHOU | GAME_TYPE_POKER<<16
	var roomType2 int32 = ROOM_TYPE_FRIEND<<8 | GAME_TYPE_POKER<<16
	var roomType3 int32 = ROOM_TYPE_NORMAL<<8 | ROOM_LEVEL_XINSHOU | GAME_TYPE_MAHJONG<<16
	var roomType4 int32 = ROOM_TYPE_DASHISAI_DAILY<<8 | ROOM_LEVEL_XINSHOU | GAME_TYPE_POKER<<16
	var roomType5 int32 = ROOM_TYPE_DASHISAI_WEEKLY<<8 | GAME_TYPE_POKER<<16
	fmt.Printf("roomType:%d, roomLevel:%d, gameType:%d\n", (roomType&ROOM_TYPE_MASK)>>8, roomType&ROOM_LEVEL_MASK, (roomType&GAME_TYPE_MASK)>>16)
	fmt.Printf("roomType2:%d, roomLevel2:%d, gameType2:%d\n", (roomType2&ROOM_TYPE_MASK)>>8, roomType2&ROOM_LEVEL_MASK, (roomType2&GAME_TYPE_MASK)>>16)
	fmt.Printf("roomType3:%d, roomLevel3:%d, gameType3:%d\n", (roomType3&ROOM_TYPE_MASK)>>8, roomType3&ROOM_LEVEL_MASK, (roomType3&GAME_TYPE_MASK)>>16)
	fmt.Printf("roomType4:%d, roomLevel4:%d, gameType4:%d\n", (roomType4&ROOM_TYPE_MASK)>>8, roomType4&ROOM_LEVEL_MASK, (roomType4&GAME_TYPE_MASK)>>16)
	fmt.Printf("roomType5:%d, roomLevel5:%d, gameType5:%d\n", (roomType5&ROOM_TYPE_MASK)>>8, roomType5&ROOM_LEVEL_MASK, (roomType5&GAME_TYPE_MASK)>>16)

	fmt.Printf("扑克-朋友场：%d\n", ROOM_TYPE_FRIEND<<8|GAME_TYPE_POKER<<16)
	fmt.Printf("扑克-金币场-新手场：%d\n", ROOM_TYPE_NORMAL<<8|ROOM_LEVEL_XINSHOU|GAME_TYPE_POKER<<16)
	fmt.Printf("扑克-金币场-普通场：%d\n", ROOM_TYPE_NORMAL<<8|ROOM_LEVEL_PUTONG|GAME_TYPE_POKER<<16)
	fmt.Printf("扑克-金币场-精英场：%d\n", ROOM_TYPE_NORMAL<<8|ROOM_LEVEL_JINGYING|GAME_TYPE_POKER<<16)
	fmt.Printf("扑克-金币场-土豪场：%d\n", ROOM_TYPE_NORMAL<<8|ROOM_LEVEL_TUHAO|GAME_TYPE_POKER<<16)

	fmt.Printf("扑克-大师赛-日循环赛-初级场：%d\n", ROOM_TYPE_DASHISAI_DAILY<<8|ROOM_LEVEL_XINSHOU|GAME_TYPE_POKER<<16)
	fmt.Printf("扑克-大师赛-日循环赛-中级场：%d\n", ROOM_TYPE_DASHISAI_DAILY<<8|ROOM_LEVEL_PUTONG|GAME_TYPE_POKER<<16)
	fmt.Printf("扑克-大师赛-日循环赛-高级场：%d\n", ROOM_TYPE_DASHISAI_DAILY<<8|ROOM_LEVEL_JINGYING|GAME_TYPE_POKER<<16)
	fmt.Printf("扑克-大师赛-周循环赛：%d\n", ROOM_TYPE_DASHISAI_WEEKLY<<8|GAME_TYPE_POKER<<16)

	fmt.Printf("麻将-朋友场：%d\n", ROOM_TYPE_FRIEND<<8|GAME_TYPE_MAHJONG<<16)
	fmt.Printf("麻将-金币场-新手场：%d\n", ROOM_TYPE_NORMAL<<8|ROOM_LEVEL_XINSHOU|GAME_TYPE_MAHJONG<<16)
	fmt.Printf("麻将-金币场-普通场：%d\n", ROOM_TYPE_NORMAL<<8|ROOM_LEVEL_PUTONG|GAME_TYPE_MAHJONG<<16)
}
