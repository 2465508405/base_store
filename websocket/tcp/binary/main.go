/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-09-07 18:27:50
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-09-08 10:12:46
 * @FilePath: /allfunc/websocket/tcp/binary/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type BinaryStru struct {
}

func (b *BinaryStru) SmallBufferSize() {

}
func main() {

	i := IntToBytes(500)

	fmt.Println(i)
	fmt.Println()
	fmt.Println("======大端序========")
	bytes := make([]byte, 16)
	binary.BigEndian.PutUint32(bytes, 500)
	fmt.Println(bytes)
	fmt.Printf("value:%d \n", binary.BigEndian.Uint32(bytes[:4]))
	binary.BigEndian.PutUint32(bytes[4:], 500)
	fmt.Println(bytes)
	fmt.Printf("value:%d \n", binary.BigEndian.Uint32(bytes[:4]))

	fmt.Println("======小端序========")

	bytes2 := make([]byte, 16)
	binary.LittleEndian.PutUint32(bytes2, 500)
	fmt.Println(bytes2)
	fmt.Printf("value:%d\n", binary.LittleEndian.Uint64(bytes2))
	binary.LittleEndian.PutUint32(bytes2[4:], 500)
	fmt.Println(bytes2)
}

func IntToBytes(n int) []byte {
	x := uint32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}
