package main

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func main() {
	firstExample()

}

func myTestCode1() {
	mspClient, err := connect()
	if err != nil {
		fmt.Printf("Failed to connect: %s\n", err)
		return
	}

	err = registerAndEnroll(mspClient, "User11", "org1.department1", "client")
	if err != nil {
		fmt.Printf("Failed to register and enroll user: %s\n", err)
		return
	}

	fmt.Println("Successfully registered and enrolled user")
}

func myTestCode2() {
	username := "User11"

	mspClient, err := connect()
	if err != nil {
		fmt.Printf("Failed to connect: %s\n", err)
		return
	}

	enrolledUser, err := mspClient.GetSigningIdentity(username)
	if err != nil {
		fmt.Printf("user not found %s\n", err)
		return
	}

	fmt.Println("Private Key: ", enrolledUser.PrivateKey())
}

func firstExample() {
	// Path to your network configuration file
	configPath := "./network-config.yaml"

	// Create a new Fabric SDK instance
	sdk, err := fabsdk.New(config.FromFile(configPath))
	if err != nil {
		fmt.Printf("Failed to create new SDK: %s\n", err)
		return
	}
	defer sdk.Close()

	// Create an MSP client
	mspClient, err := msp.New(sdk.Context())
	if err != nil {
		fmt.Printf("Failed to create MSP client: %s\n", err)
		return
	}

	// // Register a new user (User5)
	registrationRequest := &msp.RegistrationRequest{
		Name:        "User10",           // Name of the new user
		Type:        "client",           // Type of the identity
		Affiliation: "org1.department1", // Affiliation
		Attributes: []msp.Attribute{
			{Name: "role", Value: "client", ECert: true},
		},
	}

	secret, err := mspClient.Register(registrationRequest)
	if err != nil {
		fmt.Printf("Failed to register user: %s\n", err)
		return
	}

	fmt.Println("Move all file folders in credentials folder to the sub folder name admin_artifacts")

	// Enroll the user (User5) with the secret received from registration
	err = mspClient.Enroll("User10", msp.WithSecret(secret))
	if err != nil {
		fmt.Printf("Failed to enroll user: %s\n", err)
		return
	}
	fmt.Println("Successfully enrolled user")
}

func connect() (*msp.Client, error) {
	// Path to your network configuration file
	configPath := "./network-config.yaml"

	// Create a new Fabric SDK instance
	sdk, err := fabsdk.New(config.FromFile(configPath))
	if err != nil {
		fmt.Printf("Failed to create new SDK: %s\n", err)
		return nil, err
	}
	defer sdk.Close()

	// Create an MSP client
	mspClient, err := msp.New(sdk.Context())
	if err != nil {
		fmt.Printf("Failed to create MSP client: %s\n", err)
		return nil, err
	}

	return mspClient, nil
}

func register(mspClient *msp.Client, name string, affiliation string, role string) (string, error) {
	registrationRequest := &msp.RegistrationRequest{
		Name:        name,
		Type:        "client",
		Affiliation: affiliation,
		Attributes: []msp.Attribute{
			{Name: "role", Value: role, ECert: true},
		},
	}

	secret, err := mspClient.Register(registrationRequest)
	if err != nil {
		return "", fmt.Errorf("failed to register user: %w", err)
	}

	return secret, nil
}

func enroll(mspClient *msp.Client, name string, secret string) error {
	err := mspClient.Enroll(name, msp.WithSecret(secret))
	if err != nil {
		return fmt.Errorf("failed to enroll user: %w", err)
	}
	return nil
}

func registerAndEnroll(mspClient *msp.Client, name string, affiliation string, role string) error {
	secret, err := register(mspClient, name, affiliation, role)
	if err != nil {
		return err
	}

	err = enroll(mspClient, name, secret)
	if err != nil {
		return err
	}

	return nil
}
