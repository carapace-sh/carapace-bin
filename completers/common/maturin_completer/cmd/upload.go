package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/maturin_completer/cmd/action"
	"github.com/spf13/cobra"
)

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload python packages to pypi",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uploadCmd).Standalone()

	uploadCmd.Flags().Bool("non-interactive", false, "Do not interactively prompt for username/password if the required credentials are missing")
	uploadCmd.Flags().StringP("password", "p", "", "Password for pypi or your custom registry")
	uploadCmd.Flags().StringP("repository", "r", "pypi", "The repository (package index) to upload the package to")
	uploadCmd.Flags().String("repository-url", "", "The URL of the registry where the wheels are uploaded to")
	uploadCmd.Flags().Bool("skip-existing", false, "Continue uploading files if one already exists")
	uploadCmd.Flags().StringP("username", "u", "", "Username for pypi or your custom registry")
	uploadCmd.Flags().CountP("verbose", "v", "Use verbose output")
	rootCmd.AddCommand(uploadCmd)

	carapace.Gen(uploadCmd).FlagCompletion(carapace.ActionMap{
		"repository": action.ActionRepositories(),
	})

	carapace.Gen(uploadCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".whl", ".tar.gz"),
	)
}
