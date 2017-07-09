package main

import (
    "fmt"
    "strings"
)

func main() {
	var str string = "This is an example of a string"
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "Th" )
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
	fmt.Printf("%t\n", strings.HasSuffix(str, "string"))
	fmt.Printf("%t\n", strings.Contains(str, "s1tring"))
	fmt.Printf("%t\n", strings.Index(str, "a"))
	fmt.Printf("%t\n", strings.LastIndex(str, "a"))
	fmt.Println(strings.Replace(str, "a", "zz", -1))
	fmt.Println(strings.Count(str, "a"))
	fmt.Println(strings.Repeat("aa", 10))
	fmt.Println(strings.ToLower("TTTTTTTTTTsdasdDadVVVVVVV"))
	fmt.Println(strings.ToUpper("shdiahishdiashihYYYYYYYYYYYshdiahida"))
	fmt.Println(strings.TrimSpace(" sjidahishdiashdiasdiaid hsidhiashidhi hishdiahi   "))
	fmt.Println(strings.Trim("ssjidahishdiashdiasdiaid hsidhiashidhi hishdiahi   sssss", "s"))
	fmt.Println(strings.TrimLeft("shaihsidhaihdhsidhaish", "sh"))
	fmt.Println(strings.TrimRight("shaihsidhaihdhsidhaish", "sh"))
	fmt.Println(strings.Fields("shaihdias sahidhia ashidhia shidahi   hshdiahsidh   shihi"))
	str3 := strings.Split("shadhsiadshiahdiahsiashidhaishdiahsidhaishdas", "sh")
    fmt.Println(str3)
    str4 := strings.Join(str3, "OO")
	fmt.Println(strings.NewReader(str4))
	
}
