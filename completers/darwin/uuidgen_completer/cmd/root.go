package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "uuidgen",
	Short: "generates new UUID strings",
	Long:  "https://keith.github.io/xcode-manpages/uuidgen.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("hdr", "hdr", false, "Emit CoreFoundation CFUUID-based source code for using the uuid in a header")
}
