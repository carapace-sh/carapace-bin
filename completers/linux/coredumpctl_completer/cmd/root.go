package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/time"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "coredumpctl",
	Short: "List or retrieve coredumps from the journal",
	Long:  "https://www.freedesktop.org/software/systemd/man/latest/coredumpctl.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().BoolS("1", "1", false, "Show information about most recent entry only")
	rootCmd.PersistentFlags().Bool("all", false, "Look at all journal files instead of local ones")
	rootCmd.PersistentFlags().String("debugger", "", "Use the given debugger")
	rootCmd.PersistentFlags().StringP("debugger-arguments", "A", "", "Pass the given arguments to the debugger")
	rootCmd.PersistentFlags().StringP("directory", "D", "", "Use journal files from directory")
	rootCmd.PersistentFlags().StringP("field", "F", "", "List all values a certain field takes")
	rootCmd.PersistentFlags().String("file", "", "Use journal file")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Show this help")
	rootCmd.PersistentFlags().String("image", "", "Operate on disk image as filesystem root")
	rootCmd.PersistentFlags().String("image-policy", "", "Specify disk image dissection policy")
	rootCmd.PersistentFlags().String("json", "", "Generate JSON output")
	rootCmd.PersistentFlags().StringS("n", "n", "", "Show maximum number of rows")
	rootCmd.PersistentFlags().Bool("no-legend", false, "Do not print the column headers")
	rootCmd.PersistentFlags().Bool("no-pager", false, "Do not pipe output into a pager")
	rootCmd.PersistentFlags().StringP("output", "o", "", "Write output to FILE")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "Do not show info messages and privilege warning")
	rootCmd.PersistentFlags().BoolP("reverse", "r", false, "Show the newest entries first")
	rootCmd.PersistentFlags().String("root", "", "Operate on an alternate filesystem root")
	rootCmd.PersistentFlags().StringP("since", "S", "", "Only print coredumps since the date")
	rootCmd.PersistentFlags().StringP("until", "U", "", "Only print coredumps until the date")
	rootCmd.PersistentFlags().Bool("version", false, "Print version string")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"debugger": carapace.Batch(
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
		"directory": carapace.ActionDirectories(),
		"image":     carapace.ActionFiles(),
		"json":      carapace.ActionValues("pretty", "short", "off"),
		"output":    carapace.ActionFiles(),
		"root":      carapace.ActionDirectories(),
		"since":     time.ActionDate(),
		"until":     time.ActionDate(),
	})
}
