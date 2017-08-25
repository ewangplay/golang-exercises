package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	//pattern := ".*bi\\[([\\w\\s=@\\.:'$\\/\\[\\]\\-]+)\\].*"
	//url := "bi[c=1037 t=8 y=3 g=1 u=1 d=1 w=18310028269 s=100]"
	//url := "2015-08-07 21:05:03 maillogparser TESTING INFO logParser.py 143 - bi[y=2 s=200 w=wangshuye@chunbo.com ip=101.251.211.148 id=09D07A0FC time='2015-08-07 21:04:03'] relay=smtp01.chunbo.com dsn=2.6.0 status=sent desc=(250 2.6.0 <55031533-d652-4956-8a52-ed77a392b841@PEK1-MBX-01.chunbo.local> [InternalId=33328946217017, Hostname=PEK1-MBX-01.chunbo.local] Queued mail for delivery)"

	pattern := ".*bi\\[(.*)\\].*"
    url := "2016-01-31 15:56:51.480 ▶  collectioninfo TESTING INFO redirect_click.go 91 -  bi[c=1237 t=19 y=2 g=12 u=29 d=13 w=1019423882@qq.com s=401 ourl=http://app.360email.cn/admin/template.php?act=templatepublic_edit&amp;id=96#]helloworld"

	regexp, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Printf("comple regular expression fail. %v", err)
		return
	}

	matchs := regexp.FindStringSubmatch(url)
	if matchs == nil {
		fmt.Printf("not any match.\n")
		return
	}

	for _, str := range matchs {
		fmt.Printf("%v\n", str)
	}

	result := make(map[string]string, 0)
	entries := strings.Split(matchs[1], " ")
	for _, entry := range entries {
		fmt.Println(entry)
		items := strings.Split(entry, "=")

		if len(items) == 2 {
			result[items[0]] = items[1]
		}
	}

	fmt.Println(result)
}
