# abierr

A Go library for decoding Ethereum ABI-encoded error messages.

## Installation

```bash
go get github.com/phathdt/abierr
```

## Usage

```go
package main

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/phathdt/abierr"
)

func main() {
	// Example ABI with error definitions
	contractABI, err := abi.JSON(strings.NewReader(`{
		"errors": [
			{
				"name": "InsufficientBalance",
				"inputs": [
					{
						"name": "available",
						"type": "uint256"
					},
					{
						"name": "required",
						"type": "uint256"
					}
				]
			}
		]
	}`))
	if err != nil {
		panic(err)
	}

	decoder := abierr.NewDecoder(contractABI)

	// Example error data (hex string)
	errData := "0x1234567800000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002"

	// Create a mock error
	mockErr := &rpc.DataError{
		ErrorData: errData,
	}

	// Decode the error
	decoded, err := decoder.Decode(mockErr)
	if err != nil {
		panic(err)
	}

	fmt.Println(decoded) // Output: InsufficientBalance(1, 2)
}
```

## Features

- Decodes ABI-encoded error messages from Ethereum smart contracts
- Supports error messages with parameters
- Handles both single and multiple parameter errors
- Provides human-readable error messages

## License

MIT
