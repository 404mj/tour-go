package main

import (
	"encoding/xml"
	"fmt"
)

//将微信消息的xml格式转化为结构体,原文章问题比较多！
//http://blog.studygolang.com/2012/12/标准库-xml处理（一）/

type WXNewsMsg struct {
	ToName       string `xml:"ToName"`
	FromName     string `xml:"FromName"`
	CreateTime   uint64 `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	ArticleCount uint32 `xml:"ArticleCount"`
	//***Note****:HERE!
	Articles []WXArticleItem `xml:"Articles>Item"`
}
type WXArticleItem struct {
	Title       string `xml:"Title"`
	Description string `xml:"Description"`
	PicUrl      string `xml:"PicUrl"`
	Url         string `xml:"Url"`
}

//将xml转化为struct
func decodeDemo() {
	xmlstr := `
	<?xml version="1.0" encoding="UTF-8"?>
	<WXNewsMsg>
	<ToName>psq</ToName>
	<FromName>zsx</FromName>
	<CreateTime>20180127</CreateTime>
	<MsgType>say hi~</MsgType>
	<ArticleCount>2</ArticleCount>
	<Articles>
		<Item>
			<Title>标题111</Title>
			<Description>内容111</Description>
			<PicUrl>图片URL111</PicUrl>
			<Url>图文跳转链接111</Url>
		</Item>
		<Item>
			<Title>标题222</Title>
			<Description>内容222</Description>
			<PicUrl>图片URL222</PicUrl>
			<Url>图文跳转链接222</Url>
		</Item>
	</Articles>
	</WXNewsMsg>
	`

	v := WXNewsMsg{}
	err := xml.Unmarshal([]byte(xmlstr), &v)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(v)
	fmt.Println("ToUserName:", v.ToName)
	fmt.Println("MsgType:", v.MsgType)
	fmt.Println(v.Articles)
}

//将struct转化为xml串
func encodeDemo() {

}

func main() {
	decodeDemo()
}
