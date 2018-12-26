// +build ignore

包裹 主要

进口 (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"bufio"
	stat "gonum.org/v1/gonum/stat"
)

程序 主要() {
	// 平均 / 文件 ping2 jun1 wenjian

	数组 := 做([]整数, 0)

	// 加起来每个文件的线
	数组 = 遍历目录("./", 数组)

	// transform it....
	浮点数组 := 新浮点(数组)

	// sum
	总 := 加起来(浮点数组)

	// average (per file) #gonum
	平均 := stat.Mean(浮点数组, nil)

	// print the results
	fmt.Println("Results shown by File:", 数组)
	fmt.Println("Total Lines:", 总)
	fmt.Println("Average Lines per File:", 平均)


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
	文件, _ := os.Open(名字)
	文件看者 := bufio.NewScanner(文件)
	总 := 0
	循环 文件看者.Scan() {
		总++
	}
	返回 总
}

程序 新浮点(老 []整数) []浮点64 {
	新 := 做([]浮点64, len(老))
	循环 i := 0; i < len(老); i++ {
		新[i] = 浮点64(老[i])
	}
	返回 新
}

程序 加起来(数组 []浮点64) 浮点64 {
	变量 总 浮点64
	循环 i := 0; i < len(数组); i++ {
		总 += 数组[i]
	}
	返回 总
}