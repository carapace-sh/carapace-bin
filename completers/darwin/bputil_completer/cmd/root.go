package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bputil",
	Short: "modify security settings on Apple Silicon Macs",
	Long:  "https://keith.github.io/xcode-manpages/bputil.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("a", "a", false, "All")
	rootCmd.Flags().BoolS("c", "c", false, "Clear")
	rootCmd.Flags().BoolS("d", "d", false, "Display local policy container")
	rootCmd.Flags().BoolS("e", "e", false, "Enable")
	rootCmd.Flags().BoolS("f", "f", false, "Force")
	rootCmd.Flags().BoolS("g", "g", false, "Get")
	rootCmd.Flags().BoolS("j", "j", false, "Disable")
	rootCmd.Flags().BoolS("k", "k", false, "Kernel")
	rootCmd.Flags().BoolS("l", "l", false, "List all local policy containers")
	rootCmd.Flags().BoolS("m", "m", false, "Modify")
	rootCmd.Flags().BoolS("n", "n", false, "Next")
	rootCmd.Flags().StringS("p", "p", "", "Password")
	rootCmd.Flags().StringS("r", "r", "", "Recovery APFS Volume Group UUID")
	rootCmd.Flags().BoolS("s", "s", false, "Status")
	rootCmd.Flags().StringS("u", "u", "", "Username")
	rootCmd.Flags().StringS("v", "v", "", "APFS Volume Group UUID")
	rootCmd.Flags().BoolS("z", "z", false, "Zero")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"u": os.ActionUsers(),
	})
}
