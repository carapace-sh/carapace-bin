package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var objdumpCmd = &cobra.Command{
	Use:     "objdump [options] file",
	Short:   "Print information about executables and relocatables",
	GroupID: "wrapper",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(objdumpCmd).Standalone()

	rootCmd.AddCommand(objdumpCmd)

	objdumpCmd.Flags().Bool("all-headers", false, "Alias for --file-headers --linker-member=2 --member-headers --section-headers --relocs --symbols")
	objdumpCmd.Flags().String("elements", "", "Select which formatting elements are displayed")
	objdumpCmd.Flags().String("exports", "", "Display exported symbols, optionally sorted by symbol name")
	objdumpCmd.Flags().Bool("file-headers", false, "Display file-format specific headers")
	objdumpCmd.Flags().BoolP("help", "h", false, "Print this help and exit")
	objdumpCmd.Flags().Bool("imports", false, "Display imported symbols")
	objdumpCmd.Flags().String("linker-member", "", "Display contents of the specified archive linker member")
	objdumpCmd.Flags().Bool("member-headers", false, "Display archive member headers")
	objdumpCmd.Flags().StringSlice("only-member", nil, "Only consider archive members names that contain name")
	objdumpCmd.Flags().StringSlice("only-section", nil, "Only consider section names that contain name")
	objdumpCmd.Flags().StringSlice("only-symbol", nil, "Only consider symbol names that contain name")
	objdumpCmd.Flags().String("redact", "", "Redact the specified field kind")
	objdumpCmd.Flags().Bool("relocs", false, "Display relocations")
	objdumpCmd.Flags().Bool("section-headers", false, "Display section headers")
	objdumpCmd.Flags().BoolP("snapshot", "s", false, "Alias for --redact=all --elements=-all")
	objdumpCmd.Flags().Bool("strings", false, "Display string tables")
	objdumpCmd.Flags().Bool("symbols", false, "Display symbol tables")
	objdumpCmd.Flags().Bool("tls", false, "Display TLS information")

	objdumpCmd.Flags().Lookup("exports").NoOptDefVal = " "
	objdumpCmd.Flags().Lookup("linker-member").NoOptDefVal = " "

	carapace.Gen(objdumpCmd).FlagCompletion(carapace.ActionMap{
		"elements": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.Batch(
				carapace.ActionValues("file-type", "header-name", "member-path", "newlines", "table-header", "all"),
				carapace.ActionValues("-file-type", "-header-name", "-member-path", "-newlines", "-table-header"),
			).ToA()
		}),
		"exports":       carapace.ActionValues("sort"),
		"linker-member": carapace.ActionValues("1", "2", "longnames"),
		"redact":        carapace.ActionValues("rva", "va", "ord", "size", "all"),
	})

	carapace.Gen(objdumpCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
