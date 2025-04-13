package main

import (
	"fmt"
	"math/big"
	"os"

	"github.com/Phathdt/abierr/example/client"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
		os.Exit(1)
	}

	rpcURL := os.Getenv("RPC_URL")
	privateKey := os.Getenv("PRIVATE_KEY")
	contractAddress := os.Getenv("TOKEN_ADDRESS")

	if rpcURL == "" || privateKey == "" || contractAddress == "" {
		fmt.Println("Missing required environment variables: RPC_URL, PRIVATE_KEY, or TOKEN_ADDRESS")
		os.Exit(1)
	}

	erc20Client, err := client.NewERC20Client(rpcURL, privateKey, contractAddress)
	if err != nil {
		fmt.Printf("Failed to create ERC20 client: %v\n", err)
		os.Exit(1)
	}

	toAddress := os.Getenv("TO_ADDRESS")
	if toAddress == "" {
		fmt.Println("Missing required environment variable: TO_ADDRESS")
		os.Exit(1)
	}

	amount := new(big.Int).Mul(big.NewInt(1000), big.NewInt(1e18))

	txHash, err := erc20Client.Mint(toAddress, amount)
	if err != nil {
		errorMsg := err.Error()
		fmt.Println("errorMsg: ", errorMsg)
		os.Exit(1)
	}

	fmt.Printf("Successfully minted tokens. Transaction hash: %s\n", txHash.Hex())
}
