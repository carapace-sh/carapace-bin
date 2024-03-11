package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var utils_jqCmd = &cobra.Command{
	Use:   "jq",
	Short: "Parse json strings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(utils_jqCmd).Standalone()
	utilsCmd.AddCommand(utils_jqCmd)
}
