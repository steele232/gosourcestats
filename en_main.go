package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Println("你好")

	// 平均 / 文件 ping2 jun1 wenjian

	数组 := make([]int, 0)

	// 加起来每个文件的线
	数组 = 遍历目录("./", 数组)

	// 先做这个。。。
	fmt.Println("Results: ", 数组)

	// sum #gonum

	// average (per file) #gonum

	// print the results

}

// 遍历目录，返回
func 遍历目录(目录 string, 数组 []int) []int {

	// 遍历目录 bianli mulu
	文件, 错 := ioutil.ReadDir(目录 + ".")
	if 错 != nil {
		log.Fatal(错)
	}

	// 加载文件 jiazai wenjian
	for _, f  := range 文件 {
		名字 := f.Name()
		if f.IsDir() {
			数组 = 遍历目录(目录+名字+"/", 数组) //TODO only works on *nix because I'm using "/"
			continue
		}

		// skip over any non-".go" files
		if len(名字) < 3 {
			continue
		}
		if 名字[len(名字)-3:] != ".go" {
			continue
		}

		// 加起来线 jiaqilai xian4 
		这个int := 加起来线一个文件(目录+名字)
		数组 = append(数组, 这个int)
	}
	return 数组
}

func 加起来线一个文件(名字 string) int {

	// 复制了从 https://stackoverflow.com/questions/24562942/golang-how-do-i-determine-the-number-of-lines-in-a-file-efficiently


	文件bytes, 错 := ioutil.ReadFile(名字)
	if 错 != nil {
		log.Fatal(错)
	}
	文件string := string(文件bytes)
	fmt.Println(文件string)
	
	return len(文件bytes)
}
