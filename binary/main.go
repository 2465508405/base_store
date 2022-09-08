/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-09-07 18:27:50
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-09-08 14:11:04
 * @FilePath: /allfunc/websocket/tcp/binary/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

type BinaryStr struct {
}

func (b *BinaryStr) Opera() {
	buf := make([]byte, binary.MaxVarintLen64)
	fmt.Print(buf)
	// for _, x := range []uint64{1, 2, 127, 128, 255, 256} {
	// 	n := binary.PutUvarint(buf, x)
	// 	fmt.Println("num:", n)
	// 	fmt.Printf("%x\n", buf[:n])
	// }
	// for _, x := range []int64{-65, -64, -2, -1, 0, 1, 2, 63, 64} {
	// 	n := binary.PutVarint(buf, x)
	// 	fmt.Printf("%x\n", buf[:n])
	// }
	// fmt.Println(buf)

	// b.Read()
	// b.Multi()
	b.Write()

}

func (b *BinaryStr) Read() {

	var pi float64
	c := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	fmt.Println(int(c[0]))
	buf := bytes.NewReader(c)
	err := binary.Read(buf, binary.LittleEndian, &pi)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	fmt.Print(pi)
}

func (bs *BinaryStr) Multi() {
	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40, 0xff, 0x01, 0x02, 0x03, 0xbe, 0xef}
	r := bytes.NewReader(b)

	var data struct {
		PI   float64
		Uate uint8
		Mine [3]byte
		Too  uint16
	}

	if err := binary.Read(r, binary.LittleEndian, &data); err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	fmt.Println(data.PI)
	fmt.Println(data.Uate)
	fmt.Printf("% x\n", data.Mine)
	fmt.Println(data.Too)
}

func (bs BinaryStr) Write() {
	buf := new(bytes.Buffer)
	var pi float64 = math.Pi
	err := binary.Write(buf, binary.LittleEndian, pi)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	fmt.Printf("% x", buf.Bytes())
}
func main() {

	b := BinaryStr{}
	b.Opera()
	// message := "afafafafacesf"
	// var length = int32(len(message))

	// var pkg = new(bytes.Buffer)
	// err := binary.Write(pkg, binary.BigEndian, length)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(pkg.Bytes())

	// // 写入消息实体
	// err = binary.Write(pkg, binary.BigEndian, []byte(message))
	// if err != nil {
	// 	panic(err)
	// }

	// msg := pkg.Bytes()
	// fmt.Println("msg:", msg)
	// var len int32

	// //获取len类型的长度的值，返回要获取的数据的长度
	// binary.Read(pkg, binary.BigEndian, &len)
	// a := pkg.String()
	// fmt.Println(a)
	// m := a[:]
	// fmt.Println("s:", len)
	// fmt.Println("m:", string(m))
	// i := IntToBytes(500)

	// fmt.Println(i)
	// fmt.Println("======大端序========")
	// bytes := make([]byte, 16)
	// binary.BigEndian.PutUint32(bytes, 500)
	// fmt.Println(bytes)
	// fmt.Printf("value:%d \n", binary.BigEndian.Uint32(bytes[:4]))
	// binary.BigEndian.PutUint32(bytes[4:], 500)
	// fmt.Println(bytes)
	// fmt.Printf("value:%d \n", binary.BigEndian.Uint32(bytes[:4]))

	// fmt.Println("======小端序========")

	// bytes2 := make([]byte, 16)
	// binary.LittleEndian.PutUint32(bytes2, 500)
	// fmt.Println(bytes2)
	// fmt.Printf("value:%d\n", binary.LittleEndian.Uint64(bytes2))
	// binary.LittleEndian.PutUint32(bytes2[4:], 500)
	// fmt.Println(bytes2)
}

func IntToBytes(n int) []byte {
	x := uint32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}
