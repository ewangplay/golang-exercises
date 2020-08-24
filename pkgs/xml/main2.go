package main

import (
	"encoding/xml"
	"fmt"
)

type LocalWeixinPrivateMsg struct {
	CustomerId   int
	AccountId    int
	ToUserName   string  `xml:"ToUserName"`
	FromUserName string  `xml:"FromUserName"`
	CreateTime   int64   `xml:"CreateTime"`
	MsgType      string  `xml:"MsgType"`
	Content      string  `xml:"Content"`
	MsgId        string  `xml:"MsgId"`
	PicUrl       string  `xml:"PicUrl"`
	MediaId      string  `xml:"MediaId"`
	Format       string  `xml:"Format"`
	ThumbMediaId string  `xml:"ThumbMediaId"`
	Location_X   float64 `xml:"Location_X"`
	Location_Y   float64 `xml:"Location_Y"`
	Scale        int     `xml:"Scale"`
	Label        string  `xml:"Label"`
	Title        string  `xml:"Title"`
	Description  string  `xml:"Description"`
	Url          string  `xml:"Url"`
	Event        string  `xml:"Event"`
	EventKey     string  `xml:"EventKey"`
	Ticket       string  `xml:"Ticket"`
	Latitude     float64 `xml:"Latitude"`
	Longitude    float64 `xml:"Longitude"`
	Precision    float64 `xml:"Precision"`
	Recognition  string  `xml:"Recognition"`
}

func main() {
	v_text := LocalWeixinPrivateMsg{
		CustomerId: 10,
		AccountId:  20,
	}
	v_pic := LocalWeixinPrivateMsg{}
	v_voice := LocalWeixinPrivateMsg{}
	v_vedio := LocalWeixinPrivateMsg{}
	v_location := LocalWeixinPrivateMsg{}
	v_link := LocalWeixinPrivateMsg{}
	v_sub := LocalWeixinPrivateMsg{}
	v_tick := LocalWeixinPrivateMsg{}
	v_postloc := LocalWeixinPrivateMsg{}
	v_menuclick := LocalWeixinPrivateMsg{}
	v_menuview := LocalWeixinPrivateMsg{}

	data_text := `
    <xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[fromUser]]></FromUserName> 
    <CreateTime>1348831860</CreateTime>
    <MsgType><![CDATA[text]]></MsgType>
    <Content><![CDATA[this is a test]]></Content>
    <MsgId>1234567890123456</MsgId>
    </xml>
    `

	data_pic := `
    <xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[fromUser]]></FromUserName>
    <CreateTime>1348831860</CreateTime>
    <MsgType><![CDATA[image]]></MsgType>
    <PicUrl><![CDATA[this is a url]]></PicUrl>
    <MediaId><![CDATA[media_id]]></MediaId>
    <MsgId>1234567890123456</MsgId>
    </xml>
    `

	data_voice := `
    <xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[fromUser]]></FromUserName>
    <CreateTime>1357290913</CreateTime>
    <MsgType><![CDATA[voice]]></MsgType>
    <MediaId><![CDATA[media_id]]></MediaId>
    <Format><![CDATA[Format]]></Format>
    <MsgId>1234567890123456</MsgId>
    </xml>
    `

	data_vedio := `
    <xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[fromUser]]></FromUserName>
    <CreateTime>1357290913</CreateTime>
    <MsgType><![CDATA[video]]></MsgType>
    <MediaId><![CDATA[media_id]]></MediaId>
    <ThumbMediaId><![CDATA[thumb_media_id]]></ThumbMediaId>
    <MsgId>1234567890123456</MsgId>
    </xml>
    `

	data_location := `
    <xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[fromUser]]></FromUserName>
    <CreateTime>1351776360</CreateTime>
    <MsgType><![CDATA[location]]></MsgType>
    <Location_X>23.134521</Location_X>
    <Location_Y>113.358803</Location_Y>
    <Scale>20</Scale>
    <Label><![CDATA[位置信息]]></Label>
    <MsgId>1234567890123456</MsgId>
    </xml> 
    `

	data_link := `
    <xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[fromUser]]></FromUserName>
    <CreateTime>1351776360</CreateTime>
    <MsgType><![CDATA[link]]></MsgType>
    <Title><![CDATA[公众平台官网链接]]></Title>
    <Description><![CDATA[公众平台官网链接]]></Description>
    <Url><![CDATA[url]]></Url>
    <MsgId>1234567890123456</MsgId>
    </xml> 
    `

	data_sub := `
    <xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[FromUser]]></FromUserName>
    <CreateTime>123456789</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[subscribe]]></Event>
    </xml>
    `

	data_tick := `
    <xml><ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[FromUser]]></FromUserName>
    <CreateTime>123456789</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[subscribe]]></Event>
    <EventKey><![CDATA[qrscene_123123]]></EventKey>
    <Ticket><![CDATA[TICKET]]></Ticket>
    </xml>
    `

	data_postloc := `
    <xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[fromUser]]></FromUserName>
    <CreateTime>123456789</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[LOCATION]]></Event>
    <Latitude>23.137466</Latitude>
    <Longitude>113.352425</Longitude>
    <Precision>119.385040</Precision>
    </xml>
    `

	data_menuclick := `
    <xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[FromUser]]></FromUserName>
    <CreateTime>123456789</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[CLICK]]></Event>
    <EventKey><![CDATA[EVENTKEY]]></EventKey>
    </xml>
    `

	data_menuview := `
    <xml>
    <ToUserName><![CDATA[toUser]]></ToUserName>
    <FromUserName><![CDATA[FromUser]]></FromUserName>
    <CreateTime>123456789</CreateTime>
    <MsgType><![CDATA[event]]></MsgType>
    <Event><![CDATA[VIEW]]></Event>
    <EventKey><![CDATA[www.qq.com]]></EventKey>
    </xml>
    `

	var err error
	err = xml.Unmarshal([]byte(data_text), &v_text)
	if err != nil {
		fmt.Printf("text error: %v", err)
		return
	}

	err = xml.Unmarshal([]byte(data_pic), &v_pic)
	if err != nil {
		fmt.Printf("pic error: %v", err)
		return
	}

	err = xml.Unmarshal([]byte(data_voice), &v_voice)
	if err != nil {
		fmt.Printf("voice error: %v", err)
		return
	}

	err = xml.Unmarshal([]byte(data_vedio), &v_vedio)
	if err != nil {
		fmt.Printf("vedio error: %v", err)
		return
	}

	err = xml.Unmarshal([]byte(data_location), &v_location)
	if err != nil {
		fmt.Printf("location error: %v", err)
		return
	}

	err = xml.Unmarshal([]byte(data_link), &v_link)
	if err != nil {
		fmt.Printf("link error: %v", err)
		return
	}

	err = xml.Unmarshal([]byte(data_sub), &v_sub)
	if err != nil {
		fmt.Printf("sub error: %v", err)
		return
	}

	err = xml.Unmarshal([]byte(data_tick), &v_tick)
	if err != nil {
		fmt.Printf("tick error: %v", err)
		return
	}

	err = xml.Unmarshal([]byte(data_postloc), &v_postloc)
	if err != nil {
		fmt.Printf("postloc error: %v", err)
		return
	}

	err = xml.Unmarshal([]byte(data_menuclick), &v_menuclick)
	if err != nil {
		fmt.Printf("menuclick error: %v", err)
		return
	}

	err = xml.Unmarshal([]byte(data_menuview), &v_menuview)
	if err != nil {
		fmt.Printf("menuview error: %v", err)
		return
	}

	fmt.Println("text: ", v_text)
	fmt.Println("pic: ", v_pic)
	fmt.Println("voice: ", v_voice)
	fmt.Println("vedio: ", v_vedio)
	fmt.Println("location: ", v_location)
	fmt.Println("link: ", v_link)
	fmt.Println("sub: ", v_sub)
	fmt.Println("tick: ", v_tick)
	fmt.Println("postloc: ", v_postloc)
	fmt.Println("menuclick: ", v_menuclick)
	fmt.Println("menuview: ", v_menuview)
}
