package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var srumutilCmd = &cobra.Command{
	Use:   "srumutil",
	Short: "dump energy estimation data from SRUM",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(srumutilCmd).Standalone()
	rootCmd.AddCommand(srumutilCmd)
}
