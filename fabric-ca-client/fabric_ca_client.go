package main

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

// Constants definition
const (
	ConfigPath = "./config.yaml"
)

func main() {
	sdk, err := fabsdk.New(config.FromFile(ConfigPath))
	if err != nil {
		fmt.Printf("Failed to create fabsdk: %v\n", err)
		return
	}
	defer sdk.Close()

	// Create msp client
	ctx := sdk.Context()
	c, err := msp.New(ctx)
	if err != nil {
		fmt.Printf("Failed to create msp client: %v\n", err)
		return
	}

	// Enroll admin user
	err = c.Enroll("admin", msp.WithSecret("adminpw"))
	if err != nil {
		fmt.Printf("Failed to enroll admin: %v\n", err)
		return
	}

	username := "user5"
	enrollmentSecret, err := c.Register(&msp.RegistrationRequest{
		Name: username,
		Type: "client",
	})
	if err != nil {
		fmt.Printf("Failed to register user[%s]: %v\n", username, err)
		return
	}

	err = c.Enroll(username, msp.WithSecret(enrollmentSecret))
	if err != nil {
		fmt.Printf("Failed to enroll user[%s]: %v\n", username, err)
		return
	}

	fmt.Printf("enroll user succ. %s - %s\n", username, enrollmentSecret)

	enrolledUser, err := c.GetSigningIdentity(username)
	if err != nil {
		fmt.Printf("user not found %s\n", err)
		return
	}

	if enrolledUser.Identifier().ID != username {
		fmt.Println("Enrolled user name doesn't match")
		return
	}
}
