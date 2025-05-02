// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ITypesAffiliate is an auto generated low-level Go binding around an user-defined struct.
type ITypesAffiliate struct {
	AggregatedValue *big.Int
	Schema          string
	Data            []byte
}

// ITypesBundlePayment is an auto generated low-level Go binding around an user-defined struct.
type ITypesBundlePayment struct {
	TradeIds    [][32]byte
	SignedAt    uint64
	StartIdx    uint64
	PaymentTxId []byte
	Signature   []byte
}

// ITypesFailureDetails is an auto generated low-level Go binding around an user-defined struct.
type ITypesFailureDetails struct {
	Stage    *big.Int
	MsgError []byte
}

// ITypesFeeDetails is an auto generated low-level Go binding around an user-defined struct.
type ITypesFeeDetails struct {
	TotalAmount *big.Int
	PFeeAmount  *big.Int
	AFeeAmount  *big.Int
	PFeeRate    *big.Int
	AFeeRate    *big.Int
}

// ITypesPMMSelection is an auto generated low-level Go binding around an user-defined struct.
type ITypesPMMSelection struct {
	RfqInfo ITypesRFQInfo
	PmmInfo ITypesSelectedPMMInfo
}

// ITypesRFQInfo is an auto generated low-level Go binding around an user-defined struct.
type ITypesRFQInfo struct {
	MinAmountOut     *big.Int
	TradeTimeout     uint64
	RfqInfoSignature []byte
}

// ITypesRefundPresign is an auto generated low-level Go binding around an user-defined struct.
type ITypesRefundPresign struct {
	RefundAddress []byte
	Presigns      [][]byte
}

// ITypesScriptInfo is an auto generated low-level Go binding around an user-defined struct.
type ITypesScriptInfo struct {
	DepositInfo            [5][]byte
	UserEphemeralL2Address common.Address
	ScriptTimeout          uint64
}

// ITypesSelectedPMMInfo is an auto generated low-level Go binding around an user-defined struct.
type ITypesSelectedPMMInfo struct {
	AmountOut     *big.Int
	SelectedPMMId [32]byte
	Info          [2][]byte
	SigExpiry     uint64
}

// ITypesSettlementPresign is an auto generated low-level Go binding around an user-defined struct.
type ITypesSettlementPresign struct {
	PmmId          [32]byte
	PmmRecvAddress []byte
	Presigns       [][]byte
}

// ITypesTradeData is an auto generated low-level Go binding around an user-defined struct.
type ITypesTradeData struct {
	SessionId  *big.Int
	TradeInfo  ITypesTradeInfo
	ScriptInfo ITypesScriptInfo
}

// ITypesTradeFinalization is an auto generated low-level Go binding around an user-defined struct.
type ITypesTradeFinalization struct {
	BundlerHash [32]byte
	Index       *big.Int
	PaymentTxId []byte
	ReleaseTxId []byte
	RefundTxId  []byte
	IsConfirmed bool
}

// ITypesTradeInfo is an auto generated low-level Go binding around an user-defined struct.
type ITypesTradeInfo struct {
	AmountIn  *big.Int
	FromChain [3][]byte
	ToChain   [3][]byte
}

// EvmBtcMetaData contains all meta data concerning the EvmBtc contract.
var EvmBtcMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIRouter\",\"name\":\"router_\",\"type\":\"address\"},{\"internalType\":\"contractIVaultRegistry\",\"name\":\"registry_\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"intervalTime\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotRefundBeforeTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotReportBeforeTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DeadlineExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceededAffiliateFeeLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InSuspension\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientQuoteAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMPCAssetPubkey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMPCSign\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPMMSelection\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPMMSign\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPaymentIdLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedStage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentStage\",\"type\":\"uint256\"}],\"name\":\"InvalidProcedureState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRFQSign\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRefundAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTradeId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotPending\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OutdatedSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PMMAddrNotMatched\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PMMNotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PresignsMustBeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StageAlreadyFinalized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VaultAddrNotMatched\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mpc\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalFeeRate\",\"type\":\"uint256\"}],\"name\":\"DepositConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mpc\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stage\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"referenceInfo\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"msgError\",\"type\":\"bytes\"}],\"name\":\"FailureReported\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"network\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"paymentTxId\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"bundle\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startIdx\",\"type\":\"uint256\"}],\"name\":\"MadePayment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mpc\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"paymentTxId\",\"type\":\"bytes\"}],\"name\":\"PaymentConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mpc\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"refundTxId\",\"type\":\"bytes\"}],\"name\":\"Refunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"selectedPMMId\",\"type\":\"bytes32\"}],\"name\":\"SelectedPMM\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mpc\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"releaseTxId\",\"type\":\"bytes\"}],\"name\":\"SettlementConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"network\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"depositTxId\",\"type\":\"bytes\"}],\"name\":\"TradeInfoSubmitted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"depositFromList\",\"type\":\"bytes[]\"}],\"name\":\"confirmDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"confirmPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"releaseTxId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"confirmSettlement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epoch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"}],\"name\":\"getAffiliate\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"aggregatedValue\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"schema\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structITypes.Affiliate\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"}],\"name\":\"getCurrentStage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"}],\"name\":\"getDepositAddressList\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"}],\"name\":\"getFailureInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stage\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"msgError\",\"type\":\"bytes\"}],\"internalType\":\"structITypes.FailureDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"}],\"name\":\"getFeeDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"pFeeRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"aFeeRate\",\"type\":\"uint128\"}],\"internalType\":\"structITypes.FeeDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"}],\"name\":\"getLastSignedPayment\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"}],\"name\":\"getPMMSelection\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"tradeTimeout\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"rfqInfoSignature\",\"type\":\"bytes\"}],\"internalType\":\"structITypes.RFQInfo\",\"name\":\"rfqInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"selectedPMMId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[2]\",\"name\":\"info\",\"type\":\"bytes[2]\"},{\"internalType\":\"uint64\",\"name\":\"sigExpiry\",\"type\":\"uint64\"}],\"internalType\":\"structITypes.SelectedPMMInfo\",\"name\":\"pmmInfo\",\"type\":\"tuple\"}],\"internalType\":\"structITypes.PMMSelection\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochNo\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromIdx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIdx\",\"type\":\"uint256\"}],\"name\":\"getPendingTrades\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"list\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochNo\",\"type\":\"uint256\"}],\"name\":\"getPendingTradesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"}],\"name\":\"getRefundPresign\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"refundAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"presigns\",\"type\":\"bytes[]\"}],\"internalType\":\"structITypes.RefundPresign\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"}],\"name\":\"getSettlementPresign\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"pmmId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"pmmRecvAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"presigns\",\"type\":\"bytes[]\"}],\"internalType\":\"structITypes.SettlementPresign[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"}],\"name\":\"getTradeData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"bytes[3]\",\"name\":\"fromChain\",\"type\":\"bytes[3]\"},{\"internalType\":\"bytes[3]\",\"name\":\"toChain\",\"type\":\"bytes[3]\"}],\"internalType\":\"structITypes.TradeInfo\",\"name\":\"tradeInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes[5]\",\"name\":\"depositInfo\",\"type\":\"bytes[5]\"},{\"internalType\":\"address\",\"name\":\"userEphemeralL2Address\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"scriptTimeout\",\"type\":\"uint64\"}],\"internalType\":\"structITypes.ScriptInfo\",\"name\":\"scriptInfo\",\"type\":\"tuple\"}],\"internalType\":\"structITypes.TradeData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"}],\"name\":\"getTradeFinalization\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"bundlerHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymentTxId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"releaseTxId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"refundTxId\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isConfirmed\",\"type\":\"bool\"}],\"internalType\":\"structITypes.TradeFinalization\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"indexOfTrade\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"tradeIds\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint64\",\"name\":\"signedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startIdx\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"paymentTxId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structITypes.BundlePayment\",\"name\":\"bundle\",\"type\":\"tuple\"}],\"name\":\"makePayment\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"pmmId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxAffiliateFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"refundTxId\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractIVaultRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"msgError\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"report\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"contractIRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"tradeTimeout\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"rfqInfoSignature\",\"type\":\"bytes\"}],\"internalType\":\"structITypes.RFQInfo\",\"name\":\"rfqInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"selectedPMMId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[2]\",\"name\":\"info\",\"type\":\"bytes[2]\"},{\"internalType\":\"uint64\",\"name\":\"sigExpiry\",\"type\":\"uint64\"}],\"internalType\":\"structITypes.SelectedPMMInfo\",\"name\":\"pmmInfo\",\"type\":\"tuple\"}],\"internalType\":\"structITypes.PMMSelection\",\"name\":\"info\",\"type\":\"tuple\"}],\"name\":\"selectPMM\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newRate\",\"type\":\"uint256\"}],\"name\":\"setMaxAffiliateFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRouter\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRegistry\",\"type\":\"address\"}],\"name\":\"setVaultRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tradeId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"bytes[3]\",\"name\":\"fromChain\",\"type\":\"bytes[3]\"},{\"internalType\":\"bytes[3]\",\"name\":\"toChain\",\"type\":\"bytes[3]\"}],\"internalType\":\"structITypes.TradeInfo\",\"name\":\"tradeInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes[5]\",\"name\":\"depositInfo\",\"type\":\"bytes[5]\"},{\"internalType\":\"address\",\"name\":\"userEphemeralL2Address\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"scriptTimeout\",\"type\":\"uint64\"}],\"internalType\":\"structITypes.ScriptInfo\",\"name\":\"scriptInfo\",\"type\":\"tuple\"}],\"internalType\":\"structITypes.TradeData\",\"name\":\"tradeData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"aggregatedValue\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"schema\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structITypes.Affiliate\",\"name\":\"affiliateInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"pmmId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"pmmRecvAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"presigns\",\"type\":\"bytes[]\"}],\"internalType\":\"structITypes.SettlementPresign[]\",\"name\":\"settlementPresigns\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"refundAddress\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"presigns\",\"type\":\"bytes[]\"}],\"internalType\":\"structITypes.RefundPresign\",\"name\":\"refundPresign\",\"type\":\"tuple\"}],\"name\":\"submitTrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeOfHandler\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// EvmBtcABI is the input ABI used to generate the binding from.
// Deprecated: Use EvmBtcMetaData.ABI instead.
var EvmBtcABI = EvmBtcMetaData.ABI

// EvmBtc is an auto generated Go binding around an Ethereum contract.
type EvmBtc struct {
	EvmBtcCaller     // Read-only binding to the contract
	EvmBtcTransactor // Write-only binding to the contract
	EvmBtcFilterer   // Log filterer for contract events
}

// EvmBtcCaller is an auto generated read-only Go binding around an Ethereum contract.
type EvmBtcCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EvmBtcTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EvmBtcTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EvmBtcFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EvmBtcFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EvmBtcSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EvmBtcSession struct {
	Contract     *EvmBtc           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EvmBtcCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EvmBtcCallerSession struct {
	Contract *EvmBtcCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EvmBtcTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EvmBtcTransactorSession struct {
	Contract     *EvmBtcTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EvmBtcRaw is an auto generated low-level Go binding around an Ethereum contract.
type EvmBtcRaw struct {
	Contract *EvmBtc // Generic contract binding to access the raw methods on
}

// EvmBtcCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EvmBtcCallerRaw struct {
	Contract *EvmBtcCaller // Generic read-only contract binding to access the raw methods on
}

// EvmBtcTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EvmBtcTransactorRaw struct {
	Contract *EvmBtcTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEvmBtc creates a new instance of EvmBtc, bound to a specific deployed contract.
func NewEvmBtc(address common.Address, backend bind.ContractBackend) (*EvmBtc, error) {
	contract, err := bindEvmBtc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EvmBtc{EvmBtcCaller: EvmBtcCaller{contract: contract}, EvmBtcTransactor: EvmBtcTransactor{contract: contract}, EvmBtcFilterer: EvmBtcFilterer{contract: contract}}, nil
}

// NewEvmBtcCaller creates a new read-only instance of EvmBtc, bound to a specific deployed contract.
func NewEvmBtcCaller(address common.Address, caller bind.ContractCaller) (*EvmBtcCaller, error) {
	contract, err := bindEvmBtc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EvmBtcCaller{contract: contract}, nil
}

// NewEvmBtcTransactor creates a new write-only instance of EvmBtc, bound to a specific deployed contract.
func NewEvmBtcTransactor(address common.Address, transactor bind.ContractTransactor) (*EvmBtcTransactor, error) {
	contract, err := bindEvmBtc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EvmBtcTransactor{contract: contract}, nil
}

// NewEvmBtcFilterer creates a new log filterer instance of EvmBtc, bound to a specific deployed contract.
func NewEvmBtcFilterer(address common.Address, filterer bind.ContractFilterer) (*EvmBtcFilterer, error) {
	contract, err := bindEvmBtc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EvmBtcFilterer{contract: contract}, nil
}

// bindEvmBtc binds a generic wrapper to an already deployed contract.
func bindEvmBtc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EvmBtcMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EvmBtc *EvmBtcRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EvmBtc.Contract.EvmBtcCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EvmBtc *EvmBtcRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EvmBtc.Contract.EvmBtcTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EvmBtc *EvmBtcRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EvmBtc.Contract.EvmBtcTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EvmBtc *EvmBtcCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EvmBtc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EvmBtc *EvmBtcTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EvmBtc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EvmBtc *EvmBtcTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EvmBtc.Contract.contract.Transact(opts, method, params...)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint64, uint64, uint64)
func (_EvmBtc *EvmBtcCaller) Epoch(opts *bind.CallOpts) (uint64, uint64, uint64, error) {
	var out []interface{}
	err := _EvmBtc.contract.Call(opts, &out, "epoch")

	if err != nil {
		return *new(uint64), *new(uint64), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)
	out2 := *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return out0, out1, out2, err

}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint64, uint64, uint64)
func (_EvmBtc *EvmBtcSession) Epoch() (uint64, uint64, uint64, error) {
	return _EvmBtc.Contract.Epoch(&_EvmBtc.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint64, uint64, uint64)
func (_EvmBtc *EvmBtcCallerSession) Epoch() (uint64, uint64, uint64, error) {
	return _EvmBtc.Contract.Epoch(&_EvmBtc.CallOpts)
}

// GetAffiliate is a free data retrieval call binding the contract method 0xb58ec773.
//
// Solidity: function getAffiliate(bytes32 tradeId) view returns((uint256,string,bytes))
func (_EvmBtc *EvmBtcCaller) GetAffiliate(opts *bind.CallOpts, tradeId [32]byte) (ITypesAffiliate, error) {
	var out []interface{}
	err := _EvmBtc.contract.Call(opts, &out, "getAffiliate", tradeId)

	if err != nil {
		return *new(ITypesAffiliate), err
	}

	out0 := *abi.ConvertType(out[0], new(ITypesAffiliate)).(*ITypesAffiliate)

	return out0, err

}

// GetAffiliate is a free data retrieval call binding the contract method 0xb58ec773.
//
// Solidity: function getAffiliate(bytes32 tradeId) view returns((uint256,string,bytes))
func (_EvmBtc *EvmBtcSession) GetAffiliate(tradeId [32]byte) (ITypesAffiliate, error) {
	return _EvmBtc.Contract.GetAffiliate(&_EvmBtc.CallOpts, tradeId)
}

// GetAffiliate is a free data retrieval call binding the contract method 0xb58ec773.
//
// Solidity: function getAffiliate(bytes32 tradeId) view returns((uint256,string,bytes))
func (_EvmBtc *EvmBtcCallerSession) GetAffiliate(tradeId [32]byte) (ITypesAffiliate, error) {
	return _EvmBtc.Contract.GetAffiliate(&_EvmBtc.CallOpts, tradeId)
}

// GetCurrentStage is a free data retrieval call binding the contract method 0x10178310.
//
// Solidity: function getCurrentStage(bytes32 tradeId) view returns(uint256)
func (_EvmBtc *EvmBtcCaller) GetCurrentStage(opts *bind.CallOpts, tradeId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _EvmBtc.contract.Call(opts, &out, "getCurrentStage", tradeId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentStage is a free data retrieval call binding the contract method 0x10178310.
//
// Solidity: function getCurrentStage(bytes32 tradeId) view returns(uint256)
func (_EvmBtc *EvmBtcSession) GetCurrentStage(tradeId [32]byte) (*big.Int, error) {
	return _EvmBtc.Contract.GetCurrentStage(&_EvmBtc.CallOpts, tradeId)
}

// GetCurrentStage is a free data retrieval call binding the contract method 0x10178310.
//
// Solidity: function getCurrentStage(bytes32 tradeId) view returns(uint256)
func (_EvmBtc *EvmBtcCallerSession) GetCurrentStage(tradeId [32]byte) (*big.Int, error) {
	return _EvmBtc.Contract.GetCurrentStage(&_EvmBtc.CallOpts, tradeId)
}

// GetDepositAddressList is a free data retrieval call binding the contract method 0xae4637a8.
//
// Solidity: function getDepositAddressList(bytes32 tradeId) view returns(bytes[])
func (_EvmBtc *EvmBtcCaller) GetDepositAddressList(opts *bind.CallOpts, tradeId [32]byte) ([][]byte, error) {
	var out []interface{}
	err := _EvmBtc.contract.Call(opts, &out, "getDepositAddressList", tradeId)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetDepositAddressList is a free data retrieval call binding the contract method 0xae4637a8.
//
// Solidity: function getDepositAddressList(bytes32 tradeId) view returns(bytes[])
func (_EvmBtc *EvmBtcSession) GetDepositAddressList(tradeId [32]byte) ([][]byte, error) {
	return _EvmBtc.Contract.GetDepositAddressList(&_EvmBtc.CallOpts, tradeId)
}

// GetDepositAddressList is a free data retrieval call binding the contract method 0xae4637a8.
//
// Solidity: function getDepositAddressList(bytes32 tradeId) view returns(bytes[])
func (_EvmBtc *EvmBtcCallerSession) GetDepositAddressList(tradeId [32]byte) ([][]byte, error) {
	return _EvmBtc.Contract.GetDepositAddressList(&_EvmBtc.CallOpts, tradeId)
}

// GetFailureInfo is a free data retrieval call binding the contract method 0x9db4ab75.
//
// Solidity: function getFailureInfo(bytes32 tradeId) view returns((uint256,bytes))
func (_EvmBtc *EvmBtcCaller) GetFailureInfo(opts *bind.CallOpts, tradeId [32]byte) (ITypesFailureDetails, error) {
	var out []interface{}
	err := _EvmBtc.contract.Call(opts, &out, "getFailureInfo", tradeId)

	if err != nil {
		return *new(ITypesFailureDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(ITypesFailureDetails)).(*ITypesFailureDetails)

	return out0, err

}

// GetFailureInfo is a free data retrieval call binding the contract method 0x9db4ab75.
//
// Solidity: function getFailureInfo(bytes32 tradeId) view returns((uint256,bytes))
func (_EvmBtc *EvmBtcSession) GetFailureInfo(tradeId [32]byte) (ITypesFailureDetails, error) {
	return _EvmBtc.Contract.GetFailureInfo(&_EvmBtc.CallOpts, tradeId)
}

// GetFailureInfo is a free data retrieval call binding the contract method 0x9db4ab75.
//
// Solidity: function getFailureInfo(bytes32 tradeId) view returns((uint256,bytes))
func (_EvmBtc *EvmBtcCallerSession) GetFailureInfo(tradeId [32]byte) (ITypesFailureDetails, error) {
	return _EvmBtc.Contract.GetFailureInfo(&_EvmBtc.CallOpts, tradeId)
}

// GetFeeDetails is a free data retrieval call binding the contract method 0x9108e266.
//
// Solidity: function getFeeDetails(bytes32 tradeId) view returns((uint256,uint256,uint256,uint128,uint128))
func (_EvmBtc *EvmBtcCaller) GetFeeDetails(opts *bind.CallOpts, tradeId [32]byte) (ITypesFeeDetails, error) {
	var out []interface{}
	err := _EvmBtc.contract.Call(opts, &out, "getFeeDetails", tradeId)

	if err != nil {
		return *new(ITypesFeeDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(ITypesFeeDetails)).(*ITypesFeeDetails)

	return out0, err

}

// GetFeeDetails is a free data retrieval call binding the contract method 0x9108e266.
//
// Solidity: function getFeeDetails(bytes32 tradeId) view returns((uint256,uint256,uint256,uint128,uint128))
func (_EvmBtc *EvmBtcSession) GetFeeDetails(tradeId [32]byte) (ITypesFeeDetails, error) {
	return _EvmBtc.Contract.GetFeeDetails(&_EvmBtc.CallOpts, tradeId)
}

// GetFeeDetails is a free data retrieval call binding the contract method 0x9108e266.
//
// Solidity: function getFeeDetails(bytes32 tradeId) view returns((uint256,uint256,uint256,uint128,uint128))
func (_EvmBtc *EvmBtcCallerSession) GetFeeDetails(tradeId [32]byte) (ITypesFeeDetails, error) {
	return _EvmBtc.Contract.GetFeeDetails(&_EvmBtc.CallOpts, tradeId)
}

// GetLastSignedPayment is a free data retrieval call binding the contract method 0x140bf2a6.
//
// Solidity: function getLastSignedPayment(bytes32 tradeId) view returns(uint64)
func (_EvmBtc *EvmBtcCaller) GetLastSignedPayment(opts *bind.CallOpts, tradeId [32]byte) (uint64, error) {
	var out []interface{}
	err := _EvmBtc.contract.Call(opts, &out, "getLastSignedPayment", tradeId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLastSignedPayment is a free data retrieval call binding the contract method 0x140bf2a6.
//
// Solidity: function getLastSignedPayment(bytes32 tradeId) view returns(uint64)
func (_EvmBtc *EvmBtcSession) GetLastSignedPayment(tradeId [32]byte) (uint64, error) {
	return _EvmBtc.Contract.GetLastSignedPayment(&_EvmBtc.CallOpts, tradeId)
}

// GetLastSignedPayment is a free data retrieval call binding the contract method 0x140bf2a6.
//
// Solidity: function getLastSignedPayment(bytes32 tradeId) view returns(uint64)
func (_EvmBtc *EvmBtcCallerSession) GetLastSignedPayment(tradeId [32]byte) (uint64, error) {
	return _EvmBtc.Contract.GetLastSignedPayment(&_EvmBtc.CallOpts, tradeId)
}

// GetPMMSelection is a free data retrieval call binding the contract method 0x26cea89b.
//
// Solidity: function getPMMSelection(bytes32 tradeId) view returns(((uint256,uint64,bytes),(uint256,bytes32,bytes[2],uint64)))
func (_EvmBtc *EvmBtcCaller) GetPMMSelection(opts *bind.CallOpts, tradeId [32]byte) (ITypesPMMSelection, error) {
	var out []interface{}
	err := _EvmBtc.contract.Call(opts, &out, "getPMMSelection", tradeId)

	if err != nil {
		return *new(ITypesPMMSelection), err
	}

	out0 := *abi.ConvertType(out[0], new(ITypesPMMSelection)).(*ITypesPMMSelection)

	return out0, err

}

// GetPMMSelection is a free data retrieval call binding the contract method 0x26cea89b.
//
// Solidity: function getPMMSelection(bytes32 tradeId) view returns(((uint256,uint64,bytes),(uint256,bytes32,bytes[2],uint64)))
func (_EvmBtc *EvmBtcSession) GetPMMSelection(tradeId [32]byte) (ITypesPMMSelection, error) {
	return _EvmBtc.Contract.GetPMMSelection(&_EvmBtc.CallOpts, tradeId)
}

// GetPMMSelection is a free data retrieval call binding the contract method 0x26cea89b.
//
// Solidity: function getPMMSelection(bytes32 tradeId) view returns(((uint256,uint64,bytes),(uint256,bytes32,bytes[2],uint64)))
func (_EvmBtc *EvmBtcCallerSession) GetPMMSelection(tradeId [32]byte) (ITypesPMMSelection, error) {
	return _EvmBtc.Contract.GetPMMSelection(&_EvmBtc.CallOpts, tradeId)
}

// GetPendingTrades is a free data retrieval call binding the contract method 0x9a6d66ec.
//
// Solidity: function getPendingTrades(uint256 epochNo, uint256 fromIdx, uint256 toIdx) view returns(bytes32[] list)
func (_EvmBtc *EvmBtcCaller) GetPendingTrades(opts *bind.CallOpts, epochNo *big.Int, fromIdx *big.Int, toIdx *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _EvmBtc.contract.Call(opts, &out, "getPendingTrades", epochNo, fromIdx, toIdx)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetPendingTrades is a free data retrieval call binding the contract method 0x9a6d66ec.
//
// Solidity: function getPendingTrades(uint256 epochNo, uint256 fromIdx, uint256 toIdx) view returns(bytes32[] list)
func (_EvmBtc *EvmBtcSession) GetPendingTrades(epochNo *big.Int, fromIdx *big.Int, toIdx *big.Int) ([][32]byte, error) {
	return _EvmBtc.Contract.GetPendingTrades(&_EvmBtc.CallOpts, epochNo, fromIdx, toIdx)
}

// GetPendingTrades is a free data retrieval call binding the contract method 0x9a6d66ec.
//
// Solidity: function getPendingTrades(uint256 epochNo, uint256 fromIdx, uint256 toIdx) view returns(bytes32[] list)
func (_EvmBtc *EvmBtcCallerSession) GetPendingTrades(epochNo *big.Int, fromIdx *big.Int, toIdx *big.Int) ([][32]byte, error) {
	return _EvmBtc.Contract.GetPendingTrades(&_EvmBtc.CallOpts, epochNo, fromIdx, toIdx)
}

// GetPendingTradesCount is a free data retrieval call binding the contract method 0xabfb7c2d.
//
// Solidity: function getPendingTradesCount(uint256 epochNo) view returns(uint256)
func (_EvmBtc *EvmBtcCaller) GetPendingTradesCount(opts *bind.CallOpts, epochNo *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EvmBtc.contract.Call(opts, &out, "getPendingTradesCount", epochNo)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingTradesCount is a free data retrieval call binding the contract method 0xabfb7c2d.
//
// Solidity: function getPendingTradesCount(uint256 epochNo) view returns(uint256)
func (_EvmBtc *EvmBtcSession) GetPendingTradesCount(epochNo *big.Int) (*big.Int, error) {
	return _EvmBtc.Contract.GetPendingTradesCount(&_EvmBtc.CallOpts, epochNo)
}

// GetPendingTradesCount is a free data retrieval call binding the contract method 0xabfb7c2d.
//
// Solidity: function getPendingTradesCount(uint256 epochNo) view returns(uint256)
func (_EvmBtc *EvmBtcCallerSession) GetPendingTradesCount(epochNo *big.Int) (*big.Int, error) {
	return _EvmBtc.Contract.GetPendingTradesCount(&_EvmBtc.CallOpts, epochNo)
}

// GetRefundPresign is a free data retrieval call binding the contract method 0xb55350c5.
//
// Solidity: function getRefundPresign(bytes32 tradeId) view returns((bytes,bytes[]))
func (_EvmBtc *EvmBtcCaller) GetRefundPresign(opts *bind.CallOpts, tradeId [32]byte) (ITypesRefundPresign, error) {
	var out []interface{}
	err := _EvmBtc.contract.Call(opts, &out, "getRefundPresign", tradeId)

	if err != nil {
		return *new(ITypesRefundPresign), err
	}

	out0 := *abi.ConvertType(out[0], new(ITypesRefundPresign)).(*ITypesRefundPresign)

	return out0, err

}

// GetRefundPresign is a free data retrieval call binding the contract method 0xb55350c5.
//
// Solidity: function getRefundPresign(bytes32 tradeId) view returns((bytes,bytes[]))
func (_EvmBtc *EvmBtcSession) GetRefundPresign(tradeId [32]byte) (ITypesRefundPresign, error) {
	return _EvmBtc.Contract.GetRefundPresign(&_EvmBtc.CallOpts, tradeId)
}

// GetRefundPresign is a free data retrieval call binding the contract method 0xb55350c5.
//
// Solidity: function getRefundPresign(bytes32 tradeId) view returns((bytes,bytes[]))
func (_EvmBtc *EvmBtcCallerSession) GetRefundPresign(tradeId [32]byte) (ITypesRefundPresign, error) {
	return _EvmBtc.Contract.GetRefundPresign(&_EvmBtc.CallOpts, tradeId)
}

// GetSettlementPresign is a free data retrieval call binding the contract method 0x8b76747f.
//
// Solidity: function getSettlementPresign(bytes32 tradeId) view returns((bytes32,bytes,bytes[])[])
func (_EvmBtc *EvmBtcCaller) GetSettlementPresign(opts *bind.CallOpts, tradeId [32]byte) ([]ITypesSettlementPresign, error) {
	var out []interface{}
	err := _EvmBtc.contract.Call(opts, &out, "getSettlementPresign", tradeId)

	if err != nil {
		return *new([]ITypesSettlementPresign), err
	}

	out0 := *abi.ConvertType(out[0], new([]ITypesSettlementPresign)).(*[]ITypesSettlementPresign)

	return out0, err

}

// GetSettlementPresign is a free data retrieval call binding the contract method 0x8b76747f.
//
// Solidity: function getSettlementPresign(bytes32 tradeId) view returns((bytes32,bytes,bytes[])[])
func (_EvmBtc *EvmBtcSession) GetSettlementPresign(tradeId [32]byte) ([]ITypesSettlementPresign, error) {
	return _EvmBtc.Contract.GetSettlementPresign(&_EvmBtc.CallOpts, tradeId)
}

// GetSettlementPresign is a free data retrieval call binding the contract method 0x8b76747f.
//
// Solidity: function getSettlementPresign(bytes32 tradeId) view returns((bytes32,bytes,bytes[])[])
func (_EvmBtc *EvmBtcCallerSession) GetSettlementPresign(tradeId [32]byte) ([]ITypesSettlementPresign, error) {
	return _EvmBtc.Contract.GetSettlementPresign(&_EvmBtc.CallOpts, tradeId)
}

// GetTradeData is a free data retrieval call binding the contract method 0x2afdebb1.
//
// Solidity: function getTradeData(bytes32 tradeId) view returns((uint256,(uint256,bytes[3],bytes[3]),(bytes[5],address,uint64)))
func (_EvmBtc *EvmBtcCaller) GetTradeData(opts *bind.CallOpts, tradeId [32]byte) (ITypesTradeData, error) {
	var out []interface{}
	err := _EvmBtc.contract.Call(opts, &out, "getTradeData", tradeId)

	if err != nil {
		return *new(ITypesTradeData), err
	}

	out0 := *abi.ConvertType(out[0], new(ITypesTradeData)).(*ITypesTradeData)

	return out0, err

}

// GetTradeData is a free data retrieval call binding the contract method 0x2afdebb1.
//
// Solidity: function getTradeData(bytes32 tradeId) view returns((uint256,(uint256,bytes[3],bytes[3]),(bytes[5],address,uint64)))
func (_EvmBtc *EvmBtcSession) GetTradeData(tradeId [32]byte) (ITypesTradeData, error) {
	return _EvmBtc.Contract.GetTradeData(&_EvmBtc.CallOpts, tradeId)
}

// GetTradeData is a free data retrieval call binding the contract method 0x2afdebb1.
//
// Solidity: function getTradeData(bytes32 tradeId) view returns((uint256,(uint256,bytes[3],bytes[3]),(bytes[5],address,uint64)))
func (_EvmBtc *EvmBtcCallerSession) GetTradeData(tradeId [32]byte) (ITypesTradeData, error) {
	return _EvmBtc.Contract.GetTradeData(&_EvmBtc.CallOpts, tradeId)
}

// GetTradeFinalization is a free data retrieval call binding the contract method 0xb7068b76.
//
// Solidity: function getTradeFinalization(bytes32 tradeId) view returns((bytes32,uint256,bytes,bytes,bytes,bool))
func (_EvmBtc *EvmBtcCaller) GetTradeFinalization(opts *bind.CallOpts, tradeId [32]byte) (ITypesTradeFinalization, error) {
	var out []interface{}
	err := _EvmBtc.contract.Call(opts, &out, "getTradeFinalization", tradeId)

	if err != nil {
		return *new(ITypesTradeFinalization), err
	}

	out0 := *abi.ConvertType(out[0], new(ITypesTradeFinalization)).(*ITypesTradeFinalization)

	return out0, err

}

// GetTradeFinalization is a free data retrieval call binding the contract method 0xb7068b76.
//
// Solidity: function getTradeFinalization(bytes32 tradeId) view returns((bytes32,uint256,bytes,bytes,bytes,bool))
func (_EvmBtc *EvmBtcSession) GetTradeFinalization(tradeId [32]byte) (ITypesTradeFinalization, error) {
	return _EvmBtc.Contract.GetTradeFinalization(&_EvmBtc.CallOpts, tradeId)
}

// GetTradeFinalization is a free data retrieval call binding the contract method 0xb7068b76.
//
// Solidity: function getTradeFinalization(bytes32 tradeId) view returns((bytes32,uint256,bytes,bytes,bytes,bool))
func (_EvmBtc *EvmBtcCallerSession) GetTradeFinalization(tradeId [32]byte) (ITypesTradeFinalization, error) {
	return _EvmBtc.Contract.GetTradeFinalization(&_EvmBtc.CallOpts, tradeId)
}

// MaxAffiliateFeeRate is a free data retrieval call binding the contract method 0x3750e4f1.
//
// Solidity: function maxAffiliateFeeRate() view returns(uint256)
func (_EvmBtc *EvmBtcCaller) MaxAffiliateFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EvmBtc.contract.Call(opts, &out, "maxAffiliateFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAffiliateFeeRate is a free data retrieval call binding the contract method 0x3750e4f1.
//
// Solidity: function maxAffiliateFeeRate() view returns(uint256)
func (_EvmBtc *EvmBtcSession) MaxAffiliateFeeRate() (*big.Int, error) {
	return _EvmBtc.Contract.MaxAffiliateFeeRate(&_EvmBtc.CallOpts)
}

// MaxAffiliateFeeRate is a free data retrieval call binding the contract method 0x3750e4f1.
//
// Solidity: function maxAffiliateFeeRate() view returns(uint256)
func (_EvmBtc *EvmBtcCallerSession) MaxAffiliateFeeRate() (*big.Int, error) {
	return _EvmBtc.Contract.MaxAffiliateFeeRate(&_EvmBtc.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_EvmBtc *EvmBtcCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EvmBtc.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_EvmBtc *EvmBtcSession) Registry() (common.Address, error) {
	return _EvmBtc.Contract.Registry(&_EvmBtc.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_EvmBtc *EvmBtcCallerSession) Registry() (common.Address, error) {
	return _EvmBtc.Contract.Registry(&_EvmBtc.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_EvmBtc *EvmBtcCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EvmBtc.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_EvmBtc *EvmBtcSession) Router() (common.Address, error) {
	return _EvmBtc.Contract.Router(&_EvmBtc.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_EvmBtc *EvmBtcCallerSession) Router() (common.Address, error) {
	return _EvmBtc.Contract.Router(&_EvmBtc.CallOpts)
}

// TypeOfHandler is a free data retrieval call binding the contract method 0xc4c26c3a.
//
// Solidity: function typeOfHandler() pure returns(string)
func (_EvmBtc *EvmBtcCaller) TypeOfHandler(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EvmBtc.contract.Call(opts, &out, "typeOfHandler")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TypeOfHandler is a free data retrieval call binding the contract method 0xc4c26c3a.
//
// Solidity: function typeOfHandler() pure returns(string)
func (_EvmBtc *EvmBtcSession) TypeOfHandler() (string, error) {
	return _EvmBtc.Contract.TypeOfHandler(&_EvmBtc.CallOpts)
}

// TypeOfHandler is a free data retrieval call binding the contract method 0xc4c26c3a.
//
// Solidity: function typeOfHandler() pure returns(string)
func (_EvmBtc *EvmBtcCallerSession) TypeOfHandler() (string, error) {
	return _EvmBtc.Contract.TypeOfHandler(&_EvmBtc.CallOpts)
}

// ConfirmDeposit is a paid mutator transaction binding the contract method 0xe2e312ea.
//
// Solidity: function confirmDeposit(address requester, bytes32 tradeId, bytes signature, bytes[] depositFromList) returns()
func (_EvmBtc *EvmBtcTransactor) ConfirmDeposit(opts *bind.TransactOpts, requester common.Address, tradeId [32]byte, signature []byte, depositFromList [][]byte) (*types.Transaction, error) {
	return _EvmBtc.contract.Transact(opts, "confirmDeposit", requester, tradeId, signature, depositFromList)
}

// ConfirmDeposit is a paid mutator transaction binding the contract method 0xe2e312ea.
//
// Solidity: function confirmDeposit(address requester, bytes32 tradeId, bytes signature, bytes[] depositFromList) returns()
func (_EvmBtc *EvmBtcSession) ConfirmDeposit(requester common.Address, tradeId [32]byte, signature []byte, depositFromList [][]byte) (*types.Transaction, error) {
	return _EvmBtc.Contract.ConfirmDeposit(&_EvmBtc.TransactOpts, requester, tradeId, signature, depositFromList)
}

// ConfirmDeposit is a paid mutator transaction binding the contract method 0xe2e312ea.
//
// Solidity: function confirmDeposit(address requester, bytes32 tradeId, bytes signature, bytes[] depositFromList) returns()
func (_EvmBtc *EvmBtcTransactorSession) ConfirmDeposit(requester common.Address, tradeId [32]byte, signature []byte, depositFromList [][]byte) (*types.Transaction, error) {
	return _EvmBtc.Contract.ConfirmDeposit(&_EvmBtc.TransactOpts, requester, tradeId, signature, depositFromList)
}

// ConfirmPayment is a paid mutator transaction binding the contract method 0x1e053dc2.
//
// Solidity: function confirmPayment(address requester, bytes32 tradeId, bytes signature) returns()
func (_EvmBtc *EvmBtcTransactor) ConfirmPayment(opts *bind.TransactOpts, requester common.Address, tradeId [32]byte, signature []byte) (*types.Transaction, error) {
	return _EvmBtc.contract.Transact(opts, "confirmPayment", requester, tradeId, signature)
}

// ConfirmPayment is a paid mutator transaction binding the contract method 0x1e053dc2.
//
// Solidity: function confirmPayment(address requester, bytes32 tradeId, bytes signature) returns()
func (_EvmBtc *EvmBtcSession) ConfirmPayment(requester common.Address, tradeId [32]byte, signature []byte) (*types.Transaction, error) {
	return _EvmBtc.Contract.ConfirmPayment(&_EvmBtc.TransactOpts, requester, tradeId, signature)
}

// ConfirmPayment is a paid mutator transaction binding the contract method 0x1e053dc2.
//
// Solidity: function confirmPayment(address requester, bytes32 tradeId, bytes signature) returns()
func (_EvmBtc *EvmBtcTransactorSession) ConfirmPayment(requester common.Address, tradeId [32]byte, signature []byte) (*types.Transaction, error) {
	return _EvmBtc.Contract.ConfirmPayment(&_EvmBtc.TransactOpts, requester, tradeId, signature)
}

// ConfirmSettlement is a paid mutator transaction binding the contract method 0xda81261d.
//
// Solidity: function confirmSettlement(address requester, bytes32 tradeId, bytes releaseTxId, bytes signature) returns()
func (_EvmBtc *EvmBtcTransactor) ConfirmSettlement(opts *bind.TransactOpts, requester common.Address, tradeId [32]byte, releaseTxId []byte, signature []byte) (*types.Transaction, error) {
	return _EvmBtc.contract.Transact(opts, "confirmSettlement", requester, tradeId, releaseTxId, signature)
}

// ConfirmSettlement is a paid mutator transaction binding the contract method 0xda81261d.
//
// Solidity: function confirmSettlement(address requester, bytes32 tradeId, bytes releaseTxId, bytes signature) returns()
func (_EvmBtc *EvmBtcSession) ConfirmSettlement(requester common.Address, tradeId [32]byte, releaseTxId []byte, signature []byte) (*types.Transaction, error) {
	return _EvmBtc.Contract.ConfirmSettlement(&_EvmBtc.TransactOpts, requester, tradeId, releaseTxId, signature)
}

// ConfirmSettlement is a paid mutator transaction binding the contract method 0xda81261d.
//
// Solidity: function confirmSettlement(address requester, bytes32 tradeId, bytes releaseTxId, bytes signature) returns()
func (_EvmBtc *EvmBtcTransactorSession) ConfirmSettlement(requester common.Address, tradeId [32]byte, releaseTxId []byte, signature []byte) (*types.Transaction, error) {
	return _EvmBtc.Contract.ConfirmSettlement(&_EvmBtc.TransactOpts, requester, tradeId, releaseTxId, signature)
}

// MakePayment is a paid mutator transaction binding the contract method 0xbabcd962.
//
// Solidity: function makePayment(address requester, uint256 indexOfTrade, (bytes32[],uint64,uint64,bytes,bytes) bundle) returns(bytes32 pmmId)
func (_EvmBtc *EvmBtcTransactor) MakePayment(opts *bind.TransactOpts, requester common.Address, indexOfTrade *big.Int, bundle ITypesBundlePayment) (*types.Transaction, error) {
	return _EvmBtc.contract.Transact(opts, "makePayment", requester, indexOfTrade, bundle)
}

// MakePayment is a paid mutator transaction binding the contract method 0xbabcd962.
//
// Solidity: function makePayment(address requester, uint256 indexOfTrade, (bytes32[],uint64,uint64,bytes,bytes) bundle) returns(bytes32 pmmId)
func (_EvmBtc *EvmBtcSession) MakePayment(requester common.Address, indexOfTrade *big.Int, bundle ITypesBundlePayment) (*types.Transaction, error) {
	return _EvmBtc.Contract.MakePayment(&_EvmBtc.TransactOpts, requester, indexOfTrade, bundle)
}

// MakePayment is a paid mutator transaction binding the contract method 0xbabcd962.
//
// Solidity: function makePayment(address requester, uint256 indexOfTrade, (bytes32[],uint64,uint64,bytes,bytes) bundle) returns(bytes32 pmmId)
func (_EvmBtc *EvmBtcTransactorSession) MakePayment(requester common.Address, indexOfTrade *big.Int, bundle ITypesBundlePayment) (*types.Transaction, error) {
	return _EvmBtc.Contract.MakePayment(&_EvmBtc.TransactOpts, requester, indexOfTrade, bundle)
}

// Refund is a paid mutator transaction binding the contract method 0x5e9fc717.
//
// Solidity: function refund(address requester, bytes32 tradeId, bytes refundTxId, bytes signature) returns()
func (_EvmBtc *EvmBtcTransactor) Refund(opts *bind.TransactOpts, requester common.Address, tradeId [32]byte, refundTxId []byte, signature []byte) (*types.Transaction, error) {
	return _EvmBtc.contract.Transact(opts, "refund", requester, tradeId, refundTxId, signature)
}

// Refund is a paid mutator transaction binding the contract method 0x5e9fc717.
//
// Solidity: function refund(address requester, bytes32 tradeId, bytes refundTxId, bytes signature) returns()
func (_EvmBtc *EvmBtcSession) Refund(requester common.Address, tradeId [32]byte, refundTxId []byte, signature []byte) (*types.Transaction, error) {
	return _EvmBtc.Contract.Refund(&_EvmBtc.TransactOpts, requester, tradeId, refundTxId, signature)
}

// Refund is a paid mutator transaction binding the contract method 0x5e9fc717.
//
// Solidity: function refund(address requester, bytes32 tradeId, bytes refundTxId, bytes signature) returns()
func (_EvmBtc *EvmBtcTransactorSession) Refund(requester common.Address, tradeId [32]byte, refundTxId []byte, signature []byte) (*types.Transaction, error) {
	return _EvmBtc.Contract.Refund(&_EvmBtc.TransactOpts, requester, tradeId, refundTxId, signature)
}

// Report is a paid mutator transaction binding the contract method 0x4cb8d6ea.
//
// Solidity: function report(address requester, bytes32 tradeId, bytes msgError, bytes signature) returns()
func (_EvmBtc *EvmBtcTransactor) Report(opts *bind.TransactOpts, requester common.Address, tradeId [32]byte, msgError []byte, signature []byte) (*types.Transaction, error) {
	return _EvmBtc.contract.Transact(opts, "report", requester, tradeId, msgError, signature)
}

// Report is a paid mutator transaction binding the contract method 0x4cb8d6ea.
//
// Solidity: function report(address requester, bytes32 tradeId, bytes msgError, bytes signature) returns()
func (_EvmBtc *EvmBtcSession) Report(requester common.Address, tradeId [32]byte, msgError []byte, signature []byte) (*types.Transaction, error) {
	return _EvmBtc.Contract.Report(&_EvmBtc.TransactOpts, requester, tradeId, msgError, signature)
}

// Report is a paid mutator transaction binding the contract method 0x4cb8d6ea.
//
// Solidity: function report(address requester, bytes32 tradeId, bytes msgError, bytes signature) returns()
func (_EvmBtc *EvmBtcTransactorSession) Report(requester common.Address, tradeId [32]byte, msgError []byte, signature []byte) (*types.Transaction, error) {
	return _EvmBtc.Contract.Report(&_EvmBtc.TransactOpts, requester, tradeId, msgError, signature)
}

// SelectPMM is a paid mutator transaction binding the contract method 0x71987d30.
//
// Solidity: function selectPMM(address requester, bytes32 tradeId, ((uint256,uint64,bytes),(uint256,bytes32,bytes[2],uint64)) info) returns()
func (_EvmBtc *EvmBtcTransactor) SelectPMM(opts *bind.TransactOpts, requester common.Address, tradeId [32]byte, info ITypesPMMSelection) (*types.Transaction, error) {
	return _EvmBtc.contract.Transact(opts, "selectPMM", requester, tradeId, info)
}

// SelectPMM is a paid mutator transaction binding the contract method 0x71987d30.
//
// Solidity: function selectPMM(address requester, bytes32 tradeId, ((uint256,uint64,bytes),(uint256,bytes32,bytes[2],uint64)) info) returns()
func (_EvmBtc *EvmBtcSession) SelectPMM(requester common.Address, tradeId [32]byte, info ITypesPMMSelection) (*types.Transaction, error) {
	return _EvmBtc.Contract.SelectPMM(&_EvmBtc.TransactOpts, requester, tradeId, info)
}

// SelectPMM is a paid mutator transaction binding the contract method 0x71987d30.
//
// Solidity: function selectPMM(address requester, bytes32 tradeId, ((uint256,uint64,bytes),(uint256,bytes32,bytes[2],uint64)) info) returns()
func (_EvmBtc *EvmBtcTransactorSession) SelectPMM(requester common.Address, tradeId [32]byte, info ITypesPMMSelection) (*types.Transaction, error) {
	return _EvmBtc.Contract.SelectPMM(&_EvmBtc.TransactOpts, requester, tradeId, info)
}

// SetMaxAffiliateFeeRate is a paid mutator transaction binding the contract method 0x66ff6244.
//
// Solidity: function setMaxAffiliateFeeRate(uint256 newRate) returns()
func (_EvmBtc *EvmBtcTransactor) SetMaxAffiliateFeeRate(opts *bind.TransactOpts, newRate *big.Int) (*types.Transaction, error) {
	return _EvmBtc.contract.Transact(opts, "setMaxAffiliateFeeRate", newRate)
}

// SetMaxAffiliateFeeRate is a paid mutator transaction binding the contract method 0x66ff6244.
//
// Solidity: function setMaxAffiliateFeeRate(uint256 newRate) returns()
func (_EvmBtc *EvmBtcSession) SetMaxAffiliateFeeRate(newRate *big.Int) (*types.Transaction, error) {
	return _EvmBtc.Contract.SetMaxAffiliateFeeRate(&_EvmBtc.TransactOpts, newRate)
}

// SetMaxAffiliateFeeRate is a paid mutator transaction binding the contract method 0x66ff6244.
//
// Solidity: function setMaxAffiliateFeeRate(uint256 newRate) returns()
func (_EvmBtc *EvmBtcTransactorSession) SetMaxAffiliateFeeRate(newRate *big.Int) (*types.Transaction, error) {
	return _EvmBtc.Contract.SetMaxAffiliateFeeRate(&_EvmBtc.TransactOpts, newRate)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address newRouter) returns()
func (_EvmBtc *EvmBtcTransactor) SetRouter(opts *bind.TransactOpts, newRouter common.Address) (*types.Transaction, error) {
	return _EvmBtc.contract.Transact(opts, "setRouter", newRouter)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address newRouter) returns()
func (_EvmBtc *EvmBtcSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _EvmBtc.Contract.SetRouter(&_EvmBtc.TransactOpts, newRouter)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address newRouter) returns()
func (_EvmBtc *EvmBtcTransactorSession) SetRouter(newRouter common.Address) (*types.Transaction, error) {
	return _EvmBtc.Contract.SetRouter(&_EvmBtc.TransactOpts, newRouter)
}

// SetVaultRegistry is a paid mutator transaction binding the contract method 0xcb91dd7b.
//
// Solidity: function setVaultRegistry(address newRegistry) returns()
func (_EvmBtc *EvmBtcTransactor) SetVaultRegistry(opts *bind.TransactOpts, newRegistry common.Address) (*types.Transaction, error) {
	return _EvmBtc.contract.Transact(opts, "setVaultRegistry", newRegistry)
}

// SetVaultRegistry is a paid mutator transaction binding the contract method 0xcb91dd7b.
//
// Solidity: function setVaultRegistry(address newRegistry) returns()
func (_EvmBtc *EvmBtcSession) SetVaultRegistry(newRegistry common.Address) (*types.Transaction, error) {
	return _EvmBtc.Contract.SetVaultRegistry(&_EvmBtc.TransactOpts, newRegistry)
}

// SetVaultRegistry is a paid mutator transaction binding the contract method 0xcb91dd7b.
//
// Solidity: function setVaultRegistry(address newRegistry) returns()
func (_EvmBtc *EvmBtcTransactorSession) SetVaultRegistry(newRegistry common.Address) (*types.Transaction, error) {
	return _EvmBtc.Contract.SetVaultRegistry(&_EvmBtc.TransactOpts, newRegistry)
}

// SubmitTrade is a paid mutator transaction binding the contract method 0x1c6666f4.
//
// Solidity: function submitTrade(address requester, bytes32 tradeId, (uint256,(uint256,bytes[3],bytes[3]),(bytes[5],address,uint64)) tradeData, (uint256,string,bytes) affiliateInfo, (bytes32,bytes,bytes[])[] settlementPresigns, (bytes,bytes[]) refundPresign) returns()
func (_EvmBtc *EvmBtcTransactor) SubmitTrade(opts *bind.TransactOpts, requester common.Address, tradeId [32]byte, tradeData ITypesTradeData, affiliateInfo ITypesAffiliate, settlementPresigns []ITypesSettlementPresign, refundPresign ITypesRefundPresign) (*types.Transaction, error) {
	return _EvmBtc.contract.Transact(opts, "submitTrade", requester, tradeId, tradeData, affiliateInfo, settlementPresigns, refundPresign)
}

// SubmitTrade is a paid mutator transaction binding the contract method 0x1c6666f4.
//
// Solidity: function submitTrade(address requester, bytes32 tradeId, (uint256,(uint256,bytes[3],bytes[3]),(bytes[5],address,uint64)) tradeData, (uint256,string,bytes) affiliateInfo, (bytes32,bytes,bytes[])[] settlementPresigns, (bytes,bytes[]) refundPresign) returns()
func (_EvmBtc *EvmBtcSession) SubmitTrade(requester common.Address, tradeId [32]byte, tradeData ITypesTradeData, affiliateInfo ITypesAffiliate, settlementPresigns []ITypesSettlementPresign, refundPresign ITypesRefundPresign) (*types.Transaction, error) {
	return _EvmBtc.Contract.SubmitTrade(&_EvmBtc.TransactOpts, requester, tradeId, tradeData, affiliateInfo, settlementPresigns, refundPresign)
}

// SubmitTrade is a paid mutator transaction binding the contract method 0x1c6666f4.
//
// Solidity: function submitTrade(address requester, bytes32 tradeId, (uint256,(uint256,bytes[3],bytes[3]),(bytes[5],address,uint64)) tradeData, (uint256,string,bytes) affiliateInfo, (bytes32,bytes,bytes[])[] settlementPresigns, (bytes,bytes[]) refundPresign) returns()
func (_EvmBtc *EvmBtcTransactorSession) SubmitTrade(requester common.Address, tradeId [32]byte, tradeData ITypesTradeData, affiliateInfo ITypesAffiliate, settlementPresigns []ITypesSettlementPresign, refundPresign ITypesRefundPresign) (*types.Transaction, error) {
	return _EvmBtc.Contract.SubmitTrade(&_EvmBtc.TransactOpts, requester, tradeId, tradeData, affiliateInfo, settlementPresigns, refundPresign)
}

// EvmBtcDepositConfirmedIterator is returned from FilterDepositConfirmed and is used to iterate over the raw logs and unpacked data for DepositConfirmed events raised by the EvmBtc contract.
type EvmBtcDepositConfirmedIterator struct {
	Event *EvmBtcDepositConfirmed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EvmBtcDepositConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EvmBtcDepositConfirmed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EvmBtcDepositConfirmed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EvmBtcDepositConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EvmBtcDepositConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EvmBtcDepositConfirmed represents a DepositConfirmed event raised by the EvmBtc contract.
type EvmBtcDepositConfirmed struct {
	Forwarder    common.Address
	Mpc          common.Address
	TradeId      [32]byte
	TotalFeeRate *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDepositConfirmed is a free log retrieval operation binding the contract event 0xd817456a528dde0c57bcc60be2b10c106f3a9108e8c0cfa84f6b51839334127f.
//
// Solidity: event DepositConfirmed(address indexed forwarder, address indexed mpc, bytes32 indexed tradeId, uint256 totalFeeRate)
func (_EvmBtc *EvmBtcFilterer) FilterDepositConfirmed(opts *bind.FilterOpts, forwarder []common.Address, mpc []common.Address, tradeId [][32]byte) (*EvmBtcDepositConfirmedIterator, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}
	var mpcRule []interface{}
	for _, mpcItem := range mpc {
		mpcRule = append(mpcRule, mpcItem)
	}
	var tradeIdRule []interface{}
	for _, tradeIdItem := range tradeId {
		tradeIdRule = append(tradeIdRule, tradeIdItem)
	}

	logs, sub, err := _EvmBtc.contract.FilterLogs(opts, "DepositConfirmed", forwarderRule, mpcRule, tradeIdRule)
	if err != nil {
		return nil, err
	}
	return &EvmBtcDepositConfirmedIterator{contract: _EvmBtc.contract, event: "DepositConfirmed", logs: logs, sub: sub}, nil
}

// WatchDepositConfirmed is a free log subscription operation binding the contract event 0xd817456a528dde0c57bcc60be2b10c106f3a9108e8c0cfa84f6b51839334127f.
//
// Solidity: event DepositConfirmed(address indexed forwarder, address indexed mpc, bytes32 indexed tradeId, uint256 totalFeeRate)
func (_EvmBtc *EvmBtcFilterer) WatchDepositConfirmed(opts *bind.WatchOpts, sink chan<- *EvmBtcDepositConfirmed, forwarder []common.Address, mpc []common.Address, tradeId [][32]byte) (event.Subscription, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}
	var mpcRule []interface{}
	for _, mpcItem := range mpc {
		mpcRule = append(mpcRule, mpcItem)
	}
	var tradeIdRule []interface{}
	for _, tradeIdItem := range tradeId {
		tradeIdRule = append(tradeIdRule, tradeIdItem)
	}

	logs, sub, err := _EvmBtc.contract.WatchLogs(opts, "DepositConfirmed", forwarderRule, mpcRule, tradeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EvmBtcDepositConfirmed)
				if err := _EvmBtc.contract.UnpackLog(event, "DepositConfirmed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositConfirmed is a log parse operation binding the contract event 0xd817456a528dde0c57bcc60be2b10c106f3a9108e8c0cfa84f6b51839334127f.
//
// Solidity: event DepositConfirmed(address indexed forwarder, address indexed mpc, bytes32 indexed tradeId, uint256 totalFeeRate)
func (_EvmBtc *EvmBtcFilterer) ParseDepositConfirmed(log types.Log) (*EvmBtcDepositConfirmed, error) {
	event := new(EvmBtcDepositConfirmed)
	if err := _EvmBtc.contract.UnpackLog(event, "DepositConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EvmBtcFailureReportedIterator is returned from FilterFailureReported and is used to iterate over the raw logs and unpacked data for FailureReported events raised by the EvmBtc contract.
type EvmBtcFailureReportedIterator struct {
	Event *EvmBtcFailureReported // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EvmBtcFailureReportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EvmBtcFailureReported)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EvmBtcFailureReported)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EvmBtcFailureReportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EvmBtcFailureReportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EvmBtcFailureReported represents a FailureReported event raised by the EvmBtc contract.
type EvmBtcFailureReported struct {
	Forwarder     common.Address
	Mpc           common.Address
	TradeId       [32]byte
	Stage         *big.Int
	ReferenceInfo []byte
	MsgError      []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFailureReported is a free log retrieval operation binding the contract event 0x2301b4f60027232b5409ef0a59170d7983a8128643ec377fcf9fbf8df35581e6.
//
// Solidity: event FailureReported(address indexed forwarder, address indexed mpc, bytes32 indexed tradeId, uint256 stage, bytes referenceInfo, bytes msgError)
func (_EvmBtc *EvmBtcFilterer) FilterFailureReported(opts *bind.FilterOpts, forwarder []common.Address, mpc []common.Address, tradeId [][32]byte) (*EvmBtcFailureReportedIterator, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}
	var mpcRule []interface{}
	for _, mpcItem := range mpc {
		mpcRule = append(mpcRule, mpcItem)
	}
	var tradeIdRule []interface{}
	for _, tradeIdItem := range tradeId {
		tradeIdRule = append(tradeIdRule, tradeIdItem)
	}

	logs, sub, err := _EvmBtc.contract.FilterLogs(opts, "FailureReported", forwarderRule, mpcRule, tradeIdRule)
	if err != nil {
		return nil, err
	}
	return &EvmBtcFailureReportedIterator{contract: _EvmBtc.contract, event: "FailureReported", logs: logs, sub: sub}, nil
}

// WatchFailureReported is a free log subscription operation binding the contract event 0x2301b4f60027232b5409ef0a59170d7983a8128643ec377fcf9fbf8df35581e6.
//
// Solidity: event FailureReported(address indexed forwarder, address indexed mpc, bytes32 indexed tradeId, uint256 stage, bytes referenceInfo, bytes msgError)
func (_EvmBtc *EvmBtcFilterer) WatchFailureReported(opts *bind.WatchOpts, sink chan<- *EvmBtcFailureReported, forwarder []common.Address, mpc []common.Address, tradeId [][32]byte) (event.Subscription, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}
	var mpcRule []interface{}
	for _, mpcItem := range mpc {
		mpcRule = append(mpcRule, mpcItem)
	}
	var tradeIdRule []interface{}
	for _, tradeIdItem := range tradeId {
		tradeIdRule = append(tradeIdRule, tradeIdItem)
	}

	logs, sub, err := _EvmBtc.contract.WatchLogs(opts, "FailureReported", forwarderRule, mpcRule, tradeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EvmBtcFailureReported)
				if err := _EvmBtc.contract.UnpackLog(event, "FailureReported", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFailureReported is a log parse operation binding the contract event 0x2301b4f60027232b5409ef0a59170d7983a8128643ec377fcf9fbf8df35581e6.
//
// Solidity: event FailureReported(address indexed forwarder, address indexed mpc, bytes32 indexed tradeId, uint256 stage, bytes referenceInfo, bytes msgError)
func (_EvmBtc *EvmBtcFilterer) ParseFailureReported(log types.Log) (*EvmBtcFailureReported, error) {
	event := new(EvmBtcFailureReported)
	if err := _EvmBtc.contract.UnpackLog(event, "FailureReported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EvmBtcMadePaymentIterator is returned from FilterMadePayment and is used to iterate over the raw logs and unpacked data for MadePayment events raised by the EvmBtc contract.
type EvmBtcMadePaymentIterator struct {
	Event *EvmBtcMadePayment // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EvmBtcMadePaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EvmBtcMadePayment)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EvmBtcMadePayment)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EvmBtcMadePaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EvmBtcMadePaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EvmBtcMadePayment represents a MadePayment event raised by the EvmBtc contract.
type EvmBtcMadePayment struct {
	Forwarder   common.Address
	Operator    common.Address
	TradeId     [32]byte
	Network     []byte
	PaymentTxId []byte
	Bundle      [][32]byte
	StartIdx    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMadePayment is a free log retrieval operation binding the contract event 0xf42a9bc88ea9245bb0cd847b4badcc6facb969dad3a514251a600e875927ece9.
//
// Solidity: event MadePayment(address indexed forwarder, address indexed operator, bytes32 indexed tradeId, bytes network, bytes paymentTxId, bytes32[] bundle, uint256 startIdx)
func (_EvmBtc *EvmBtcFilterer) FilterMadePayment(opts *bind.FilterOpts, forwarder []common.Address, operator []common.Address, tradeId [][32]byte) (*EvmBtcMadePaymentIterator, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tradeIdRule []interface{}
	for _, tradeIdItem := range tradeId {
		tradeIdRule = append(tradeIdRule, tradeIdItem)
	}

	logs, sub, err := _EvmBtc.contract.FilterLogs(opts, "MadePayment", forwarderRule, operatorRule, tradeIdRule)
	if err != nil {
		return nil, err
	}
	return &EvmBtcMadePaymentIterator{contract: _EvmBtc.contract, event: "MadePayment", logs: logs, sub: sub}, nil
}

// WatchMadePayment is a free log subscription operation binding the contract event 0xf42a9bc88ea9245bb0cd847b4badcc6facb969dad3a514251a600e875927ece9.
//
// Solidity: event MadePayment(address indexed forwarder, address indexed operator, bytes32 indexed tradeId, bytes network, bytes paymentTxId, bytes32[] bundle, uint256 startIdx)
func (_EvmBtc *EvmBtcFilterer) WatchMadePayment(opts *bind.WatchOpts, sink chan<- *EvmBtcMadePayment, forwarder []common.Address, operator []common.Address, tradeId [][32]byte) (event.Subscription, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var tradeIdRule []interface{}
	for _, tradeIdItem := range tradeId {
		tradeIdRule = append(tradeIdRule, tradeIdItem)
	}

	logs, sub, err := _EvmBtc.contract.WatchLogs(opts, "MadePayment", forwarderRule, operatorRule, tradeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EvmBtcMadePayment)
				if err := _EvmBtc.contract.UnpackLog(event, "MadePayment", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMadePayment is a log parse operation binding the contract event 0xf42a9bc88ea9245bb0cd847b4badcc6facb969dad3a514251a600e875927ece9.
//
// Solidity: event MadePayment(address indexed forwarder, address indexed operator, bytes32 indexed tradeId, bytes network, bytes paymentTxId, bytes32[] bundle, uint256 startIdx)
func (_EvmBtc *EvmBtcFilterer) ParseMadePayment(log types.Log) (*EvmBtcMadePayment, error) {
	event := new(EvmBtcMadePayment)
	if err := _EvmBtc.contract.UnpackLog(event, "MadePayment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EvmBtcPaymentConfirmedIterator is returned from FilterPaymentConfirmed and is used to iterate over the raw logs and unpacked data for PaymentConfirmed events raised by the EvmBtc contract.
type EvmBtcPaymentConfirmedIterator struct {
	Event *EvmBtcPaymentConfirmed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EvmBtcPaymentConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EvmBtcPaymentConfirmed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EvmBtcPaymentConfirmed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EvmBtcPaymentConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EvmBtcPaymentConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EvmBtcPaymentConfirmed represents a PaymentConfirmed event raised by the EvmBtc contract.
type EvmBtcPaymentConfirmed struct {
	Forwarder   common.Address
	Mpc         common.Address
	TradeId     [32]byte
	PaymentTxId []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPaymentConfirmed is a free log retrieval operation binding the contract event 0xba0272c2d3764773d2b2074085e81b83fee550ad6a21a1a72ab08dfbef22af6f.
//
// Solidity: event PaymentConfirmed(address indexed forwarder, address indexed mpc, bytes32 indexed tradeId, bytes paymentTxId)
func (_EvmBtc *EvmBtcFilterer) FilterPaymentConfirmed(opts *bind.FilterOpts, forwarder []common.Address, mpc []common.Address, tradeId [][32]byte) (*EvmBtcPaymentConfirmedIterator, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}
	var mpcRule []interface{}
	for _, mpcItem := range mpc {
		mpcRule = append(mpcRule, mpcItem)
	}
	var tradeIdRule []interface{}
	for _, tradeIdItem := range tradeId {
		tradeIdRule = append(tradeIdRule, tradeIdItem)
	}

	logs, sub, err := _EvmBtc.contract.FilterLogs(opts, "PaymentConfirmed", forwarderRule, mpcRule, tradeIdRule)
	if err != nil {
		return nil, err
	}
	return &EvmBtcPaymentConfirmedIterator{contract: _EvmBtc.contract, event: "PaymentConfirmed", logs: logs, sub: sub}, nil
}

// WatchPaymentConfirmed is a free log subscription operation binding the contract event 0xba0272c2d3764773d2b2074085e81b83fee550ad6a21a1a72ab08dfbef22af6f.
//
// Solidity: event PaymentConfirmed(address indexed forwarder, address indexed mpc, bytes32 indexed tradeId, bytes paymentTxId)
func (_EvmBtc *EvmBtcFilterer) WatchPaymentConfirmed(opts *bind.WatchOpts, sink chan<- *EvmBtcPaymentConfirmed, forwarder []common.Address, mpc []common.Address, tradeId [][32]byte) (event.Subscription, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}
	var mpcRule []interface{}
	for _, mpcItem := range mpc {
		mpcRule = append(mpcRule, mpcItem)
	}
	var tradeIdRule []interface{}
	for _, tradeIdItem := range tradeId {
		tradeIdRule = append(tradeIdRule, tradeIdItem)
	}

	logs, sub, err := _EvmBtc.contract.WatchLogs(opts, "PaymentConfirmed", forwarderRule, mpcRule, tradeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EvmBtcPaymentConfirmed)
				if err := _EvmBtc.contract.UnpackLog(event, "PaymentConfirmed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaymentConfirmed is a log parse operation binding the contract event 0xba0272c2d3764773d2b2074085e81b83fee550ad6a21a1a72ab08dfbef22af6f.
//
// Solidity: event PaymentConfirmed(address indexed forwarder, address indexed mpc, bytes32 indexed tradeId, bytes paymentTxId)
func (_EvmBtc *EvmBtcFilterer) ParsePaymentConfirmed(log types.Log) (*EvmBtcPaymentConfirmed, error) {
	event := new(EvmBtcPaymentConfirmed)
	if err := _EvmBtc.contract.UnpackLog(event, "PaymentConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EvmBtcRefundedIterator is returned from FilterRefunded and is used to iterate over the raw logs and unpacked data for Refunded events raised by the EvmBtc contract.
type EvmBtcRefundedIterator struct {
	Event *EvmBtcRefunded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EvmBtcRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EvmBtcRefunded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EvmBtcRefunded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EvmBtcRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EvmBtcRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EvmBtcRefunded represents a Refunded event raised by the EvmBtc contract.
type EvmBtcRefunded struct {
	Forwarder  common.Address
	Mpc        common.Address
	TradeId    [32]byte
	RefundTxId []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRefunded is a free log retrieval operation binding the contract event 0x243fca6277ae6674a807f2c94ee5af8f7a00976fffc1100b24fc727358c03521.
//
// Solidity: event Refunded(address indexed forwarder, address indexed mpc, bytes32 indexed tradeId, bytes refundTxId)
func (_EvmBtc *EvmBtcFilterer) FilterRefunded(opts *bind.FilterOpts, forwarder []common.Address, mpc []common.Address, tradeId [][32]byte) (*EvmBtcRefundedIterator, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}
	var mpcRule []interface{}
	for _, mpcItem := range mpc {
		mpcRule = append(mpcRule, mpcItem)
	}
	var tradeIdRule []interface{}
	for _, tradeIdItem := range tradeId {
		tradeIdRule = append(tradeIdRule, tradeIdItem)
	}

	logs, sub, err := _EvmBtc.contract.FilterLogs(opts, "Refunded", forwarderRule, mpcRule, tradeIdRule)
	if err != nil {
		return nil, err
	}
	return &EvmBtcRefundedIterator{contract: _EvmBtc.contract, event: "Refunded", logs: logs, sub: sub}, nil
}

// WatchRefunded is a free log subscription operation binding the contract event 0x243fca6277ae6674a807f2c94ee5af8f7a00976fffc1100b24fc727358c03521.
//
// Solidity: event Refunded(address indexed forwarder, address indexed mpc, bytes32 indexed tradeId, bytes refundTxId)
func (_EvmBtc *EvmBtcFilterer) WatchRefunded(opts *bind.WatchOpts, sink chan<- *EvmBtcRefunded, forwarder []common.Address, mpc []common.Address, tradeId [][32]byte) (event.Subscription, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}
	var mpcRule []interface{}
	for _, mpcItem := range mpc {
		mpcRule = append(mpcRule, mpcItem)
	}
	var tradeIdRule []interface{}
	for _, tradeIdItem := range tradeId {
		tradeIdRule = append(tradeIdRule, tradeIdItem)
	}

	logs, sub, err := _EvmBtc.contract.WatchLogs(opts, "Refunded", forwarderRule, mpcRule, tradeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EvmBtcRefunded)
				if err := _EvmBtc.contract.UnpackLog(event, "Refunded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRefunded is a log parse operation binding the contract event 0x243fca6277ae6674a807f2c94ee5af8f7a00976fffc1100b24fc727358c03521.
//
// Solidity: event Refunded(address indexed forwarder, address indexed mpc, bytes32 indexed tradeId, bytes refundTxId)
func (_EvmBtc *EvmBtcFilterer) ParseRefunded(log types.Log) (*EvmBtcRefunded, error) {
	event := new(EvmBtcRefunded)
	if err := _EvmBtc.contract.UnpackLog(event, "Refunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EvmBtcSelectedPMMIterator is returned from FilterSelectedPMM and is used to iterate over the raw logs and unpacked data for SelectedPMM events raised by the EvmBtc contract.
type EvmBtcSelectedPMMIterator struct {
	Event *EvmBtcSelectedPMM // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EvmBtcSelectedPMMIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EvmBtcSelectedPMM)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EvmBtcSelectedPMM)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EvmBtcSelectedPMMIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EvmBtcSelectedPMMIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EvmBtcSelectedPMM represents a SelectedPMM event raised by the EvmBtc contract.
type EvmBtcSelectedPMM struct {
	Forwarder     common.Address
	Solver        common.Address
	TradeId       [32]byte
	SelectedPMMId [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSelectedPMM is a free log retrieval operation binding the contract event 0x1b14e32c3455ce347f72e29e08fd17fee05ef7b5743137d370bd9bc344fecf13.
//
// Solidity: event SelectedPMM(address indexed forwarder, address indexed solver, bytes32 indexed tradeId, bytes32 selectedPMMId)
func (_EvmBtc *EvmBtcFilterer) FilterSelectedPMM(opts *bind.FilterOpts, forwarder []common.Address, solver []common.Address, tradeId [][32]byte) (*EvmBtcSelectedPMMIterator, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}
	var solverRule []interface{}
	for _, solverItem := range solver {
		solverRule = append(solverRule, solverItem)
	}
	var tradeIdRule []interface{}
	for _, tradeIdItem := range tradeId {
		tradeIdRule = append(tradeIdRule, tradeIdItem)
	}

	logs, sub, err := _EvmBtc.contract.FilterLogs(opts, "SelectedPMM", forwarderRule, solverRule, tradeIdRule)
	if err != nil {
		return nil, err
	}
	return &EvmBtcSelectedPMMIterator{contract: _EvmBtc.contract, event: "SelectedPMM", logs: logs, sub: sub}, nil
}

// WatchSelectedPMM is a free log subscription operation binding the contract event 0x1b14e32c3455ce347f72e29e08fd17fee05ef7b5743137d370bd9bc344fecf13.
//
// Solidity: event SelectedPMM(address indexed forwarder, address indexed solver, bytes32 indexed tradeId, bytes32 selectedPMMId)
func (_EvmBtc *EvmBtcFilterer) WatchSelectedPMM(opts *bind.WatchOpts, sink chan<- *EvmBtcSelectedPMM, forwarder []common.Address, solver []common.Address, tradeId [][32]byte) (event.Subscription, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}
	var solverRule []interface{}
	for _, solverItem := range solver {
		solverRule = append(solverRule, solverItem)
	}
	var tradeIdRule []interface{}
	for _, tradeIdItem := range tradeId {
		tradeIdRule = append(tradeIdRule, tradeIdItem)
	}

	logs, sub, err := _EvmBtc.contract.WatchLogs(opts, "SelectedPMM", forwarderRule, solverRule, tradeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EvmBtcSelectedPMM)
				if err := _EvmBtc.contract.UnpackLog(event, "SelectedPMM", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSelectedPMM is a log parse operation binding the contract event 0x1b14e32c3455ce347f72e29e08fd17fee05ef7b5743137d370bd9bc344fecf13.
//
// Solidity: event SelectedPMM(address indexed forwarder, address indexed solver, bytes32 indexed tradeId, bytes32 selectedPMMId)
func (_EvmBtc *EvmBtcFilterer) ParseSelectedPMM(log types.Log) (*EvmBtcSelectedPMM, error) {
	event := new(EvmBtcSelectedPMM)
	if err := _EvmBtc.contract.UnpackLog(event, "SelectedPMM", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EvmBtcSettlementConfirmedIterator is returned from FilterSettlementConfirmed and is used to iterate over the raw logs and unpacked data for SettlementConfirmed events raised by the EvmBtc contract.
type EvmBtcSettlementConfirmedIterator struct {
	Event *EvmBtcSettlementConfirmed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EvmBtcSettlementConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EvmBtcSettlementConfirmed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EvmBtcSettlementConfirmed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EvmBtcSettlementConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EvmBtcSettlementConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EvmBtcSettlementConfirmed represents a SettlementConfirmed event raised by the EvmBtc contract.
type EvmBtcSettlementConfirmed struct {
	Forwarder   common.Address
	Mpc         common.Address
	TradeId     [32]byte
	ReleaseTxId []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSettlementConfirmed is a free log retrieval operation binding the contract event 0x1285aa4b99e3eaad5567e85fe29f0938417a42e3cae2b260122fe0565de9cb95.
//
// Solidity: event SettlementConfirmed(address indexed forwarder, address indexed mpc, bytes32 indexed tradeId, bytes releaseTxId)
func (_EvmBtc *EvmBtcFilterer) FilterSettlementConfirmed(opts *bind.FilterOpts, forwarder []common.Address, mpc []common.Address, tradeId [][32]byte) (*EvmBtcSettlementConfirmedIterator, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}
	var mpcRule []interface{}
	for _, mpcItem := range mpc {
		mpcRule = append(mpcRule, mpcItem)
	}
	var tradeIdRule []interface{}
	for _, tradeIdItem := range tradeId {
		tradeIdRule = append(tradeIdRule, tradeIdItem)
	}

	logs, sub, err := _EvmBtc.contract.FilterLogs(opts, "SettlementConfirmed", forwarderRule, mpcRule, tradeIdRule)
	if err != nil {
		return nil, err
	}
	return &EvmBtcSettlementConfirmedIterator{contract: _EvmBtc.contract, event: "SettlementConfirmed", logs: logs, sub: sub}, nil
}

// WatchSettlementConfirmed is a free log subscription operation binding the contract event 0x1285aa4b99e3eaad5567e85fe29f0938417a42e3cae2b260122fe0565de9cb95.
//
// Solidity: event SettlementConfirmed(address indexed forwarder, address indexed mpc, bytes32 indexed tradeId, bytes releaseTxId)
func (_EvmBtc *EvmBtcFilterer) WatchSettlementConfirmed(opts *bind.WatchOpts, sink chan<- *EvmBtcSettlementConfirmed, forwarder []common.Address, mpc []common.Address, tradeId [][32]byte) (event.Subscription, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}
	var mpcRule []interface{}
	for _, mpcItem := range mpc {
		mpcRule = append(mpcRule, mpcItem)
	}
	var tradeIdRule []interface{}
	for _, tradeIdItem := range tradeId {
		tradeIdRule = append(tradeIdRule, tradeIdItem)
	}

	logs, sub, err := _EvmBtc.contract.WatchLogs(opts, "SettlementConfirmed", forwarderRule, mpcRule, tradeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EvmBtcSettlementConfirmed)
				if err := _EvmBtc.contract.UnpackLog(event, "SettlementConfirmed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSettlementConfirmed is a log parse operation binding the contract event 0x1285aa4b99e3eaad5567e85fe29f0938417a42e3cae2b260122fe0565de9cb95.
//
// Solidity: event SettlementConfirmed(address indexed forwarder, address indexed mpc, bytes32 indexed tradeId, bytes releaseTxId)
func (_EvmBtc *EvmBtcFilterer) ParseSettlementConfirmed(log types.Log) (*EvmBtcSettlementConfirmed, error) {
	event := new(EvmBtcSettlementConfirmed)
	if err := _EvmBtc.contract.UnpackLog(event, "SettlementConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EvmBtcTradeInfoSubmittedIterator is returned from FilterTradeInfoSubmitted and is used to iterate over the raw logs and unpacked data for TradeInfoSubmitted events raised by the EvmBtc contract.
type EvmBtcTradeInfoSubmittedIterator struct {
	Event *EvmBtcTradeInfoSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EvmBtcTradeInfoSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EvmBtcTradeInfoSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EvmBtcTradeInfoSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EvmBtcTradeInfoSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EvmBtcTradeInfoSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EvmBtcTradeInfoSubmitted represents a TradeInfoSubmitted event raised by the EvmBtc contract.
type EvmBtcTradeInfoSubmitted struct {
	Forwarder   common.Address
	Solver      common.Address
	TradeId     [32]byte
	Network     []byte
	DepositTxId []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTradeInfoSubmitted is a free log retrieval operation binding the contract event 0x5e542c7c404fe0d7491036af0741943f863337ae466beda73ac86e94adb86557.
//
// Solidity: event TradeInfoSubmitted(address indexed forwarder, address indexed solver, bytes32 indexed tradeId, bytes network, bytes depositTxId)
func (_EvmBtc *EvmBtcFilterer) FilterTradeInfoSubmitted(opts *bind.FilterOpts, forwarder []common.Address, solver []common.Address, tradeId [][32]byte) (*EvmBtcTradeInfoSubmittedIterator, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}
	var solverRule []interface{}
	for _, solverItem := range solver {
		solverRule = append(solverRule, solverItem)
	}
	var tradeIdRule []interface{}
	for _, tradeIdItem := range tradeId {
		tradeIdRule = append(tradeIdRule, tradeIdItem)
	}

	logs, sub, err := _EvmBtc.contract.FilterLogs(opts, "TradeInfoSubmitted", forwarderRule, solverRule, tradeIdRule)
	if err != nil {
		return nil, err
	}
	return &EvmBtcTradeInfoSubmittedIterator{contract: _EvmBtc.contract, event: "TradeInfoSubmitted", logs: logs, sub: sub}, nil
}

// WatchTradeInfoSubmitted is a free log subscription operation binding the contract event 0x5e542c7c404fe0d7491036af0741943f863337ae466beda73ac86e94adb86557.
//
// Solidity: event TradeInfoSubmitted(address indexed forwarder, address indexed solver, bytes32 indexed tradeId, bytes network, bytes depositTxId)
func (_EvmBtc *EvmBtcFilterer) WatchTradeInfoSubmitted(opts *bind.WatchOpts, sink chan<- *EvmBtcTradeInfoSubmitted, forwarder []common.Address, solver []common.Address, tradeId [][32]byte) (event.Subscription, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}
	var solverRule []interface{}
	for _, solverItem := range solver {
		solverRule = append(solverRule, solverItem)
	}
	var tradeIdRule []interface{}
	for _, tradeIdItem := range tradeId {
		tradeIdRule = append(tradeIdRule, tradeIdItem)
	}

	logs, sub, err := _EvmBtc.contract.WatchLogs(opts, "TradeInfoSubmitted", forwarderRule, solverRule, tradeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EvmBtcTradeInfoSubmitted)
				if err := _EvmBtc.contract.UnpackLog(event, "TradeInfoSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTradeInfoSubmitted is a log parse operation binding the contract event 0x5e542c7c404fe0d7491036af0741943f863337ae466beda73ac86e94adb86557.
//
// Solidity: event TradeInfoSubmitted(address indexed forwarder, address indexed solver, bytes32 indexed tradeId, bytes network, bytes depositTxId)
func (_EvmBtc *EvmBtcFilterer) ParseTradeInfoSubmitted(log types.Log) (*EvmBtcTradeInfoSubmitted, error) {
	event := new(EvmBtcTradeInfoSubmitted)
	if err := _EvmBtc.contract.UnpackLog(event, "TradeInfoSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
