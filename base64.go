package main

import (
	"fmt"
	"git.jiaxianghudong.com/go/utils"
)

var decodeStr = `
CguSVgQY2Df4LxG0UT/xwA==
`

var encodeStr = `
package=weile431
`

func main() {
	fmt.Println("base64 decode:\n" + string(utils.Base64Decode(decodeStr)))
	//fmt.Println("base64 encode:\n" + utils.Base64Encode([]byte(encodeStr)))
}
