package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/qemu"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qemu-keymap",
	Short: "QEMU reverse keymap generator",
	Long:  "https://www.qemu.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("file", "f", "", "set output file (default: stdout)")
	rootCmd.Flags().BoolP("help", "h", false, "print this text")
	rootCmd.Flags().StringP("layout", "l", "", "set kbd layout (default: us)")
	rootCmd.Flags().StringP("model", "m", "", "set kbd model (default: pc105)")
	rootCmd.Flags().StringP("options", "o", "", "set kbd options")
	rootCmd.Flags().StringP("variant", "v", "", "set kbd variant")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"file":    carapace.ActionFiles(),
		"layout":  qemu.ActionKeymapLayouts(),
		"model":   qemu.ActionKeymapModels(),
		"options": qemu.ActionKeymapOptions(),
		"variant": qemu.ActionKeymapVariants(),
	})
}
