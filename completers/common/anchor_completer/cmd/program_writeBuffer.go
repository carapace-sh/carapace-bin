package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/anchor"
	"github.com/spf13/cobra"
)

var program_writeBufferCmd = &cobra.Command{
	Use:   "write-buffer",
	Short: "Write a program into a buffer account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(program_writeBufferCmd).Standalone()

	program_writeBufferCmd.Flags().String("buffer", "", "Buffer account keypair (defaults to new keypair)")
	program_writeBufferCmd.Flags().String("buffer-authority", "", "Buffer authority (defaults to configured wallet)")
	program_writeBufferCmd.Flags().BoolP("help", "h", false, "Print help")
	program_writeBufferCmd.Flags().String("max-len", "", "Maximum transaction length")
	program_writeBufferCmd.Flags().StringP("program-name", "p", "", "Program name to write (from workspace). Used when program_filepath is not provided")
	programCmd.AddCommand(program_writeBufferCmd)

	carapace.Gen(program_writeBufferCmd).FlagCompletion(carapace.ActionMap{
		"buffer-authority": carapace.ActionFiles(),
		"program-name":     anchor.ActionPrograms(),
	})

	carapace.Gen(program_writeBufferCmd).PositionalCompletion(
		carapace.ActionFiles(".so"),
	)
}
