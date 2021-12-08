import { ContractInterface } from "ethers";

export const MASTERCHEF_PARTIAL_ABI: ContractInterface = [
  {
      "inputs": [
      {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
      }
      ],
      "name": "poolInfo",
      "outputs": [
      {
          "internalType": "contract IERC20",
          "name": "lpToken",
          "type": "address"
      },
      {
          "internalType": "uint256",
          "name": "accJoePerShare",
          "type": "uint256"
      },
      {
          "internalType": "uint256",
          "name": "lastRewardTimestamp",
          "type": "uint256"
      },
      {
          "internalType": "uint256",
          "name": "allocPoint",
          "type": "uint256"
      },
      {
          "internalType": "contract IRewarder",
          "name": "rewarder",
          "type": "address"
      }
      ],
      "stateMutability": "view",
      "type": "function"
  },
  {
      "inputs": [],
      "name": "poolLength",
      "outputs": [
      {
          "internalType": "uint256",
          "name": "pools",
          "type": "uint256"
      }
      ],
      "stateMutability": "view",
      "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      },
      {
        "internalType": "address",
        "name": "",
        "type": "address"
      }
    ],
    "name": "userInfo",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "amount",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "rewardDebt",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  }
]
