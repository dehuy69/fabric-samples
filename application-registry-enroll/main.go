package main

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func main() {
	// Initialize the Fabric SDK
	configPath := "../test-network/organizations/peerOrganizations/org1.example.com/connection-org1.json"
	sdk, err := fabsdk.New(config.FromFile(configPath))
	if err != nil {
		fmt.Printf("Failed to create new SDK: %s\n", err)
		return
	}
	defer sdk.Close()

	// Create a new MSP client
	mspClient, err := msp.New(sdk.Context(), msp.WithOrg("Org1"))
	if err != nil {
		fmt.Printf("Failed to create MSP client: %s\n", err)
		return
	}

	// Register the user
	enrollmentSecret, err := mspClient.Register(&msp.RegistrationRequest{
		Name:           "User5",
		Type:           "client",
		Affiliation:    "org1.department1",
		MaxEnrollments: -1, // No limit on re-enrollments
	})
	if err != nil {
		fmt.Printf("Failed to register user: %s\n", err)
		return
	}

	// Enroll the user
	err = mspClient.Enroll("User5", msp.WithSecret(enrollmentSecret))
	if err != nil {
		fmt.Printf("Failed to enroll user: %s\n", err)
		return
	}
	if err != nil {
		fmt.Printf("Failed to enroll user: %s\n", err)
		return
	}

	fmt.Println("Successfully registered and enrolled user.")
	// You can now use enrollmentResponse.Key and enrollmentResponse.Cert
}
