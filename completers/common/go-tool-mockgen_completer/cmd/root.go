package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-tool-mockgen",
	Long:  "https://github.com/uber-go/mock",
	Short: "mock interfaces",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().SetInterspersed(false)

	rootCmd.Flags().StringS("aux_files", "aux_files", "", "Comma-separated pkg=path pairs of auxiliary Go source files")
	rootCmd.Flags().StringS("build_constraint", "build_constraint", "", "If non-empty, added as //go:build <constraint>")
	rootCmd.Flags().StringS("build_flags", "build_flags", "", "Additional flags for go build")
	rootCmd.Flags().StringS("copyright_file", "copyright_file", "", "Copyright file used to add copyright header")
	rootCmd.Flags().BoolS("debug_parser", "debug_parser", false, "Print out parser results only.")
	rootCmd.Flags().StringS("destination", "destination", "", "Output file")
	rootCmd.Flags().StringS("exclude_interfaces", "exclude_interfaces", "", "Comma-separated names of interfaces to be exclude")
	rootCmd.Flags().StringS("exec_only", "exec_only", "", "DEPRECATED  If set, execute this reflection program")
	rootCmd.Flags().StringS("imports", "imports", "", "Comma-separated name=path pairs of explicit imports to use")
	rootCmd.Flags().StringS("mock_names", "mock_names", "", "Comma-separated interfaceName=mockName pairs of explicit mock names to use")
	rootCmd.Flags().StringS("model_gob", "model_gob", "", "Skip package/source loading entirely and use the gob encoded model.Package at the given path")
	rootCmd.Flags().StringS("package", "package", "", "Package of the generated code")
	rootCmd.Flags().BoolS("prog_only", "prog_only", false, "DEPRECATED  Only generate the reflection progra")
	rootCmd.Flags().StringS("self_package", "self_package", "", "The full package import path for the generated code")
	rootCmd.Flags().StringS("source", "source", "", "Input Go source fil")
	rootCmd.Flags().BoolS("typed", "typed", false, "Generate Type-safe 'Return', 'Do', 'DoAndReturn' function")
	rootCmd.Flags().BoolS("version", "version", false, "Print version")
	rootCmd.Flags().BoolS("write_command_comment", "write_command_comment", false, "Writes the command used as a comment if true")
	rootCmd.Flags().BoolS("write_generate_directive", "write_generate_directive", false, "Add //go:generate directive to regenerate the mock")
	rootCmd.Flags().BoolS("write_package_comment", "write_package_comment", false, "Writes package documentation comment  if true")
	rootCmd.Flags().BoolS("write_source_comment", "write_source_comment", false, "Writes original file  or interface names  comment if true")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"aux_files": carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
			if len(c.Parts) == 0 {
				return carapace.ActionValues()
			}
			return carapace.ActionFiles()
		}).List(","),
		"build_flags":    bridge.ActionCarapaceBin("go", "build").Split(),
		"copyright_file": carapace.ActionFiles(),
		"destination":    carapace.ActionFiles(),
		"exec_only": carapace.Batch(
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
		"source": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		golang.ActionPackages(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return golang.ActionSymbols(golang.SymbolOpts{Package: c.Args[0], Unexported: true}) // TODO only interfaces
		}),
	)
}
