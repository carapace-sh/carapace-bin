package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sdshowCmd = &cobra.Command{
	Use:   "sdshow",
	Short: "display a service's security descriptor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sdshowCmd).Standalone()
	rootCmd.AddCommand(sdshowCmd)
}
