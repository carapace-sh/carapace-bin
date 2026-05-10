package zig

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

func newFmtCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fmt [options] [files]",
		Short: "Reformat Zig source code",
	}

	cmd.Flags().Bool("check", false, "List non-conforming files and exit with an error if the list is non-empty")
	cmd.Flags().Bool("ast-check", false, "Run zig ast-check on every file")
	cmd.Flags().Bool("stdin", false, "Format code from stdin; output to stdout")
	cmd.Flags().String("color", "", "Enable or disable colored error messages")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"color": carapace.ActionValues("auto", "off", "on").StyleF(style.ForKeyword),
	})

	carapace.Gen(cmd).PositionalAnyCompletion(
		carapace.ActionFiles(".zig", ".zon"),
	)

	return cmd
}
