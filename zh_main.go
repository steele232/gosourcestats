// +build ignore

包裹 主要

进口 (
	"fmt"
	"io/ioutil"
	"log"
)

程序 主要() {
	fmt.Println("你好")

	// 加载文件 jiazai wenjian

	// 遍历目录 bianli mulu
	_, 错 := ioutil.ReadDir("./")
	如果 错 != nil {
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
	返回

}

// 现在返回字节
程序 加起来线一个文件(名字 字符串) 整数 {

	文件bytes, 错 := ioutil.ReadFile(名字)
	如果 错 != nil {
		log.Fatal(错)
	}
	文件字符串 := 字符串(文件bytes)
	fmt.Println(文件字符串)
	
	返回 len(文件bytes)
}
