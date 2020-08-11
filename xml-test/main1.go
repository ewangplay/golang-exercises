package main

import (
    "fmt"
    "encoding/xml"
)

type Result struct {
    SPNumber string     `xml:"spnumber"`
    Phone string        `xml:"phone"`
    Status string       `xml:"status"`
    SendTime string     `xml:"sendtime"`

}
type SMSResult struct {
    XMLName xml.Name     `xml:"smsResult"`
    Results []Result     `xml:"result"`
}

func main() {
    v := SMSResult{}

    data := `
        <?xml version="1.0" encoding="UTF-8"?>
        <smsResult>
            <result>
                <spnumber>短信流水号</spnumber>
                <phone>手机号码</phone>
                <status>DELIVRD</status>
                <sendtime>2008-12-24 23:58:24</sendtime>
            </result>
            <result>
                <spnumber>0</spnumber>
                <phone>15313235171</phone>
                <status>UNDELIV</status>
                <sendtime>20121106112751</sendtime>
            </result>
        </smsResult>
    `
    err := xml.Unmarshal([]byte(data), &v)
    if err != nil {
        fmt.Printf("error: %v", err)
        return
    }

    fmt.Printf("XMLName: %#v\n", v.XMLName)
    for _, v := range v.Results {
        fmt.Println(v.SPNumber)
        fmt.Println(v.Phone)
        fmt.Println(v.Status)
        fmt.Println(v.SendTime)
        fmt.Println()
    }
}
