package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var fmtCmd = &cobra.Command{
	Use:     "fmt [options] [files]",
	Short:   "Reformat Zig source into canonical form",
	GroupID: "source",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fmtCmd).Standalone()

	rootCmd.AddCommand(fmtCmd)

	fmtCmd.Flags().Bool("ast-check", false, "Run zig ast-check on every file")
	fmtCmd.Flags().Bool("check", false, "List non-conforming files and exit with an error if the list is non-empty")
	fmtCmd.Flags().String("color", "", "Enable or disable colored error messages")
	fmtCmd.Flags().StringSlice("exclude", nil, "Exclude file or directory from formatting")
	fmtCmd.Flags().BoolP("help", "h", false, "Print help")
	fmtCmd.Flags().Bool("stdin", false, "Format code from stdin; output to stdout")
	fmtCmd.Flags().Bool("zon", false, "Treat all input files as ZON, regardless of file extension")

	carapace.Gen(fmtCmd).FlagCompletion(carapace.ActionMap{
		"color":   carapace.ActionValues("auto", "off", "on").StyleF(style.ForKeyword),
		"exclude": carapace.ActionFiles(),
	})

	carapace.Gen(fmtCmd).PositionalAnyCompletion(
		carapace.ActionFiles().FilterArgs(),
	)
}
