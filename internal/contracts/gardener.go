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
)

// GardenerMetaData contains all meta data concerning the Gardener contract.
var GardenerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_govToken\",\"internalType\":\"contractJewelToken\",\"type\":\"address\"},{\"type\":\"address\",\"internalType\":\"address\",\"name\":\"_devaddr\"},{\"name\":\"_liquidityaddr\",\"internalType\":\"address\",\"type\":\"address\"},{\"name\":\"_comfundaddr\",\"internalType\":\"address\",\"type\":\"address\"},{\"name\":\"_founderaddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardPerBlock\",\"type\":\"uint256\"},{\"type\":\"uint256\",\"internalType\":\"uint256\",\"name\":\"_startBlock\"},{\"internalType\":\"uint256\",\"name\":\"_halvingAfterBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_userDepFee\",\"type\":\"uint256\"},{\"name\":\"_devDepFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_rewardMultiplier\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"type\":\"uint256[]\",\"name\":\"_blockDeltaStartStage\",\"internalType\":\"uint256[]\"},{\"type\":\"uint256[]\",\"internalType\":\"uint256[]\",\"name\":\"_blockDeltaEndStage\"},{\"type\":\"uint256[]\",\"name\":\"_userFeeStage\",\"internalType\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_devFeeStage\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\"},{\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"uint256\",\"indexed\":true,\"name\":\"pid\",\"internalType\":\"uint256\"},{\"indexed\":false,\"type\":\"uint256\",\"internalType\":\"uint256\",\"name\":\"amount\"}],\"name\":\"Deposit\",\"type\":\"event\",\"anonymous\":false},{\"anonymous\":false,\"name\":\"EmergencyWithdraw\",\"type\":\"event\",\"inputs\":[{\"name\":\"user\",\"indexed\":true,\"internalType\":\"address\",\"type\":\"address\"},{\"name\":\"pid\",\"type\":\"uint256\",\"internalType\":\"uint256\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"amount\",\"indexed\":false,\"internalType\":\"uint256\"}]},{\"inputs\":[{\"internalType\":\"address\",\"type\":\"address\",\"indexed\":true,\"name\":\"previousOwner\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\",\"anonymous\":false},{\"anonymous\":false,\"type\":\"event\",\"name\":\"SendGovernanceTokenReward\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"pid\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false},{\"indexed\":false,\"type\":\"uint256\",\"internalType\":\"uint256\",\"name\":\"lockAmount\"}]},{\"inputs\":[{\"name\":\"user\",\"internalType\":\"address\",\"indexed\":true,\"type\":\"address\"},{\"indexed\":true,\"name\":\"pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"internalType\":\"uint256\",\"type\":\"uint256\"}],\"anonymous\":false,\"name\":\"Withdraw\",\"type\":\"event\"},{\"name\":\"FINISH_BONUS_AT_BLOCK\",\"stateMutability\":\"view\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"constant\":true,\"signature\":\"0x980c2a98\"},{\"type\":\"function\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"HALVING_AT_BLOCK\",\"stateMutability\":\"view\",\"inputs\":[{\"type\":\"uint256\",\"internalType\":\"uint256\",\"name\":\"\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"name\":\"\",\"internalType\":\"uint256\",\"type\":\"uint256\"}],\"name\":\"PERCENT_FOR_COM\",\"inputs\":[],\"constant\":true,\"signature\":\"0xa02306f9\"},{\"stateMutability\":\"view\",\"inputs\":[],\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"type\":\"function\",\"name\":\"PERCENT_FOR_DEV\",\"constant\":true,\"signature\":\"0xed9bdeda\"},{\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"type\":\"function\",\"inputs\":[],\"stateMutability\":\"view\",\"name\":\"PERCENT_FOR_FOUNDERS\",\"constant\":true,\"signature\":\"0xc6929e53\"},{\"name\":\"PERCENT_FOR_LP\",\"outputs\":[{\"type\":\"uint256\",\"internalType\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"type\":\"function\",\"stateMutability\":\"view\",\"constant\":true,\"signature\":\"0x0a67d518\"},{\"name\":\"PERCENT_LOCK_BONUS_REWARD\",\"inputs\":[{\"name\":\"\",\"internalType\":\"uint256\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"type\":\"function\",\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REWARD_MULTIPLIER\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\",\"inputs\":[{\"name\":\"\",\"internalType\":\"uint256\",\"type\":\"uint256\"}]},{\"name\":\"REWARD_PER_BLOCK\",\"inputs\":[],\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"internalType\":\"uint256\",\"name\":\"\"}],\"constant\":true,\"signature\":\"0x975532dc\"},{\"stateMutability\":\"view\",\"name\":\"START_BLOCK\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"constant\":true,\"signature\":\"0x39b3e826\"},{\"name\":\"addAuthorized\",\"inputs\":[{\"internalType\":\"address\",\"name\":\"_toAdd\",\"type\":\"address\"}],\"outputs\":[],\"type\":\"function\",\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"authorized\",\"stateMutability\":\"view\",\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"outputs\":[{\"type\":\"bool\",\"internalType\":\"bool\",\"name\":\"\"}]},{\"inputs\":[{\"type\":\"uint256\",\"internalType\":\"uint256\",\"name\":\"\"}],\"name\":\"blockDeltaEndStage\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"type\":\"function\",\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blockDeltaStartStage\",\"stateMutability\":\"view\",\"inputs\":[{\"type\":\"uint256\",\"internalType\":\"uint256\",\"name\":\"\"}],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"stateMutability\":\"view\",\"name\":\"comfundaddr\",\"constant\":true,\"signature\":\"0x3c9d9267\"},{\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"devDepFee\",\"type\":\"function\",\"inputs\":[],\"constant\":true,\"signature\":\"0xc56a10ff\"},{\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"name\":\"devFeeStage\",\"stateMutability\":\"view\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"stateMutability\":\"view\",\"name\":\"devaddr\",\"inputs\":[],\"outputs\":[{\"type\":\"address\",\"internalType\":\"address\",\"name\":\"\"}],\"type\":\"function\",\"constant\":true,\"signature\":\"0xd49e77cd\"},{\"inputs\":[],\"stateMutability\":\"view\",\"name\":\"founderaddr\",\"type\":\"function\",\"outputs\":[{\"internalType\":\"address\",\"type\":\"address\",\"name\":\"\"}],\"constant\":true,\"signature\":\"0xec12173d\"},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"internalType\":\"contractJewelToken\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"govToken\",\"inputs\":[],\"constant\":true,\"signature\":\"0x05268cff\"},{\"type\":\"function\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"liquidityaddr\",\"inputs\":[],\"stateMutability\":\"view\",\"constant\":true,\"signature\":\"0x22a376b0\"},{\"name\":\"owner\",\"type\":\"function\",\"inputs\":[],\"stateMutability\":\"view\",\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"constant\":true,\"signature\":\"0x8da5cb5b\"},{\"type\":\"function\",\"name\":\"poolExistence\",\"stateMutability\":\"view\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"inputs\":[{\"name\":\"\",\"internalType\":\"contractIERC20\",\"type\":\"address\"}]},{\"outputs\":[{\"internalType\":\"uint256\",\"type\":\"uint256\",\"name\":\"\"}],\"stateMutability\":\"view\",\"name\":\"poolId1\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"internalType\":\"address\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"uint256\",\"type\":\"uint256\",\"name\":\"\"}],\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"name\":\"lastRewardBlock\",\"internalType\":\"uint256\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"type\":\"uint256\",\"name\":\"accGovTokenPerShare\"}],\"type\":\"function\",\"stateMutability\":\"view\",\"name\":\"poolInfo\"},{\"name\":\"removeAuthorized\",\"stateMutability\":\"nonpayable\",\"inputs\":[{\"type\":\"address\",\"name\":\"_toRemove\",\"internalType\":\"address\"}],\"outputs\":[],\"type\":\"function\"},{\"stateMutability\":\"nonpayable\",\"inputs\":[],\"outputs\":[],\"type\":\"function\",\"name\":\"renounceOwnership\"},{\"outputs\":[{\"type\":\"uint256\",\"internalType\":\"uint256\",\"name\":\"\"}],\"stateMutability\":\"view\",\"inputs\":[],\"type\":\"function\",\"name\":\"totalAllocPoint\",\"constant\":true,\"signature\":\"0x17caf6f1\"},{\"type\":\"function\",\"outputs\":[],\"name\":\"transferOwnership\",\"inputs\":[{\"type\":\"address\",\"internalType\":\"address\",\"name\":\"newOwner\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"name\":\"\",\"internalType\":\"address\",\"type\":\"address\"}],\"inputs\":[],\"name\":\"usdOracle\",\"constant\":true,\"signature\":\"0xc8a4271f\"},{\"stateMutability\":\"view\",\"name\":\"userDepFee\",\"type\":\"function\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"inputs\":[],\"constant\":true,\"signature\":\"0x82796e98\"},{\"name\":\"userFeeStage\",\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"name\":\"\",\"internalType\":\"uint256\",\"type\":\"uint256\"}],\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"userGlobalInfo\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"type\":\"uint256\",\"name\":\"globalAmount\",\"internalType\":\"uint256\"},{\"internalType\":\"uint256\",\"type\":\"uint256\",\"name\":\"totalReferals\"},{\"name\":\"globalRefAmount\",\"internalType\":\"uint256\",\"type\":\"uint256\"}],\"constant\":true,\"signature\":\"0xd9608d8a\"},{\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rewardDebt\",\"internalType\":\"uint256\",\"type\":\"uint256\"},{\"name\":\"rewardDebtAtBlock\",\"internalType\":\"uint256\",\"type\":\"uint256\"},{\"type\":\"uint256\",\"internalType\":\"uint256\",\"name\":\"lastWithdrawBlock\"},{\"internalType\":\"uint256\",\"name\":\"firstDepositBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockdelta\",\"type\":\"uint256\"},{\"name\":\"lastDepositBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"name\":\"userInfo\",\"type\":\"function\",\"stateMutability\":\"view\",\"inputs\":[{\"name\":\"\",\"internalType\":\"uint256\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"constant\":true,\"signature\":\"0x93f1a40b\"},{\"outputs\":[{\"type\":\"uint256\",\"internalType\":\"uint256\",\"name\":\"\"}],\"name\":\"poolLength\",\"stateMutability\":\"view\",\"inputs\":[],\"type\":\"function\",\"constant\":true,\"signature\":\"0x081e3eda\"},{\"outputs\":[],\"stateMutability\":\"nonpayable\",\"name\":\"add\",\"inputs\":[{\"name\":\"_allocPoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"internalType\":\"contractIERC20\",\"name\":\"_lpToken\"},{\"type\":\"bool\",\"name\":\"_withUpdate\",\"internalType\":\"bool\"}],\"type\":\"function\"},{\"type\":\"function\",\"outputs\":[],\"name\":\"set\",\"stateMutability\":\"nonpayable\",\"inputs\":[{\"internalType\":\"uint256\",\"type\":\"uint256\",\"name\":\"_pid\"},{\"name\":\"_allocPoint\",\"internalType\":\"uint256\",\"type\":\"uint256\"},{\"name\":\"_withUpdate\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"outputs\":[],\"name\":\"massUpdatePools\",\"type\":\"function\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"name\":\"updatePool\",\"inputs\":[{\"type\":\"uint256\",\"internalType\":\"uint256\",\"name\":\"_pid\"}],\"outputs\":[]},{\"name\":\"getMultiplier\",\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"inputs\":[{\"name\":\"_from\",\"internalType\":\"uint256\",\"type\":\"uint256\"},{\"name\":\"_to\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"name\":\"getLockPercentage\",\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"type\":\"uint256\",\"name\":\"_to\"}]},{\"name\":\"getPoolReward\",\"inputs\":[{\"type\":\"uint256\",\"internalType\":\"uint256\",\"name\":\"_from\"},{\"name\":\"_to\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"}],\"type\":\"function\",\"outputs\":[{\"internalType\":\"uint256\",\"type\":\"uint256\",\"name\":\"forDev\"},{\"name\":\"forFarmer\",\"internalType\":\"uint256\",\"type\":\"uint256\"},{\"name\":\"forLP\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"forCom\",\"internalType\":\"uint256\",\"type\":\"uint256\"},{\"name\":\"forFounders\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"stateMutability\":\"view\",\"inputs\":[{\"type\":\"uint256\",\"internalType\":\"uint256\",\"name\":\"_pid\"},{\"name\":\"_user\",\"internalType\":\"address\",\"type\":\"address\"}],\"outputs\":[{\"type\":\"uint256\",\"internalType\":\"uint256\",\"name\":\"\"}],\"name\":\"pendingReward\",\"type\":\"function\"},{\"name\":\"claimRewards\",\"stateMutability\":\"nonpayable\",\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_pids\",\"type\":\"uint256[]\"}],\"type\":\"function\",\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"outputs\":[],\"inputs\":[{\"name\":\"_pid\",\"internalType\":\"uint256\",\"type\":\"uint256\"}],\"name\":\"claimReward\",\"type\":\"function\"},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"internalType\":\"uint256\",\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getGlobalAmount\"},{\"inputs\":[{\"name\":\"_user\",\"internalType\":\"address\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getGlobalRefAmount\"},{\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"type\":\"function\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"}],\"name\":\"getTotalRefs\",\"stateMutability\":\"view\"},{\"inputs\":[{\"type\":\"address\",\"name\":\"_user\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"_user2\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"type\":\"function\",\"stateMutability\":\"view\",\"name\":\"getRefValueOf\"},{\"name\":\"deposit\",\"stateMutability\":\"nonpayable\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_ref\",\"internalType\":\"address\",\"type\":\"address\"}],\"type\":\"function\",\"outputs\":[]},{\"inputs\":[{\"name\":\"_pid\",\"internalType\":\"uint256\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_ref\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"name\":\"withdraw\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"emergencyWithdraw\",\"type\":\"function\",\"outputs\":[],\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"inputs\":[{\"name\":\"_devaddr\",\"internalType\":\"address\",\"type\":\"address\"}],\"name\":\"dev\",\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"stateMutability\":\"nonpayable\",\"inputs\":[{\"name\":\"_newFinish\",\"internalType\":\"uint256\",\"type\":\"uint256\"}],\"name\":\"bonusFinishUpdate\",\"type\":\"function\",\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_newHalving\",\"type\":\"uint256[]\"}],\"type\":\"function\",\"outputs\":[],\"name\":\"halvingUpdate\"},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"outputs\":[],\"name\":\"lpUpdate\",\"inputs\":[{\"name\":\"_newLP\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"comUpdate\",\"type\":\"function\",\"inputs\":[{\"name\":\"_newCom\",\"internalType\":\"address\",\"type\":\"address\"}]},{\"outputs\":[],\"type\":\"function\",\"inputs\":[{\"name\":\"_newFounder\",\"internalType\":\"address\",\"type\":\"address\"}],\"name\":\"founderUpdate\",\"stateMutability\":\"nonpayable\"},{\"name\":\"rewardUpdate\",\"type\":\"function\",\"inputs\":[{\"name\":\"_newReward\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"stateMutability\":\"nonpayable\",\"name\":\"rewardMulUpdate\",\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_newMulReward\",\"type\":\"uint256[]\"}],\"type\":\"function\",\"outputs\":[]},{\"type\":\"function\",\"name\":\"lockUpdate\",\"inputs\":[{\"name\":\"_newlock\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"outputs\":[]},{\"inputs\":[{\"type\":\"uint256\",\"name\":\"_newdevlock\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\",\"name\":\"lockdevUpdate\",\"type\":\"function\"},{\"stateMutability\":\"nonpayable\",\"outputs\":[],\"type\":\"function\",\"name\":\"locklpUpdate\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_newlplock\",\"internalType\":\"uint256\"}]},{\"name\":\"lockcomUpdate\",\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"outputs\":[],\"inputs\":[{\"internalType\":\"uint256\",\"type\":\"uint256\",\"name\":\"_newcomlock\"}]},{\"type\":\"function\",\"outputs\":[],\"name\":\"lockfounderUpdate\",\"stateMutability\":\"nonpayable\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_newfounderlock\",\"internalType\":\"uint256\"}]},{\"name\":\"starblockUpdate\",\"outputs\":[],\"type\":\"function\",\"inputs\":[{\"name\":\"_newstarblock\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"stateMutability\":\"view\",\"type\":\"function\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getNewRewardPerBlock\",\"inputs\":[{\"internalType\":\"uint256\",\"type\":\"uint256\",\"name\":\"pid1\"}]},{\"name\":\"userDelta\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"reviseWithdraw\",\"stateMutability\":\"nonpayable\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_block\",\"internalType\":\"uint256\"}],\"type\":\"function\",\"outputs\":[]},{\"name\":\"reviseDeposit\",\"inputs\":[{\"name\":\"_pid\",\"internalType\":\"uint256\",\"type\":\"uint256\"},{\"type\":\"address\",\"internalType\":\"address\",\"name\":\"_user\"},{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"outputs\":[]},{\"outputs\":[],\"inputs\":[{\"name\":\"_blockStarts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"setStageStarts\"},{\"inputs\":[{\"name\":\"_blockEnds\",\"internalType\":\"uint256[]\",\"type\":\"uint256[]\"}],\"name\":\"setStageEnds\",\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"inputs\":[{\"name\":\"_userFees\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"type\":\"function\",\"outputs\":[],\"name\":\"setUserFeeStage\"},{\"name\":\"setDevFeeStage\",\"inputs\":[{\"internalType\":\"uint256[]\",\"type\":\"uint256[]\",\"name\":\"_devFees\"}],\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[]},{\"name\":\"setDevDepFee\",\"stateMutability\":\"nonpayable\",\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_devDepFees\",\"type\":\"uint256\"}],\"outputs\":[],\"type\":\"function\"},{\"type\":\"function\",\"outputs\":[],\"inputs\":[{\"internalType\":\"uint256\",\"type\":\"uint256\",\"name\":\"_usrDepFees\"}],\"name\":\"setUserDepFee\",\"stateMutability\":\"nonpayable\"},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"name\":\"reclaimTokenOwnership\"}]",
}

// GardenerABI is the input ABI used to generate the binding from.
// Deprecated: Use GardenerMetaData.ABI instead.
var GardenerABI = GardenerMetaData.ABI

// Gardener is an auto generated Go binding around an Ethereum contract.
type Gardener struct {
	GardenerCaller     // Read-only binding to the contract
	GardenerTransactor // Write-only binding to the contract
	GardenerFilterer   // Log filterer for contract events
}

// GardenerCaller is an auto generated read-only Go binding around an Ethereum contract.
type GardenerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GardenerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GardenerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GardenerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GardenerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GardenerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GardenerSession struct {
	Contract     *Gardener         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GardenerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GardenerCallerSession struct {
	Contract *GardenerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// GardenerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GardenerTransactorSession struct {
	Contract     *GardenerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GardenerRaw is an auto generated low-level Go binding around an Ethereum contract.
type GardenerRaw struct {
	Contract *Gardener // Generic contract binding to access the raw methods on
}

// GardenerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GardenerCallerRaw struct {
	Contract *GardenerCaller // Generic read-only contract binding to access the raw methods on
}

// GardenerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GardenerTransactorRaw struct {
	Contract *GardenerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGardener creates a new instance of Gardener, bound to a specific deployed contract.
func NewGardener(address common.Address, backend bind.ContractBackend) (*Gardener, error) {
	contract, err := bindGardener(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gardener{GardenerCaller: GardenerCaller{contract: contract}, GardenerTransactor: GardenerTransactor{contract: contract}, GardenerFilterer: GardenerFilterer{contract: contract}}, nil
}

// NewGardenerCaller creates a new read-only instance of Gardener, bound to a specific deployed contract.
func NewGardenerCaller(address common.Address, caller bind.ContractCaller) (*GardenerCaller, error) {
	contract, err := bindGardener(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GardenerCaller{contract: contract}, nil
}

// NewGardenerTransactor creates a new write-only instance of Gardener, bound to a specific deployed contract.
func NewGardenerTransactor(address common.Address, transactor bind.ContractTransactor) (*GardenerTransactor, error) {
	contract, err := bindGardener(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GardenerTransactor{contract: contract}, nil
}

// NewGardenerFilterer creates a new log filterer instance of Gardener, bound to a specific deployed contract.
func NewGardenerFilterer(address common.Address, filterer bind.ContractFilterer) (*GardenerFilterer, error) {
	contract, err := bindGardener(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GardenerFilterer{contract: contract}, nil
}

// bindGardener binds a generic wrapper to an already deployed contract.
func bindGardener(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GardenerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gardener *GardenerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gardener.Contract.GardenerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gardener *GardenerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gardener.Contract.GardenerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gardener *GardenerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gardener.Contract.GardenerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gardener *GardenerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gardener.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gardener *GardenerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gardener.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gardener *GardenerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gardener.Contract.contract.Transact(opts, method, params...)
}

// FINISHBONUSATBLOCK is a free data retrieval call binding the contract method 0x980c2a98.
//
// Solidity: function FINISH_BONUS_AT_BLOCK() view returns(uint256)
func (_Gardener *GardenerCaller) FINISHBONUSATBLOCK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "FINISH_BONUS_AT_BLOCK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FINISHBONUSATBLOCK is a free data retrieval call binding the contract method 0x980c2a98.
//
// Solidity: function FINISH_BONUS_AT_BLOCK() view returns(uint256)
func (_Gardener *GardenerSession) FINISHBONUSATBLOCK() (*big.Int, error) {
	return _Gardener.Contract.FINISHBONUSATBLOCK(&_Gardener.CallOpts)
}

// FINISHBONUSATBLOCK is a free data retrieval call binding the contract method 0x980c2a98.
//
// Solidity: function FINISH_BONUS_AT_BLOCK() view returns(uint256)
func (_Gardener *GardenerCallerSession) FINISHBONUSATBLOCK() (*big.Int, error) {
	return _Gardener.Contract.FINISHBONUSATBLOCK(&_Gardener.CallOpts)
}

// HALVINGATBLOCK is a free data retrieval call binding the contract method 0x4179b4fb.
//
// Solidity: function HALVING_AT_BLOCK(uint256 ) view returns(uint256)
func (_Gardener *GardenerCaller) HALVINGATBLOCK(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "HALVING_AT_BLOCK", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HALVINGATBLOCK is a free data retrieval call binding the contract method 0x4179b4fb.
//
// Solidity: function HALVING_AT_BLOCK(uint256 ) view returns(uint256)
func (_Gardener *GardenerSession) HALVINGATBLOCK(arg0 *big.Int) (*big.Int, error) {
	return _Gardener.Contract.HALVINGATBLOCK(&_Gardener.CallOpts, arg0)
}

// HALVINGATBLOCK is a free data retrieval call binding the contract method 0x4179b4fb.
//
// Solidity: function HALVING_AT_BLOCK(uint256 ) view returns(uint256)
func (_Gardener *GardenerCallerSession) HALVINGATBLOCK(arg0 *big.Int) (*big.Int, error) {
	return _Gardener.Contract.HALVINGATBLOCK(&_Gardener.CallOpts, arg0)
}

// PERCENTFORCOM is a free data retrieval call binding the contract method 0xa02306f9.
//
// Solidity: function PERCENT_FOR_COM() view returns(uint256)
func (_Gardener *GardenerCaller) PERCENTFORCOM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "PERCENT_FOR_COM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PERCENTFORCOM is a free data retrieval call binding the contract method 0xa02306f9.
//
// Solidity: function PERCENT_FOR_COM() view returns(uint256)
func (_Gardener *GardenerSession) PERCENTFORCOM() (*big.Int, error) {
	return _Gardener.Contract.PERCENTFORCOM(&_Gardener.CallOpts)
}

// PERCENTFORCOM is a free data retrieval call binding the contract method 0xa02306f9.
//
// Solidity: function PERCENT_FOR_COM() view returns(uint256)
func (_Gardener *GardenerCallerSession) PERCENTFORCOM() (*big.Int, error) {
	return _Gardener.Contract.PERCENTFORCOM(&_Gardener.CallOpts)
}

// PERCENTFORDEV is a free data retrieval call binding the contract method 0xed9bdeda.
//
// Solidity: function PERCENT_FOR_DEV() view returns(uint256)
func (_Gardener *GardenerCaller) PERCENTFORDEV(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "PERCENT_FOR_DEV")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PERCENTFORDEV is a free data retrieval call binding the contract method 0xed9bdeda.
//
// Solidity: function PERCENT_FOR_DEV() view returns(uint256)
func (_Gardener *GardenerSession) PERCENTFORDEV() (*big.Int, error) {
	return _Gardener.Contract.PERCENTFORDEV(&_Gardener.CallOpts)
}

// PERCENTFORDEV is a free data retrieval call binding the contract method 0xed9bdeda.
//
// Solidity: function PERCENT_FOR_DEV() view returns(uint256)
func (_Gardener *GardenerCallerSession) PERCENTFORDEV() (*big.Int, error) {
	return _Gardener.Contract.PERCENTFORDEV(&_Gardener.CallOpts)
}

// PERCENTFORFOUNDERS is a free data retrieval call binding the contract method 0xc6929e53.
//
// Solidity: function PERCENT_FOR_FOUNDERS() view returns(uint256)
func (_Gardener *GardenerCaller) PERCENTFORFOUNDERS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "PERCENT_FOR_FOUNDERS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PERCENTFORFOUNDERS is a free data retrieval call binding the contract method 0xc6929e53.
//
// Solidity: function PERCENT_FOR_FOUNDERS() view returns(uint256)
func (_Gardener *GardenerSession) PERCENTFORFOUNDERS() (*big.Int, error) {
	return _Gardener.Contract.PERCENTFORFOUNDERS(&_Gardener.CallOpts)
}

// PERCENTFORFOUNDERS is a free data retrieval call binding the contract method 0xc6929e53.
//
// Solidity: function PERCENT_FOR_FOUNDERS() view returns(uint256)
func (_Gardener *GardenerCallerSession) PERCENTFORFOUNDERS() (*big.Int, error) {
	return _Gardener.Contract.PERCENTFORFOUNDERS(&_Gardener.CallOpts)
}

// PERCENTFORLP is a free data retrieval call binding the contract method 0x0a67d518.
//
// Solidity: function PERCENT_FOR_LP() view returns(uint256)
func (_Gardener *GardenerCaller) PERCENTFORLP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "PERCENT_FOR_LP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PERCENTFORLP is a free data retrieval call binding the contract method 0x0a67d518.
//
// Solidity: function PERCENT_FOR_LP() view returns(uint256)
func (_Gardener *GardenerSession) PERCENTFORLP() (*big.Int, error) {
	return _Gardener.Contract.PERCENTFORLP(&_Gardener.CallOpts)
}

// PERCENTFORLP is a free data retrieval call binding the contract method 0x0a67d518.
//
// Solidity: function PERCENT_FOR_LP() view returns(uint256)
func (_Gardener *GardenerCallerSession) PERCENTFORLP() (*big.Int, error) {
	return _Gardener.Contract.PERCENTFORLP(&_Gardener.CallOpts)
}

// PERCENTLOCKBONUSREWARD is a free data retrieval call binding the contract method 0xa82859c5.
//
// Solidity: function PERCENT_LOCK_BONUS_REWARD(uint256 ) view returns(uint256)
func (_Gardener *GardenerCaller) PERCENTLOCKBONUSREWARD(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "PERCENT_LOCK_BONUS_REWARD", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PERCENTLOCKBONUSREWARD is a free data retrieval call binding the contract method 0xa82859c5.
//
// Solidity: function PERCENT_LOCK_BONUS_REWARD(uint256 ) view returns(uint256)
func (_Gardener *GardenerSession) PERCENTLOCKBONUSREWARD(arg0 *big.Int) (*big.Int, error) {
	return _Gardener.Contract.PERCENTLOCKBONUSREWARD(&_Gardener.CallOpts, arg0)
}

// PERCENTLOCKBONUSREWARD is a free data retrieval call binding the contract method 0xa82859c5.
//
// Solidity: function PERCENT_LOCK_BONUS_REWARD(uint256 ) view returns(uint256)
func (_Gardener *GardenerCallerSession) PERCENTLOCKBONUSREWARD(arg0 *big.Int) (*big.Int, error) {
	return _Gardener.Contract.PERCENTLOCKBONUSREWARD(&_Gardener.CallOpts, arg0)
}

// REWARDMULTIPLIER is a free data retrieval call binding the contract method 0x2fda7735.
//
// Solidity: function REWARD_MULTIPLIER(uint256 ) view returns(uint256)
func (_Gardener *GardenerCaller) REWARDMULTIPLIER(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "REWARD_MULTIPLIER", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REWARDMULTIPLIER is a free data retrieval call binding the contract method 0x2fda7735.
//
// Solidity: function REWARD_MULTIPLIER(uint256 ) view returns(uint256)
func (_Gardener *GardenerSession) REWARDMULTIPLIER(arg0 *big.Int) (*big.Int, error) {
	return _Gardener.Contract.REWARDMULTIPLIER(&_Gardener.CallOpts, arg0)
}

// REWARDMULTIPLIER is a free data retrieval call binding the contract method 0x2fda7735.
//
// Solidity: function REWARD_MULTIPLIER(uint256 ) view returns(uint256)
func (_Gardener *GardenerCallerSession) REWARDMULTIPLIER(arg0 *big.Int) (*big.Int, error) {
	return _Gardener.Contract.REWARDMULTIPLIER(&_Gardener.CallOpts, arg0)
}

// REWARDPERBLOCK is a free data retrieval call binding the contract method 0x975532dc.
//
// Solidity: function REWARD_PER_BLOCK() view returns(uint256)
func (_Gardener *GardenerCaller) REWARDPERBLOCK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "REWARD_PER_BLOCK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REWARDPERBLOCK is a free data retrieval call binding the contract method 0x975532dc.
//
// Solidity: function REWARD_PER_BLOCK() view returns(uint256)
func (_Gardener *GardenerSession) REWARDPERBLOCK() (*big.Int, error) {
	return _Gardener.Contract.REWARDPERBLOCK(&_Gardener.CallOpts)
}

// REWARDPERBLOCK is a free data retrieval call binding the contract method 0x975532dc.
//
// Solidity: function REWARD_PER_BLOCK() view returns(uint256)
func (_Gardener *GardenerCallerSession) REWARDPERBLOCK() (*big.Int, error) {
	return _Gardener.Contract.REWARDPERBLOCK(&_Gardener.CallOpts)
}

// STARTBLOCK is a free data retrieval call binding the contract method 0x39b3e826.
//
// Solidity: function START_BLOCK() view returns(uint256)
func (_Gardener *GardenerCaller) STARTBLOCK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "START_BLOCK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STARTBLOCK is a free data retrieval call binding the contract method 0x39b3e826.
//
// Solidity: function START_BLOCK() view returns(uint256)
func (_Gardener *GardenerSession) STARTBLOCK() (*big.Int, error) {
	return _Gardener.Contract.STARTBLOCK(&_Gardener.CallOpts)
}

// STARTBLOCK is a free data retrieval call binding the contract method 0x39b3e826.
//
// Solidity: function START_BLOCK() view returns(uint256)
func (_Gardener *GardenerCallerSession) STARTBLOCK() (*big.Int, error) {
	return _Gardener.Contract.STARTBLOCK(&_Gardener.CallOpts)
}

// Authorized is a free data retrieval call binding the contract method 0xb9181611.
//
// Solidity: function authorized(address ) view returns(bool)
func (_Gardener *GardenerCaller) Authorized(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "authorized", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Authorized is a free data retrieval call binding the contract method 0xb9181611.
//
// Solidity: function authorized(address ) view returns(bool)
func (_Gardener *GardenerSession) Authorized(arg0 common.Address) (bool, error) {
	return _Gardener.Contract.Authorized(&_Gardener.CallOpts, arg0)
}

// Authorized is a free data retrieval call binding the contract method 0xb9181611.
//
// Solidity: function authorized(address ) view returns(bool)
func (_Gardener *GardenerCallerSession) Authorized(arg0 common.Address) (bool, error) {
	return _Gardener.Contract.Authorized(&_Gardener.CallOpts, arg0)
}

// BlockDeltaEndStage is a free data retrieval call binding the contract method 0xcb0b8ca1.
//
// Solidity: function blockDeltaEndStage(uint256 ) view returns(uint256)
func (_Gardener *GardenerCaller) BlockDeltaEndStage(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "blockDeltaEndStage", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockDeltaEndStage is a free data retrieval call binding the contract method 0xcb0b8ca1.
//
// Solidity: function blockDeltaEndStage(uint256 ) view returns(uint256)
func (_Gardener *GardenerSession) BlockDeltaEndStage(arg0 *big.Int) (*big.Int, error) {
	return _Gardener.Contract.BlockDeltaEndStage(&_Gardener.CallOpts, arg0)
}

// BlockDeltaEndStage is a free data retrieval call binding the contract method 0xcb0b8ca1.
//
// Solidity: function blockDeltaEndStage(uint256 ) view returns(uint256)
func (_Gardener *GardenerCallerSession) BlockDeltaEndStage(arg0 *big.Int) (*big.Int, error) {
	return _Gardener.Contract.BlockDeltaEndStage(&_Gardener.CallOpts, arg0)
}

// BlockDeltaStartStage is a free data retrieval call binding the contract method 0x6245f084.
//
// Solidity: function blockDeltaStartStage(uint256 ) view returns(uint256)
func (_Gardener *GardenerCaller) BlockDeltaStartStage(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "blockDeltaStartStage", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockDeltaStartStage is a free data retrieval call binding the contract method 0x6245f084.
//
// Solidity: function blockDeltaStartStage(uint256 ) view returns(uint256)
func (_Gardener *GardenerSession) BlockDeltaStartStage(arg0 *big.Int) (*big.Int, error) {
	return _Gardener.Contract.BlockDeltaStartStage(&_Gardener.CallOpts, arg0)
}

// BlockDeltaStartStage is a free data retrieval call binding the contract method 0x6245f084.
//
// Solidity: function blockDeltaStartStage(uint256 ) view returns(uint256)
func (_Gardener *GardenerCallerSession) BlockDeltaStartStage(arg0 *big.Int) (*big.Int, error) {
	return _Gardener.Contract.BlockDeltaStartStage(&_Gardener.CallOpts, arg0)
}

// Comfundaddr is a free data retrieval call binding the contract method 0x3c9d9267.
//
// Solidity: function comfundaddr() view returns(address)
func (_Gardener *GardenerCaller) Comfundaddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "comfundaddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Comfundaddr is a free data retrieval call binding the contract method 0x3c9d9267.
//
// Solidity: function comfundaddr() view returns(address)
func (_Gardener *GardenerSession) Comfundaddr() (common.Address, error) {
	return _Gardener.Contract.Comfundaddr(&_Gardener.CallOpts)
}

// Comfundaddr is a free data retrieval call binding the contract method 0x3c9d9267.
//
// Solidity: function comfundaddr() view returns(address)
func (_Gardener *GardenerCallerSession) Comfundaddr() (common.Address, error) {
	return _Gardener.Contract.Comfundaddr(&_Gardener.CallOpts)
}

// DevDepFee is a free data retrieval call binding the contract method 0xc56a10ff.
//
// Solidity: function devDepFee() view returns(uint256)
func (_Gardener *GardenerCaller) DevDepFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "devDepFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DevDepFee is a free data retrieval call binding the contract method 0xc56a10ff.
//
// Solidity: function devDepFee() view returns(uint256)
func (_Gardener *GardenerSession) DevDepFee() (*big.Int, error) {
	return _Gardener.Contract.DevDepFee(&_Gardener.CallOpts)
}

// DevDepFee is a free data retrieval call binding the contract method 0xc56a10ff.
//
// Solidity: function devDepFee() view returns(uint256)
func (_Gardener *GardenerCallerSession) DevDepFee() (*big.Int, error) {
	return _Gardener.Contract.DevDepFee(&_Gardener.CallOpts)
}

// DevFeeStage is a free data retrieval call binding the contract method 0xbeff2a53.
//
// Solidity: function devFeeStage(uint256 ) view returns(uint256)
func (_Gardener *GardenerCaller) DevFeeStage(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "devFeeStage", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DevFeeStage is a free data retrieval call binding the contract method 0xbeff2a53.
//
// Solidity: function devFeeStage(uint256 ) view returns(uint256)
func (_Gardener *GardenerSession) DevFeeStage(arg0 *big.Int) (*big.Int, error) {
	return _Gardener.Contract.DevFeeStage(&_Gardener.CallOpts, arg0)
}

// DevFeeStage is a free data retrieval call binding the contract method 0xbeff2a53.
//
// Solidity: function devFeeStage(uint256 ) view returns(uint256)
func (_Gardener *GardenerCallerSession) DevFeeStage(arg0 *big.Int) (*big.Int, error) {
	return _Gardener.Contract.DevFeeStage(&_Gardener.CallOpts, arg0)
}

// Devaddr is a free data retrieval call binding the contract method 0xd49e77cd.
//
// Solidity: function devaddr() view returns(address)
func (_Gardener *GardenerCaller) Devaddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "devaddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Devaddr is a free data retrieval call binding the contract method 0xd49e77cd.
//
// Solidity: function devaddr() view returns(address)
func (_Gardener *GardenerSession) Devaddr() (common.Address, error) {
	return _Gardener.Contract.Devaddr(&_Gardener.CallOpts)
}

// Devaddr is a free data retrieval call binding the contract method 0xd49e77cd.
//
// Solidity: function devaddr() view returns(address)
func (_Gardener *GardenerCallerSession) Devaddr() (common.Address, error) {
	return _Gardener.Contract.Devaddr(&_Gardener.CallOpts)
}

// Founderaddr is a free data retrieval call binding the contract method 0xec12173d.
//
// Solidity: function founderaddr() view returns(address)
func (_Gardener *GardenerCaller) Founderaddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "founderaddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Founderaddr is a free data retrieval call binding the contract method 0xec12173d.
//
// Solidity: function founderaddr() view returns(address)
func (_Gardener *GardenerSession) Founderaddr() (common.Address, error) {
	return _Gardener.Contract.Founderaddr(&_Gardener.CallOpts)
}

// Founderaddr is a free data retrieval call binding the contract method 0xec12173d.
//
// Solidity: function founderaddr() view returns(address)
func (_Gardener *GardenerCallerSession) Founderaddr() (common.Address, error) {
	return _Gardener.Contract.Founderaddr(&_Gardener.CallOpts)
}

// GetGlobalAmount is a free data retrieval call binding the contract method 0x929c6971.
//
// Solidity: function getGlobalAmount(address _user) view returns(uint256)
func (_Gardener *GardenerCaller) GetGlobalAmount(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "getGlobalAmount", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGlobalAmount is a free data retrieval call binding the contract method 0x929c6971.
//
// Solidity: function getGlobalAmount(address _user) view returns(uint256)
func (_Gardener *GardenerSession) GetGlobalAmount(_user common.Address) (*big.Int, error) {
	return _Gardener.Contract.GetGlobalAmount(&_Gardener.CallOpts, _user)
}

// GetGlobalAmount is a free data retrieval call binding the contract method 0x929c6971.
//
// Solidity: function getGlobalAmount(address _user) view returns(uint256)
func (_Gardener *GardenerCallerSession) GetGlobalAmount(_user common.Address) (*big.Int, error) {
	return _Gardener.Contract.GetGlobalAmount(&_Gardener.CallOpts, _user)
}

// GetGlobalRefAmount is a free data retrieval call binding the contract method 0x5c5e490b.
//
// Solidity: function getGlobalRefAmount(address _user) view returns(uint256)
func (_Gardener *GardenerCaller) GetGlobalRefAmount(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "getGlobalRefAmount", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGlobalRefAmount is a free data retrieval call binding the contract method 0x5c5e490b.
//
// Solidity: function getGlobalRefAmount(address _user) view returns(uint256)
func (_Gardener *GardenerSession) GetGlobalRefAmount(_user common.Address) (*big.Int, error) {
	return _Gardener.Contract.GetGlobalRefAmount(&_Gardener.CallOpts, _user)
}

// GetGlobalRefAmount is a free data retrieval call binding the contract method 0x5c5e490b.
//
// Solidity: function getGlobalRefAmount(address _user) view returns(uint256)
func (_Gardener *GardenerCallerSession) GetGlobalRefAmount(_user common.Address) (*big.Int, error) {
	return _Gardener.Contract.GetGlobalRefAmount(&_Gardener.CallOpts, _user)
}

// GetLockPercentage is a free data retrieval call binding the contract method 0xf930e770.
//
// Solidity: function getLockPercentage(uint256 _from, uint256 _to) view returns(uint256)
func (_Gardener *GardenerCaller) GetLockPercentage(opts *bind.CallOpts, _from *big.Int, _to *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "getLockPercentage", _from, _to)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLockPercentage is a free data retrieval call binding the contract method 0xf930e770.
//
// Solidity: function getLockPercentage(uint256 _from, uint256 _to) view returns(uint256)
func (_Gardener *GardenerSession) GetLockPercentage(_from *big.Int, _to *big.Int) (*big.Int, error) {
	return _Gardener.Contract.GetLockPercentage(&_Gardener.CallOpts, _from, _to)
}

// GetLockPercentage is a free data retrieval call binding the contract method 0xf930e770.
//
// Solidity: function getLockPercentage(uint256 _from, uint256 _to) view returns(uint256)
func (_Gardener *GardenerCallerSession) GetLockPercentage(_from *big.Int, _to *big.Int) (*big.Int, error) {
	return _Gardener.Contract.GetLockPercentage(&_Gardener.CallOpts, _from, _to)
}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) view returns(uint256)
func (_Gardener *GardenerCaller) GetMultiplier(opts *bind.CallOpts, _from *big.Int, _to *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "getMultiplier", _from, _to)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) view returns(uint256)
func (_Gardener *GardenerSession) GetMultiplier(_from *big.Int, _to *big.Int) (*big.Int, error) {
	return _Gardener.Contract.GetMultiplier(&_Gardener.CallOpts, _from, _to)
}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) view returns(uint256)
func (_Gardener *GardenerCallerSession) GetMultiplier(_from *big.Int, _to *big.Int) (*big.Int, error) {
	return _Gardener.Contract.GetMultiplier(&_Gardener.CallOpts, _from, _to)
}

// GetNewRewardPerBlock is a free data retrieval call binding the contract method 0x1d465c82.
//
// Solidity: function getNewRewardPerBlock(uint256 pid1) view returns(uint256)
func (_Gardener *GardenerCaller) GetNewRewardPerBlock(opts *bind.CallOpts, pid1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "getNewRewardPerBlock", pid1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNewRewardPerBlock is a free data retrieval call binding the contract method 0x1d465c82.
//
// Solidity: function getNewRewardPerBlock(uint256 pid1) view returns(uint256)
func (_Gardener *GardenerSession) GetNewRewardPerBlock(pid1 *big.Int) (*big.Int, error) {
	return _Gardener.Contract.GetNewRewardPerBlock(&_Gardener.CallOpts, pid1)
}

// GetNewRewardPerBlock is a free data retrieval call binding the contract method 0x1d465c82.
//
// Solidity: function getNewRewardPerBlock(uint256 pid1) view returns(uint256)
func (_Gardener *GardenerCallerSession) GetNewRewardPerBlock(pid1 *big.Int) (*big.Int, error) {
	return _Gardener.Contract.GetNewRewardPerBlock(&_Gardener.CallOpts, pid1)
}

// GetPoolReward is a free data retrieval call binding the contract method 0xc8ed7680.
//
// Solidity: function getPoolReward(uint256 _from, uint256 _to, uint256 _allocPoint) view returns(uint256 forDev, uint256 forFarmer, uint256 forLP, uint256 forCom, uint256 forFounders)
func (_Gardener *GardenerCaller) GetPoolReward(opts *bind.CallOpts, _from *big.Int, _to *big.Int, _allocPoint *big.Int) (struct {
	ForDev      *big.Int
	ForFarmer   *big.Int
	ForLP       *big.Int
	ForCom      *big.Int
	ForFounders *big.Int
}, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "getPoolReward", _from, _to, _allocPoint)

	outstruct := new(struct {
		ForDev      *big.Int
		ForFarmer   *big.Int
		ForLP       *big.Int
		ForCom      *big.Int
		ForFounders *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ForDev = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ForFarmer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ForLP = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ForCom = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ForFounders = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPoolReward is a free data retrieval call binding the contract method 0xc8ed7680.
//
// Solidity: function getPoolReward(uint256 _from, uint256 _to, uint256 _allocPoint) view returns(uint256 forDev, uint256 forFarmer, uint256 forLP, uint256 forCom, uint256 forFounders)
func (_Gardener *GardenerSession) GetPoolReward(_from *big.Int, _to *big.Int, _allocPoint *big.Int) (struct {
	ForDev      *big.Int
	ForFarmer   *big.Int
	ForLP       *big.Int
	ForCom      *big.Int
	ForFounders *big.Int
}, error) {
	return _Gardener.Contract.GetPoolReward(&_Gardener.CallOpts, _from, _to, _allocPoint)
}

// GetPoolReward is a free data retrieval call binding the contract method 0xc8ed7680.
//
// Solidity: function getPoolReward(uint256 _from, uint256 _to, uint256 _allocPoint) view returns(uint256 forDev, uint256 forFarmer, uint256 forLP, uint256 forCom, uint256 forFounders)
func (_Gardener *GardenerCallerSession) GetPoolReward(_from *big.Int, _to *big.Int, _allocPoint *big.Int) (struct {
	ForDev      *big.Int
	ForFarmer   *big.Int
	ForLP       *big.Int
	ForCom      *big.Int
	ForFounders *big.Int
}, error) {
	return _Gardener.Contract.GetPoolReward(&_Gardener.CallOpts, _from, _to, _allocPoint)
}

// GetRefValueOf is a free data retrieval call binding the contract method 0xdd77b9fd.
//
// Solidity: function getRefValueOf(address _user, address _user2) view returns(uint256)
func (_Gardener *GardenerCaller) GetRefValueOf(opts *bind.CallOpts, _user common.Address, _user2 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "getRefValueOf", _user, _user2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRefValueOf is a free data retrieval call binding the contract method 0xdd77b9fd.
//
// Solidity: function getRefValueOf(address _user, address _user2) view returns(uint256)
func (_Gardener *GardenerSession) GetRefValueOf(_user common.Address, _user2 common.Address) (*big.Int, error) {
	return _Gardener.Contract.GetRefValueOf(&_Gardener.CallOpts, _user, _user2)
}

// GetRefValueOf is a free data retrieval call binding the contract method 0xdd77b9fd.
//
// Solidity: function getRefValueOf(address _user, address _user2) view returns(uint256)
func (_Gardener *GardenerCallerSession) GetRefValueOf(_user common.Address, _user2 common.Address) (*big.Int, error) {
	return _Gardener.Contract.GetRefValueOf(&_Gardener.CallOpts, _user, _user2)
}

// GetTotalRefs is a free data retrieval call binding the contract method 0x36ebcde6.
//
// Solidity: function getTotalRefs(address _user) view returns(uint256)
func (_Gardener *GardenerCaller) GetTotalRefs(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "getTotalRefs", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalRefs is a free data retrieval call binding the contract method 0x36ebcde6.
//
// Solidity: function getTotalRefs(address _user) view returns(uint256)
func (_Gardener *GardenerSession) GetTotalRefs(_user common.Address) (*big.Int, error) {
	return _Gardener.Contract.GetTotalRefs(&_Gardener.CallOpts, _user)
}

// GetTotalRefs is a free data retrieval call binding the contract method 0x36ebcde6.
//
// Solidity: function getTotalRefs(address _user) view returns(uint256)
func (_Gardener *GardenerCallerSession) GetTotalRefs(_user common.Address) (*big.Int, error) {
	return _Gardener.Contract.GetTotalRefs(&_Gardener.CallOpts, _user)
}

// GovToken is a free data retrieval call binding the contract method 0x05268cff.
//
// Solidity: function govToken() view returns(address)
func (_Gardener *GardenerCaller) GovToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "govToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovToken is a free data retrieval call binding the contract method 0x05268cff.
//
// Solidity: function govToken() view returns(address)
func (_Gardener *GardenerSession) GovToken() (common.Address, error) {
	return _Gardener.Contract.GovToken(&_Gardener.CallOpts)
}

// GovToken is a free data retrieval call binding the contract method 0x05268cff.
//
// Solidity: function govToken() view returns(address)
func (_Gardener *GardenerCallerSession) GovToken() (common.Address, error) {
	return _Gardener.Contract.GovToken(&_Gardener.CallOpts)
}

// Liquidityaddr is a free data retrieval call binding the contract method 0x22a376b0.
//
// Solidity: function liquidityaddr() view returns(address)
func (_Gardener *GardenerCaller) Liquidityaddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "liquidityaddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Liquidityaddr is a free data retrieval call binding the contract method 0x22a376b0.
//
// Solidity: function liquidityaddr() view returns(address)
func (_Gardener *GardenerSession) Liquidityaddr() (common.Address, error) {
	return _Gardener.Contract.Liquidityaddr(&_Gardener.CallOpts)
}

// Liquidityaddr is a free data retrieval call binding the contract method 0x22a376b0.
//
// Solidity: function liquidityaddr() view returns(address)
func (_Gardener *GardenerCallerSession) Liquidityaddr() (common.Address, error) {
	return _Gardener.Contract.Liquidityaddr(&_Gardener.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Gardener *GardenerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Gardener *GardenerSession) Owner() (common.Address, error) {
	return _Gardener.Contract.Owner(&_Gardener.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Gardener *GardenerCallerSession) Owner() (common.Address, error) {
	return _Gardener.Contract.Owner(&_Gardener.CallOpts)
}

// PendingReward is a free data retrieval call binding the contract method 0x98969e82.
//
// Solidity: function pendingReward(uint256 _pid, address _user) view returns(uint256)
func (_Gardener *GardenerCaller) PendingReward(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "pendingReward", _pid, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingReward is a free data retrieval call binding the contract method 0x98969e82.
//
// Solidity: function pendingReward(uint256 _pid, address _user) view returns(uint256)
func (_Gardener *GardenerSession) PendingReward(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _Gardener.Contract.PendingReward(&_Gardener.CallOpts, _pid, _user)
}

// PendingReward is a free data retrieval call binding the contract method 0x98969e82.
//
// Solidity: function pendingReward(uint256 _pid, address _user) view returns(uint256)
func (_Gardener *GardenerCallerSession) PendingReward(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _Gardener.Contract.PendingReward(&_Gardener.CallOpts, _pid, _user)
}

// PoolExistence is a free data retrieval call binding the contract method 0xcbd258b5.
//
// Solidity: function poolExistence(address ) view returns(bool)
func (_Gardener *GardenerCaller) PoolExistence(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "poolExistence", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PoolExistence is a free data retrieval call binding the contract method 0xcbd258b5.
//
// Solidity: function poolExistence(address ) view returns(bool)
func (_Gardener *GardenerSession) PoolExistence(arg0 common.Address) (bool, error) {
	return _Gardener.Contract.PoolExistence(&_Gardener.CallOpts, arg0)
}

// PoolExistence is a free data retrieval call binding the contract method 0xcbd258b5.
//
// Solidity: function poolExistence(address ) view returns(bool)
func (_Gardener *GardenerCallerSession) PoolExistence(arg0 common.Address) (bool, error) {
	return _Gardener.Contract.PoolExistence(&_Gardener.CallOpts, arg0)
}

// PoolId1 is a free data retrieval call binding the contract method 0xce2529c9.
//
// Solidity: function poolId1(address ) view returns(uint256)
func (_Gardener *GardenerCaller) PoolId1(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "poolId1", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolId1 is a free data retrieval call binding the contract method 0xce2529c9.
//
// Solidity: function poolId1(address ) view returns(uint256)
func (_Gardener *GardenerSession) PoolId1(arg0 common.Address) (*big.Int, error) {
	return _Gardener.Contract.PoolId1(&_Gardener.CallOpts, arg0)
}

// PoolId1 is a free data retrieval call binding the contract method 0xce2529c9.
//
// Solidity: function poolId1(address ) view returns(uint256)
func (_Gardener *GardenerCallerSession) PoolId1(arg0 common.Address) (*big.Int, error) {
	return _Gardener.Contract.PoolId1(&_Gardener.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accGovTokenPerShare)
func (_Gardener *GardenerCaller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LpToken             common.Address
	AllocPoint          *big.Int
	LastRewardBlock     *big.Int
	AccGovTokenPerShare *big.Int
}, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		LpToken             common.Address
		AllocPoint          *big.Int
		LastRewardBlock     *big.Int
		AccGovTokenPerShare *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LpToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AllocPoint = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastRewardBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AccGovTokenPerShare = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accGovTokenPerShare)
func (_Gardener *GardenerSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken             common.Address
	AllocPoint          *big.Int
	LastRewardBlock     *big.Int
	AccGovTokenPerShare *big.Int
}, error) {
	return _Gardener.Contract.PoolInfo(&_Gardener.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accGovTokenPerShare)
func (_Gardener *GardenerCallerSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken             common.Address
	AllocPoint          *big.Int
	LastRewardBlock     *big.Int
	AccGovTokenPerShare *big.Int
}, error) {
	return _Gardener.Contract.PoolInfo(&_Gardener.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Gardener *GardenerCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Gardener *GardenerSession) PoolLength() (*big.Int, error) {
	return _Gardener.Contract.PoolLength(&_Gardener.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Gardener *GardenerCallerSession) PoolLength() (*big.Int, error) {
	return _Gardener.Contract.PoolLength(&_Gardener.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_Gardener *GardenerCaller) TotalAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "totalAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_Gardener *GardenerSession) TotalAllocPoint() (*big.Int, error) {
	return _Gardener.Contract.TotalAllocPoint(&_Gardener.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_Gardener *GardenerCallerSession) TotalAllocPoint() (*big.Int, error) {
	return _Gardener.Contract.TotalAllocPoint(&_Gardener.CallOpts)
}

// UsdOracle is a free data retrieval call binding the contract method 0xc8a4271f.
//
// Solidity: function usdOracle() view returns(address)
func (_Gardener *GardenerCaller) UsdOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "usdOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsdOracle is a free data retrieval call binding the contract method 0xc8a4271f.
//
// Solidity: function usdOracle() view returns(address)
func (_Gardener *GardenerSession) UsdOracle() (common.Address, error) {
	return _Gardener.Contract.UsdOracle(&_Gardener.CallOpts)
}

// UsdOracle is a free data retrieval call binding the contract method 0xc8a4271f.
//
// Solidity: function usdOracle() view returns(address)
func (_Gardener *GardenerCallerSession) UsdOracle() (common.Address, error) {
	return _Gardener.Contract.UsdOracle(&_Gardener.CallOpts)
}

// UserDelta is a free data retrieval call binding the contract method 0x09ae4d2c.
//
// Solidity: function userDelta(uint256 _pid) view returns(uint256)
func (_Gardener *GardenerCaller) UserDelta(opts *bind.CallOpts, _pid *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "userDelta", _pid)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserDelta is a free data retrieval call binding the contract method 0x09ae4d2c.
//
// Solidity: function userDelta(uint256 _pid) view returns(uint256)
func (_Gardener *GardenerSession) UserDelta(_pid *big.Int) (*big.Int, error) {
	return _Gardener.Contract.UserDelta(&_Gardener.CallOpts, _pid)
}

// UserDelta is a free data retrieval call binding the contract method 0x09ae4d2c.
//
// Solidity: function userDelta(uint256 _pid) view returns(uint256)
func (_Gardener *GardenerCallerSession) UserDelta(_pid *big.Int) (*big.Int, error) {
	return _Gardener.Contract.UserDelta(&_Gardener.CallOpts, _pid)
}

// UserDepFee is a free data retrieval call binding the contract method 0x82796e98.
//
// Solidity: function userDepFee() view returns(uint256)
func (_Gardener *GardenerCaller) UserDepFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "userDepFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserDepFee is a free data retrieval call binding the contract method 0x82796e98.
//
// Solidity: function userDepFee() view returns(uint256)
func (_Gardener *GardenerSession) UserDepFee() (*big.Int, error) {
	return _Gardener.Contract.UserDepFee(&_Gardener.CallOpts)
}

// UserDepFee is a free data retrieval call binding the contract method 0x82796e98.
//
// Solidity: function userDepFee() view returns(uint256)
func (_Gardener *GardenerCallerSession) UserDepFee() (*big.Int, error) {
	return _Gardener.Contract.UserDepFee(&_Gardener.CallOpts)
}

// UserFeeStage is a free data retrieval call binding the contract method 0xd007db29.
//
// Solidity: function userFeeStage(uint256 ) view returns(uint256)
func (_Gardener *GardenerCaller) UserFeeStage(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "userFeeStage", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserFeeStage is a free data retrieval call binding the contract method 0xd007db29.
//
// Solidity: function userFeeStage(uint256 ) view returns(uint256)
func (_Gardener *GardenerSession) UserFeeStage(arg0 *big.Int) (*big.Int, error) {
	return _Gardener.Contract.UserFeeStage(&_Gardener.CallOpts, arg0)
}

// UserFeeStage is a free data retrieval call binding the contract method 0xd007db29.
//
// Solidity: function userFeeStage(uint256 ) view returns(uint256)
func (_Gardener *GardenerCallerSession) UserFeeStage(arg0 *big.Int) (*big.Int, error) {
	return _Gardener.Contract.UserFeeStage(&_Gardener.CallOpts, arg0)
}

// UserGlobalInfo is a free data retrieval call binding the contract method 0xd9608d8a.
//
// Solidity: function userGlobalInfo(address ) view returns(uint256 globalAmount, uint256 totalReferals, uint256 globalRefAmount)
func (_Gardener *GardenerCaller) UserGlobalInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	GlobalAmount    *big.Int
	TotalReferals   *big.Int
	GlobalRefAmount *big.Int
}, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "userGlobalInfo", arg0)

	outstruct := new(struct {
		GlobalAmount    *big.Int
		TotalReferals   *big.Int
		GlobalRefAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GlobalAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalReferals = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.GlobalRefAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserGlobalInfo is a free data retrieval call binding the contract method 0xd9608d8a.
//
// Solidity: function userGlobalInfo(address ) view returns(uint256 globalAmount, uint256 totalReferals, uint256 globalRefAmount)
func (_Gardener *GardenerSession) UserGlobalInfo(arg0 common.Address) (struct {
	GlobalAmount    *big.Int
	TotalReferals   *big.Int
	GlobalRefAmount *big.Int
}, error) {
	return _Gardener.Contract.UserGlobalInfo(&_Gardener.CallOpts, arg0)
}

// UserGlobalInfo is a free data retrieval call binding the contract method 0xd9608d8a.
//
// Solidity: function userGlobalInfo(address ) view returns(uint256 globalAmount, uint256 totalReferals, uint256 globalRefAmount)
func (_Gardener *GardenerCallerSession) UserGlobalInfo(arg0 common.Address) (struct {
	GlobalAmount    *big.Int
	TotalReferals   *big.Int
	GlobalRefAmount *big.Int
}, error) {
	return _Gardener.Contract.UserGlobalInfo(&_Gardener.CallOpts, arg0)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt, uint256 rewardDebtAtBlock, uint256 lastWithdrawBlock, uint256 firstDepositBlock, uint256 blockdelta, uint256 lastDepositBlock)
func (_Gardener *GardenerCaller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount            *big.Int
	RewardDebt        *big.Int
	RewardDebtAtBlock *big.Int
	LastWithdrawBlock *big.Int
	FirstDepositBlock *big.Int
	Blockdelta        *big.Int
	LastDepositBlock  *big.Int
}, error) {
	var out []interface{}
	err := _Gardener.contract.Call(opts, &out, "userInfo", arg0, arg1)

	outstruct := new(struct {
		Amount            *big.Int
		RewardDebt        *big.Int
		RewardDebtAtBlock *big.Int
		LastWithdrawBlock *big.Int
		FirstDepositBlock *big.Int
		Blockdelta        *big.Int
		LastDepositBlock  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RewardDebtAtBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LastWithdrawBlock = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.FirstDepositBlock = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Blockdelta = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.LastDepositBlock = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt, uint256 rewardDebtAtBlock, uint256 lastWithdrawBlock, uint256 firstDepositBlock, uint256 blockdelta, uint256 lastDepositBlock)
func (_Gardener *GardenerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount            *big.Int
	RewardDebt        *big.Int
	RewardDebtAtBlock *big.Int
	LastWithdrawBlock *big.Int
	FirstDepositBlock *big.Int
	Blockdelta        *big.Int
	LastDepositBlock  *big.Int
}, error) {
	return _Gardener.Contract.UserInfo(&_Gardener.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt, uint256 rewardDebtAtBlock, uint256 lastWithdrawBlock, uint256 firstDepositBlock, uint256 blockdelta, uint256 lastDepositBlock)
func (_Gardener *GardenerCallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount            *big.Int
	RewardDebt        *big.Int
	RewardDebtAtBlock *big.Int
	LastWithdrawBlock *big.Int
	FirstDepositBlock *big.Int
	Blockdelta        *big.Int
	LastDepositBlock  *big.Int
}, error) {
	return _Gardener.Contract.UserInfo(&_Gardener.CallOpts, arg0, arg1)
}

// Add is a paid mutator transaction binding the contract method 0x1eaaa045.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _withUpdate) returns()
func (_Gardener *GardenerTransactor) Add(opts *bind.TransactOpts, _allocPoint *big.Int, _lpToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "add", _allocPoint, _lpToken, _withUpdate)
}

// Add is a paid mutator transaction binding the contract method 0x1eaaa045.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _withUpdate) returns()
func (_Gardener *GardenerSession) Add(_allocPoint *big.Int, _lpToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _Gardener.Contract.Add(&_Gardener.TransactOpts, _allocPoint, _lpToken, _withUpdate)
}

// Add is a paid mutator transaction binding the contract method 0x1eaaa045.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _withUpdate) returns()
func (_Gardener *GardenerTransactorSession) Add(_allocPoint *big.Int, _lpToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _Gardener.Contract.Add(&_Gardener.TransactOpts, _allocPoint, _lpToken, _withUpdate)
}

// AddAuthorized is a paid mutator transaction binding the contract method 0xcf1c316a.
//
// Solidity: function addAuthorized(address _toAdd) returns()
func (_Gardener *GardenerTransactor) AddAuthorized(opts *bind.TransactOpts, _toAdd common.Address) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "addAuthorized", _toAdd)
}

// AddAuthorized is a paid mutator transaction binding the contract method 0xcf1c316a.
//
// Solidity: function addAuthorized(address _toAdd) returns()
func (_Gardener *GardenerSession) AddAuthorized(_toAdd common.Address) (*types.Transaction, error) {
	return _Gardener.Contract.AddAuthorized(&_Gardener.TransactOpts, _toAdd)
}

// AddAuthorized is a paid mutator transaction binding the contract method 0xcf1c316a.
//
// Solidity: function addAuthorized(address _toAdd) returns()
func (_Gardener *GardenerTransactorSession) AddAuthorized(_toAdd common.Address) (*types.Transaction, error) {
	return _Gardener.Contract.AddAuthorized(&_Gardener.TransactOpts, _toAdd)
}

// BonusFinishUpdate is a paid mutator transaction binding the contract method 0x245b211d.
//
// Solidity: function bonusFinishUpdate(uint256 _newFinish) returns()
func (_Gardener *GardenerTransactor) BonusFinishUpdate(opts *bind.TransactOpts, _newFinish *big.Int) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "bonusFinishUpdate", _newFinish)
}

// BonusFinishUpdate is a paid mutator transaction binding the contract method 0x245b211d.
//
// Solidity: function bonusFinishUpdate(uint256 _newFinish) returns()
func (_Gardener *GardenerSession) BonusFinishUpdate(_newFinish *big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.BonusFinishUpdate(&_Gardener.TransactOpts, _newFinish)
}

// BonusFinishUpdate is a paid mutator transaction binding the contract method 0x245b211d.
//
// Solidity: function bonusFinishUpdate(uint256 _newFinish) returns()
func (_Gardener *GardenerTransactorSession) BonusFinishUpdate(_newFinish *big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.BonusFinishUpdate(&_Gardener.TransactOpts, _newFinish)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xae169a50.
//
// Solidity: function claimReward(uint256 _pid) returns()
func (_Gardener *GardenerTransactor) ClaimReward(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "claimReward", _pid)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xae169a50.
//
// Solidity: function claimReward(uint256 _pid) returns()
func (_Gardener *GardenerSession) ClaimReward(_pid *big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.ClaimReward(&_Gardener.TransactOpts, _pid)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xae169a50.
//
// Solidity: function claimReward(uint256 _pid) returns()
func (_Gardener *GardenerTransactorSession) ClaimReward(_pid *big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.ClaimReward(&_Gardener.TransactOpts, _pid)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x5eac6239.
//
// Solidity: function claimRewards(uint256[] _pids) returns()
func (_Gardener *GardenerTransactor) ClaimRewards(opts *bind.TransactOpts, _pids []*big.Int) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "claimRewards", _pids)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x5eac6239.
//
// Solidity: function claimRewards(uint256[] _pids) returns()
func (_Gardener *GardenerSession) ClaimRewards(_pids []*big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.ClaimRewards(&_Gardener.TransactOpts, _pids)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x5eac6239.
//
// Solidity: function claimRewards(uint256[] _pids) returns()
func (_Gardener *GardenerTransactorSession) ClaimRewards(_pids []*big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.ClaimRewards(&_Gardener.TransactOpts, _pids)
}

// ComUpdate is a paid mutator transaction binding the contract method 0xc749d614.
//
// Solidity: function comUpdate(address _newCom) returns()
func (_Gardener *GardenerTransactor) ComUpdate(opts *bind.TransactOpts, _newCom common.Address) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "comUpdate", _newCom)
}

// ComUpdate is a paid mutator transaction binding the contract method 0xc749d614.
//
// Solidity: function comUpdate(address _newCom) returns()
func (_Gardener *GardenerSession) ComUpdate(_newCom common.Address) (*types.Transaction, error) {
	return _Gardener.Contract.ComUpdate(&_Gardener.TransactOpts, _newCom)
}

// ComUpdate is a paid mutator transaction binding the contract method 0xc749d614.
//
// Solidity: function comUpdate(address _newCom) returns()
func (_Gardener *GardenerTransactorSession) ComUpdate(_newCom common.Address) (*types.Transaction, error) {
	return _Gardener.Contract.ComUpdate(&_Gardener.TransactOpts, _newCom)
}

// Deposit is a paid mutator transaction binding the contract method 0x8dbdbe6d.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount, address _ref) returns()
func (_Gardener *GardenerTransactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int, _ref common.Address) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "deposit", _pid, _amount, _ref)
}

// Deposit is a paid mutator transaction binding the contract method 0x8dbdbe6d.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount, address _ref) returns()
func (_Gardener *GardenerSession) Deposit(_pid *big.Int, _amount *big.Int, _ref common.Address) (*types.Transaction, error) {
	return _Gardener.Contract.Deposit(&_Gardener.TransactOpts, _pid, _amount, _ref)
}

// Deposit is a paid mutator transaction binding the contract method 0x8dbdbe6d.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount, address _ref) returns()
func (_Gardener *GardenerTransactorSession) Deposit(_pid *big.Int, _amount *big.Int, _ref common.Address) (*types.Transaction, error) {
	return _Gardener.Contract.Deposit(&_Gardener.TransactOpts, _pid, _amount, _ref)
}

// Dev is a paid mutator transaction binding the contract method 0x8d88a90e.
//
// Solidity: function dev(address _devaddr) returns()
func (_Gardener *GardenerTransactor) Dev(opts *bind.TransactOpts, _devaddr common.Address) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "dev", _devaddr)
}

// Dev is a paid mutator transaction binding the contract method 0x8d88a90e.
//
// Solidity: function dev(address _devaddr) returns()
func (_Gardener *GardenerSession) Dev(_devaddr common.Address) (*types.Transaction, error) {
	return _Gardener.Contract.Dev(&_Gardener.TransactOpts, _devaddr)
}

// Dev is a paid mutator transaction binding the contract method 0x8d88a90e.
//
// Solidity: function dev(address _devaddr) returns()
func (_Gardener *GardenerTransactorSession) Dev(_devaddr common.Address) (*types.Transaction, error) {
	return _Gardener.Contract.Dev(&_Gardener.TransactOpts, _devaddr)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_Gardener *GardenerTransactor) EmergencyWithdraw(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "emergencyWithdraw", _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_Gardener *GardenerSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.EmergencyWithdraw(&_Gardener.TransactOpts, _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_Gardener *GardenerTransactorSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.EmergencyWithdraw(&_Gardener.TransactOpts, _pid)
}

// FounderUpdate is a paid mutator transaction binding the contract method 0x30fb8e0e.
//
// Solidity: function founderUpdate(address _newFounder) returns()
func (_Gardener *GardenerTransactor) FounderUpdate(opts *bind.TransactOpts, _newFounder common.Address) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "founderUpdate", _newFounder)
}

// FounderUpdate is a paid mutator transaction binding the contract method 0x30fb8e0e.
//
// Solidity: function founderUpdate(address _newFounder) returns()
func (_Gardener *GardenerSession) FounderUpdate(_newFounder common.Address) (*types.Transaction, error) {
	return _Gardener.Contract.FounderUpdate(&_Gardener.TransactOpts, _newFounder)
}

// FounderUpdate is a paid mutator transaction binding the contract method 0x30fb8e0e.
//
// Solidity: function founderUpdate(address _newFounder) returns()
func (_Gardener *GardenerTransactorSession) FounderUpdate(_newFounder common.Address) (*types.Transaction, error) {
	return _Gardener.Contract.FounderUpdate(&_Gardener.TransactOpts, _newFounder)
}

// HalvingUpdate is a paid mutator transaction binding the contract method 0x354affb7.
//
// Solidity: function halvingUpdate(uint256[] _newHalving) returns()
func (_Gardener *GardenerTransactor) HalvingUpdate(opts *bind.TransactOpts, _newHalving []*big.Int) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "halvingUpdate", _newHalving)
}

// HalvingUpdate is a paid mutator transaction binding the contract method 0x354affb7.
//
// Solidity: function halvingUpdate(uint256[] _newHalving) returns()
func (_Gardener *GardenerSession) HalvingUpdate(_newHalving []*big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.HalvingUpdate(&_Gardener.TransactOpts, _newHalving)
}

// HalvingUpdate is a paid mutator transaction binding the contract method 0x354affb7.
//
// Solidity: function halvingUpdate(uint256[] _newHalving) returns()
func (_Gardener *GardenerTransactorSession) HalvingUpdate(_newHalving []*big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.HalvingUpdate(&_Gardener.TransactOpts, _newHalving)
}

// LockUpdate is a paid mutator transaction binding the contract method 0x27e68928.
//
// Solidity: function lockUpdate(uint256[] _newlock) returns()
func (_Gardener *GardenerTransactor) LockUpdate(opts *bind.TransactOpts, _newlock []*big.Int) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "lockUpdate", _newlock)
}

// LockUpdate is a paid mutator transaction binding the contract method 0x27e68928.
//
// Solidity: function lockUpdate(uint256[] _newlock) returns()
func (_Gardener *GardenerSession) LockUpdate(_newlock []*big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.LockUpdate(&_Gardener.TransactOpts, _newlock)
}

// LockUpdate is a paid mutator transaction binding the contract method 0x27e68928.
//
// Solidity: function lockUpdate(uint256[] _newlock) returns()
func (_Gardener *GardenerTransactorSession) LockUpdate(_newlock []*big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.LockUpdate(&_Gardener.TransactOpts, _newlock)
}

// LockcomUpdate is a paid mutator transaction binding the contract method 0xc663baa6.
//
// Solidity: function lockcomUpdate(uint256 _newcomlock) returns()
func (_Gardener *GardenerTransactor) LockcomUpdate(opts *bind.TransactOpts, _newcomlock *big.Int) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "lockcomUpdate", _newcomlock)
}

// LockcomUpdate is a paid mutator transaction binding the contract method 0xc663baa6.
//
// Solidity: function lockcomUpdate(uint256 _newcomlock) returns()
func (_Gardener *GardenerSession) LockcomUpdate(_newcomlock *big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.LockcomUpdate(&_Gardener.TransactOpts, _newcomlock)
}

// LockcomUpdate is a paid mutator transaction binding the contract method 0xc663baa6.
//
// Solidity: function lockcomUpdate(uint256 _newcomlock) returns()
func (_Gardener *GardenerTransactorSession) LockcomUpdate(_newcomlock *big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.LockcomUpdate(&_Gardener.TransactOpts, _newcomlock)
}

// LockdevUpdate is a paid mutator transaction binding the contract method 0xb6066962.
//
// Solidity: function lockdevUpdate(uint256 _newdevlock) returns()
func (_Gardener *GardenerTransactor) LockdevUpdate(opts *bind.TransactOpts, _newdevlock *big.Int) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "lockdevUpdate", _newdevlock)
}

// LockdevUpdate is a paid mutator transaction binding the contract method 0xb6066962.
//
// Solidity: function lockdevUpdate(uint256 _newdevlock) returns()
func (_Gardener *GardenerSession) LockdevUpdate(_newdevlock *big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.LockdevUpdate(&_Gardener.TransactOpts, _newdevlock)
}

// LockdevUpdate is a paid mutator transaction binding the contract method 0xb6066962.
//
// Solidity: function lockdevUpdate(uint256 _newdevlock) returns()
func (_Gardener *GardenerTransactorSession) LockdevUpdate(_newdevlock *big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.LockdevUpdate(&_Gardener.TransactOpts, _newdevlock)
}

// LockfounderUpdate is a paid mutator transaction binding the contract method 0xde988524.
//
// Solidity: function lockfounderUpdate(uint256 _newfounderlock) returns()
func (_Gardener *GardenerTransactor) LockfounderUpdate(opts *bind.TransactOpts, _newfounderlock *big.Int) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "lockfounderUpdate", _newfounderlock)
}

// LockfounderUpdate is a paid mutator transaction binding the contract method 0xde988524.
//
// Solidity: function lockfounderUpdate(uint256 _newfounderlock) returns()
func (_Gardener *GardenerSession) LockfounderUpdate(_newfounderlock *big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.LockfounderUpdate(&_Gardener.TransactOpts, _newfounderlock)
}

// LockfounderUpdate is a paid mutator transaction binding the contract method 0xde988524.
//
// Solidity: function lockfounderUpdate(uint256 _newfounderlock) returns()
func (_Gardener *GardenerTransactorSession) LockfounderUpdate(_newfounderlock *big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.LockfounderUpdate(&_Gardener.TransactOpts, _newfounderlock)
}

// LocklpUpdate is a paid mutator transaction binding the contract method 0xf2ffc22c.
//
// Solidity: function locklpUpdate(uint256 _newlplock) returns()
func (_Gardener *GardenerTransactor) LocklpUpdate(opts *bind.TransactOpts, _newlplock *big.Int) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "locklpUpdate", _newlplock)
}

// LocklpUpdate is a paid mutator transaction binding the contract method 0xf2ffc22c.
//
// Solidity: function locklpUpdate(uint256 _newlplock) returns()
func (_Gardener *GardenerSession) LocklpUpdate(_newlplock *big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.LocklpUpdate(&_Gardener.TransactOpts, _newlplock)
}

// LocklpUpdate is a paid mutator transaction binding the contract method 0xf2ffc22c.
//
// Solidity: function locklpUpdate(uint256 _newlplock) returns()
func (_Gardener *GardenerTransactorSession) LocklpUpdate(_newlplock *big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.LocklpUpdate(&_Gardener.TransactOpts, _newlplock)
}

// LpUpdate is a paid mutator transaction binding the contract method 0x4dbf85ca.
//
// Solidity: function lpUpdate(address _newLP) returns()
func (_Gardener *GardenerTransactor) LpUpdate(opts *bind.TransactOpts, _newLP common.Address) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "lpUpdate", _newLP)
}

// LpUpdate is a paid mutator transaction binding the contract method 0x4dbf85ca.
//
// Solidity: function lpUpdate(address _newLP) returns()
func (_Gardener *GardenerSession) LpUpdate(_newLP common.Address) (*types.Transaction, error) {
	return _Gardener.Contract.LpUpdate(&_Gardener.TransactOpts, _newLP)
}

// LpUpdate is a paid mutator transaction binding the contract method 0x4dbf85ca.
//
// Solidity: function lpUpdate(address _newLP) returns()
func (_Gardener *GardenerTransactorSession) LpUpdate(_newLP common.Address) (*types.Transaction, error) {
	return _Gardener.Contract.LpUpdate(&_Gardener.TransactOpts, _newLP)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_Gardener *GardenerTransactor) MassUpdatePools(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "massUpdatePools")
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_Gardener *GardenerSession) MassUpdatePools() (*types.Transaction, error) {
	return _Gardener.Contract.MassUpdatePools(&_Gardener.TransactOpts)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_Gardener *GardenerTransactorSession) MassUpdatePools() (*types.Transaction, error) {
	return _Gardener.Contract.MassUpdatePools(&_Gardener.TransactOpts)
}

// ReclaimTokenOwnership is a paid mutator transaction binding the contract method 0xeda67048.
//
// Solidity: function reclaimTokenOwnership(address _newOwner) returns()
func (_Gardener *GardenerTransactor) ReclaimTokenOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "reclaimTokenOwnership", _newOwner)
}

// ReclaimTokenOwnership is a paid mutator transaction binding the contract method 0xeda67048.
//
// Solidity: function reclaimTokenOwnership(address _newOwner) returns()
func (_Gardener *GardenerSession) ReclaimTokenOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Gardener.Contract.ReclaimTokenOwnership(&_Gardener.TransactOpts, _newOwner)
}

// ReclaimTokenOwnership is a paid mutator transaction binding the contract method 0xeda67048.
//
// Solidity: function reclaimTokenOwnership(address _newOwner) returns()
func (_Gardener *GardenerTransactorSession) ReclaimTokenOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Gardener.Contract.ReclaimTokenOwnership(&_Gardener.TransactOpts, _newOwner)
}

// RemoveAuthorized is a paid mutator transaction binding the contract method 0x485d7d94.
//
// Solidity: function removeAuthorized(address _toRemove) returns()
func (_Gardener *GardenerTransactor) RemoveAuthorized(opts *bind.TransactOpts, _toRemove common.Address) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "removeAuthorized", _toRemove)
}

// RemoveAuthorized is a paid mutator transaction binding the contract method 0x485d7d94.
//
// Solidity: function removeAuthorized(address _toRemove) returns()
func (_Gardener *GardenerSession) RemoveAuthorized(_toRemove common.Address) (*types.Transaction, error) {
	return _Gardener.Contract.RemoveAuthorized(&_Gardener.TransactOpts, _toRemove)
}

// RemoveAuthorized is a paid mutator transaction binding the contract method 0x485d7d94.
//
// Solidity: function removeAuthorized(address _toRemove) returns()
func (_Gardener *GardenerTransactorSession) RemoveAuthorized(_toRemove common.Address) (*types.Transaction, error) {
	return _Gardener.Contract.RemoveAuthorized(&_Gardener.TransactOpts, _toRemove)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Gardener *GardenerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Gardener *GardenerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Gardener.Contract.RenounceOwnership(&_Gardener.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Gardener *GardenerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Gardener.Contract.RenounceOwnership(&_Gardener.TransactOpts)
}

// ReviseDeposit is a paid mutator transaction binding the contract method 0x82386d58.
//
// Solidity: function reviseDeposit(uint256 _pid, address _user, uint256 _block) returns()
func (_Gardener *GardenerTransactor) ReviseDeposit(opts *bind.TransactOpts, _pid *big.Int, _user common.Address, _block *big.Int) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "reviseDeposit", _pid, _user, _block)
}

// ReviseDeposit is a paid mutator transaction binding the contract method 0x82386d58.
//
// Solidity: function reviseDeposit(uint256 _pid, address _user, uint256 _block) returns()
func (_Gardener *GardenerSession) ReviseDeposit(_pid *big.Int, _user common.Address, _block *big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.ReviseDeposit(&_Gardener.TransactOpts, _pid, _user, _block)
}

// ReviseDeposit is a paid mutator transaction binding the contract method 0x82386d58.
//
// Solidity: function reviseDeposit(uint256 _pid, address _user, uint256 _block) returns()
func (_Gardener *GardenerTransactorSession) ReviseDeposit(_pid *big.Int, _user common.Address, _block *big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.ReviseDeposit(&_Gardener.TransactOpts, _pid, _user, _block)
}

// ReviseWithdraw is a paid mutator transaction binding the contract method 0x6066debd.
//
// Solidity: function reviseWithdraw(uint256 _pid, address _user, uint256 _block) returns()
func (_Gardener *GardenerTransactor) ReviseWithdraw(opts *bind.TransactOpts, _pid *big.Int, _user common.Address, _block *big.Int) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "reviseWithdraw", _pid, _user, _block)
}

// ReviseWithdraw is a paid mutator transaction binding the contract method 0x6066debd.
//
// Solidity: function reviseWithdraw(uint256 _pid, address _user, uint256 _block) returns()
func (_Gardener *GardenerSession) ReviseWithdraw(_pid *big.Int, _user common.Address, _block *big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.ReviseWithdraw(&_Gardener.TransactOpts, _pid, _user, _block)
}

// ReviseWithdraw is a paid mutator transaction binding the contract method 0x6066debd.
//
// Solidity: function reviseWithdraw(uint256 _pid, address _user, uint256 _block) returns()
func (_Gardener *GardenerTransactorSession) ReviseWithdraw(_pid *big.Int, _user common.Address, _block *big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.ReviseWithdraw(&_Gardener.TransactOpts, _pid, _user, _block)
}

// RewardMulUpdate is a paid mutator transaction binding the contract method 0x46664064.
//
// Solidity: function rewardMulUpdate(uint256[] _newMulReward) returns()
func (_Gardener *GardenerTransactor) RewardMulUpdate(opts *bind.TransactOpts, _newMulReward []*big.Int) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "rewardMulUpdate", _newMulReward)
}

// RewardMulUpdate is a paid mutator transaction binding the contract method 0x46664064.
//
// Solidity: function rewardMulUpdate(uint256[] _newMulReward) returns()
func (_Gardener *GardenerSession) RewardMulUpdate(_newMulReward []*big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.RewardMulUpdate(&_Gardener.TransactOpts, _newMulReward)
}

// RewardMulUpdate is a paid mutator transaction binding the contract method 0x46664064.
//
// Solidity: function rewardMulUpdate(uint256[] _newMulReward) returns()
func (_Gardener *GardenerTransactorSession) RewardMulUpdate(_newMulReward []*big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.RewardMulUpdate(&_Gardener.TransactOpts, _newMulReward)
}

// RewardUpdate is a paid mutator transaction binding the contract method 0x3d479cd5.
//
// Solidity: function rewardUpdate(uint256 _newReward) returns()
func (_Gardener *GardenerTransactor) RewardUpdate(opts *bind.TransactOpts, _newReward *big.Int) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "rewardUpdate", _newReward)
}

// RewardUpdate is a paid mutator transaction binding the contract method 0x3d479cd5.
//
// Solidity: function rewardUpdate(uint256 _newReward) returns()
func (_Gardener *GardenerSession) RewardUpdate(_newReward *big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.RewardUpdate(&_Gardener.TransactOpts, _newReward)
}

// RewardUpdate is a paid mutator transaction binding the contract method 0x3d479cd5.
//
// Solidity: function rewardUpdate(uint256 _newReward) returns()
func (_Gardener *GardenerTransactorSession) RewardUpdate(_newReward *big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.RewardUpdate(&_Gardener.TransactOpts, _newReward)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_Gardener *GardenerTransactor) Set(opts *bind.TransactOpts, _pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "set", _pid, _allocPoint, _withUpdate)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_Gardener *GardenerSession) Set(_pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _Gardener.Contract.Set(&_Gardener.TransactOpts, _pid, _allocPoint, _withUpdate)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_Gardener *GardenerTransactorSession) Set(_pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _Gardener.Contract.Set(&_Gardener.TransactOpts, _pid, _allocPoint, _withUpdate)
}

// SetDevDepFee is a paid mutator transaction binding the contract method 0xfb075433.
//
// Solidity: function setDevDepFee(uint256 _devDepFees) returns()
func (_Gardener *GardenerTransactor) SetDevDepFee(opts *bind.TransactOpts, _devDepFees *big.Int) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "setDevDepFee", _devDepFees)
}

// SetDevDepFee is a paid mutator transaction binding the contract method 0xfb075433.
//
// Solidity: function setDevDepFee(uint256 _devDepFees) returns()
func (_Gardener *GardenerSession) SetDevDepFee(_devDepFees *big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.SetDevDepFee(&_Gardener.TransactOpts, _devDepFees)
}

// SetDevDepFee is a paid mutator transaction binding the contract method 0xfb075433.
//
// Solidity: function setDevDepFee(uint256 _devDepFees) returns()
func (_Gardener *GardenerTransactorSession) SetDevDepFee(_devDepFees *big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.SetDevDepFee(&_Gardener.TransactOpts, _devDepFees)
}

// SetDevFeeStage is a paid mutator transaction binding the contract method 0x8cc883ce.
//
// Solidity: function setDevFeeStage(uint256[] _devFees) returns()
func (_Gardener *GardenerTransactor) SetDevFeeStage(opts *bind.TransactOpts, _devFees []*big.Int) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "setDevFeeStage", _devFees)
}

// SetDevFeeStage is a paid mutator transaction binding the contract method 0x8cc883ce.
//
// Solidity: function setDevFeeStage(uint256[] _devFees) returns()
func (_Gardener *GardenerSession) SetDevFeeStage(_devFees []*big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.SetDevFeeStage(&_Gardener.TransactOpts, _devFees)
}

// SetDevFeeStage is a paid mutator transaction binding the contract method 0x8cc883ce.
//
// Solidity: function setDevFeeStage(uint256[] _devFees) returns()
func (_Gardener *GardenerTransactorSession) SetDevFeeStage(_devFees []*big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.SetDevFeeStage(&_Gardener.TransactOpts, _devFees)
}

// SetStageEnds is a paid mutator transaction binding the contract method 0x4af0e3e1.
//
// Solidity: function setStageEnds(uint256[] _blockEnds) returns()
func (_Gardener *GardenerTransactor) SetStageEnds(opts *bind.TransactOpts, _blockEnds []*big.Int) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "setStageEnds", _blockEnds)
}

// SetStageEnds is a paid mutator transaction binding the contract method 0x4af0e3e1.
//
// Solidity: function setStageEnds(uint256[] _blockEnds) returns()
func (_Gardener *GardenerSession) SetStageEnds(_blockEnds []*big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.SetStageEnds(&_Gardener.TransactOpts, _blockEnds)
}

// SetStageEnds is a paid mutator transaction binding the contract method 0x4af0e3e1.
//
// Solidity: function setStageEnds(uint256[] _blockEnds) returns()
func (_Gardener *GardenerTransactorSession) SetStageEnds(_blockEnds []*big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.SetStageEnds(&_Gardener.TransactOpts, _blockEnds)
}

// SetStageStarts is a paid mutator transaction binding the contract method 0x847bdaa4.
//
// Solidity: function setStageStarts(uint256[] _blockStarts) returns()
func (_Gardener *GardenerTransactor) SetStageStarts(opts *bind.TransactOpts, _blockStarts []*big.Int) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "setStageStarts", _blockStarts)
}

// SetStageStarts is a paid mutator transaction binding the contract method 0x847bdaa4.
//
// Solidity: function setStageStarts(uint256[] _blockStarts) returns()
func (_Gardener *GardenerSession) SetStageStarts(_blockStarts []*big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.SetStageStarts(&_Gardener.TransactOpts, _blockStarts)
}

// SetStageStarts is a paid mutator transaction binding the contract method 0x847bdaa4.
//
// Solidity: function setStageStarts(uint256[] _blockStarts) returns()
func (_Gardener *GardenerTransactorSession) SetStageStarts(_blockStarts []*big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.SetStageStarts(&_Gardener.TransactOpts, _blockStarts)
}

// SetUserDepFee is a paid mutator transaction binding the contract method 0x7c39c9c9.
//
// Solidity: function setUserDepFee(uint256 _usrDepFees) returns()
func (_Gardener *GardenerTransactor) SetUserDepFee(opts *bind.TransactOpts, _usrDepFees *big.Int) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "setUserDepFee", _usrDepFees)
}

// SetUserDepFee is a paid mutator transaction binding the contract method 0x7c39c9c9.
//
// Solidity: function setUserDepFee(uint256 _usrDepFees) returns()
func (_Gardener *GardenerSession) SetUserDepFee(_usrDepFees *big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.SetUserDepFee(&_Gardener.TransactOpts, _usrDepFees)
}

// SetUserDepFee is a paid mutator transaction binding the contract method 0x7c39c9c9.
//
// Solidity: function setUserDepFee(uint256 _usrDepFees) returns()
func (_Gardener *GardenerTransactorSession) SetUserDepFee(_usrDepFees *big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.SetUserDepFee(&_Gardener.TransactOpts, _usrDepFees)
}

// SetUserFeeStage is a paid mutator transaction binding the contract method 0x2a7b0107.
//
// Solidity: function setUserFeeStage(uint256[] _userFees) returns()
func (_Gardener *GardenerTransactor) SetUserFeeStage(opts *bind.TransactOpts, _userFees []*big.Int) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "setUserFeeStage", _userFees)
}

// SetUserFeeStage is a paid mutator transaction binding the contract method 0x2a7b0107.
//
// Solidity: function setUserFeeStage(uint256[] _userFees) returns()
func (_Gardener *GardenerSession) SetUserFeeStage(_userFees []*big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.SetUserFeeStage(&_Gardener.TransactOpts, _userFees)
}

// SetUserFeeStage is a paid mutator transaction binding the contract method 0x2a7b0107.
//
// Solidity: function setUserFeeStage(uint256[] _userFees) returns()
func (_Gardener *GardenerTransactorSession) SetUserFeeStage(_userFees []*big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.SetUserFeeStage(&_Gardener.TransactOpts, _userFees)
}

// StarblockUpdate is a paid mutator transaction binding the contract method 0x24b95c3e.
//
// Solidity: function starblockUpdate(uint256 _newstarblock) returns()
func (_Gardener *GardenerTransactor) StarblockUpdate(opts *bind.TransactOpts, _newstarblock *big.Int) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "starblockUpdate", _newstarblock)
}

// StarblockUpdate is a paid mutator transaction binding the contract method 0x24b95c3e.
//
// Solidity: function starblockUpdate(uint256 _newstarblock) returns()
func (_Gardener *GardenerSession) StarblockUpdate(_newstarblock *big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.StarblockUpdate(&_Gardener.TransactOpts, _newstarblock)
}

// StarblockUpdate is a paid mutator transaction binding the contract method 0x24b95c3e.
//
// Solidity: function starblockUpdate(uint256 _newstarblock) returns()
func (_Gardener *GardenerTransactorSession) StarblockUpdate(_newstarblock *big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.StarblockUpdate(&_Gardener.TransactOpts, _newstarblock)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Gardener *GardenerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Gardener *GardenerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Gardener.Contract.TransferOwnership(&_Gardener.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Gardener *GardenerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Gardener.Contract.TransferOwnership(&_Gardener.TransactOpts, newOwner)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_Gardener *GardenerTransactor) UpdatePool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "updatePool", _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_Gardener *GardenerSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.UpdatePool(&_Gardener.TransactOpts, _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_Gardener *GardenerTransactorSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _Gardener.Contract.UpdatePool(&_Gardener.TransactOpts, _pid)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0ad58d2f.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount, address _ref) returns()
func (_Gardener *GardenerTransactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int, _ref common.Address) (*types.Transaction, error) {
	return _Gardener.contract.Transact(opts, "withdraw", _pid, _amount, _ref)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0ad58d2f.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount, address _ref) returns()
func (_Gardener *GardenerSession) Withdraw(_pid *big.Int, _amount *big.Int, _ref common.Address) (*types.Transaction, error) {
	return _Gardener.Contract.Withdraw(&_Gardener.TransactOpts, _pid, _amount, _ref)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0ad58d2f.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount, address _ref) returns()
func (_Gardener *GardenerTransactorSession) Withdraw(_pid *big.Int, _amount *big.Int, _ref common.Address) (*types.Transaction, error) {
	return _Gardener.Contract.Withdraw(&_Gardener.TransactOpts, _pid, _amount, _ref)
}

// GardenerDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Gardener contract.
type GardenerDepositIterator struct {
	Event *GardenerDeposit // Event containing the contract specifics and raw log

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
func (it *GardenerDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GardenerDeposit)
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
		it.Event = new(GardenerDeposit)
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
func (it *GardenerDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GardenerDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GardenerDeposit represents a Deposit event raised by the Gardener contract.
type GardenerDeposit struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_Gardener *GardenerFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*GardenerDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Gardener.contract.FilterLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &GardenerDepositIterator{contract: _Gardener.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_Gardener *GardenerFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *GardenerDeposit, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Gardener.contract.WatchLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GardenerDeposit)
				if err := _Gardener.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_Gardener *GardenerFilterer) ParseDeposit(log types.Log) (*GardenerDeposit, error) {
	event := new(GardenerDeposit)
	if err := _Gardener.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GardenerEmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the Gardener contract.
type GardenerEmergencyWithdrawIterator struct {
	Event *GardenerEmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *GardenerEmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GardenerEmergencyWithdraw)
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
		it.Event = new(GardenerEmergencyWithdraw)
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
func (it *GardenerEmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GardenerEmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GardenerEmergencyWithdraw represents a EmergencyWithdraw event raised by the Gardener contract.
type GardenerEmergencyWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Gardener *GardenerFilterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*GardenerEmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Gardener.contract.FilterLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &GardenerEmergencyWithdrawIterator{contract: _Gardener.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Gardener *GardenerFilterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *GardenerEmergencyWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Gardener.contract.WatchLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GardenerEmergencyWithdraw)
				if err := _Gardener.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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

// ParseEmergencyWithdraw is a log parse operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Gardener *GardenerFilterer) ParseEmergencyWithdraw(log types.Log) (*GardenerEmergencyWithdraw, error) {
	event := new(GardenerEmergencyWithdraw)
	if err := _Gardener.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GardenerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Gardener contract.
type GardenerOwnershipTransferredIterator struct {
	Event *GardenerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GardenerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GardenerOwnershipTransferred)
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
		it.Event = new(GardenerOwnershipTransferred)
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
func (it *GardenerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GardenerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GardenerOwnershipTransferred represents a OwnershipTransferred event raised by the Gardener contract.
type GardenerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Gardener *GardenerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GardenerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Gardener.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GardenerOwnershipTransferredIterator{contract: _Gardener.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Gardener *GardenerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GardenerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Gardener.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GardenerOwnershipTransferred)
				if err := _Gardener.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Gardener *GardenerFilterer) ParseOwnershipTransferred(log types.Log) (*GardenerOwnershipTransferred, error) {
	event := new(GardenerOwnershipTransferred)
	if err := _Gardener.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GardenerSendGovernanceTokenRewardIterator is returned from FilterSendGovernanceTokenReward and is used to iterate over the raw logs and unpacked data for SendGovernanceTokenReward events raised by the Gardener contract.
type GardenerSendGovernanceTokenRewardIterator struct {
	Event *GardenerSendGovernanceTokenReward // Event containing the contract specifics and raw log

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
func (it *GardenerSendGovernanceTokenRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GardenerSendGovernanceTokenReward)
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
		it.Event = new(GardenerSendGovernanceTokenReward)
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
func (it *GardenerSendGovernanceTokenRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GardenerSendGovernanceTokenRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GardenerSendGovernanceTokenReward represents a SendGovernanceTokenReward event raised by the Gardener contract.
type GardenerSendGovernanceTokenReward struct {
	User       common.Address
	Pid        *big.Int
	Amount     *big.Int
	LockAmount *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSendGovernanceTokenReward is a free log retrieval operation binding the contract event 0x3887f2857beaaf367eb618dfb5e22c1ebd74425affb0602c2e9fe126e3f860eb.
//
// Solidity: event SendGovernanceTokenReward(address indexed user, uint256 indexed pid, uint256 amount, uint256 lockAmount)
func (_Gardener *GardenerFilterer) FilterSendGovernanceTokenReward(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*GardenerSendGovernanceTokenRewardIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Gardener.contract.FilterLogs(opts, "SendGovernanceTokenReward", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &GardenerSendGovernanceTokenRewardIterator{contract: _Gardener.contract, event: "SendGovernanceTokenReward", logs: logs, sub: sub}, nil
}

// WatchSendGovernanceTokenReward is a free log subscription operation binding the contract event 0x3887f2857beaaf367eb618dfb5e22c1ebd74425affb0602c2e9fe126e3f860eb.
//
// Solidity: event SendGovernanceTokenReward(address indexed user, uint256 indexed pid, uint256 amount, uint256 lockAmount)
func (_Gardener *GardenerFilterer) WatchSendGovernanceTokenReward(opts *bind.WatchOpts, sink chan<- *GardenerSendGovernanceTokenReward, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Gardener.contract.WatchLogs(opts, "SendGovernanceTokenReward", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GardenerSendGovernanceTokenReward)
				if err := _Gardener.contract.UnpackLog(event, "SendGovernanceTokenReward", log); err != nil {
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

// ParseSendGovernanceTokenReward is a log parse operation binding the contract event 0x3887f2857beaaf367eb618dfb5e22c1ebd74425affb0602c2e9fe126e3f860eb.
//
// Solidity: event SendGovernanceTokenReward(address indexed user, uint256 indexed pid, uint256 amount, uint256 lockAmount)
func (_Gardener *GardenerFilterer) ParseSendGovernanceTokenReward(log types.Log) (*GardenerSendGovernanceTokenReward, error) {
	event := new(GardenerSendGovernanceTokenReward)
	if err := _Gardener.contract.UnpackLog(event, "SendGovernanceTokenReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GardenerWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Gardener contract.
type GardenerWithdrawIterator struct {
	Event *GardenerWithdraw // Event containing the contract specifics and raw log

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
func (it *GardenerWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GardenerWithdraw)
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
		it.Event = new(GardenerWithdraw)
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
func (it *GardenerWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GardenerWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GardenerWithdraw represents a Withdraw event raised by the Gardener contract.
type GardenerWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Gardener *GardenerFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*GardenerWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Gardener.contract.FilterLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &GardenerWithdrawIterator{contract: _Gardener.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Gardener *GardenerFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *GardenerWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Gardener.contract.WatchLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GardenerWithdraw)
				if err := _Gardener.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Gardener *GardenerFilterer) ParseWithdraw(log types.Log) (*GardenerWithdraw, error) {
	event := new(GardenerWithdraw)
	if err := _Gardener.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
