package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upload_anacondaCmd = &cobra.Command{
	Use:   "anaconda",
	Short: "Upload to an Anaconda.org server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upload_anacondaCmd).Standalone()

	upload_anacondaCmd.Flags().StringP("api-key", "a", "", "The Anaconda API key")
	upload_anacondaCmd.Flags().StringP("channel", "c", "", "The channel to upload the package to")
	upload_anacondaCmd.Flags().BoolP("force", "f", false, "Force upload even if the package already exists")
	upload_anacondaCmd.Flags().StringP("owner", "o", "", "The owner of the package")
	upload_anacondaCmd.Flags().StringP("url", "u", "", "The URL to the Anaconda server")
	uploadCmd.AddCommand(upload_anacondaCmd)
}
