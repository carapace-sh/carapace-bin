package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upload_condaForgeCmd = &cobra.Command{
	Use:   "conda-forge",
	Short: "Upload to conda-forge",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upload_condaForgeCmd).Standalone()

	upload_condaForgeCmd.Flags().String("anaconda-url", "", "The Anaconda URL")
	upload_condaForgeCmd.Flags().Bool("dry-run", false, "Perform a dry run without uploading")
	upload_condaForgeCmd.Flags().String("feedstock", "", "The feedstock name")
	upload_condaForgeCmd.Flags().String("feedstock-token", "", "The feedstock token")
	upload_condaForgeCmd.Flags().String("provider", "", "The provider")
	upload_condaForgeCmd.Flags().String("staging-channel", "", "The staging channel")
	upload_condaForgeCmd.Flags().String("staging-token", "", "The staging token")
	upload_condaForgeCmd.Flags().String("validation-endpoint", "", "The validation endpoint")
	uploadCmd.AddCommand(upload_condaForgeCmd)
}
