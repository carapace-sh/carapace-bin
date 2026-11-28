package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func AddSizeFlags(cmd *cobra.Command) {
	cmd.Flags().String("diff", "", "Show the differences in comparison with another project")
	cmd.Flags().String("format", "", "Specify output format")
	cmd.Flags().BoolP("legacy", "l", false, "Use legacy esp-idf-size version")
	cmd.Flags().String("output-file", "", "Print output to the specified file instead of stdout")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"diff":        carapace.ActionFiles(),
		"format":      carapace.ActionValues("csv", "default", "json", "json2", "raw", "text", "tree"),
		"output-file": carapace.ActionFiles(),
	})
}

func AddFlashFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("all", "a", false, "Flash all data")
	cmd.Flags().String("extra-args", "", "Pass extra arguments to esptool separated by space")
	cmd.Flags().Bool("force", false, "Force write, skip security and compatibility checks")
	cmd.Flags().Bool("trace", false, "Enable trace-level output of flasher tool interactions")
	cmd.Flags().BoolP("trust-flash-content", "t", false, "Skip MD5 verification of files which do not need reflashing")
}

func AddEfuseFlags(cmd *cobra.Command) {
	cmd.Flags().String("before", "", "What to do before connecting to the chip")
	cmd.Flags().BoolP("debug", "d", false, "Print debug information")
	cmd.Flags().Bool("do-not-confirm", false, "Do not pause for confirmation before permanently writing eFuses")
	cmd.Flags().Bool("virt", false, "Work in virtual mode without connecting to a chip")

	cmd.Flag("virt").Hidden = true

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"before": carapace.ActionValues("default-reset", "no-reset", "no-reset-no-sync", "usb-reset"),
	})
}

func AddGdbFlags(cmd *cobra.Command) {
	cmd.Flags().String("ex", "", "Execute given GDB command")
	cmd.Flags().String("gdb-commands", "", "Command line arguments for gdb")
	cmd.Flags().StringP("gdbinit", "x", "", "Specify the name of gdbinit file to use")
	cmd.Flags().Bool("require-openocd", false, "Fail this target if openocd failed")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"gdbinit": carapace.ActionFiles(),
	})
}

func AddCoredumpFlags(cmd *cobra.Command) {
	cmd.Flags().String("chip-rev", "", "Specify the chip revision")
	cmd.Flags().StringP("core", "c", "", "Path to core dump file")
	cmd.Flags().StringP("save-core", "s", "", "Save core to file")

	cmd.Flag("chip-rev").Hidden = true

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"core":      carapace.ActionFiles(),
		"save-core": carapace.ActionFiles(),
	})
}
