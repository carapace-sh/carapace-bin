package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "diskutil",
	Short: "modify, verify and repair local disks",
	Long:  "https://keith.github.io/xcode-manpages/diskutil.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().BoolP("acls", "a", false, "Enable ACL support on the volume")
	rootCmd.PersistentFlags().BoolP("force", "f", false, "Force operation")
	rootCmd.PersistentFlags().BoolP("interactive", "i", false, "Interactive mode")
	rootCmd.PersistentFlags().BoolP("plist", "p", false, "Output in plist format")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "Quiet mode")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose mode")
}
