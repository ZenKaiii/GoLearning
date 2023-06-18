package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is abc@qq.com@123
email1 is abc@def.org
email2 is hello@gmail.com
`

func main() {
	//re := regexp.MustCompile(`.+@.+\..+`) // .+ : 匹配任何字符（.）一个或者多个（+） .* : 0个或者一个（*）但会匹配空格
	//re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+`)
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`) // 提取功能
	//match := re.FindString(text) // 找一个
	//match1 := re.FindAllString(text, -1) // 找所有
	match2 := re.FindAllStringSubmatch(text, -1)
	//println(match)
	//for _, s := range match1 {
	//	println(s)
	//}
	for _, m := range match2 {
		fmt.Println(m)
	}

}
