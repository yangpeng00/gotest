package main

import "fmt"

/*
map的增删改查， 循环
*/
func main() {
	var countryMap map[string]string
	countryMap = make(map[string]string)
	countryMap["fasd"] = "马裤"
	countryMap["sfad"] = "十分动感"
	countryMap["hgljk"] = "官服"
	countryMap["tyui"] = "额外"
	for country := range countryMap {
		fmt.Println(country, "首都是", countryMap[country])
	}
	cap, ok := countryMap["美国"]
	if ok {
		fmt.Println("美国的首都是", cap)
	} else {
		fmt.Println("美国的首都不存在")
	}
}
