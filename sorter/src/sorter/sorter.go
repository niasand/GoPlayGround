/*
 * @Script: sorter.go
 * @Author: Zhiwei.Yang
 * @Email: tencrance@gmail.com
 * @Create At: 2018-07-30 22:09:54
 * @Last Modified By: Zhiwei.Yang
 * @Last Modified At: 2018-07-31 00:11:05
 * @Description: This is description.
 */

package main

import (
	"algorithms/bubblesort" // 导入算法文件夹下面的冒泡排序包，这样下面的函数就可以调用的到
	"algorithms/qsort"      // 导入算法文件夹下面的快速排序包，这样下面的函数就可以调用的到
	"bufio"                 //系统自带的读写文件的包
	"flag"                  //类似python的argparse
	"fmt"                   // 打印
	"io"                    //io操作
	"os"                    // 打开 关闭
	"strconv"               //转换字符串到int等类型转换
	"time"                  //时间包 类似python的datetime
)

var infile *string = flag.String("i", "unsorted.dat", "File contains values for sorting")

// 输入文件，参数为i，默认值为unsorted.dat，后面一段话是help的介绍
var outfile *string = flag.String("o", "outfile.dat", "File to receive sorted values")

// 输出文件，参数为o，默认值为outfile.dat，后面一段话是help的介绍
var algorithm *string = flag.String("a", "qsort", "sort algorithms")

//算法选型，参数为a，默认是快排，后面一段话是介绍
func readValues(infile string) (values []int, err error) {
	//定义一个读文件的函数，文件名是infile是字符串类型，然后返回值是数组和err
	file, err := os.Open(infile) //打开一个文件
	if err != nil {              //判断是否有错误，如果err不为空，说明有err，那么就打印失败的信息并返回
		fmt.Println("Failed to open the input file:", infile)
		return
	}
	defer file.Close()          //最后一步执行，相当于清理环境
	br := bufio.NewReader(file) //new一个对象，是去读取文件
	values = make([]int, 0)     // 新建一个cap为0的空数组，后面去append操作

	for { //搞一个无限循环，条件是读完所有行就停止
		line, isPrefix, err1 := br.ReadLine() //开始一行一行的读
		if err1 != nil {                      //判断是否有错误，如果不为空
			if err1 != io.EOF { //判断是否读到最后了
				err = err1 //将err1的值赋值给err
			}
			break //中断循环
		}
		if isPrefix {
			fmt.Println("A too long line, seems unexpected!")
			return
		}
		str := string(line)              //转换字符组为字符串
		value, err1 := strconv.Atoi(str) //看atoi源码，是转换str为int64
		if err1 != nil {
			err = err1
			// return
		}

		values = append(values, value) //append读取到的数值到数组。要是python，几行就搞定了

	}
	return
}

func writeValues(values []int, outfile string) error {
	//定一个一个写函数
	file, err := os.Create(outfile) //新建一个文件
	if err != nil {                 //如果有错误，则打印错误消息
		fmt.Println("Failed to create the output file", outfile)
		return err
	}

	defer file.Close()             //最后一步操作 关闭文件，很优雅
	for _, value := range values { //循环数组，并将数组中的数字编程字符串，写入文件
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}

func main() {
	flag.Parse()
	if infile != nil { //main函数，判断infile是否为空，如果不为空 则报错
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
	}
	values, err := readValues(*infile) //调用读取函数
	if err == nil {                    //如果没有error，就开始下面的流程
		fmt.Println("Read values:", values) // 打印读取到的值，可以当调试用
		t1 := time.Now()                    //打印当前时间
		switch *algorithm {                 //转换器 两种算法，如果输入的参数是qsort，则用qsort
		case "qsort":
			qsort.QuickSort(values)
		case "bubblesort":
			bubblesort.BubbleSort(values)
		default: //其他输入就报错
			fmt.Println("Sorting algorithms", *algorithm, "is either unkonw or unsupported")
		}
		t2 := time.Now() //为了算执行时间用的
		fmt.Println("The sort process costs", t2.Sub(t1), "to complete")
		writeValues(values, *outfile) //写到输出的文件里头

	} else {
		fmt.Println(err) //如果有错误，就跑到这里打印下。
	}
}
