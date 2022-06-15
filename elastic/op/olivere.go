package op

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"reflect"

	"github.com/olivere/elastic/v7"
)

//这里使用的是版本5，最新的是6，有改动

var client *elastic.Client
var host = "http://127.0.0.1:9200/"

type Employee struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	About     string   `json:"about"`
	Interests []string `json:"interests"`
}

//初始化
func init() {
	errorlog := log.New(os.Stdout, "ELASTIC", log.LstdFlags)
	var err error
	client, err = elastic.NewClient( // elasticsearch 服务地址，多个服务地址使用逗号分隔
		elastic.SetURL(host),
		//不加上elastic.SetSniff(false) 会连接不上
		elastic.SetSniff(false),
		// 基于http base auth验证机制的账号和密码
		elastic.SetBasicAuth("elastic", "123456"),
		// 启用gzip压缩
		elastic.SetGzip(true),
		// 设置监控检查时间间隔
		elastic.SetHealthcheckInterval(10*time.Second),
		// 设置请求失败最大重试次数
		// elastic.SetMaxRetries(5),
		// 设置错误日志输出
		elastic.SetErrorLog(errorlog),
		// 设置info日志输出
		elastic.SetInfoLog(log.New(os.Stdout, "APP", log.LstdFlags)),
	)
	if err != nil {
		panic(err)
	}
	info, code, err := client.Ping(host).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	esversion, err := client.ElasticsearchVersion(host)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

}

/*下面是简单的CURD*/

//创建
func Create() {

	//使用结构体
	e1 := Employee{"Jane", "Smith", 32, "I like to collect rock albums", []string{"music"}}
	put1, err := client.Index().
		Index("megacorp").
		Id("1").
		BodyJson(e1).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put1.Id, put1.Index, put1.Type)

	//使用字符串
	e2 := `{"first_name":"John","last_name":"Smith","age":25,"about":"I love to go rock climbing","interests":["sports","music"]}`
	put2, err := client.Index().
		Index("megacorp").
		Id("2").
		BodyJson(e2).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put2.Id, put2.Index, put2.Type)

	e3 := `{"first_name":"Douglas","last_name":"Fir","age":35,"about":"I like to build cabinets","interests":["forestry"]}`
	put3, err := client.Index().
		Index("megacorp").
		Id("3").
		BodyJson(e3).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put3.Id, put3.Index, put3.Type)

}

//删除
func Delete() {

	res, err := client.Delete().Index("megacorp").
		Id("1").
		Do(context.Background())
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("delete result %s\n", res.Result)
}

//修改
func Update() {
	res, err := client.Update().
		Index("megacorp").
		Id("2").
		Doc(map[string]interface{}{"age": 89}).
		Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	fmt.Printf("update age %s\n", res.Result)

}

//查找
func Gets() {
	//通过id查找
	get1, err := client.Get().Index("megacorp").Id("2").Do(context.Background())
	if err != nil {
		panic(err)
	}
	if get1.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
	}
	data, _ := get1.Source.MarshalJSON()
	emp := Employee{}
	json.Unmarshal(data, &emp)
	fmt.Println(emp)
}

//搜索
func Query() {
	var res *elastic.SearchResult
	var err error
	//取所有
	res, err = client.Search("megacorp").Do(context.Background())
	PrintEmployee(res, err)

	//字段相等
	q := elastic.NewQueryStringQuery("last_name:Smith")
	res, err = client.Search("megacorp").Query(q).Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	PrintEmployee(res, err)

	//条件查询
	//年龄大于30岁的
	boolQ := elastic.NewBoolQuery()
	boolQ.Must(elastic.NewMatchQuery("last_name", "smith"))
	boolQ.Filter(elastic.NewRangeQuery("age").Gt(30))
	res, err = client.Search("megacorp").Query(q).Do(context.Background())
	PrintEmployee(res, err)

	//短语搜索 搜索about字段中有 rock climbing
	matchPhraseQuery := elastic.NewMatchPhraseQuery("about", "rock climbing")
	res, err = client.Search("megacorp").Query(matchPhraseQuery).Do(context.Background())
	PrintEmployee(res, err)

	// 执行ES请求需要提供一个上下文对象
	ctx := context.Background()

	// 创建Value Count指标聚合
	aggs := elastic.NewValueCountAggregation().
		Field("age") // 设置统计字段

	searchResult, err := client.Search().
		Index("megacorp").                 // 设置索引名
		Query(elastic.NewMatchAllQuery()). // 设置查询条件
		Aggregation("total", aggs).        // 设置聚合条件，并为聚合条件设置一个名字, 支持添加多个聚合条件，命名不一样即可。
		Size(0).                           // 设置分页参数 - 每页大小,设置为0代表不返回搜索结果，仅返回聚合分析结果
		Do(ctx)                            // 执行请求

	if err != nil {
		// Handle error
		panic(err)
	}

	// 使用ValueCount函数和前面定义的聚合条件名称，查询结果
	agg, found := searchResult.Aggregations.ValueCount("total")
	if found {
		// 打印结果，注意：这里使用的是取值运算符
		fmt.Println(*agg.Value)
	}

	//分析 interests
	// aggsT := elastic.NewTermsAggregation().Field("interests")
	// res, err = client.Search().Index("megacorp").Query(elastic.NewMatchAllQuery()).Aggregation("alinterests", aggsT).Size(0).Do(context.Background())
	// PrintEmployee(res, err)

	// agg, found = res.Aggregations.ValueCount("alinterests")
	// if !found {
	// 	panic(found)
	// }
	// fmt.Println(*agg.Value)

}

//简单分页
func List(size, page int) {
	if size < 0 || page < 1 {
		fmt.Printf("param error")
		return
	}
	res, err := client.Search("megacorp").
		Size(size).
		From((page - 1) * size).
		Do(context.Background())
	PrintEmployee(res, err)

}

//打印查询到的Employee
func PrintEmployee(res *elastic.SearchResult, err error) {
	if err != nil {
		print(err.Error())
		return
	}
	var typ Employee
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(Employee)
		fmt.Printf("%#v\n", t)
	}
	fmt.Println("==========")
}
