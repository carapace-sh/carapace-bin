package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pub_pubCmd = &cobra.Command{
	Use:   "pub",
	Short: "Pass the remaining arguments to Dart's \"pub\" tool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_pubCmd).Standalone()

	pub_pubCmd.Flags().Bool("color", false, "Use colors in terminal output.")
	pub_pubCmd.Flags().StringP("directory", "C", "", "Run the subcommand in the directory <dir>.")
	pub_pubCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pub_pubCmd.Flags().Bool("no-color", false, "Do not use colors in terminal output.")
	pub_pubCmd.Flags().BoolP("verbose", "v", false, "Print detailed logging.")
	pubCmd.AddCommand(pub_pubCmd)
}
