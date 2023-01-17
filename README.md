# chainlink-tally

This repository contains all the required code for uploading L2 election tallies to a smartcontract on EVM based chains.

- Relayer: The main piece for interacting with the EVM smart contract. Connects to an EVM endpoint and is able to interact with the EVM smart contract with Golang auto-generated bindings.
  - EVMClient: Is the client for connecting to a given EVM endpoint for interacting with the smart contract.
  - EVMContracts: Contain all the smart contract code. Uses Hardhat as the base for development.

## Compile & Run

Compile from source in a golang environment (Go>1.19 required):

```bash
git clone https://github.com/vocdoni/chainlink-tally.git
cd chainlink-tally/relayer/evmcontracts
make all
cd ../../
go build ./cmd/main.go
./main --help
```

[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v1.4%20adopted-ff69b4.svg)](code-of-conduct.md) [![License: AGPL v3](https://img.shields.io/badge/License-AGPL%20v3-blue.svg)](https://www.gnu.org/licenses/agpl-3.0)
