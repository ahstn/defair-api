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

// JoeChefMetadata contains the metadata for TraderJoe's Chef contract.
var JoeChefMetadata = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"accJoePerShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"contractIRewarder\",\"name\":\"rewarder\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pools\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DFKChefMetadata contains the metadata for Defi Kingdoms' Gardener contract.
var DFKChefMetadata = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_govToken\",\"internalType\":\"contractJewelToken\",\"type\":\"address\"},{\"type\":\"address\",\"internalType\":\"address\",\"name\":\"_devaddr\"},{\"name\":\"_liquidityaddr\",\"internalType\":\"address\",\"type\":\"address\"},{\"name\":\"_comfundaddr\",\"internalType\":\"address\",\"type\":\"address\"},{\"name\":\"_founderaddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardPerBlock\",\"type\":\"uint256\"},{\"type\":\"uint256\",\"internalType\":\"uint256\",\"name\":\"_startBlock\"},{\"internalType\":\"uint256\",\"name\":\"_halvingAfterBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_userDepFee\",\"type\":\"uint256\"},{\"name\":\"_devDepFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_rewardMultiplier\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"type\":\"uint256[]\",\"name\":\"_blockDeltaStartStage\",\"internalType\":\"uint256[]\"},{\"type\":\"uint256[]\",\"internalType\":\"uint256[]\",\"name\":\"_blockDeltaEndStage\"},{\"type\":\"uint256[]\",\"name\":\"_userFeeStage\",\"internalType\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_devFeeStage\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\"},{\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"uint256\",\"indexed\":true,\"name\":\"pid\",\"internalType\":\"uint256\"},{\"indexed\":false,\"type\":\"uint256\",\"internalType\":\"uint256\",\"name\":\"amount\"}],\"name\":\"Deposit\",\"type\":\"event\",\"anonymous\":false},{\"anonymous\":false,\"name\":\"EmergencyWithdraw\",\"type\":\"event\",\"inputs\":[{\"name\":\"user\",\"indexed\":true,\"internalType\":\"address\",\"type\":\"address\"},{\"name\":\"pid\",\"type\":\"uint256\",\"internalType\":\"uint256\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"amount\",\"indexed\":false,\"internalType\":\"uint256\"}]},{\"inputs\":[{\"internalType\":\"address\",\"type\":\"address\",\"indexed\":true,\"name\":\"previousOwner\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\",\"anonymous\":false},{\"anonymous\":false,\"type\":\"event\",\"name\":\"SendGovernanceTokenReward\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"pid\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false},{\"indexed\":false,\"type\":\"uint256\",\"internalType\":\"uint256\",\"name\":\"lockAmount\"}]},{\"inputs\":[{\"name\":\"user\",\"internalType\":\"address\",\"indexed\":true,\"type\":\"address\"},{\"indexed\":true,\"name\":\"pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"internalType\":\"uint256\",\"type\":\"uint256\"}],\"anonymous\":false,\"name\":\"Withdraw\",\"type\":\"event\"},{\"name\":\"FINISH_BONUS_AT_BLOCK\",\"stateMutability\":\"view\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"constant\":true,\"signature\":\"0x980c2a98\"},{\"type\":\"function\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"HALVING_AT_BLOCK\",\"stateMutability\":\"view\",\"inputs\":[{\"type\":\"uint256\",\"internalType\":\"uint256\",\"name\":\"\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"name\":\"\",\"internalType\":\"uint256\",\"type\":\"uint256\"}],\"name\":\"PERCENT_FOR_COM\",\"inputs\":[],\"constant\":true,\"signature\":\"0xa02306f9\"},{\"stateMutability\":\"view\",\"inputs\":[],\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"type\":\"function\",\"name\":\"PERCENT_FOR_DEV\",\"constant\":true,\"signature\":\"0xed9bdeda\"},{\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"type\":\"function\",\"inputs\":[],\"stateMutability\":\"view\",\"name\":\"PERCENT_FOR_FOUNDERS\",\"constant\":true,\"signature\":\"0xc6929e53\"},{\"name\":\"PERCENT_FOR_LP\",\"outputs\":[{\"type\":\"uint256\",\"internalType\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"type\":\"function\",\"stateMutability\":\"view\",\"constant\":true,\"signature\":\"0x0a67d518\"},{\"name\":\"PERCENT_LOCK_BONUS_REWARD\",\"inputs\":[{\"name\":\"\",\"internalType\":\"uint256\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"type\":\"function\",\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REWARD_MULTIPLIER\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\",\"inputs\":[{\"name\":\"\",\"internalType\":\"uint256\",\"type\":\"uint256\"}]},{\"name\":\"REWARD_PER_BLOCK\",\"inputs\":[],\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"internalType\":\"uint256\",\"name\":\"\"}],\"constant\":true,\"signature\":\"0x975532dc\"},{\"stateMutability\":\"view\",\"name\":\"START_BLOCK\",\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"constant\":true,\"signature\":\"0x39b3e826\"},{\"name\":\"addAuthorized\",\"inputs\":[{\"internalType\":\"address\",\"name\":\"_toAdd\",\"type\":\"address\"}],\"outputs\":[],\"type\":\"function\",\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"authorized\",\"stateMutability\":\"view\",\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"outputs\":[{\"type\":\"bool\",\"internalType\":\"bool\",\"name\":\"\"}]},{\"inputs\":[{\"type\":\"uint256\",\"internalType\":\"uint256\",\"name\":\"\"}],\"name\":\"blockDeltaEndStage\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"type\":\"function\",\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blockDeltaStartStage\",\"stateMutability\":\"view\",\"inputs\":[{\"type\":\"uint256\",\"internalType\":\"uint256\",\"name\":\"\"}],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"stateMutability\":\"view\",\"name\":\"comfundaddr\",\"constant\":true,\"signature\":\"0x3c9d9267\"},{\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"devDepFee\",\"type\":\"function\",\"inputs\":[],\"constant\":true,\"signature\":\"0xc56a10ff\"},{\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"name\":\"devFeeStage\",\"stateMutability\":\"view\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"stateMutability\":\"view\",\"name\":\"devaddr\",\"inputs\":[],\"outputs\":[{\"type\":\"address\",\"internalType\":\"address\",\"name\":\"\"}],\"type\":\"function\",\"constant\":true,\"signature\":\"0xd49e77cd\"},{\"inputs\":[],\"stateMutability\":\"view\",\"name\":\"founderaddr\",\"type\":\"function\",\"outputs\":[{\"internalType\":\"address\",\"type\":\"address\",\"name\":\"\"}],\"constant\":true,\"signature\":\"0xec12173d\"},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"internalType\":\"contractJewelToken\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"govToken\",\"inputs\":[],\"constant\":true,\"signature\":\"0x05268cff\"},{\"type\":\"function\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"liquidityaddr\",\"inputs\":[],\"stateMutability\":\"view\",\"constant\":true,\"signature\":\"0x22a376b0\"},{\"name\":\"owner\",\"type\":\"function\",\"inputs\":[],\"stateMutability\":\"view\",\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"constant\":true,\"signature\":\"0x8da5cb5b\"},{\"type\":\"function\",\"name\":\"poolExistence\",\"stateMutability\":\"view\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"inputs\":[{\"name\":\"\",\"internalType\":\"contractIERC20\",\"type\":\"address\"}]},{\"outputs\":[{\"internalType\":\"uint256\",\"type\":\"uint256\",\"name\":\"\"}],\"stateMutability\":\"view\",\"name\":\"poolId1\",\"type\":\"function\",\"inputs\":[{\"name\":\"\",\"internalType\":\"address\",\"type\":\"address\"}]},{\"inputs\":[{\"internalType\":\"uint256\",\"type\":\"uint256\",\"name\":\"\"}],\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"name\":\"lastRewardBlock\",\"internalType\":\"uint256\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"type\":\"uint256\",\"name\":\"accGovTokenPerShare\"}],\"type\":\"function\",\"stateMutability\":\"view\",\"name\":\"poolInfo\"},{\"name\":\"removeAuthorized\",\"stateMutability\":\"nonpayable\",\"inputs\":[{\"type\":\"address\",\"name\":\"_toRemove\",\"internalType\":\"address\"}],\"outputs\":[],\"type\":\"function\"},{\"stateMutability\":\"nonpayable\",\"inputs\":[],\"outputs\":[],\"type\":\"function\",\"name\":\"renounceOwnership\"},{\"outputs\":[{\"type\":\"uint256\",\"internalType\":\"uint256\",\"name\":\"\"}],\"stateMutability\":\"view\",\"inputs\":[],\"type\":\"function\",\"name\":\"totalAllocPoint\",\"constant\":true,\"signature\":\"0x17caf6f1\"},{\"type\":\"function\",\"outputs\":[],\"name\":\"transferOwnership\",\"inputs\":[{\"type\":\"address\",\"internalType\":\"address\",\"name\":\"newOwner\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"name\":\"\",\"internalType\":\"address\",\"type\":\"address\"}],\"inputs\":[],\"name\":\"usdOracle\",\"constant\":true,\"signature\":\"0xc8a4271f\"},{\"stateMutability\":\"view\",\"name\":\"userDepFee\",\"type\":\"function\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"inputs\":[],\"constant\":true,\"signature\":\"0x82796e98\"},{\"name\":\"userFeeStage\",\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"name\":\"\",\"internalType\":\"uint256\",\"type\":\"uint256\"}],\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"userGlobalInfo\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"type\":\"uint256\",\"name\":\"globalAmount\",\"internalType\":\"uint256\"},{\"internalType\":\"uint256\",\"type\":\"uint256\",\"name\":\"totalReferals\"},{\"name\":\"globalRefAmount\",\"internalType\":\"uint256\",\"type\":\"uint256\"}],\"constant\":true,\"signature\":\"0xd9608d8a\"},{\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rewardDebt\",\"internalType\":\"uint256\",\"type\":\"uint256\"},{\"name\":\"rewardDebtAtBlock\",\"internalType\":\"uint256\",\"type\":\"uint256\"},{\"type\":\"uint256\",\"internalType\":\"uint256\",\"name\":\"lastWithdrawBlock\"},{\"internalType\":\"uint256\",\"name\":\"firstDepositBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockdelta\",\"type\":\"uint256\"},{\"name\":\"lastDepositBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"name\":\"userInfo\",\"type\":\"function\",\"stateMutability\":\"view\",\"inputs\":[{\"name\":\"\",\"internalType\":\"uint256\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"constant\":true,\"signature\":\"0x93f1a40b\"},{\"outputs\":[{\"type\":\"uint256\",\"internalType\":\"uint256\",\"name\":\"\"}],\"name\":\"poolLength\",\"stateMutability\":\"view\",\"inputs\":[],\"type\":\"function\",\"constant\":true,\"signature\":\"0x081e3eda\"},{\"outputs\":[],\"stateMutability\":\"nonpayable\",\"name\":\"add\",\"inputs\":[{\"name\":\"_allocPoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"internalType\":\"contractIERC20\",\"name\":\"_lpToken\"},{\"type\":\"bool\",\"name\":\"_withUpdate\",\"internalType\":\"bool\"}],\"type\":\"function\"},{\"type\":\"function\",\"outputs\":[],\"name\":\"set\",\"stateMutability\":\"nonpayable\",\"inputs\":[{\"internalType\":\"uint256\",\"type\":\"uint256\",\"name\":\"_pid\"},{\"name\":\"_allocPoint\",\"internalType\":\"uint256\",\"type\":\"uint256\"},{\"name\":\"_withUpdate\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"outputs\":[],\"name\":\"massUpdatePools\",\"type\":\"function\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"name\":\"updatePool\",\"inputs\":[{\"type\":\"uint256\",\"internalType\":\"uint256\",\"name\":\"_pid\"}],\"outputs\":[]},{\"name\":\"getMultiplier\",\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"inputs\":[{\"name\":\"_from\",\"internalType\":\"uint256\",\"type\":\"uint256\"},{\"name\":\"_to\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"name\":\"getLockPercentage\",\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"type\":\"uint256\",\"name\":\"_to\"}]},{\"name\":\"getPoolReward\",\"inputs\":[{\"type\":\"uint256\",\"internalType\":\"uint256\",\"name\":\"_from\"},{\"name\":\"_to\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"}],\"type\":\"function\",\"outputs\":[{\"internalType\":\"uint256\",\"type\":\"uint256\",\"name\":\"forDev\"},{\"name\":\"forFarmer\",\"internalType\":\"uint256\",\"type\":\"uint256\"},{\"name\":\"forLP\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"forCom\",\"internalType\":\"uint256\",\"type\":\"uint256\"},{\"name\":\"forFounders\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"stateMutability\":\"view\",\"inputs\":[{\"type\":\"uint256\",\"internalType\":\"uint256\",\"name\":\"_pid\"},{\"name\":\"_user\",\"internalType\":\"address\",\"type\":\"address\"}],\"outputs\":[{\"type\":\"uint256\",\"internalType\":\"uint256\",\"name\":\"\"}],\"name\":\"pendingReward\",\"type\":\"function\"},{\"name\":\"claimRewards\",\"stateMutability\":\"nonpayable\",\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_pids\",\"type\":\"uint256[]\"}],\"type\":\"function\",\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"outputs\":[],\"inputs\":[{\"name\":\"_pid\",\"internalType\":\"uint256\",\"type\":\"uint256\"}],\"name\":\"claimReward\",\"type\":\"function\"},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"internalType\":\"uint256\",\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getGlobalAmount\"},{\"inputs\":[{\"name\":\"_user\",\"internalType\":\"address\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getGlobalRefAmount\"},{\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"type\":\"function\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"}],\"name\":\"getTotalRefs\",\"stateMutability\":\"view\"},{\"inputs\":[{\"type\":\"address\",\"name\":\"_user\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"_user2\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"type\":\"function\",\"stateMutability\":\"view\",\"name\":\"getRefValueOf\"},{\"name\":\"deposit\",\"stateMutability\":\"nonpayable\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_ref\",\"internalType\":\"address\",\"type\":\"address\"}],\"type\":\"function\",\"outputs\":[]},{\"inputs\":[{\"name\":\"_pid\",\"internalType\":\"uint256\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_ref\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"name\":\"withdraw\",\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"emergencyWithdraw\",\"type\":\"function\",\"outputs\":[],\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"inputs\":[{\"name\":\"_devaddr\",\"internalType\":\"address\",\"type\":\"address\"}],\"name\":\"dev\",\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"stateMutability\":\"nonpayable\",\"inputs\":[{\"name\":\"_newFinish\",\"internalType\":\"uint256\",\"type\":\"uint256\"}],\"name\":\"bonusFinishUpdate\",\"type\":\"function\",\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_newHalving\",\"type\":\"uint256[]\"}],\"type\":\"function\",\"outputs\":[],\"name\":\"halvingUpdate\"},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"outputs\":[],\"name\":\"lpUpdate\",\"inputs\":[{\"name\":\"_newLP\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"comUpdate\",\"type\":\"function\",\"inputs\":[{\"name\":\"_newCom\",\"internalType\":\"address\",\"type\":\"address\"}]},{\"outputs\":[],\"type\":\"function\",\"inputs\":[{\"name\":\"_newFounder\",\"internalType\":\"address\",\"type\":\"address\"}],\"name\":\"founderUpdate\",\"stateMutability\":\"nonpayable\"},{\"name\":\"rewardUpdate\",\"type\":\"function\",\"inputs\":[{\"name\":\"_newReward\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"stateMutability\":\"nonpayable\",\"name\":\"rewardMulUpdate\",\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_newMulReward\",\"type\":\"uint256[]\"}],\"type\":\"function\",\"outputs\":[]},{\"type\":\"function\",\"name\":\"lockUpdate\",\"inputs\":[{\"name\":\"_newlock\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"outputs\":[]},{\"inputs\":[{\"type\":\"uint256\",\"name\":\"_newdevlock\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\",\"name\":\"lockdevUpdate\",\"type\":\"function\"},{\"stateMutability\":\"nonpayable\",\"outputs\":[],\"type\":\"function\",\"name\":\"locklpUpdate\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_newlplock\",\"internalType\":\"uint256\"}]},{\"name\":\"lockcomUpdate\",\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"outputs\":[],\"inputs\":[{\"internalType\":\"uint256\",\"type\":\"uint256\",\"name\":\"_newcomlock\"}]},{\"type\":\"function\",\"outputs\":[],\"name\":\"lockfounderUpdate\",\"stateMutability\":\"nonpayable\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_newfounderlock\",\"internalType\":\"uint256\"}]},{\"name\":\"starblockUpdate\",\"outputs\":[],\"type\":\"function\",\"inputs\":[{\"name\":\"_newstarblock\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"stateMutability\":\"view\",\"type\":\"function\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getNewRewardPerBlock\",\"inputs\":[{\"internalType\":\"uint256\",\"type\":\"uint256\",\"name\":\"pid1\"}]},{\"name\":\"userDelta\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"reviseWithdraw\",\"stateMutability\":\"nonpayable\",\"inputs\":[{\"name\":\"_pid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_block\",\"internalType\":\"uint256\"}],\"type\":\"function\",\"outputs\":[]},{\"name\":\"reviseDeposit\",\"inputs\":[{\"name\":\"_pid\",\"internalType\":\"uint256\",\"type\":\"uint256\"},{\"type\":\"address\",\"internalType\":\"address\",\"name\":\"_user\"},{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"outputs\":[]},{\"outputs\":[],\"inputs\":[{\"name\":\"_blockStarts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"setStageStarts\"},{\"inputs\":[{\"name\":\"_blockEnds\",\"internalType\":\"uint256[]\",\"type\":\"uint256[]\"}],\"name\":\"setStageEnds\",\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"inputs\":[{\"name\":\"_userFees\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"type\":\"function\",\"outputs\":[],\"name\":\"setUserFeeStage\"},{\"name\":\"setDevFeeStage\",\"inputs\":[{\"internalType\":\"uint256[]\",\"type\":\"uint256[]\",\"name\":\"_devFees\"}],\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[]},{\"name\":\"setDevDepFee\",\"stateMutability\":\"nonpayable\",\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_devDepFees\",\"type\":\"uint256\"}],\"outputs\":[],\"type\":\"function\"},{\"type\":\"function\",\"outputs\":[],\"inputs\":[{\"internalType\":\"uint256\",\"type\":\"uint256\",\"name\":\"_usrDepFees\"}],\"name\":\"setUserDepFee\",\"stateMutability\":\"nonpayable\"},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"name\":\"reclaimTokenOwnership\"}]",
}

// Chef is an auto generated Go binding around an Ethereum contract.
type Chef struct {
	ChefCaller
}

type ChefCaller struct {
	contract *bind.BoundContract
}

// ChefPoolInfo houses the data returned by PoolInfo
type ChefPoolInfo struct {
	// LpToken is the address of the Liquidity Pool Token.
	LpToken common.Address
	// LastRewardTimestamp is the last timestamp for reward distribution.
	LastRewardTimestamp *big.Int
	// AllocPoint is how many allocation points are assigned to this pool. (Token to distribute per second).
	AllocPoint *big.Int
	// Rewarder is the address of the Rewarding contract.
	Rewarder common.Address
	// AccJoePerShare is the accumulated reward token per LP share.
	AccJoePerShare *big.Int
}

// ChefUserInfo houses the data returned by UserInfo
type ChefUserInfo struct {
	Amount     *big.Int
	RewardDebt *big.Int
}

// NewChef creates a new instance of Chef, bound to a specific deployed contract.
func NewChef(address common.Address, backend bind.ContractBackend) (*Chef, error) {
	contract, err := bindChef(address, JoeChefMetadata.ABI, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Chef{ChefCaller: ChefCaller{contract: contract}}, nil
}

func NewChefWithABI(address common.Address, abiString string, backend bind.ContractBackend) (*Chef, error) {
	contract, err := bindChef(address, abiString, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Chef{ChefCaller: ChefCaller{contract: contract}}, nil
}

// bindChef binds a generic wrapper to an already deployed contract.
func bindChef(
	address common.Address,
	abiString string,
	caller bind.ContractCaller,
	transactor bind.ContractTransactor,
	filterer bind.ContractFilterer,
) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(abiString))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
func (_Chef *ChefCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Chef.contract.Call(opts, &out, "poolLength")
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
func (_Chef *ChefCaller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (ChefPoolInfo, error) {
	var out []interface{}
	err := _Chef.contract.Call(opts, &out, "poolInfo", arg0)
	if err != nil {
		return ChefPoolInfo{}, err
	}

	if len(out) > 4 {
		return ChefPoolInfo{
			LpToken:             *abi.ConvertType(out[0], new(common.Address)).(*common.Address),
			AccJoePerShare:      *abi.ConvertType(out[1], new(*big.Int)).(**big.Int),
			LastRewardTimestamp: *abi.ConvertType(out[2], new(*big.Int)).(**big.Int),
			AllocPoint:          *abi.ConvertType(out[3], new(*big.Int)).(**big.Int),
			Rewarder:            *abi.ConvertType(out[4], new(common.Address)).(*common.Address),
		}, err
	}

	return ChefPoolInfo{
		LpToken:             *abi.ConvertType(out[0], new(common.Address)).(*common.Address),
		LastRewardTimestamp: *abi.ConvertType(out[2], new(*big.Int)).(**big.Int),
		AllocPoint:          *abi.ConvertType(out[1], new(*big.Int)).(**big.Int),
		AccJoePerShare:      *abi.ConvertType(out[3], new(*big.Int)).(**big.Int),
	}, nil
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
func (_Chef *ChefCaller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (ChefUserInfo, error) {
	var out []interface{}
	err := _Chef.contract.Call(opts, &out, "userInfo", arg0, arg1)
	if err != nil {
		return ChefUserInfo{}, err
	}

	return ChefUserInfo{
		Amount:     *abi.ConvertType(out[0], new(*big.Int)).(**big.Int),
		RewardDebt: *abi.ConvertType(out[1], new(*big.Int)).(**big.Int),
	}, nil
}