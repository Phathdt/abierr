package client

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/Phathdt/abierr"
	"github.com/Phathdt/abierr/example/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ERC20Client struct {
	contract   *contracts.Erc20Owner
	privateKey *ecdsa.PrivateKey
	chainID    *big.Int
	address    common.Address
	decoder    *abierr.Decoder
}

func NewERC20Client(rpcURL string, privateKeyStr string, contractAddress string) (*ERC20Client, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %w", err)
	}

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %w", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %w", err)
	}

	contractAddr := common.HexToAddress(contractAddress)
	contract, err := contracts.NewErc20Owner(contractAddr, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create contract instance: %w", err)
	}

	decoder, err := abierr.NewDecoder(contracts.Erc20MetaData.ABI, contracts.Erc20OwnerMetaData.ABI)
	if err != nil {
		return nil, fmt.Errorf("failed to create decoder: %w", err)
	}

	return &ERC20Client{
		contract:   contract,
		privateKey: privateKey,
		chainID:    chainID,
		address:    contractAddr,
		decoder:    decoder,
	}, nil
}

func (c *ERC20Client) Mint(toAddress string, amount *big.Int) (common.Hash, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(c.privateKey, c.chainID)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to create transactor: %w", err)
	}

	to := common.HexToAddress(toAddress)
	tx, err := c.contract.Mint(auth, to, amount)
	if err != nil {
		decodedErr, decodeErr := c.decoder.Decode(err)
		if decodeErr == nil {
			return common.Hash{}, fmt.Errorf("contract error: %s", decodedErr)
		}
		return common.Hash{}, decodeErr
	}

	return tx.Hash(), nil
}
