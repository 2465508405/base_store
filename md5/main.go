package main

import (
	"fmt"
	"net/url"
)

/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-07-22 13:50:38
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-07-23 15:25:22
 * 加密数据与  php 保持一致
 */
type stu struct {
	condition map[string][]string
}

func main() {
	s := stu{}
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
	// s.condition["sign"] = []string{s.getSign(s.condition)}
	// params, err := json.Marshal(s.condition)
	// if err != nil {
	// 	panic(err)
	// }
	v := url.Values{}
	account := "{hid@neq}0"
	v.Set("account", account)
	// fmt.Println(s.condition)
	fmt.Println(v)

}
