package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kcl",
	Short: "KCL command line interface",
	Long:  "https://www.kcl-lang.io/docs/tools/cli/kcl/overview",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error { return rootCmd.Execute() }

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "help for kcl")
	rootCmd.Flags().BoolP("version", "v", false, "version for kcl")

	commands := map[string]string{
		"clean":      "KCL clean tool",
		"completion": "Generate the autocompletion script for the specified shell",
		"doc":        "KCL document tool",
		"fmt":        "KCL format tool",
		"help":       "Help about any command",
		"import":     "KCL import tool",
		"lint":       "Lint KCL codes",
		"mod":        "KCL module management",
		"play":       "Open the KCL playground in the browser",
		"registry":   "KCL registry management",
		"run":        "Run KCL codes",
		"server":     "Run a KCL server",
		"test":       "KCL test tool",
		"version":    "Show version of the KCL CLI",
		"vet":        "KCL validation tool",
	}
	for name, desc := range commands {
		cmd := &cobra.Command{Use: name, Short: desc, DisableFlagParsing: true, Run: func(cmd *cobra.Command, args []string) {}}
		carapace.Gen(cmd).PositionalAnyCompletion(bridge.ActionCarapaceBin("kcl", name))
		rootCmd.AddCommand(cmd)
	}

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"clean", commands["clean"],
			"completion", commands["completion"],
			"doc", commands["doc"],
			"fmt", commands["fmt"],
			"help", commands["help"],
			"import", commands["import"],
			"lint", commands["lint"],
			"mod", commands["mod"],
			"play", commands["play"],
			"registry", commands["registry"],
			"run", commands["run"],
			"server", commands["server"],
			"test", commands["test"],
			"version", commands["version"],
			"vet", commands["vet"],
		).StyleF(style.ForKeyword),
	)
}
