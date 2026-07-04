package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dyld_info",
	Short: "Display information used by dyld in programs and dylibs",
	Long:  "https://keith.github.io/xcode-manpages/dyld_info.1.html",
	Run:   func(*cobra.Command, []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("all_dir", "", "start at specified directory, recurse to find all mach-o files")
	rootCmd.Flags().Bool("all_dyld_cache", false, "show info about all dylibs in the dyld cache")
	rootCmd.Flags().Bool("all_sections", false, "print content of all sections, formatted by section type")
	rootCmd.Flags().Bool("all_sections_bytes", false, "print content of all sections as raw hex bytes")
	rootCmd.Flags().String("arch", "", "limit to specified architecture(s)")
	rootCmd.Flags().Bool("disassemble", false, "print all code sections using disassembler")
	rootCmd.Flags().Bool("exports", false, "print all exported symbols")
	rootCmd.Flags().Bool("fixup_chain_details", false, "print detailed info about every fixup in chain")
	rootCmd.Flags().Bool("fixup_chain_header", false, "print detailed info about the fixup chains header")
	rootCmd.Flags().Bool("fixup_chains", false, "print info about chain format and starts")
	rootCmd.Flags().Bool("fixups", false, "print locations dyld will rebase/bind")
	rootCmd.Flags().Bool("function_starts", false, "print function starts information")
	rootCmd.Flags().Bool("function_variants", false, "print info on function variants in binary")
	rootCmd.Flags().Bool("imports", false, "print all symbols needed from other dylibs")
	rootCmd.Flags().Bool("inits", false, "print initializers")
	rootCmd.Flags().Bool("linked_dylibs", false, "print all dylibs this image links against")
	rootCmd.Flags().Bool("load_commands", false, "print load commands")
	rootCmd.Flags().Bool("no_validate", false, "don't check for malformedness about file(s)")
	rootCmd.Flags().Bool("objc", false, "print objc classes, categories, etc")
	rootCmd.Flags().Bool("opcodes", false, "print opcodes information")
	rootCmd.Flags().Bool("platform", false, "print platform")
	rootCmd.Flags().StringSlice("section", nil, "print content of section (segment section)")
	rootCmd.Flags().StringSlice("section_bytes", nil, "print content of section as raw hex bytes (segment section)")
	rootCmd.Flags().Bool("segments", false, "print segments")
	rootCmd.Flags().Bool("shared_region", false, "print shared cache (split seg) info")
	rootCmd.Flags().Bool("symbolic_fixups", false, "print ranges of each atom of DATA with symbol name and fixups")
	rootCmd.Flags().Bool("uuid", false, "print UUID of binary")
	rootCmd.Flags().Bool("validate_only", false, "only print malformedness about file(s)")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"all_dir": carapace.ActionDirectories(),
		"arch":    carapace.ActionValues("arm64", "arm64e", "x86_64", "x86_64h"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
