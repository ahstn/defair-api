import { ContractInterface } from "ethers";

export const TOKEN_PAIR_ABI: ContractInterface = [
  {
      "inputs": [
      {
          "internalType": "address",
          "name": "owner",
          "type": "address"
      }
      ],
      "name": "balanceOf",
      "outputs": [
      {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
      }
      ],
      "stateMutability": "view",
      "type": "function"
  },
  {
      "inputs": [],
      "name": "price0CumulativeLast",
      "outputs": [
      {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
      }
      ],
      "stateMutability": "view",
      "type": "function"
  },
  {
      "inputs": [],
      "name": "price1CumulativeLast",
      "outputs": [
      {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
      }
      ],
      "stateMutability": "view",
      "type": "function"
  },
  {
      "inputs": [],
      "name": "symbol",
      "outputs": [
      {
          "internalType": "string",
          "name": "",
          "type": "string"
      }
      ],
      "stateMutability": "pure",
      "type": "function"
  },
  {
      "inputs": [],
      "name": "token0",
      "outputs": [
      {
          "internalType": "address",
          "name": "",
          "type": "address"
      }
      ],
      "stateMutability": "view",
      "type": "function"
  },
  {
      "inputs": [],
      "name": "token1",
      "outputs": [
      {
          "internalType": "address",
          "name": "",
          "type": "address"
      }
      ],
      "stateMutability": "view",
      "type": "function"
  },
  {
      "inputs": [],
      "name": "totalSupply",
      "outputs": [
      {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
      }
      ],
      "stateMutability": "view",
      "type": "function"
  },
]
