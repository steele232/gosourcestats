package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Println("你好")

	// 加载文件 jiazai wenjian

	// 遍历目录 bianli mulu
	_, 错 := ioutil.ReadDir("./")
	if 错 != nil {
		log.Fatal(错)
	}


	// 加起来线 jiaqilai xian4 

	// 平均 / 文件 ping2 jun1 wenjian

	// iterate through all files,
	// for _, f := range files {
	// 	if f.IsDir() {
	// 		continue
	// 	}
	// 	name := f.Name()
	// 	if len(name) < 6 {
	// 		continue
	// 	}
	// 	if name[:3] != "zh_" {
	// 		continue
	// 	}
	// 	if name[len(name)-3:] != ".go" {
	// 		continue
	// 	}
	// 	cleanUpZhFile(name) //clean up file BEFORE converting it
	// 	newfilename := "en_" + name[3:]
	// 	writeConvertedFile(name, newfilename)
	// }





	fmt.Println(
		加起来线一个文件("en_main.go"),
	)
	return

}

// 现在return字节
func 加起来线一个文件(名字 string) int {

	文件bytes, 错 := ioutil.ReadFile(名字)
	if 错 != nil {
		log.Fatal(错)
	}
	文件string := string(文件bytes)
	fmt.Println(文件string)
	
	return len(文件bytes)
}
