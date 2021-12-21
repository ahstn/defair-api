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
