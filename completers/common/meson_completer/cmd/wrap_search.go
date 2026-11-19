package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wrap_searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search the db by name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wrap_searchCmd).Standalone()

	wrap_searchCmd.Flags().Bool("allow-insecure", false, "Allow insecure server connections.")
	applyCommonFlagCompletion(wrap_searchCmd)
	wrapCmd.AddCommand(wrap_searchCmd)
}
