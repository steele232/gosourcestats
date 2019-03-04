// +build ignore

包裹 主要

进口 (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"bufio"
	"sort"
	stat "gonum.org/v1/gonum/stat"
)

函数 主要() {

	数组 := 初始化([]整数, 0)

	// 加起来每个文件的线
	数组 = 遍历目录("./", 数组)

	// 变换类型
	浮点数组 := 新浮点数组(数组)

	// 加起来总计的线
	总 := 加起来(浮点数组)

	// 平均，使用gonum库
	平均 := stat.Mean(浮点数组, nil)

	// 中位数, 使用gonum库
	sort.Float64s(浮点数组)
	中位数 := stat.Quantile(0.5, stat.Empirical, 浮点数组, nil)

	//
	最大 := 找到最大(浮点数组)

	//
	最小 := 找到最小(浮点数组, 最大)

	// print the results
	// fmt.Println("Results shown by File:", 数组)
	fmt.Println("**********************")
	fmt.Println("* Total Number of Files:", len(数组))
	fmt.Println("* Total Lines:", 总)
	fmt.Println("* Mean Lines per File:", 平均)
	fmt.Println("* Median Lines per File:", 中位数)
	fmt.Println("* Largest Go File:", 最大, "lines")
	fmt.Println("* Smallest Go File:", 最小, "lines")
	fmt.Println("**********************")
}

// 遍历目录，返回
函数 遍历目录(目录 字符串, 数组 []整数) []整数 {

	// 遍历目录
	文件, 错 := ioutil.ReadDir(目录 + ".")
	如果 错 != nil {
		log.Fatal(错)
	}

	// 加载文件
	循环 _, f  := 范围 文件 {
		名字 := f.Name()
		如果 f.IsDir() {
			数组 = 遍历目录(目录+名字+"/", 数组) //TODO：应该支持所有操作系统
			继续
		}

		// 跳过都不".go"文件
		如果 len(名字) < 3 {
			继续
		}
		如果 名字[len(名字)-3:] != ".go" {
			继续
		}

		// 加起来线
		这个整数 := 加起来线一个文件(目录+名字)
		数组 = append(数组, 这个整数)
	}
	返回 数组
}

函数 加起来线一个文件(名字 字符串) 整数 {

	// 复制了从 https://stackoverflow.com/questions/24562942/golang-how-do-i-determine-the-number-of-lines-in-a-file-efficiently
	文件, 错 := os.Open(名字)
	如果 错 != nil {
		log.Fatal(错)
	}
	推迟 文件.Close()
	文件看者 := bufio.NewScanner(文件)

	总 := 0
	循环 文件看者.Scan() {
		总++
	}
	返回 总
}

函数 新浮点数组(老 []整数) []浮点数64 {
	新 := 初始化([]浮点数64, len(老))
	循环 i := 0; i < len(老); i++ {
		新[i] = 浮点数64(老[i])
	}
	返回 新
}

函数 加起来(数组 []浮点数64) 浮点数64 {
	变量 总 浮点数64
	循环 i := 0; i < len(数组); i++ {
		总 += 数组[i]
	}
	返回 总
}

函数 找到最大(数组 []浮点数64) 浮点数64 {
	变量 最大 浮点数64 // 从0开始
	循环 i := 0; i < len(数组); i++ {
		如果 数组[i] > 最大 {
			最大 = 数组[i]
		}
	}
	返回 最大
}

函数 找到最小(数组 []浮点数64, 最大 浮点数64) 浮点数64 {
	最小 := 最大
	循环 i := 0; i < len(数组); i++ {
		如果 数组[i] < 最小 {
			最小 = 数组[i]
		}
	}
	返回 最小
}