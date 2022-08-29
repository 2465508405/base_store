package main

import (
	"fmt"
	"reflect"
)

//反射获取interface类型信息

func reflect_type(a interface{}) {

	t := reflect.TypeOf(a)
	fmt.Println("类型是：", t)
	// kind()可以获取具体类型
	k := t.Kind()
	fmt.Println(k)
	switch k {
	case reflect.Float64:
		fmt.Printf("a is float64\n")
	case reflect.String:
		fmt.Println("string")
	}
}

//反射获取interface值信息

func reflect_value(a interface{}) {
	v := reflect.ValueOf(a)
	fmt.Println(7777)
	fmt.Println(v)
	k := v.Kind()
	fmt.Println(k)
	switch k {
	case reflect.Float64:
		fmt.Println("a是：", v.Float())
	case reflect.Float32:
		fmt.Println("a 是： ", v.Float())
	case reflect.String:
		fmt.Println("a 是：", v.String())
	case reflect.Map:
		iter := v.MapRange()
		fmt.Println("a 是：", iter)
		for iter.Next() {
			k := iter.Key()
			v := iter.Value()
			fmt.Printf("key:%s,val:%s\n", k, v)
		}
	case reflect.Slice:
		fmt.Println("a 是：", v.Slice(0, v.Len()))
	}
}

//反射修改值
/**
*可修改条件，1.可寻址，2。可寻址的类型(1:指针指向的具体元素,2:slice的元素,3:可寻址的结构体的字段(指向结构体的指针) 4:可寻址的数组的元素(指向数组的指针))
 */
func reflect_set_value(a interface{}) {
	v := reflect.ValueOf(a)
	fmt.Println(777)
	fmt.Println(v)
	k := v.Kind()
	switch k {
	case reflect.Float64:
		// 反射修改值
		v.SetFloat(6.9)
		fmt.Println("a is ", v.Float())
	case reflect.String:

	case reflect.Slice:
		fmt.Println("can addr:", v.CanAddr())
		e := v.Index(0)
		e.SetInt(777)
		fmt.Println("after set:", v)
	case reflect.Struct:
		fmt.Println("can addr:", v.CanAddr())
	case reflect.Ptr:
		// Elem()获取地址指向的值
		// v.Elem().SetFloat(7.9)
		nk := v.Elem().Kind() //获取数据类型
		switch nk {
		case reflect.String:
			v.Elem().SetString("hahahhhaa")
			fmt.Println("case:", v.Elem().String())
		case reflect.Float64:
			ele := v.Elem()
			ele.SetFloat(6.9)
			fmt.Println("a is ", ele.Float())
		case reflect.Array:
			ele := v.Elem()
			e := ele.Index(0)
			fmt.Println("e type:", e)
			fmt.Println(" e can set:", e.CanSet())
			e.SetInt(7777)
			fmt.Println("after set :", v)
		case reflect.Struct:
			fmt.Println("can addr:", v.CanSet())
			val := v.Elem()
			age := val.FieldByName("Age")
			fmt.Println("age can set:", age.CanSet())
			age.SetInt(22333)

		}
		//地址
		fmt.Println(v.Pointer()) //It panics if v's Kind is not Chan, Func, Map, Ptr, Slice, or UnsafePointer.
	}
}

// 定义结构体
type User struct {
	Id   int
	Name string
	Age  int
}

// 匿名字段
type Boy struct {
	User
	Addr string
}

// 绑方法
// func (u User) Hello() {
// 	fmt.Println("Hello")
// }

func (u User) Hello(name string) {
	fmt.Println("Hello：", name)
}

// 传入interface{}
func Poni(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("类型：", t)
	fmt.Println("字符串类型：", t.Name())
	// 获取值
	v := reflect.ValueOf(o)
	fmt.Println(v)
	// 可以获取所有属性
	// 获取结构体字段个数：t.NumField()//字段个数
	for i := 0; i < t.NumField(); i++ {
		// 取每个字段
		f := t.Field(i)
		fmt.Printf("%s : %v\n", f.Name, f.Type)
		// 获取字段的值信息
		// Interface()：获取字段对应的值
		val := v.Field(i).Interface()
		fmt.Println("val :", val)
	}
	fmt.Println("=================方法====================")
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name)
		fmt.Println(m.Type)
	}

}

// 修改结构体值
func SetValue(o interface{}) {
	v := reflect.ValueOf(o)
	// 获取指针指向的元素
	v = v.Elem()
	// 取字段
	f := v.FieldByName("Name")
	if f.Kind() == reflect.String {
		f.SetString("kuteng")
	}
}

type Student struct {
	Name string `json:"name1" db:"name2"`
}

func main() {
	// x := make(map[string]interface{})
	// // x = map[string]interface{}{"name":"ykk", "age":20}
	// x["name"] = "ykk"
	// x["age"] = 20
	// x := make([]int, 3)
	// x[0] = 1
	// x[1] = 2
	// fmt.Println(len(x))
	// reflect_type(x)

	// reflect_value(x)

	// 反射认为下面是指针类型，不是float类型
	// x := []int{1, 2, 4}

	// x := User{Id: 1, Name: "ykk", Age: 20}
	// x := [2]int{1, 3}
	// var x float64 = 22.22
	// reflect_set_value(&x)
	// fmt.Println("main:", x)

	// u := User{1, "zs", 20}
	// // u := "sfaff"
	// Poni(u)

	// m := Boy{User{1, "zs", 20}, "bj"}
	m := Student{Name: "ykk"}
	t := reflect.TypeOf(m)
	fmt.Printf("%#v", t)
	// Anonymous：匿名
	fmt.Printf("%#v\n", t.Field(0))

	//等价操作：
	/*
		t := reflect.TypeOf(m)

		等价于：reflect.ValueOf(m).Type()

	*/
	// 值信息
	fmt.Printf("%#v\n", reflect.ValueOf(m).Type())

	// u := User{1, "5lmh.com", 20}
	// SetValue(&u)
	// fmt.Println(u)

	//通过反射方式调用方法
	// u := User{1, "5lmh.com", 20}
	// v := reflect.ValueOf(u)
	// // 获取方法
	// m := v.MethodByName("Hello")
	// // 构建一些参数
	// args := []reflect.Value{reflect.ValueOf("6666")}
	// // 没参数的情况下：var args2 []reflect.Value
	// // 调用方法，需要传入方法的参数
	// m.Call(args)

	// var s Student
	// v := reflect.ValueOf(&s)
	// // 类型
	// t := v.Type()
	// fmt.Println("type:", t)
	// // 获取字段
	// f := t.Elem().Field(0)
	// fmt.Println(f.Tag.Get("json"))
	// fmt.Println(f.Tag.Get("db"))
}
