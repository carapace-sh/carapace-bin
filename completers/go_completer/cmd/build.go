package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/golang"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "compile packages and dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	buildCmd.Flags().BoolS("a", "a", false, "force rebuilding of packages that are already up-to-date.")
	buildCmd.Flags().StringS("asmflags", "asmflags", "", "arguments to pass on each go tool asm invocation")
	buildCmd.Flags().StringS("buildmode", "buildmode", "", "build mode to use")
	buildCmd.Flags().StringS("compiler", "compiler", "", "name of compiler to use")
	buildCmd.Flags().StringS("gccgoflags", "gccgoflags", "", "arguments to pass on each gccgo compiler/linker invocation")
	buildCmd.Flags().StringS("gcflags", "gcflags", "", "arguments to pass on each go tool compile invocation.")
	buildCmd.Flags().BoolS("i", "i", false, "install the packages that are dependencies of the target")
	buildCmd.Flags().StringS("installsuffix", "installsuffix", "", "a suffix to use in the name of the package installation directory")
	buildCmd.Flags().StringS("ldflags", "ldflags", "", "arguments to pass on each go tool link invocation")
	buildCmd.Flags().BoolS("linkshared", "linkshared", false, "build code that will be linked against shared libraries")
	buildCmd.Flags().StringS("mod", "mod", "", "module download mode to use")
	buildCmd.Flags().BoolS("modcacherw", "modcacherw", false, "leave newly-created directories in the module cache read-write")
	buildCmd.Flags().StringS("modfile", "modfile", "", "read and possibly write an alternate go.mod file")
	buildCmd.Flags().BoolS("msan", "msan", false, "enable interoperation with memory sanitizer")
	buildCmd.Flags().BoolS("n", "n", false, "print the commands but do not run them.")
	buildCmd.Flags().StringS("o", "o", "", "set output file or directory")
	buildCmd.Flags().StringS("p", "p", "", "the number of programs to run in parallel")
	buildCmd.Flags().StringS("pkgdir", "pkgdir", "", "install and load all packages from dir")
	buildCmd.Flags().BoolS("race", "race", false, "enable data race detection")
	buildCmd.Flags().StringS("tags", "tags", "", "a comma-separated list of build tags to consider satisfied during the")
	buildCmd.Flags().StringS("toolexec", "toolexec", "", "a program to use to invoke toolchain programs like vet and asm")
	buildCmd.Flags().BoolS("trimpath", "trimpath", false, "remove all file system paths from the resulting executable")
	buildCmd.Flags().BoolS("v", "v", false, "print the names of packages as they are compiled")
	buildCmd.Flags().BoolS("work", "work", false, "print the name of the temporary work directory")
	buildCmd.Flags().BoolS("x", "x", false, "print the commands.")
	rootCmd.AddCommand(buildCmd)

	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		"buildmode": carapace.ActionValues("archive", "c-archive", "c-shared", "default", "shared", "exe", "pie", "plugin"),
		"compiler":  carapace.ActionValues("gccgo", "gc"),
		"mod":       carapace.ActionValues("readonly", "vendor", "mod"),
		"modfile":   carapace.ActionFiles(".mod"),
		"n":         carapace.ActionValues("1", "2", "3", "4", "5", "6", "7", "8"),
		"o":         carapace.ActionFiles(),
		"pkgdir":    carapace.ActionDirectories(),
		"tags": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return golang.ActionBuildTags().Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
