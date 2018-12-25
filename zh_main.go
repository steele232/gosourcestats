// +build ignore

包裹 主要

进口 (
	"fmt"
	"io/ioutil"
	"log"
)

程序 主要() {
	fmt.Println("你好")

	// 平均 / 文件 ping2 jun1 wenjian

	数组 := 做([]整数, 0)

	// 加起来每个文件的线
	数组 = 遍历目录("./", 数组)

	// 先做这个。。。
	fmt.Println("Results: ", 数组)

	// sum #gonum

	// average (per file) #gonum

	// print the results

}

// 遍历目录，返回
程序 遍历目录(目录 字符串, 数组 []整数) []整数 {

	// 遍历目录 bianli mulu
	文件, 错 := ioutil.ReadDir(目录 + ".")
	如果 错 != nil {
		log.Fatal(错)
	}

	// 加载文件 jiazai wenjian
	循环 _, f  := 范围 文件 {
		名字 := f.Name()
		如果 f.IsDir() {
			数组 = 遍历目录(目录+名字+"/", 数组) //TODO only works on *nix because I'm using "/"
			前进
		}

		// skip over any non-".go" files
		如果 len(名字) < 3 {
			前进
		}
		如果 名字[len(名字)-3:] != ".go" {
			前进
		}

		// 加起来线 jiaqilai xian4 
		这个整数 := 加起来线一个文件(目录+名字)
		数组 = append(数组, 这个整数)
	}
	return 数组
}

程序 加起来线一个文件(名字 字符串) 整数 {

	// 复制了从 https://stackoverflow.com/questions/24562942/golang-how-do-i-determine-the-number-of-lines-in-a-file-efficiently


	文件bytes, 错 := ioutil.ReadFile(名字)
	如果 错 != nil {
		log.Fatal(错)
	}
	文件字符串 := 字符串(文件bytes)
	fmt.Println(文件字符串)
	
	返回 len(文件bytes)
}
