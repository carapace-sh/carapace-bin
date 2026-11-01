package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apiListCmd = &cobra.Command{
	Use:     "list",
	Short:   "List server API methods",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apiListCmd).Standalone()

	apiListCmd.Flags().Bool("json", false, "output the API help document as JSON")
	apiListCmd.Flags().BoolP("quiet", "q", false, "print method names only")

	apiCmd.AddCommand(apiListCmd)
}
