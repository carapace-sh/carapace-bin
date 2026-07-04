package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dotnet",
	Short: ".NET CLI for building, running, and managing .NET projects",
	Long:  "https://learn.microsoft.com/en-us/dotnet/core/tools/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"build", "build a .NET project",
			"run", "run a .NET project",
			"publish", "publish a .NET project for deployment",
			"pack", "create a NuGet package",
			"test", "run unit tests",
			"restore", "restore NuGet packages",
			"new", "create a new .NET project",
			"add", "add a reference to a project",
			"remove", "remove a reference from a project",
			"list", "list project references",
			"nuget", "manage NuGet packages",
			"msbuild", "run MSBuild",
			"vstest", "run tests with VSTest",
			"clean", "clean a .NET project",
			"sln", "modify solution files",
			"help", "display help",
			"tool", "manage .NET tools",
			"format", "format code",
			"ef", "Entity Framework Core tools",
		),
	)
}
