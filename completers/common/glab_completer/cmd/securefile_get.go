package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
	"github.com/spf13/cobra"
)

var securefile_getCmd = &cobra.Command{
	Use:     "get <id>",
	Short:   "Get details of a secure file by ID.",
	Aliases: []string{"show"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securefile_getCmd).Standalone()

	securefile_getCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	securefileCmd.AddCommand(securefile_getCmd)

	carapace.Gen(securefile_getCmd).FlagCompletion(carapace.ActionMap{
		"jq": jq.ActionFilters(),
	})

	// TODO positional completion
}
