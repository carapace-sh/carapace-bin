package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pub_unpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "Download and extract a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_unpackCmd).Standalone()

	pub_unpackCmd.Flags().BoolP("force", "f", false, "Overwrite existing conflicting folders.")
	pub_unpackCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pub_unpackCmd.Flags().Bool("no-resolve", false, "Do not run \"pub get\" after unpacking.")
	pub_unpackCmd.Flags().StringP("output", "o", "", "Output directory.")
	pub_unpackCmd.Flags().Bool("resolve", false, "Run \"pub get\" after unpacking.")
	pubCmd.AddCommand(pub_unpackCmd)

	carapace.Gen(pub_unpackCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionDirectories(),
	})
}
