package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/rust"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rustc",
	Short: "compiler for the Rust programming language",
	Long:  "https://doc.rust-lang.org/rustc/index.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringArrayS("L", "L", nil, "Add a directory to the library search path")
	rootCmd.Flags().BoolS("O", "O", false, "Equivalent to -C opt-level=3")
	rootCmd.Flags().StringP("allow", "A", "", "Set lint allowed")
	rootCmd.Flags().String("cap-lints", "", "Set the most restrictive lint level")
	rootCmd.Flags().String("cfg", "", "Configure the compilation environment")
	rootCmd.Flags().String("check-cfg", "", "Provide list of expected cfgs for checking")
	rootCmd.Flags().StringP("codegen", "C", "", "Set a codegen option")
	rootCmd.Flags().String("crate-name", "", "Specify the name of the crate being built")
	rootCmd.Flags().String("crate-type", "", "Comma separated list of types of crates for the compiler to emit")
	rootCmd.Flags().StringP("deny", "D", "", "Set lint denied")
	rootCmd.Flags().String("edition", "", "Specify which edition of the compiler to use when compiling code")
	rootCmd.Flags().String("emit", "", "Comma separated list of types of output for the compiler to emit")
	rootCmd.Flags().String("explain", "", "Provide a detailed explanation of an error message")
	rootCmd.Flags().StringP("forbid", "F", "", "Set lint forbidden")
	rootCmd.Flags().String("force-warn", "", "Set lint force-warn")
	rootCmd.Flags().BoolS("g", "g", false, "Equivalent to -C debuginfo=2")
	rootCmd.Flags().BoolP("help", "h", false, "Display this message")
	rootCmd.Flags().StringArrayS("l", "l", nil, "Link the generated crate(s) to the specified native library NAME")
	rootCmd.Flags().StringS("o", "o", "", "Write output to FILENAME")
	rootCmd.Flags().String("out-dir", "", "Write output to compiler-chosen filename in DIR")
	rootCmd.Flags().String("print", "", "Compiler information to print on stdout")
	rootCmd.Flags().String("target", "", "Target triple for which the code is compiled")
	rootCmd.Flags().Bool("test", false, "Build a test harness")
	rootCmd.Flags().BoolP("verbose", "v", false, "Use verbose output")
	rootCmd.Flags().BoolP("version", "V", false, "Print version info and exit")
	rootCmd.Flags().StringP("warn", "W", "", "Set lint warnings")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"L": carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues("dependency", "crate", "native", "framework", "all").Suffix("=")
			default:
				return carapace.ActionFiles()
			}
		}),
		"allow":      carapace.ActionValues(), // TODO
		"cap-lints":  rust.ActionLintLevels(),
		"cfg":        carapace.ActionValues(), // TODO
		"check-cfg":  carapace.ActionValues(), // TODO
		"codegen":    rust.ActionCodegenOptions(),
		"crate-type": rust.ActionCrateTypes().UniqueList(","),
		"deny":       carapace.ActionValues(), // TODO
		"edition":    carapace.ActionValues("2015", "2018", "2021"),
		"emit":       carapace.ActionValues("asm", "llvm-bc", "llvm-ir", "obj", "metadata", "link", "dep-info", "mir").UniqueList(","),
		"forbid":     carapace.ActionValues(), // TODO
		"force-warn": carapace.ActionValues(), // TODO
		"l": carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues("static", "framework", "dylib").Suffix("=")
			default:
				return carapace.ActionFiles()
			}
		}),
		"o":       carapace.ActionFiles(),
		"out-dir": carapace.ActionDirectories(),
		"print":   carapace.ActionValues("crate-name", "file-names", "sysroot", "target-libdir", "cfg", "target-list", "target-cpus", "target-features", "relocation-models", "code-models", "tls-models", "target-spec-json", "native-static-libs"),
		"target":  carapace.ActionValues(), // TODO
		"warn":    carapace.ActionValues(), // TODO
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(".rs"),
	)
}
