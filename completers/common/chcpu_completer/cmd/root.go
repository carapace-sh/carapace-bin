package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chcpu",
	Short: "configure CPUs",
	Long:  "https://man7.org/linux/man-pages/man8/chcpu.8.html",
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

	rootCmd.Flags().StringP("configure", "c", "", "configure cpus")
	rootCmd.Flags().StringP("deconfigure", "g", "", "deconfigure cpus")
	rootCmd.Flags().StringP("disable", "d", "", "disable cpus")
	rootCmd.Flags().StringP("dispatch", "p", "", "set dispatching mode")
	rootCmd.Flags().StringP("enable", "e", "", "enable cpus")
	rootCmd.Flags().BoolP("help", "h", false, "display this help")
	rootCmd.Flags().BoolP("rescan", "r", false, "trigger rescan of cpus")
	rootCmd.Flags().BoolP("version", "V", false, "display version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"dispatch": carapace.ActionValues("horizontal", "vertical"),
	})
}
