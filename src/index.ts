import {
  BigNumber,
  Contract,
  ethers,
  providers,
  utils,
} from "ethers"

import { MASTERCHEF_PARTIAL_ABI } from './abi/JoeMasterChef';

const AVAX_MAINNET_RPC = 'https://api.avax.network/ext/bc/C/rpc'
const JOE_CHEFS = [
  '0xd6a4f121ca35509af06a0be99093d08462f53052',
  '0x188bED1968b795d5c9022F6a0bb5931Ac4c18F00'
]

type PoolInfo = {
  readonly lpToken: string
  readonly allocPoint: bigint
  readonly lastRewardTimestamp: bigint
  readonly accJoePerShare: bigint
  readonly rewarder: string
}

const main = async(): Promise<any> => {
  const address: string = (process.env.ADDRESS as string)
  const provider: providers.JsonRpcProvider = new providers.JsonRpcProvider(AVAX_MAINNET_RPC)

  const balance: BigNumber = await provider.getBalance(address)
  const formatted: string = utils.formatEther(balance)
  console.log(`Current Balance: ${formatted} AVAX`)

  await Promise.all(JOE_CHEFS.map(async(joeAddress): Promise<any> => {
    const chef: Contract = new ethers.Contract(joeAddress, MASTERCHEF_PARTIAL_ABI, provider)
    const poolLength : BigNumber  = await chef.poolLength()

    /* eslint functional/no-loop-statement: 0 */
    /* eslint functional/no-let: 0 */
    for (let i = 0; i < poolLength.toNumber(); i++) {
      const pool : PoolInfo = await chef.poolInfo(i)
      const [balance, rewards] = await chef.userInfo(i, address) as readonly [BigNumber, BigNumber]
      if (!balance.isZero()) {
        console.log(`LP balance=[${utils.formatEther(balance)}] token=[${pool.lpToken}] ${rewards}`)
      }
    }
  }))
}

main()
.then(() => process.exit(0))
.catch(error => {
  console.error(error)
  process.exit(1)
})
