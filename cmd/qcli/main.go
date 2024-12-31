package main

import (
	"fmt"
	"os"

	"github.com/MarouaneBouaricha/quantis/fs"
	"github.com/spf13/cobra"
)

const flagDataDir = "datadir"
const flagMiner = "miner"
const flagIP = "ip"
const flagPort = "port"

func main() {
	var qCmd = &cobra.Command{
		Use:   "qcli",
		Short: "The Blockchain Quantis CLI",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	qCmd.AddCommand(versionCmd)
	qCmd.AddCommand(balancesCmd())
	qCmd.AddCommand(runCmd())

	err := qCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func addDefaultRequiredFlags(cmd *cobra.Command) {
	cmd.Flags().String(flagDataDir, "", "Absolute path to the node data dir where the DB will/is stored")
	cmd.MarkFlagRequired(flagDataDir)
}

func getDataDirFromCmd(cmd *cobra.Command) string {
	dataDir, _ := cmd.Flags().GetString(flagDataDir)

	return fs.ExpandPath(dataDir)
}

func incorrectUsageErr() error {
	return fmt.Errorf("incorrect usage")
}
