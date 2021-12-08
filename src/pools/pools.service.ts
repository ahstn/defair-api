import { Injectable } from '@nestjs/common'

@Injectable()
export class PoolsService {
  findAll() {
    return `This action returns all pools`;
  }

  findOne(id: number) {
    return `This action returns a #${id} pool`;
  }
}
