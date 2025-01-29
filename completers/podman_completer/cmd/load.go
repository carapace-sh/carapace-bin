package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var loadCmd = &cobra.Command{
	Use:   "load [options]",
	Short: "Load image(s) from a tar archive",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loadCmd).Standalone()

	loadCmd.Flags().StringP("input", "i", "", "Read from specified archive file (default: stdin)")
	loadCmd.Flags().BoolP("quiet", "q", false, "Suppress the output")
	loadCmd.Flags().String("signature-policy", "", "Pathname of signature policy file")
	loadCmd.Flag("signature-policy").Hidden = true
	rootCmd.AddCommand(loadCmd)
}
