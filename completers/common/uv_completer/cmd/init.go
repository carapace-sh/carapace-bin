package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a new project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().Bool("app", false, "Create a project for an application")
	initCmd.Flags().Bool("application", false, "Create a project for an application")
	initCmd.Flags().String("author-from", "", "Fill in the `authors` field in the `pyproject.toml`")
	initCmd.Flags().Bool("backend", false, "Invalid option name for build backend")
	initCmd.Flags().Bool("bare", false, "Only create a `pyproject.toml`")
	initCmd.Flags().String("build-backend", "", "Initialize a build-backend of choice for the project")
	initCmd.Flags().String("description", "", "Set the project description")
	initCmd.Flags().Bool("lib", false, "Create a project for a library")
	initCmd.Flags().Bool("library", false, "Create a project for a library")
	initCmd.Flags().String("name", "", "The name of the project")
	initCmd.Flags().Bool("no-description", false, "Disable the description for the project")
	initCmd.Flags().Bool("no-package", false, "Do not set up the project to be built as a Python package")
	initCmd.Flags().Bool("no-pin-python", false, "Do not create a `.python-version` file for the project")
	initCmd.Flags().Bool("no-project", false, "Avoid discovering a workspace and create a standalone project")
	initCmd.Flags().Bool("no-readme", false, "Do not create a `README.md` file")
	initCmd.Flags().Bool("no-workspace", false, "Avoid discovering a workspace and create a standalone project")
	initCmd.Flags().Bool("package", false, "Set up the project to be built as a Python package")
	initCmd.Flags().Bool("pin-python", false, "Create a `.python-version` file for the project")
	initCmd.Flags().StringP("python", "p", "", "The Python interpreter to use to determine the minimum supported Python version.")
	initCmd.Flags().Bool("script", false, "Create a script")
	initCmd.Flags().String("vcs", "", "Initialize a version control system for the project")
	initCmd.Flags().Bool("virtual", false, "Create a virtual project, rather than a package")
	initCmd.Flag("application").Hidden = true
	initCmd.Flag("backend").Hidden = true
	initCmd.Flag("library").Hidden = true
	initCmd.Flag("no-project").Hidden = true
	initCmd.Flag("pin-python").Hidden = true
	initCmd.Flag("virtual").Hidden = true
	rootCmd.AddCommand(initCmd)
	carapace.Gen(initCmd).FlagCompletion(carapace.ActionMap{
		"author-from": carapace.ActionValuesDescribed(
			"auto", "Fetch the author information from some sources (e.g., Git) automatically",
			"git", "Fetch the author information from Git configuration only",
			"none", "Do not infer the author information",
		),
		"build-backend": carapace.ActionValuesDescribed(
			"uv", "Use uv as the project build backend",
			"hatch", "Use [hatchling](https://pypi.org/project/hatchling) as the project build backend",
			"flit", "Use [flit-core](https://pypi.org/project/flit-core) as the project build backend",
			"pdm", "Use [pdm-backend](https://pypi.org/project/pdm-backend) as the project build backend",
			"poetry", "Use [poetry-core](https://pypi.org/project/poetry-core) as the project build backend",
			"setuptools", "Use [setuptools](https://pypi.org/project/setuptools) as the project build backend",
			"maturin", "Use [maturin](https://pypi.org/project/maturin) as the project build backend",
			"scikit", "Use [scikit-build-core](https://pypi.org/project/scikit-build-core) as the project build backend",
		),
		"python": uv.ActionPythonInstallations(uv.InstallationsOpts{}),
		"vcs": carapace.ActionValuesDescribed(
			"git", "Use Git for version control",
			"none", "Do not use any version control system",
		),
	})
	carapace.Gen(initCmd).PositionalCompletion(carapace.ActionDirectories())
}
