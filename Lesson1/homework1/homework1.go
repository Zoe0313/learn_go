/*
给定一个字符串数组 ["I","am","stupid","and","weak"] 用 for 循环遍历该数组并修改为["I","am","smart","and","strong"]
*/

package main

import "fmt"

func main() {
	var words = []string{"I", "am", "stupid", "and", "weak"}
	for i, word := range words {
		if word == "stupid" {
			words[i] = "smart"
		} else if word == "weak" {
			words[i] = "strong"
		}
	}
	fmt.Println(words)
}
