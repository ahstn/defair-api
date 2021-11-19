

## Contributing
Adding a new Smart Contract or ABI:
```bash
# Type being arbitrary and an indentifer for the Go Type generated:
abigen --abi JoeMasterChef.abi --pkg main --type MasterChef --out joe-masterchef-v2.go
abigen --abi TokenPair.abi  --pkg contracts --type TokenPair --out token-pair.go
```