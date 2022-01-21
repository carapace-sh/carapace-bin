package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/golang"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "compile and install packages and dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().BoolS("a", "a", false, "force rebuilding of packages that are already up-to-date.")
	installCmd.Flags().String("asmflags", "", "arguments to pass on each go tool asm invocation")
	installCmd.Flags().String("buildmode", "", "build mode to use")
	installCmd.Flags().String("compiler", "", "name of compiler to use")
	installCmd.Flags().String("gccgoflags", "", "arguments to pass on each gccgo compiler/linker invocation")
	installCmd.Flags().String("gcflags", "", "arguments to pass on each go tool compile invocation.")
	installCmd.Flags().BoolS("i", "i", false, "install the packages that are dependencies of the target")
	installCmd.Flags().String("installsuffix", "", "a suffix to use in the name of the package installation directory")
	installCmd.Flags().String("ldflags", "", "arguments to pass on each go tool link invocation")
	installCmd.Flags().Bool("linkshared", false, "build code that will be linked against shared libraries")
	installCmd.Flags().String("mod", "", "module download mode to use")
	installCmd.Flags().Bool("modcacherw", false, "leave newly-created directories in the module cache read-write")
	installCmd.Flags().String("modfile", "", "read and possibly write an alternate go.mod file")
	installCmd.Flags().Bool("msan", false, "enable interoperation with memory sanitizer")
	installCmd.Flags().BoolS("n", "n", false, "print the commands but do not run them.")
	installCmd.Flags().StringS("o", "o", "", "set output file or directory")
	installCmd.Flags().StringS("p", "p", "", "the number of programs to run in parallel")
	installCmd.Flags().String("pkgdir", "", "install and load all packages from dir")
	installCmd.Flags().Bool("race", false, "enable data race detection")
	installCmd.Flags().String("tags", "", "a comma-separated list of build tags to consider satisfied during the")
	installCmd.Flags().String("toolexec", "", "a program to use to invoke toolchain programs like vet and asm")
	installCmd.Flags().Bool("trimpath", false, "remove all file system paths from the resulting executable")
	installCmd.Flags().BoolS("v", "v", false, "print the names of packages as they are compiled")
	installCmd.Flags().Bool("work", false, "print the name of the temporary work directory")
	installCmd.Flags().BoolS("x", "x", false, "print the commands.")
	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
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
