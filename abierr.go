package abierr

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/rpc"
)

type Decoder struct {
	abi abi.ABI
}

func NewDecoder(contractABI abi.ABI) *Decoder {
	return &Decoder{
		abi: contractABI,
	}
}

func (d *Decoder) Decode(err error) (string, error) {
	dataErr, ok := err.(rpc.DataError)
	if !ok {
		return "", fmt.Errorf("not a data error: %w", err)
	}

	errorData := dataErr.ErrorData()
	if errorData == nil {
		return "", fmt.Errorf("no error data")
	}

	hexStr, ok := errorData.(string)
	if !ok {
		return "", fmt.Errorf("error data is not string")
	}

	hexStr = strings.TrimPrefix(hexStr, "0x")
	errorBytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", fmt.Errorf("failed to decode error data: %w", err)
	}

	if len(errorBytes) < 4 {
		return "", fmt.Errorf("invalid error data length")
	}

	errorSelector := hex.EncodeToString(errorBytes[:4])
	for _, abiError := range d.abi.Errors {
		if hex.EncodeToString(abiError.ID[:4]) == errorSelector {
			// Try to unpack the error data
			unpacked, err := abiError.Unpack(errorBytes[4:])
			if err != nil {
				return abiError.Name, nil // Return just the error name if we can't unpack params
			}

			// Format the error with parameters
			return fmt.Sprintf("contract error: %s with params: %v", abiError.Name, unpacked), nil
		}
	}

	return "", fmt.Errorf("unknown error selector: %s", errorSelector)
}
