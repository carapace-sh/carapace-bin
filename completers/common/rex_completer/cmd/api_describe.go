package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/rex"
	"github.com/spf13/cobra"
)

var apiDescribeCmd = &cobra.Command{
	Use:   "describe METHOD",
	Short: "Describe a server API method",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apiDescribeCmd).Standalone()

	apiDescribeCmd.Flags().Bool("json", false, "output the method description as JSON")

	apiCmd.AddCommand(apiDescribeCmd)

	carapace.Gen(apiDescribeCmd).PositionalCompletion(
		rex.ActionApiMethods(),
	)
}
