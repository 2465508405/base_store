/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-09-16 14:13:15
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-09-16 14:13:30
 * @FilePath: /allfunc/gtk_example/window/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"os"

	"github.com/mattn/go-gtk/gtk"
)

func main() {
	gtk.Init(&os.Args) //环境初始化

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL) //创建窗口
	window.SetPosition(gtk.WIN_POS_CENTER)       //设置窗口居中显示
	window.SetTitle("GTK Go!")                   //设置标题
	window.SetSizeRequest(300, 200)              //设置窗口的宽度和高度

	window.Show() //显示窗口

	gtk.Main() //主事件循环，等待用户操作
}
