package main

import (
	"context"
	"fmt"

	bcs "github.com/arxanchain/bc-service/protos"
	commutils "github.com/arxanchain/sdk-go-common/utils"
	"google.golang.org/grpc"
)

func main() {
	// Init the bcs client
	bcsHost := "192.168.20.59:8030"
	client, err := InitBCSClient(bcsHost)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Batch test
	err = BatchTest(client)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func InitBCSClient(host string) (bcs.BCClient, error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := bcs.NewBCClient(conn)
	return client, nil
}

func BatchTest(client bcs.BCClient) error {
	for i := 0; i < 1000; i++ {
		assetDid := fmt.Sprintf("did:axn:%v", commutils.GenerateUUID())
		assetName := fmt.Sprintf("AssetName-%v", commutils.GenerateShortID())
		assetHash := commutils.ComputeSha256(fmt.Sprintf("%v%v", assetDid, assetName))
		poeReq := &bcs.CreatePoeRequest{
			Did:      assetDid,
			Name:     assetName,
			Metadata: assetHash,
			BaasMode: "bassMode",
		}
		resp, err := client.CreatePoe(context.Background(), poeReq)
		if err != nil {
			return err
		}
		fmt.Println(resp)
	}
	return nil
}
