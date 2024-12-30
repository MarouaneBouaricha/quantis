package database

import (
	"encoding/json"
	"os"
)

var genesisJson = `
{
  "genesis_time": "2024-12-30T00:00:00.000000000Z",
  "chain_id": "the-blockchain-quantis-ledger",
  "balances": {
    "marouane": 1000000
  }
}`

type genesis struct {
	Balances map[Account]uint `json:"balances"`
}

func loadGenesis(path string) (genesis, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return genesis{}, err
	}

	var loadedGenesis genesis
	err = json.Unmarshal(content, &loadedGenesis)
	if err != nil {
		return genesis{}, err
	}

	return loadedGenesis, nil
}

func writeGenesisToDisk(path string) error {
	return os.WriteFile(path, []byte(genesisJson), 0644)
}
