package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-tool-cgo",
	Short: "Cgo enables the creation of Go packages that call C code",
	Long:  "https://pkg.go.dev/cmd/cgo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("V", "V", false, "print version and exit")
	rootCmd.Flags().BoolS("debug-define", "debug-define", false, "print relevant #defines")
	rootCmd.Flags().BoolS("debug-gcc", "debug-gcc", false, "print gcc invocations")
	rootCmd.Flags().StringS("dynimport", "dynimport", "", "if non-empty, print dynamic import data for that file")
	rootCmd.Flags().BoolS("dynlinker", "dynlinker", false, "record dynamic linker information in -dynimport mode")
	rootCmd.Flags().StringS("dynout", "dynout", "", "write -dynimport output to this file")
	rootCmd.Flags().StringS("dynpackage", "dynpackage", "", "set Go package for -dynimport output (default \"main\")")
	rootCmd.Flags().StringS("exportheader", "exportheader", "", "where to write export header if any exported functions")
	rootCmd.Flags().BoolS("gccgo", "gccgo", false, "generate files for use with gccgo")
	rootCmd.Flags().BoolS("gccgo_define_cgoincomplete", "gccgo_define_cgoincomplete", false, "define cgo.Incomplete for older gccgo/GoLLVM")
	rootCmd.Flags().StringS("gccgopkgpath", "gccgopkgpath", "", "-fgo-pkgpath option used with gccgo")
	rootCmd.Flags().StringS("gccgoprefix", "gccgoprefix", "", "-fgo-prefix option used with gccgo")
	rootCmd.Flags().BoolS("godefs", "godefs", false, "for bootstrap: write Go definitions for C file to standard output")
	rootCmd.Flags().BoolS("import_runtime_cgo", "import_runtime_cgo", false, "import runtime/cgo in generated code (default true)")
	rootCmd.Flags().BoolS("import_syscall", "import_syscall", false, "import syscall in generated code (default true)")
	rootCmd.Flags().StringS("importpath", "importpath", "", "import path of package being built (for comments in generated files)")
	rootCmd.Flags().StringS("objdir", "objdir", "", "object directory")
	rootCmd.Flags().StringS("srcdir", "srcdir", "", "source directory")
	rootCmd.Flags().StringS("trimpath", "trimpath", "", "applies supplied rewrites or trims prefixes to recorded source file paths")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"dynimport":  carapace.ActionFiles(),
		"dynout":     carapace.ActionFiles(),
		"dynpackage": golang.ActionPackages(),
		"objdir":     carapace.ActionDirectories(),
		"srcdir":     carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".go"),
	)
}
