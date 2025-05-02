package abierr

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/rpc"
)

type Decoder struct {
	errors map[string]abi.Error
}

func NewDecoder(abis ...string) (*Decoder, error) {
	if len(abis) == 0 {
		return nil, fmt.Errorf("at least one ABI must be provided")
	}

	errors := make(map[string]abi.Error)

	for _, abiStr := range abis {
		parsedABI, err := abi.JSON(strings.NewReader(abiStr))
		if err != nil {
			return nil, fmt.Errorf("failed to parse ABI: %w", err)
		}

		for _, err := range parsedABI.Errors {
			selector := hex.EncodeToString(err.ID[:4])
			errors[selector] = err
		}
	}

	return &Decoder{
		errors: errors,
	}, nil
}

func (d *Decoder) DecodeError(errorCode string) (string, error) {
	if !strings.HasPrefix(errorCode, "0x") {
		errorCode = "0x" + errorCode
	}

	errorBytes, err := hex.DecodeString(strings.TrimPrefix(errorCode, "0x"))
	if err != nil {
		return "", fmt.Errorf("failed to decode error code: %w", err)
	}

	if len(errorBytes) < 4 {
		return "", fmt.Errorf("invalid error code length")
	}

	errorSelector := hex.EncodeToString(errorBytes[:4])
	if abiError, exists := d.errors[errorSelector]; exists {
		// Try to unpack the error data
		unpacked, err := abiError.Unpack(errorBytes[4:])
		if err != nil {
			return abiError.Name, nil // Return just the error name if we can't unpack params
		}

		// Format the error with parameters
		return fmt.Sprintf("contract error: %s with params: %v", abiError.Name, unpacked), nil
	}

	return "", fmt.Errorf("unknown error: 0x%s", errorSelector)
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

	return d.DecodeError(hexStr)
}
