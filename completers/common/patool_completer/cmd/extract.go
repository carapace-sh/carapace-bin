package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "extract one or more archives",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(extractCmd).Standalone()

	extractCmd.Flags().String("outdir", "", "output directory to extract to")
	extractCmd.Flags().String("password", "", "password for encrypted files")
	rootCmd.AddCommand(extractCmd)

	carapace.Gen(extractCmd).FlagCompletion(carapace.ActionMap{
		"outdir": carapace.ActionDirectories(),
	})

	carapace.Gen(extractCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
