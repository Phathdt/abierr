package main

import (
	"fmt"
	"os"

	"github.com/Phathdt/abierr"
	"github.com/Phathdt/abierr/example/optimex/contracts"
)

func main() {
	// Create a new decoder with the EvmBtc ABI
	decoder, err := abierr.NewDecoder(contracts.EvmBtcMetaData.ABI)
	if err != nil {
		fmt.Printf("Failed to create decoder: %v\n", err)
		os.Exit(1)
	}

	// Example error code
	errorCode := "0xe2fe2726"
	errorMsg, err := decoder.DecodeError(errorCode)
	if err != nil {
		fmt.Printf("Failed to decode error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Decoded error: %s\n", errorMsg)
}
