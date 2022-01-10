# defair

<p align="center" markdown="1">
  <img src=".github/logo-banner.png" alt="logo"/>
  <a href="https://github.com/ahstn/defair-api/actions">
    <img src="https://github.com/ahstn/defair-api/actions/workflows/main.yml/badge.svg?branch=main" alt="Build Status" />
  </a>
  <a href="https://codecov.io/gh/ahstn/defair-api">
    <img src="https://codecov.io/gh/ahstn/defair-api/branch/main/graph/badge.svg?token=X6OFD7PLEL" />
  </a>
</p>

> Aiming to make navigating & tracking decentralised finance fair for all.

**NB:** Currently learning & experimenting with Rust to investigate it's performance benefits, and due to the increase of it's usage in Crypto. If it works out better, I'll maybe keep some of the code in this repo for a single service, i.e. aggregating multiple token lists across chains to provide token data and prices. 

## Contributing
Building & running Go application:
```bash
# fetch dependencies
go mod download

# run server
go run ./cmd/defair/defair.go server

# fetch liquidity pools
# curl -s "localhost:9000/pools/${YOUR ADDRESS}"
```

Adding a new Smart Contract or ABI:
```bash
# '--type' being arbitrary and an indentifer for the Go Type generated:
abigen --abi JoeMasterChef.abi --pkg main --type MasterChef --out joe-masterchef-v2.go
abigen --abi TokenPair.abi  --pkg contracts --type TokenPair --out token-pair.go
```
