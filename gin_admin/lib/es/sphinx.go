/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-07-18 23:00:30
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-10 10:58:52
 * @FilePath: /allfunc/leju_test/lib/es/sphinx.go
 */
package es

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"project/allfunc/gin_admin/lib/curl"
)

type FwSphinx struct {
	app_type  string
	appid     string
	app_key   string
	condition map[string][]string
	keyword   string
	page      int32
	pcount    int32
	count     int32
	order     string
	time_out  int32
}

const ApiUrl = "http://info.leju.com/search/default/index"

func NewClient() *FwSphinx {

	return &FwSphinx{}
}

func (s *FwSphinx) Type(apiType string) *FwSphinx {

	switch apiType {
	case "house": //楼盘
		s.appid = "2016112598"
		s.app_key = "871196a770a984d5882c5a5df5a7494e"
		s.app_type = "house"
	case "data_set": //数据聚合
		s.appid = "2018070959"
		s.app_key = "30d85f29ae458ae20bf53623eed7cdbb"
		s.app_type = "data_set"
	default:
		s.appid = "2016112598"
		s.app_key = "871196a770a984d5882c5a5df5a7494e"
		s.app_type = "house"
	}
	s.pcount = 10
	s.page = 1
	s.count = 1
	s.time_out = 20
	return s
}

func (s *FwSphinx) Where(con map[string][]string) *FwSphinx {
	for key, c := range con {
		s.condition[key] = c
	}
	return s
}

func (s *FwSphinx) WhereIn(keyword string) *FwSphinx {
	s.keyword = keyword
	return s
}

func (s *FwSphinx) Page(page int32) *FwSphinx {
	s.page = page
	return s
}

func (s *FwSphinx) Limit(size int32) *FwSphinx {
	s.pcount = size
	return s
}

func (s *FwSphinx) Count(size int32) *FwSphinx {
	s.pcount = size
	s.page = 1
	return s
}

func (s *FwSphinx) Orderby(order string) *FwSphinx {
	s.order = order
	return s
}

func (s *FwSphinx) Get() []byte {

	s.condition = map[string][]string{}
	s.condition["type"] = []string{"house"}
	s.condition["appid"] = []string{"2016112598"}
	s.condition["pcount"] = []string{"10"}
	s.condition["page"] = []string{"1"}
	s.condition["count"] = []string{"1"}
	s.condition["field"] = []string{"keyid|hid|site|house_id|name|price_avg|price_display|coordx2|coordy2|distance|site|pic_s|tags_id|esf_id2|salestate|salestate_name|main_housetype|phone_extension|district|district_name|area|area_name|address|developer|special|licence|main_housetype|hometype_name|cover|saleaddress|opentime|delivertime|salestate_name|city_cn|pic_s320|subway|pricerange|hometype|hometype_name|hometype|hometype_name|price_display_info|opentime_desc|building_area|price_sum|totalrange|circlelocation|pic_hx_on_sale_area_range|is_vr|is_video|is_live|house_unique_id|house_brand_hall_ids|house_vr_search_sign|house_excellent_search_sign|house_excellent_search_sign|spring_work_start|spring_work_end|saleshouse_range_hours|salestateorder|updatetime|coupon_activity|is_house_review|functional_district_name|house_ticket_multi_search_sign|house_ticket_multi_search_platform"}
	s.condition["filter1"] = []string{"{status@eq}1"}
	s.condition["filter2"] = []string{"{hid@neq}0"}
	s.condition["filter3"] = []string{"{salestate@eq}1|2|3|10|4|11"}
	s.condition["filter4"] = []string{"{relation_city@eq}ab"}
	s.condition["order"] = []string{"{salestateorder@desc|house_hot_score@desc|updatetime@desc}desc"}
	s.condition["sign"] = []string{s.getSign(s.condition)}
	// params, err := json.Marshal(s.condition)
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println(s.condition)
	// params := s.condition
	params := s.condition
	res, err := curl.HTTPDo("POST", ApiUrl, params)
	if err != nil {

		panic(err)
	}
	return res
}

func (s *FwSphinx) getSign(params map[string][]string) string {
	str := ""

	for _, val := range params {
		fmt.Println(val)

		str = str + val[0]
		// fmt.Println(str)
	}
	has := []byte(str + s.app_key)

	md5 := md5.New()
	md5.Write(has)
	hex_string_data := hex.EncodeToString(md5.Sum(nil))
	// md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	// hex_string_data := base64.StdEncoding.EncodeToString([]byte(md5str1))
	// fmt.Println(base64.StdEncoding.EncodeToString([]byte(md5str1)))
	// fmt.Println(hex_string_data)
	return hex_string_data
}
