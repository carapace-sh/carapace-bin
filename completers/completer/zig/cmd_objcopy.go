package zig

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func newObjcopyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "objcopy [options] input output",
		Short: "Use Zig as a drop-in objcopy",
	}

	cmd.Flags().String("output-target", "", "Format of the output file")
	cmd.Flags().String("only-section", "", "Remove all sections except specified ones")
	cmd.Flags().String("pad-to", "", "Pad the last section to the specified address")
	cmd.Flags().String("gap-fill", "", "Fill gaps between sections with the specified byte")
	cmd.Flags().Bool("strip-debug", false, "Remove all debug sections from the output")
	cmd.Flags().Bool("strip-all", false, "Remove all debug sections and symbol table from the output")
	cmd.Flags().String("add-section", "", "Add a section to the output file")
	cmd.Flags().String("set-section-alignment", "", "Set alignment of section")
	cmd.Flags().String("set-section-flags", "", "Set flags of section")
	cmd.Flags().Bool("compress-debug-sections", false, "Compress DWARF debug sections")
	cmd.Flags().String("decompress-debug-sections", "", "Decompress DWARF debug sections")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"output-target": carapace.ActionValues(
			"elf", "binary", "hex", "ihex",
		),
	})

	carapace.Gen(cmd).PositionalCompletion(
		carapace.ActionFiles(".o", ".elf", ".bin", ".hex"),
		carapace.ActionFiles(".o", ".elf", ".bin", ".hex"),
	)

	return cmd
}
