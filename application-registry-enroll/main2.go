package main

import (
	"fmt"
	"log"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func main() {
	// configPath := "../test-network/organizations/peerOrganizations/org1.example.com/connection-org1.json"
	configPath := "connection-org1.json"

	// Initialize the SDK
	sdk, err := fabsdk.New(config.FromFile(configPath))
	if err != nil {
		log.Fatalf("Failed to create new SDK: %s", err)
	}
	defer sdk.Close()

	// Create an MSP client
	// mspClient, err := msp.New(sdk.Context(), msp.WithOrg("Org1"))
	// if err != nil {
	// 	log.Fatalf("Failed to create MSP client: %s", err)
	// }

	// Create an MSP client for Org1 using the admin identity
	mspClient, err := msp.New(sdk.Context(fabsdk.WithUser("admin"), fabsdk.WithOrg("Org1")))
	if err != nil {
		log.Fatalf("Failed to create MSP client: %s", err)
	}

	// Register the user
	registrationRequest := &msp.RegistrationRequest{
		Name:        "User5",
		Type:        "client",
		Affiliation: "org1.department1",
		Attributes:  []msp.Attribute{{Name: "abac.creator", Value: "true"}},
	}

	enrollmentSecret, err := mspClient.Register(registrationRequest)
	if err != nil {
		log.Fatalf("Failed to register user: %v", err)
	}

	// Enroll the user
	err = mspClient.Enroll("User5", msp.WithSecret(enrollmentSecret))
	if err != nil {
		log.Fatalf("Failed to enroll user: %v", err)
	}

	fmt.Println("User enrolled successfully")
}
