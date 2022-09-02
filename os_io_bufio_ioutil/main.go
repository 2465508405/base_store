/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-29 16:47:36
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-31 18:31:53
 * @FilePath: /allfunc/os_io/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"project/allfunc/os_io_bufio_ioutil/relate"
)

func file_open() {
	// 只读方式打开当前目录下的main.go文件
	file, err := os.Open("./main.txt")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	fmt.Println("open succ")
	// 关闭文件
	defer file.Close()
}

func main() {

	get_pwd() //获取文件目录
	// os_stat()//文件是否存在
	// file_open()
	// file_write()
	// file_remove()
	// file_read()
	// file_copy()
	// bufio_re()
	// bufio_wr()
	// ioutil_re()

	//模拟cat命令
	// flag.Parse() // 解析命令行参数
	// if flag.NArg() == 0 {
	// 	// 如果没有参数默认从标准输入读取内容
	// 	cat(bufio.NewReader(os.Stdin))
	// }
	// fmt.Println(flag.Arg(0))
	// // 依次读取每个指定文件的内容并打印到终端
	// for i := 0; i < flag.NArg(); i++ {
	// 	f, err := os.Open(flag.Arg(i))
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stdout, "reading from %s failed, err:%v\n", flag.Arg(i), err)
	// 		continue
	// 	}

	// 	cat(bufio.NewReader(f))
	// }
}

//文件是否存在
func os_stat() {

	stat, err := os.Stat("./main.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(stat.Mode())
}

//获取文件目录

func get_pwd() {
	// dir, _ := os.Getwd() //获取当前目录

	relate.GetPwd()
}

// cat命令实现
func cat(r *bufio.Reader) {
	for {
		buf, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
}

func ioutil_wr() {
	err := ioutil.WriteFile("./main.txt", []byte("www.5lmh.com"), 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ioutil_re() {
	content, err := ioutil.ReadFile("./main.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}

func bufio_wr() {
	// 参数2：打开模式，所有模式d都在上面
	// 参数3是权限控制
	// w写 r读 x执行   w  2   r  4   x  1
	file, err := os.OpenFile("./main.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	// 获取writer对象
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello\n")
	}
	// 刷新缓冲区，强制写出
	writer.Flush()
}

func bufio_re() {
	file, err := os.Open("./main.txt")
	if err != nil {
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		fmt.Println(string(line))
	}

}

//拷贝文件
func file_copy() {
	// 打开源文件
	srcFile, err := os.Open("./main.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 创建新文件
	dstFile, err2 := os.Create("./main_bak.txt")
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	// 缓冲读取
	buf := make([]byte, 1024)
	for {
		// 从源文件读数据
		n, err := srcFile.Read(buf)
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		//写出去
		dstFile.Write(buf[:n])
	}
	srcFile.Close()
	dstFile.Close()

}

//读文件
func file_read() {
	file, err := os.Open("./main.txt")
	if err != nil {
		fmt.Println("open file err :", err)
		return
	}
	defer file.Close()

	// 定义接收文件读取的字节数组
	var buf [128]byte
	var content []byte
	//从某个位置开始写
	// tmp := make([]byte, 128)
	// file.ReadAt(tmp, 3)
	// fmt.Println(string(tmp))
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			// 读取结束
			break
		}
		if err != nil {
			fmt.Println("read file err ", err)
			return
		}
		content = append(content, buf[:n]...)
	}
	fmt.Println(string(content))

}

//写文件
func file_write() {
	// 新建文件
	// file, err := os.Create("./main.txt")
	file, err := os.OpenFile("./main.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	for i := 0; i < 5; i++ {
		file.WriteString("ab\n")
		file.Write([]byte("cd\n"))
		fmt.Println(77)

	}

	/**
	*文件开启追加写，会报错
	*If file was opened with the O_APPEND flag, WriteAt returns an error.
	 */
	file.WriteAt([]byte("gg\n"), int64(1))

}

func file_remove() {
	os.Remove("./main.txt")
}
