package main

import "fmt"

func main() {
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	//测试删除功能
	//delete(countryCapitalMap, "France")
	for country := range countryCapitalMap {
		fmt.Println("Capital of:", country, "is" ,countryCapitalMap[country])
	}


	//查看元素在集合中是否存在
	captial, ok := countryCapitalMap["United States"]
	//如果ok是true则存在，否则不存在
	if(ok){
		fmt.Println("Capital of United States is", captial)
	}else {
		fmt.Println("Capital of United States is not present")
	}
}
