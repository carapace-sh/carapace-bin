package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "btop",
	Short: "A monitor of resources",
	Long:  "https://github.com/aristocratos/btop",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("debug", false, "start in DEBUG mode")
	rootCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.Flags().BoolP("low-color", "lc", false, "disable truecolor")
	rootCmd.Flags().StringP("preset", "p", "", "start with preset")
	rootCmd.Flags().Bool("tty_off", false, "force (OFF) tty mode")
	rootCmd.Flags().BoolP("tty_on", "t", false, "force (ON) tty mode")
	rootCmd.Flags().StringP("update", "u", "", "set the program update rate in milliseconds")
	rootCmd.Flags().Bool("utf-force", false, "force start even if no UTF-8 locale was detected")
	rootCmd.Flags().BoolP("version", "v", false, "show version info and exit")
}
