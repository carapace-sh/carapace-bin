package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pmap",
	Short: "report memory map of a process",
	Long:  "https://man7.org/linux/man-pages/man1/pmap.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("X", "X", false, "show even more details")
	rootCmd.Flags().BoolS("XX", "XX", false, "show everything the kernel provides")
	rootCmd.Flags().BoolP("create-rc", "n", false, "create new default rc")
	rootCmd.Flags().StringP("create-rc-to", "N", "", "create new rc to file")
	rootCmd.Flags().BoolP("device", "d", false, "show the device format")
	rootCmd.Flags().BoolP("extended", "x", false, "show details")
	rootCmd.Flags().BoolP("help", "h", false, "display this help and exit")
	rootCmd.Flags().BoolP("quiet", "q", false, "do not display header and footer")
	rootCmd.Flags().StringP("range", "A", "", "limit results to the given range")
	rootCmd.Flags().BoolP("read-rc", "c", false, "read the default rc")
	rootCmd.Flags().StringP("read-rc-from", "C", "", "read the rc from file")
	rootCmd.Flags().BoolP("show-path", "p", false, "show path in the mapping")
	rootCmd.Flags().BoolP("version", "V", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"create-rc-to": carapace.ActionFiles(),
		"read-rc-from": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		ps.ActionProcessIds(),
	)
}
