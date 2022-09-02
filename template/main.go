/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-31 14:00:15
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-31 17:23:07
 * @FilePath: /allfunc/template/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

type Product struct {
	Name string
	Age  int64
}

var (
	//生成的Html保存目录
	htmlOutPath = "./htmlProductShow/"
	//静态文件模版目录
	templatePath = "./views/template/"
)

func GetGenerateHtml() {

	//1.获取模版
	contenstTmp, err := template.ParseFiles(filepath.Join(templatePath, "product.html"))
	if err != nil {
		fmt.Println(err)
	}
	//2.获取html生成路径
	fileName := filepath.Join(htmlOutPath, "htmlProduct.html")
	//3.获取模版渲染数据
	product := Product{Name: "ykk", Age: 232}
	//4.生成静态文件
	generateStaticHtml(contenstTmp, fileName, product)
}

//生成html静态文件
func generateStaticHtml(template *template.Template, fileName string, product Product) {
	//1.判断静态文件是否存在
	if exist(fileName) {
		err := os.Remove(fileName)
		if err != nil {
			fmt.Println(err)
		}
	}

	//2.生成静态文件
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	template.Execute(file, &product)
}

//判断文件是否存在
func exist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
}

func main() {
	GetGenerateHtml()
}
