package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	epssdk "github.com/gfacloud/gfa-eps-sdk-go"
	uuid "github.com/satori/go.uuid"
)

func main() {
	// 新建客户端
	opts := &epssdk.Options{
		AppID:     "632ed384-0b77-47af-835e-fb2b9b300fbc",
		AppKey:    "d26da1e334ace0f00d7a2f925d3268263ab4b9fe3107d33be8bc460bfb07da67",
		AppSecret: "05df6c85f1bda0131f36e2b28ad08ab47e76de32b8dfb6f133638008d6a88153d26da1e334ace0f00d7a2f925d3268263ab4b9fe3107d33be8bc460bfb07da67",
		Addr:      "127.0.0.1:8888",
	}
	c, err := epssdk.NewClient(opts)
	if err != nil {
		fmt.Printf("NewClient failed: %v\n", err)
		os.Exit(1)
	}

	// 读取测试文件
	contentBase64, fileType, err := readTestFile("../files/test1.pdf")
	if err != nil {
		fmt.Printf("Read test file failed: %v\n", err)
		os.Exit(1)
	}

	// 构建证据数据并提交
	e := &epssdk.Evidence{
		CollectID: uuid.NewV4().String(),
		Name:      "张三的劳动合同",
		Materials: []epssdk.Material{
			epssdk.Material{
				ID:            uuid.NewV4().String(),
				Type:          fileType,
				ContentBase64: contentBase64,
			},
			epssdk.Material{
				ID:            uuid.NewV4().String(),
				Type:          fileType,
				ContentBase64: contentBase64,
			},
		},
	}
	err = c.CreateEvidence(e)
	if err != nil {
		fmt.Printf("CreateEvidence failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("证据ID: ", e.ID)
}

func readTestFile(filename string) (contentBase64 string, fileType string, err error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	contentBase64 = base64.StdEncoding.EncodeToString(content)

	items := strings.Split(filename, ".")
	if len(items) >= 2 {
		fileType = items[len(items)-1]
	} else {
		fileType = "normal"
	}

	return
}
