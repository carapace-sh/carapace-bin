package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var utils_yqCmd = &cobra.Command{
	Use:   "yq",
	Short: "Parse json strings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(utils_yqCmd).Standalone()
	utilsCmd.AddCommand(utils_yqCmd)
}
