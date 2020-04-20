package cmd

import (
	"github.com/spf13/cobra"
)

var submoduleCmd = &cobra.Command{
	Use: "submodule",
	Short: "Initialize, update or inspect submodules",
    Run: func(cmd *cobra.Command, args []string) {
    },
}

func init() {

    rootCmd.AddCommand(submoduleCmd)
}
