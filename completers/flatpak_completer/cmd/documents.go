package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/flatpak"
	"github.com/spf13/cobra"
)

var documentsCmd = &cobra.Command{
	Use:     "documents [OPTION…] [APPID]",
	Short:   "List exported files",
	GroupID: "access",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(documentsCmd).Standalone()

	documentsCmd.Flags().String("columns", "", "What information to show")
	documentsCmd.Flags().BoolP("help", "h", false, "Show help options")
	documentsCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	documentsCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(documentsCmd)

	carapace.Gen(documentsCmd).FlagCompletion(carapace.ActionMap{
		"columns": columns(flatpak.ActionDocumentColums()),
	})

	carapace.Gen(documentsCmd).PositionalCompletion(
		flatpak.ActionApplications(),
	)
}
