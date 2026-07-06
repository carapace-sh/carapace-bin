package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qemu-edid",
	Short: "QEMU EDID test tool",
	Long:  "https://www.qemu.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("dpi", "d", "", "set display resolution")
	rootCmd.Flags().BoolP("help", "h", false, "print this text")
	rootCmd.Flags().StringP("maxx", "X", "", "set maximum width")
	rootCmd.Flags().StringP("maxy", "Y", "", "set maximum height")
	rootCmd.Flags().StringP("name", "n", "", "set monitor name")
	rootCmd.Flags().StringP("output", "o", "", "set output file (stdout by default)")
	rootCmd.Flags().StringP("prefx", "x", "", "set preferred width")
	rootCmd.Flags().StringP("prefy", "y", "", "set preferred height")
	rootCmd.Flags().StringP("serial", "s", "", "set monitor serial")
	rootCmd.Flags().StringP("vendor", "v", "", "set monitor vendor (three letters)")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
		"vendor": carapace.ActionValues(),
	})
}
