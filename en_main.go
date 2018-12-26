package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"bufio"
	stat "gonum.org/v1/gonum/stat"
)

func main() {
	// 平均 / 文件 ping2 jun1 wenjian

	数组 := make([]int, 0)

	// 加起来每个文件的线
	数组 = 遍历目录("./", 数组)

	// transform it....
	float数组 := 新float数组(数组)

	// sum
	总 := 加起来(float数组)

	// average (per file) #gonum
	平均 := stat.Mean(float数组, nil)

	// print the results
	fmt.Println("Results shown by File:", 数组)
	fmt.Println("Total Lines:", 总)
	fmt.Println("Average Lines per File:", 平均)
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
	文件, _ := os.Open(名字)
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