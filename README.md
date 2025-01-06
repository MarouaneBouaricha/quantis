# The Blockchain Quantis
A basic peer-to-peer distributed ledger in the attempt of creating a blockchain from scratch.

## Installation
```shell
Go 1.23
```
```shell
go install ./cmd/
```
## How to use
```The Blockchain Quantis CLI

Usage:
  qcli [flags]
  qcli [command]

Available Commands:
  balances    Interact with balances (list...).
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  run         Launches the Quantis node and its HTTP API.
  version     Describes version.
  wallet      Manages blockchain accounts and keys.

Flags:
  -h, --help   help for qcli

Use "qcli [command] --help" for more information about a command.
```

## Mining
A Proof of Work algorithm is implemented when creating new blocks and the account who mined the new block get a 100 unit as reward.