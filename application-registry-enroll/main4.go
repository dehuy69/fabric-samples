package main

import (
	"fmt"
	"log"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func main() {
	// Load the network configuration
	configPath := "network-config-remote-host.yaml"
	// configPath := "network-config-main4.yaml" // Thay đổi đường dẫn đến file config.yaml
	sdk, err := fabsdk.New(config.FromFile(configPath))
	if err != nil {
		log.Fatalf("Failed to create SDK: %v", err)
	}
	defer sdk.Close()

	// Create a new MSP client instance
	mspClient, err := msp.New(sdk.Context(), msp.WithOrg("Org1")) // Sử dụng Org1 hoặc tổ chức của bạn
	if err != nil {
		log.Fatalf("Failed to create MSP client: %v", err)
	}

	// Register a new identity
	registerRequest := &msp.RegistrationRequest{
		Name:        "newUser1",
		Type:        "client",
		Affiliation: "org1.department1",
		Secret:      "newUserpw",
		Attributes:  []msp.Attribute{{Name: "abac.creator", Value: "true"}},
	}

	secret, err := mspClient.Register(registerRequest)
	if err != nil {
		log.Fatalf("Failed to register identity: %v", err)
	}
	fmt.Printf("Identity registered successfully with secret: %s\n", secret)

	// Enroll the new identity
	err = mspClient.Enroll("newUser1", msp.WithSecret(secret))
	if err != nil {
		log.Fatalf("Failed to enroll identity: %v", err)
	}
	fmt.Println("Identity enrolled successfully!")
}
