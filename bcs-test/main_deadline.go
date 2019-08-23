package main

import (
	"context"
	"fmt"
	"time"

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
	ctx, _ := context.WithTimeout(context.Background(), time.Minute*2)
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			fmt.Println("batch test exit...")
			return nil
		default:
			assetDid := fmt.Sprintf("did:axn:%v", commutils.GenerateUUID())
			assetName := fmt.Sprintf("AssetName-%v", commutils.GenerateShortID())
			assetHash := commutils.ComputeSha256(fmt.Sprintf("%v%v", assetDid, assetName))
			poeReq := &bcs.CreatePoeRequest{
				Did:      assetDid,
				Name:     assetName,
				Metadata: assetHash,
				BaasMode: "bassMode",
			}
			resp, err := client.CreatePoe(ctx, poeReq)
			if err != nil {
				return err
			}
			fmt.Println(resp)
		}
		//time.Sleep(1 * time.Second)
	}

	return nil
}
