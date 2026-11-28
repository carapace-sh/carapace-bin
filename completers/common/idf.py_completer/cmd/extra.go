package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func addGdbFlags(cmd *cobra.Command) {
	cmd.Flags().String("ex", "", "Execute given GDB command")
	cmd.Flags().String("gdb-commands", "", "Command line arguments for gdb")
	cmd.Flags().StringP("gdbinit", "x", "", "Specify the name of gdbinit file to use")
	cmd.Flags().Bool("require-openocd", false, "Fail this target if openocd failed")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"gdbinit": carapace.ActionFiles(),
	})
}

func addCoredumpFlags(cmd *cobra.Command) {
	cmd.Flags().String("chip-rev", "", "Specify the chip revision")
	cmd.Flags().StringP("core", "c", "", "Path to core dump file")
	cmd.Flags().StringP("save-core", "s", "", "Save core to file")

	cmd.Flag("chip-rev").Hidden = true

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"core":      carapace.ActionFiles(),
		"save-core": carapace.ActionFiles(),
	})
}

func init() {
	buildFileCmd := &cobra.Command{
		Use:   "build-file",
		Short: "Build a standalone C file without project boilerplate",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(buildFileCmd).Standalone()
	buildFileCmd.Flags().Bool("clean-cache", false, "Remove the cached container project before building")
	carapace.Gen(buildFileCmd).PositionalCompletion(carapace.ActionFiles(".c", ".cpp"))
	rootCmd.AddCommand(buildFileCmd)

	coredumpDebugCmd := &cobra.Command{
		Use:   "coredump-debug",
		Short: "Create core dump ELF file and run GDB debug session",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(coredumpDebugCmd).Standalone()
	addCoredumpFlags(coredumpDebugCmd)
	rootCmd.AddCommand(coredumpDebugCmd)

	coredumpInfoCmd := &cobra.Command{
		Use:   "coredump-info",
		Short: "Print crashed task's registers, callstack, and memory regions",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(coredumpInfoCmd).Standalone()
	addCoredumpFlags(coredumpInfoCmd)
	coredumpInfoCmd.Flags().String("gdb-timeout-sec", "", "Overwrite the default internal delay for gdb responses")
	coredumpInfoCmd.Flags().Bool("print-all-vars", false, "Print all variables tagged with COREDUMP_*_ATTR")
	coredumpInfoCmd.Flags().StringArray("print-var", nil, "Print a variable or expression")
	rootCmd.AddCommand(coredumpInfoCmd)

	createComponentCmd := &cobra.Command{
		Use:   "create-component",
		Short: "Create a new component",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(createComponentCmd).Standalone()
	rootCmd.AddCommand(createComponentCmd)

	createProjectCmd := &cobra.Command{
		Use:   "create-project",
		Short: "Create a new project",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(createProjectCmd).Standalone()
	createProjectCmd.Flags().Bool("cpp", false, "Create a C++ main source file instead of C")
	createProjectCmd.Flags().String("path", "", "Set the path for the new project")
	carapace.Gen(createProjectCmd).FlagCompletion(carapace.ActionMap{
		"path": carapace.ActionDirectories(),
	})
	rootCmd.AddCommand(createProjectCmd)

	dfuCmd := &cobra.Command{
		Use:   "dfu",
		Short: "Build the DFU binary",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(dfuCmd).Standalone()
	dfuCmd.Flags().String("part-size", "", "Overwrite default partition size of mkdfu.py")
	rootCmd.AddCommand(dfuCmd)

	dfuFlashCmd := &cobra.Command{
		Use:   "dfu-flash",
		Short: "Flash the DFU binary",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(dfuFlashCmd).Standalone()
	dfuFlashCmd.Flags().String("path", "", "Specify path to DFU device")
	rootCmd.AddCommand(dfuFlashCmd)

	dfuListCmd := &cobra.Command{
		Use:   "dfu-list",
		Short: "List DFU capable devices",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(dfuListCmd).Standalone()
	rootCmd.AddCommand(dfuListCmd)

	diagCmd := &cobra.Command{
		Use:   "diag",
		Short: "Create diagnostic report",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(diagCmd).Standalone()
	diagCmd.Flags().BoolP("append", "a", false, "Use recipes specified with -r in combination with built-in recipes")
	diagCmd.Flags().BoolP("check", "c", false, "Validate recipes")
	diagCmd.Flags().BoolP("debug", "d", false, "Print debug information")
	diagCmd.Flags().BoolP("force", "f", false, "Delete target file or directory if it already exists")
	diagCmd.Flags().BoolP("list", "l", false, "Show information about available recipes")
	diagCmd.Flags().Bool("log-prefix", false, "Add a severity character at the beginning of log messages")
	diagCmd.Flags().Bool("no-color", false, "Do not emit ANSI color codes")
	diagCmd.Flags().StringP("output", "o", "", "Diagnostic report directory or zip file archive path")
	diagCmd.Flags().String("purge", "", "Use a custom purge file to remove sensitive information")
	diagCmd.Flags().StringArrayP("recipe", "r", nil, "Recipe to use")
	diagCmd.Flags().StringArrayP("tag", "t", nil, "Consider only recipes containing TAG")
	diagCmd.Flags().StringP("zip", "z", "", "Create zip archive for diagnostic report in PATH")
	carapace.Gen(diagCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
		"purge":  carapace.ActionFiles(),
		"recipe": carapace.ActionFiles(),
		"zip":    carapace.ActionFiles(),
	})
	rootCmd.AddCommand(diagCmd)

	gdbCmd := &cobra.Command{
		Use:   "gdb",
		Short: "Run the GDB",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(gdbCmd).Standalone()
	addGdbFlags(gdbCmd)
	gdbCmd.Flags().Bool("batch", false, "Exit after processing gdbinit")
	gdbCmd.Flags().String("gdb-tui", "", "Run gdb in TUI mode")
	gdbCmd.Flag("batch").Hidden = true
	rootCmd.AddCommand(gdbCmd)

	gdbguiCmd := &cobra.Command{
		Use:   "gdbgui",
		Short: "GDB UI in default browser",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(gdbguiCmd).Standalone()
	addGdbFlags(gdbguiCmd)
	gdbguiCmd.Flags().String("gdbgui-port", "", "Port on which gdbgui will be hosted")
	rootCmd.AddCommand(gdbguiCmd)

	gdbtuiCmd := &cobra.Command{
		Use:   "gdbtui",
		Short: "GDB TUI mode",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(gdbtuiCmd).Standalone()
	addGdbFlags(gdbtuiCmd)
	rootCmd.AddCommand(gdbtuiCmd)

	mcpServerCmd := &cobra.Command{
		Use:   "mcp-server",
		Short: "Start MCP server for AI integration",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(mcpServerCmd).Standalone()
	rootCmd.AddCommand(mcpServerCmd)

	openocdCmd := &cobra.Command{
		Use:   "openocd",
		Short: "Run openocd from current path",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(openocdCmd).Standalone()
	openocdCmd.Flags().String("openocd-commands", "", "Command line arguments for openocd")
	openocdCmd.Flags().String("openocd-scripts", "", "Script directory for openocd cfg files")
	carapace.Gen(openocdCmd).FlagCompletion(carapace.ActionMap{
		"openocd-scripts": carapace.ActionDirectories(),
	})
	rootCmd.AddCommand(openocdCmd)

	postDebugCmd := &cobra.Command{
		Use:   "post-debug",
		Short: "Run post-debug actions",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(postDebugCmd).Standalone()
	postDebugCmd.Flags().String("block", "", "Set to 1 for blocking console on outputs of async debug actions")
	rootCmd.AddCommand(postDebugCmd)

	qemuCmd := &cobra.Command{
		Use:   "qemu",
		Short: "Run QEMU",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(qemuCmd).Standalone()
	qemuCmd.Flags().String("efuse-file", "", "File used to store efuse values")
	qemuCmd.Flags().String("flash-file", "", "File used as the qemu flash image")
	qemuCmd.Flags().BoolP("gdb", "d", false, "Wait for gdb to connect")
	qemuCmd.Flags().BoolP("graphics", "g", false, "Enable graphical window")
	qemuCmd.Flags().String("qemu-extra-args", "", "Extra arguments to pass to QEMU")
	carapace.Gen(qemuCmd).FlagCompletion(carapace.ActionMap{
		"efuse-file": carapace.ActionFiles(),
		"flash-file": carapace.ActionFiles(),
	})
	rootCmd.AddCommand(qemuCmd)

	uf2Cmd := &cobra.Command{
		Use:   "uf2",
		Short: "Generate the UF2 binary with all binaries included",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(uf2Cmd).Standalone()
	uf2Cmd.Flags().Bool("md5-disable", false, "Disable MD5 checksum")
	rootCmd.AddCommand(uf2Cmd)

	uf2AppCmd := &cobra.Command{
		Use:   "uf2-app",
		Short: "Generate a UF2 binary for the application only",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(uf2AppCmd).Standalone()
	uf2AppCmd.Flags().Bool("md5-disable", false, "Disable MD5 checksum")
	rootCmd.AddCommand(uf2AppCmd)
}
