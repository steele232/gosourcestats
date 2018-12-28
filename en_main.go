package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"bufio"
	"sort"
	stat "gonum.org/v1/gonum/stat"
)

func main() {

	数组 := make([]int, 0)

	// 加起来每个文件的线
	数组 = 遍历目录("./", 数组)

	// 变换类型
	float数组 := 新float数组(数组)

	// 加起来总计的线
	总 := 加起来(float数组)

	// 平均，使用gonum库
	平均 := stat.Mean(float数组, nil)

	// 中位数, 使用gonum库
	sort.Float64s(float数组)
	中位数 := stat.Quantile(0.5, stat.Empirical, float数组, nil)

	//
	最大 := 找到最大(float数组)

	//
	最小 := 找到最小(float数组, 最大)

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
func 遍历目录(目录 string, 数组 []int) []int {

	// 遍历目录
	文件, 错 := ioutil.ReadDir(目录 + ".")
	if 错 != nil {
		log.Fatal(错)
	}

	// 加载文件
	for _, f  := range 文件 {
		名字 := f.Name()
		if f.IsDir() {
			数组 = 遍历目录(目录+名字+"/", 数组) //TODO：应该支持所有操作系统
			continue
		}

		// 跳过都不".go"文件
		if len(名字) < 3 {
			continue
		}
		if 名字[len(名字)-3:] != ".go" {
			continue
		}

		// 加起来线
		这个int := 加起来线一个文件(目录+名字)
		数组 = append(数组, 这个int)
	}
	return 数组
}

func 加起来线一个文件(名字 string) int {

	// 复制了从 https://stackoverflow.com/questions/24562942/golang-how-do-i-determine-the-number-of-lines-in-a-file-efficiently
	文件, 错 := os.Open(名字)
	if 错 != nil {
		log.Fatal(错)
	}
	defer 文件.Close()
	文件看者 := bufio.NewScanner(文件)

	总 := 0
	for 文件看者.Scan() {
		总++
	}
	return 总
}

func 新float数组(老 []int) []float64 {
	新 := make([]float64, len(老))
	for i := 0; i < len(老); i++ {
		新[i] = float64(老[i])
	}
	return 新
}

func 加起来(数组 []float64) float64 {
	var 总 float64
	for i := 0; i < len(数组); i++ {
		总 += 数组[i]
	}
	return 总
}

func 找到最大(数组 []float64) float64 {
	var 最大 float64 // 开始在0
	for i := 0; i < len(数组); i++ {
		if 数组[i] > 最大 {
			最大 = 数组[i]
		}
	}
	return 最大
}

func 找到最小(数组 []float64, 最大 float64) float64 {
	最小 := 最大
	for i := 0; i < len(数组); i++ {
		if 数组[i] < 最小 {
			最小 = 数组[i]
		}
	}
	return 最小
}