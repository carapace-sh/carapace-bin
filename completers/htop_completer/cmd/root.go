package cmd

import (
	"os/exec"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "htop",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Root()
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("delay", "d", "", "Set the delay between updates, in tenths of seconds")
	rootCmd.Flags().BoolP("help", "h", false, "Print this help screen")
	rootCmd.Flags().BoolP("no-color", "C", false, "Use a monochrome color scheme")
	rootCmd.Flags().BoolP("no-mouse", "m", false, "Disable the mouse")
	rootCmd.Flags().BoolP("no-unicode", "U", false, "Do not use unicode but plain ASCII")
	rootCmd.Flags().StringP("pid", "p", "", "Show only the given PIDs")
	rootCmd.Flags().StringP("sort-key", "s", "", "Sort by COLUMN (try --sort-key=help for a list)")
	rootCmd.Flags().BoolP("tree", "t", false, "Show the tree view by default")
	rootCmd.Flags().StringP("user", "u", "", "Show only processes for a given user (or $USER)")
	rootCmd.Flags().BoolP("version", "v", false, "Print version info")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"pid": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return os.ActionProcessIds().Invoke(c).Filter(c.Parts).ToA()
		}),
		"sort-key": ActionSortKeys(),
		"user":     os.ActionUsers(),
	})
}

func ActionSortKeys() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if output, err := exec.Command("htop", "--sort-key", "help").Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[:len(lines)-1]...)
		}
	})
}
