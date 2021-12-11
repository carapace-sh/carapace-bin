package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/tinygo_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tinygo",
	Short: "TinyGo is a Go compiler for small places",
	Long:  "https://tinygo.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	carapace.Override(carapace.Opts{
		LongShorthand: true,
	})
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("cpuprofile", "", "cpuprofile output")
	rootCmd.Flags().Bool("dumpssa", false, "dump internal Go SSA")
	rootCmd.Flags().String("gc", "", "garbage collector to use")
	rootCmd.Flags().String("ldflags", "", "Go link tool compatible ldflags")
	rootCmd.Flags().String("llvm-features", "", "comma separated LLVM features to enable")
	rootCmd.Flags().Bool("no-debug", false, "strip debug information")
	rootCmd.Flags().String("o", "", "output")
	rootCmd.Flags().String("ocd-commands", "", "OpenOCD commands, overriding target spec")
	rootCmd.Flags().Bool("ocd-output", false, "print OCD daemon output during debug")
	rootCmd.Flags().String("opt", "", "optimization level")
	rootCmd.Flags().String("panic", "", "panic strategy")
	rootCmd.Flags().String("port", "", "flash port")
	rootCmd.Flags().String("print-allocs", "", "regular expression of functions for which heap allocations should be printed")
	rootCmd.Flags().Bool("print-stacks", false, "print stack sizes of goroutines")
	rootCmd.Flags().Bool("printir", false, "print LLVM IR")
	rootCmd.Flags().String("programmer", "", "which hardware programmer to use")
	rootCmd.Flags().String("scheduler", "", "which scheduler to use")
	rootCmd.Flags().String("serial", "", "which serial output to use")
	rootCmd.Flags().String("size", "", "print sizes")
	rootCmd.Flags().String("tags", "", "a space-separated list of extra build tags")
	rootCmd.Flags().String("target", "", "chip/board name or JSON target specification file")
	rootCmd.Flags().Bool("verifyir", false, "run extra verification steps on LLVM IR")
	rootCmd.Flags().String("wasm-abi", "", "WebAssembly ABI conventions")
	rootCmd.Flags().Bool("x", false, "Print commands")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"gc":    carapace.ActionValues("none", "leaking", "extalloc", "conservative"),
		"o":     carapace.ActionFiles(),
		"opt":   carapace.ActionValues("0", "1", "2", "s", "z"),
		"panic": carapace.ActionValues("print", "trap"),
		"port": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return net.ActionPorts().Invoke(c).Filter(c.Parts).ToA()
		}),
		"programmer": carapace.ActionValues("openocd", "msd", "command", "bmp"),
		"scheduler":  carapace.ActionValues("none", "coroutines", "tasks"),
		"serial":     carapace.ActionValues("none", "uart", "usb"),
		"size":       carapace.ActionValues("none", "short", "full"),
		"target": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if util.HasPathPrefix(c.CallbackValue) {
				return carapace.ActionFiles()
			}
			return action.ActionTargets()
		}),
		"wasm-abi": carapace.ActionValues("js", "generic"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"build", "compile packages and dependencies",
			"run", "compile and run immediately",
			"test", "test packages",
			"flash", "compile and flash to the device",
			"gdb", "run/flash and immediately enter GDB",
			"env", "list environment variables used during build",
			"list", "run go list using the TinyGo root",
			"clean", "empty cache directory",
			"help", "print this help text",
		),
		carapace.ActionFiles(),
	)
}
