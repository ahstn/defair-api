import { Module } from '@nestjs/common'

import { PoolsModule } from './pools/pools.module'

@Module({
  imports: [PoolsModule],
})
export class AppModule {}
