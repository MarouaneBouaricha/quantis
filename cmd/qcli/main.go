package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var qCmd = &cobra.Command{
		Use:   "qcli",
		Short: "The Blockchain Quantis CLI",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	qCmd.AddCommand(versionCmd)

	err := qCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
