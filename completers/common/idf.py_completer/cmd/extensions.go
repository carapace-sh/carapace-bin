package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/espidf"
	"github.com/spf13/cobra"
)

func init() {
	addDependencyCmd := &cobra.Command{
		Use:   "add-dependency",
		Short: "Add dependency to the manifest file",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
	addDependencyCmd.Flags().String("component", "", "Name of the component in the project")
	addDependencyCmd.Flags().String("git", "", "Git URL of the component")
	addDependencyCmd.Flags().String("git-path", "", "Path to the component in the git repository")
	addDependencyCmd.Flags().String("git-ref", "", "Git reference (branch, tag, commit SHA)")
	addDependencyCmd.Flags().String("path", "", "Path to the component where the dependency will be added")
	carapace.Gen(addDependencyCmd).FlagCompletion(carapace.ActionMap{
		"path": carapace.ActionDirectories(),
	})
	rootCmd.AddCommand(addDependencyCmd)

	clangCheckCmd := &cobra.Command{
		Use:   "clang-check",
		Short: "Run clang-tidy check under current folder",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
	clangCheckCmd.Flags().Bool("all-files", false, "Run clang-tidy with all files")
	clangCheckCmd.Flags().StringArray("exclude-paths", nil, "Exclude extra files besides of the project dir")
	clangCheckCmd.Flags().Bool("exit-code", false, "Exit with code based on the results of the code analysis")
	clangCheckCmd.Flags().StringArray("include-paths", nil, "Include extra files besides of the project dir")
	clangCheckCmd.Flags().String("run-clang-tidy-options", "", "Optional arguments passed to run-clang-tidy.py")
	clangCheckCmd.Flags().String("run-clang-tidy-py", "", "Path to run-clang-tidy.py")
	carapace.Gen(clangCheckCmd).FlagCompletion(carapace.ActionMap{
		"exclude-paths":     carapace.ActionFiles(),
		"include-paths":     carapace.ActionFiles(),
		"run-clang-tidy-py": carapace.ActionFiles(),
	})
	rootCmd.AddCommand(clangCheckCmd)

	clangHtmlReportCmd := &cobra.Command{
		Use:   "clang-html-report",
		Short: "Generate html report to html_report folder",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
	rootCmd.AddCommand(clangHtmlReportCmd)

	createManifestCmd := &cobra.Command{
		Use:   "create-manifest",
		Short: "Create manifest for specified component",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
	createManifestCmd.Flags().String("component", "", "Name of the component in the project")
	createManifestCmd.Flags().String("path", "", "Path to the component")
	carapace.Gen(createManifestCmd).FlagCompletion(carapace.ActionMap{
		"path": carapace.ActionDirectories(),
	})
	rootCmd.AddCommand(createManifestCmd)

	createProjectFromExampleCmd := &cobra.Command{
		Use:   "create-project-from-example",
		Short: "Create a project from an example in the ESP Component Registry",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
	createProjectFromExampleCmd.Flags().String("path", "", "Set the path for the new project")
	carapace.Gen(createProjectFromExampleCmd).FlagCompletion(carapace.ActionMap{
		"path": carapace.ActionDirectories(),
	})
	rootCmd.AddCommand(createProjectFromExampleCmd)

	removeDependencyCmd := &cobra.Command{
		Use:   "remove-dependency",
		Short: "Remove specified dependency from all manifest files",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
	carapace.Gen(removeDependencyCmd).PositionalCompletion(
		espidf.ActionDependencies(),
	)
	rootCmd.AddCommand(removeDependencyCmd)

	sbomCheckCmd := &cobra.Command{
		Use:   "sbom-check",
		Short: "Check application SBOM file or path for vulnerabilities",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
	sbomCheckCmd.Flags().String("file", "", "SBOM file to check")
	sbomCheckCmd.Flags().String("format", "", "Report format")
	sbomCheckCmd.Flags().Bool("no-sync-db", false, "Prevent automatic update of local NVD database")
	sbomCheckCmd.Flags().Bool("nvd-api", false, "Use NVD REST API for vulnerabilities scan")
	sbomCheckCmd.Flags().String("path", "", "Path to recursively search for SBOM YAML manifest files")
	sbomCheckCmd.Flags().String("report-file", "", "Report output file")
	sbomCheckCmd.Flags().String("spdx-file", "", "SBOM file to check (deprecated alias for --file)")
	sbomCheckCmd.Flag("spdx-file").Hidden = true
	carapace.Gen(sbomCheckCmd).FlagCompletion(carapace.ActionMap{
		"file":        carapace.ActionFiles(),
		"format":      carapace.ActionValues("csv", "json", "markdown", "table"),
		"path":        carapace.ActionDirectories(),
		"report-file": carapace.ActionFiles(),
		"spdx-file":   carapace.ActionFiles(),
	})
	rootCmd.AddCommand(sbomCheckCmd)

	sbomCreateCmd := &cobra.Command{
		Use:   "sbom-create",
		Short: "Create application SBOM in one of the supported formats",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
	sbomCreateCmd.Flags().String("format", "", "Output SBOM format")
	sbomCreateCmd.Flags().String("output-file", "", "Output SBOM file")
	sbomCreateCmd.Flags().String("spdx-file", "", "Output SBOM file (deprecated alias for --output-file)")
	sbomCreateCmd.Flags().String("vex-format", "", "Vulnerability information format")
	sbomCreateCmd.Flags().String("vex-output", "", "Output VEX file path")
	sbomCreateCmd.Flag("spdx-file").Hidden = true
	carapace.Gen(sbomCreateCmd).FlagCompletion(carapace.ActionMap{
		"format":      carapace.ActionValues("cyclonedx-json", "spdx-json", "spdx-json-ld", "spdx-tag-value"),
		"output-file": carapace.ActionFiles(),
		"spdx-file":   carapace.ActionFiles(),
		"vex-format":  carapace.ActionValues("cyclonedx-json", "embed", "none", "openvex"),
		"vex-output":  carapace.ActionFiles(),
	})
	rootCmd.AddCommand(sbomCreateCmd)

	updateDependenciesCmd := &cobra.Command{
		Use:   "update-dependencies",
		Short: "Update dependencies of the project",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
	rootCmd.AddCommand(updateDependenciesCmd)
}
