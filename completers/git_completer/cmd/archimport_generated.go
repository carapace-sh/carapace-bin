package cmd

import (
	"github.com/spf13/cobra"
)

var archimportCmd = &cobra.Command{
	Use: "archimport",
	Short: "Import a GNU Arch repository into Git",
    Run: func(cmd *cobra.Command, args []string) {
    },
}

func init() {

    rootCmd.AddCommand(archimportCmd)
}
