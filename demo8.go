// Ghttps://www.runoob.com/go/go-map.html -- o 语言Map(集合)

package main

import "fmt"

func mapFunc() {
	// 创建集合
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string) // 如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"

	/*使用键输出地图值 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	/*查看元素在集合中是否存在 */
	capital, ok := countryCapitalMap["American"] /*如果确定是真实的, 则存在, 否则不存在 */
	/*fmt.Println(capital) */
	fmt.Println(ok)
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}
}

// France 首都是 巴黎
// Italy 首都是 罗马
// Japan 首都是 东京
// India  首都是 新德里
// false
// American 的首都不存在

// delete() 函数
func deleteMap() {
	/* 创建map */
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

	fmt.Println("原始地图")

	/* 打印地图 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	/*删除元素*/
	delete(countryCapitalMap, "France")
	fmt.Println("法国条目被删除")

	fmt.Println("删除元素后地图")

	/*打印地图*/
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
}

// 原始地图
// India 首都是 New delhi
// France 首都是 Paris
// Italy 首都是 Rome
// Japan 首都是 Tokyo
// 法国条目被删除
// 删除元素后地图
// Italy 首都是 Rome
// Japan 首都是 Tokyo
// India 首都是 New delhi
