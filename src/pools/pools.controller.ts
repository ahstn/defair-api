import { Controller, Get, Param } from '@nestjs/common'

import { PoolsService } from './pools.service'

@Controller('pools')
export class PoolsController {
  constructor(private readonly poolsService: PoolsService) {}

  @Get()
  findAll() {
    return this.poolsService.findAll();
  }

  @Get(':id')
  findOne(@Param('id') id: string) {
    return this.poolsService.findOne(+id);
  }
}
