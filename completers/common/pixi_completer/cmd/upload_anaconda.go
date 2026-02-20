package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upload_anacondaCmd = &cobra.Command{
	Use:   "anaconda",
	Short: "Options for uploading to a Anaconda.org server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upload_anacondaCmd).Standalone()

	upload_anacondaCmd.Flags().StringP("api-key", "a", "", "The Anaconda API key, if none is provided, the token is read from the keychain / auth-file")
	upload_anacondaCmd.Flags().StringSliceP("channel", "c", nil, "The channel / label to upload the package to (e.g. main / rc)")
	upload_anacondaCmd.Flags().BoolP("force", "f", false, "Replace files on conflict")
	upload_anacondaCmd.Flags().StringP("owner", "o", "", "The owner of the distribution (e.g. conda-forge or your username)")
	upload_anacondaCmd.Flags().StringP("url", "u", "", "The URL to the Anaconda server")
	upload_anacondaCmd.MarkFlagRequired("owner")
	uploadCmd.AddCommand(upload_anacondaCmd)
}
