package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/top_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "top",
	Short: "display Linux processes",
	Long:  "https://man7.org/linux/man-pages/man1/top.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("accum-time-toggle", "S", false, "Starts top with the last remembered 'S' state reversed.")
	rootCmd.Flags().BoolP("apply-defaults", "A", false, "Run top with build defaults only, ignoring all configuration files except /etc/toprc.")
	rootCmd.Flags().BoolP("batch", "b", false, "Starts top in Batch mode")
	rootCmd.Flags().BoolP("cmdline-toggle", "c", false, "Starts top with the last remembered 'c' state reversed.")
	rootCmd.Flags().StringP("delay", "d", "", "Specifies the delay between screen updates")
	rootCmd.Flags().StringP("filter-any-user", "U", "", "Display only processes with a user id or user name matching that given.")
	rootCmd.Flags().StringP("filter-only-euser", "u", "", "Display only processes with a user id or user name matching that given.")
	rootCmd.Flags().BoolP("help", "h", false, "Display usage help text, then quit.")
	rootCmd.Flags().BoolP("idle-toggle", "i", false, "Starts top with the last remembered 'i' state reversed.")
	rootCmd.Flags().StringP("iterations", "n", "", "Specifies the maximum number of iterations, or frames, top should produce before ending.")
	rootCmd.Flags().StringP("list-fields", "O", "", "print available field names")
	rootCmd.Flags().StringP("pid", "p", "", "Monitor only processes with specified process IDs.")
	rootCmd.Flags().StringP("scale-summary-mem", "E", "", "Instructs top to force summary area memory to be scaled as")
	rootCmd.Flags().StringP("scale-task-mem", "e", "", "Instructs top to force task area memory to be scaled as")
	rootCmd.Flags().BoolP("secure-mode", "s", false, "Starts top with secure mode forced, even for root.")
	rootCmd.Flags().BoolP("single-cpu-toggle", "1", false, "Starts top with the last remembered Cpu States portion of the summary area reversed.")
	rootCmd.Flags().StringP("sort-override", "o", "", "Specifies the name of the field on which tasks will be sorted")
	rootCmd.Flags().BoolP("threads-show", "H", false, "Instructs top to display individual threads.")
	rootCmd.Flags().BoolP("version", "V", false, "Display version information, then quit.")
	rootCmd.Flags().BoolP("width", "w", false, "In Batch mode, when used without an argument top will format output using the COLUMNS= and LINES= environment variables, if set.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"filter-any-user":   os.ActionUsers(),
		"filter-only-euser": os.ActionUsers(),
		"pid":               ps.ActionProcessIds().UniqueList(","),
		"scale-summary-mem": carapace.ActionValuesDescribed(
			"k", "kibibytes",
			"m", "mebibytes",
			"g", "gibibytes",
			"t", "tebibytes",
			"p", "pebibytes",
			"e", "exbibytes",
		),
		"scale-task-mem": carapace.ActionValuesDescribed(
			"k", "kibibytes",
			"m", "mebibytes",
			"g", "gibibytes",
			"t", "tebibytes",
			"p", "pebibytes",
		),
		"sort-override": action.ActionFields(),
	})
}
