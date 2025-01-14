package database

import (
	"encoding/json"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

var genesisJson = `
{
  "genesis_time": "2025-01-04T00:00:00.000000000Z",
  "chain_id": "the-blockchain-quantis-ledger",
  "balances": {
    "0x22ba1F80452E6220c7cc6ea2D1e3EEDDaC5F694A": 1000000
  }
}`

type Genesis struct {
	Balances map[common.Address]uint `json:"balances"`
}

func loadGenesis(path string) (Genesis, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return Genesis{}, err
	}

	var loadedGenesis Genesis
	err = json.Unmarshal(content, &loadedGenesis)
	if err != nil {
		return Genesis{}, err
	}

	return loadedGenesis, nil
}

func writeGenesisToDisk(path string, genesis []byte) error {
	return os.WriteFile(path, genesis, 0644)
}
