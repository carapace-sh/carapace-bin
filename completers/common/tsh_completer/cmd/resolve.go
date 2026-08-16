package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resolveCmd = &cobra.Command{
	Use:   "resolve",
	Short: "Resolves an SSH host.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resolveCmd).Standalone()

	resolveCmd.Flags().StringP("format", "f", "text", "Format output (text, json, yaml).")
	resolveCmd.Flags().Bool("no-quiet", false, "Quiet mode.")
	resolveCmd.Flags().BoolP("quiet", "q", false, "Quiet mode.")
	resolveCmd.Flag("no-quiet").Hidden = true
	rootCmd.AddCommand(resolveCmd)
}
