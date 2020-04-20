package cmd

import (
	"github.com/spf13/cobra"
)

var sparse_checkoutCmd = &cobra.Command{
	Use: "sparse-checkout",
	Short: "Initialize and modify the sparse-checkout",
    Run: func(cmd *cobra.Command, args []string) {
    },
}

func init() {

    rootCmd.AddCommand(sparse_checkoutCmd)
}
