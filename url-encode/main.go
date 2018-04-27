package main

import (
	"fmt"
	"net/url"
)

func main() {
	//sweet := "香蕉朱古力蛋糕.mp4"
	sweet := "【5分鐘】棉花糖朱古力蛋糕食譜 Marshmallow Chocolate Mug Cake recipe  ＊Happy Amy-hMCIfRG2DMg.mp4"
	fmt.Println(sweet)
	encode_sweet := url.PathEscape(sweet)
	fmt.Printf("%s\n", encode_sweet)
	decode_sweet, _ := url.PathUnescape("%E3%80%905%E5%88%86%E9%90%98%E3%80%91%E6%A3%89%E8%8A%B1%E7%B3%96%E6%9C%B1%E5%8F%A4%E5%8A%9B%E8%9B%8B%E7%B3%95%E9%A3%9F%E8%AD%9C+Marshmallow+Chocolate+Mug+Cake+recipe++%EF%BC%8AHappy+Amy-hMCIfRG2DMg.mp4")
	fmt.Printf("%s\n", decode_sweet)
	parse_sweet, _ := url.Parse("%E3%80%905%E5%88%86%E9%90%98%E3%80%91%E6%A3%89%E8%8A%B1%E7%B3%96%E6%9C%B1%E5%8F%A4%E5%8A%9B%E8%9B%8B%E7%B3%95%E9%A3%9F%E8%AD%9C+Marshmallow+Chocolate+Mug+Cake+recipe++%EF%BC%8AHappy+Amy-hMCIfRG2DMg.mp4")
	fmt.Printf("%s\n", parse_sweet.Path)
	query_sweet, _ := url.QueryUnescape("%E3%80%905%E5%88%86%E9%90%98%E3%80%91%E6%A3%89%E8%8A%B1%E7%B3%96%E6%9C%B1%E5%8F%A4%E5%8A%9B%E8%9B%8B%E7%B3%95%E9%A3%9F%E8%AD%9C+Marshmallow+Chocolate+Mug+Cake+recipe++%EF%BC%8AHappy+Amy-hMCIfRG2DMg.mp4")
	fmt.Printf("%s\n", query_sweet)
}
