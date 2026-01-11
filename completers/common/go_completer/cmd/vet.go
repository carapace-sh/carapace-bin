package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var vetCmd = &cobra.Command{
	Use:   "vet",
	Short: "report likely mistakes in packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vetCmd).Standalone()
	vetCmd.Flags().SetInterspersed(false)

	vetCmd.Flags().StringS("C", "C", "", "Change to dir before running the command")
	vetCmd.Flags().BoolS("a", "a", false, "force rebuilding of packages that are already up-to-date")
	vetCmd.Flags().BoolS("asan", "asan", false, "enable interoperation with address sanitizer")
	vetCmd.Flags().StringS("asmflags", "asmflags", "", "arguments to pass on each go tool asm invocation")
	vetCmd.Flags().StringS("buildmode", "buildmode", "", "build mode to use")
	vetCmd.Flags().StringS("buildvcs", "buildvcs", "", "whether to stamp binaries with version control information")
	vetCmd.Flags().StringS("compiler", "compiler", "", "name of compiler to use")
	vetCmd.Flags().StringS("gccgoflags", "gccgoflags", "", "arguments to pass on each gccgo compiler/linker invocation")
	vetCmd.Flags().StringS("gcflags", "gcflags", "", "arguments to pass on each go tool compile invocation")
	vetCmd.Flags().StringS("installsuffix", "installsuffix", "", "a suffix to use in the name of the package installation directory")
	vetCmd.Flags().StringS("ldflags", "ldflags", "", "arguments to pass on each go tool link invocation")
	vetCmd.Flags().BoolS("linkshared", "linkshared", false, "build code that will be linked against shared libraries")
	vetCmd.Flags().BoolS("msan", "msan", false, "enable interoperation with memory sanitizer")
	vetCmd.Flags().BoolS("n", "n", false, "print the commands but do not run them")
	vetCmd.Flags().StringS("overlay", "overlay", "", "read a JSON config file that provides an overlay for build operations")
	vetCmd.Flags().StringS("p", "p", "", "the number of programs to run in parallel")
	vetCmd.Flags().StringS("pgo", "pgo", "", "specify the file path of a profile for profile-guided optimization")
	vetCmd.Flags().StringS("pkgdir", "pkgdir", "", "install and load all packages from dir")
	vetCmd.Flags().BoolS("race", "race", false, "enable data race detection")
	vetCmd.Flags().StringS("tags", "tags", "", "a comma-separated list of build tags to consider satisfied during the")
	vetCmd.Flags().StringS("toolexec", "toolexec", "", "a program to use to invoke toolchain programs like vet and asm")
	vetCmd.Flags().BoolS("trimpath", "trimpath", false, "remove all file system paths from the resulting executable")
	vetCmd.Flags().BoolS("v", "v", false, "print the names of packages as they are compiled")
	vetCmd.Flags().BoolS("work", "work", false, "print the name of the temporary work directory")
	vetCmd.Flags().BoolS("x", "x", false, "print the commands")
	rootCmd.AddCommand(vetCmd)

	vetCmd.Flag("buildvcs").NoOptDefVal = "auto"

	carapace.Gen(vetCmd).FlagCompletion(carapace.ActionMap{
		"C":         carapace.ActionDirectories(),
		"asmflags":  bridge.ActionCarapaceBin("go", "tool", "asm").Split(),
		"buildmode": golang.ActionBuildmodes(),
		"buildvcs":  carapace.ActionValues("true", "false", "auto").StyleF(style.ForKeyword),
		"compiler":  carapace.ActionValues("gccgo", "gc"),
		"gcflags":   bridge.ActionCarapaceBin("go", "tool", "compile").Split(),
		"ldflags":   bridge.ActionCarapaceBin("go", "tool", "link").Split(),
		"overlay":   carapace.ActionFiles(".json"),
		"pgo":       carapace.ActionFiles(".pgo"),
		"pkgdir":    carapace.ActionDirectories(),
		"tags":      golang.ActionBuildTags().UniqueList(","),
		"toolexec":  bridge.ActionCarapaceBin().Split(),
	})
}
