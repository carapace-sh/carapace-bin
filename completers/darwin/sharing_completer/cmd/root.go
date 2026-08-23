package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sharing",
	Short: "create share points for smb services",
	Long:  "https://keith.github.io/xcode-manpages/sharing.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("add", "a", "", "Add a new share point for the specified path")
	rootCmd.Flags().StringP("edit", "e", "", "Edit the share point record")
	rootCmd.Flags().StringP("encrypted", "E", "", "Make share encrypted for smb v3+ (0/1)")
	rootCmd.Flags().StringP("format", "f", "", "Output format for listing (json)")
	rootCmd.Flags().StringP("guest-flags", "g", "", "Enable/disable guest access for smb (001/000)")
	rootCmd.Flags().BoolP("list", "l", false, "List all existing share point records")
	rootCmd.Flags().StringP("name", "n", "", "Customized record name")
	rootCmd.Flags().StringP("read-only", "R", "", "Make share read only for smb (0/1)")
	rootCmd.Flags().StringP("remove", "r", "", "Delete the share point record")
	rootCmd.Flags().StringP("smb-flags", "s", "", "Enable/disable sharing for smb (001/000)")
	rootCmd.Flags().StringP("smb-name", "S", "", "Customized name for smb")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"add":         carapace.ActionDirectories(),
		"encrypted":   carapace.ActionValues("0", "1"),
		"format":      carapace.ActionValues("json"),
		"guest-flags": carapace.ActionValues("000", "001"),
		"read-only":   carapace.ActionValues("0", "1"),
		"smb-flags":   carapace.ActionValues("000", "001"),
	})
}
