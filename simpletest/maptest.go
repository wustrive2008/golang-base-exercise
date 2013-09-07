package main

//统计句子中单词的出现次数，存放到map中
import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	var stmap map[string]int = make(map[string]int)
	strs := strings.Fields(s)
	for _, str := range strs {

		index, ok := stmap[str]

		if ok {
			index++
			stmap[str] = index
		} else {
			stmap[str] = 1
		}

	}
	return stmap
}

func main() {
	fmt.Println(WordCount("gocode src gotest map test src gocode gotest src src"))
}
