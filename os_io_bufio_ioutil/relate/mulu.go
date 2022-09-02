/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-31 18:17:49
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-09-01 10:01:45
 * @FilePath: /allfunc/os_io_bufio_ioutil/relate/mulu.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package relate

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

/**

go run会将源代码编译到系统TEMP或TMP环境变量目录中并启动执行,会获取到临时文件目录地址
而go build只会在当前目录编译出可执行文件，并不会自动执行。

*/

func GetPwd() {
	dir, _ := filepath.Abs("./") //获取绝对目录
	fmt.Println(dir)

	dir1, _ := os.Getwd() //获取绝对目录
	fmt.Println(dir1)

	//获取当前文件所在绝对路径
	// var abPath string
	// _, filename, _, ok := runtime.Caller(0)
	// if ok {
	// 	abPath = path.Dir(filename)
	// }
	// fmt.Println(abPath)

	di := getCurrentAbPath()
	fmt.Println(di)
	// return res
}

// 最终方案-全兼容
func getCurrentAbPath() string {
	dir := getCurrentAbPathByExecutable()
	if !strings.Contains(dir, getTmpDir()) {
		return getCurrentAbPathByCaller()
	}
	fmt.Println(dir)
	return dir
}

// 获取系统临时目录，兼容go run
func getTmpDir() string {
	dir := os.Getenv("TEMP")
	fmt.Println(77777)
	fmt.Println(dir)
	if dir == "" {
		dir = os.Getenv("TMP")
	}
	res, _ := filepath.EvalSymlinks(dir)
	fmt.Println(333)
	fmt.Println(res)
	return res
}

// 获取当前执行文件绝对路径
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	fmt.Println(9999)
	fmt.Println(exePath)
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	fmt.Println(res)
	return res
}

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	fmt.Println(222222)
	fmt.Println(abPath)
	return abPath
}
