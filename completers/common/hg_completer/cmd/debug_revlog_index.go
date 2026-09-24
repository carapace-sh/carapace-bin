package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugRevlogIndexCmd = &cobra.Command{
	Use:    "debug-revlog-index",
	Short:  "-c|-m|FILE",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugRevlogIndexCmd).Standalone()

	debugRevlogIndexCmd.Flags().BoolP("changelog", "c", false, "open changelog")
	debugRevlogIndexCmd.Flags().String("dir", "", "open directory manifest")
	debugRevlogIndexCmd.Flags().BoolP("manifest", "m", false, "open manifest")
	debugRevlogIndexCmd.Flags().StringP("template", "T", "", "display with template")
	rootCmd.AddCommand(debugRevlogIndexCmd)
}
