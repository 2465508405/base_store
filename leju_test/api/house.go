/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-07-17 12:53:10
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-07-19 22:49:52
 * @FilePath: /allfunc/leju_test/api/user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project/allfunc/leju_test/lib/es"

	"github.com/gin-gonic/gin"
)

func GetHouseList(c *gin.Context) {

	data := map[string]interface{}{"name": "ykk", "age": 20} //声明并初始化

	bytesData, _ := json.Marshal(data)
	fmt.Println(bytesData)

	esClient := es.NewClient()

	houseInfo := esClient.Get()
	// res, err := curl.HTTPJson("GET", "http://www.baidu.com/", bytesData)

	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Println(res)

	c.JSON(http.StatusOK, gin.H{
		"message": string(houseInfo),
	})
}
