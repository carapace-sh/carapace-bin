package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:     "build",
	Aliases: []string{"b"},
	Short:   "build",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	buildCmd.Flags().String("allow", "", "Allow extra privileged entitlement")
	buildCmd.Flags().String("debug-json-cache-metrics", "", "Where to output json cache metrics")
	buildCmd.Flags().String("export-cache", "", "Export build cache")
	buildCmd.Flags().String("frontend", "", "Define frontend used for build")
	buildCmd.Flags().String("import-cache", "", "Import build cache")
	buildCmd.Flags().String("local", "", "Allow build access to the local directory")
	buildCmd.Flags().String("metadata-file", "", "Output build metadata (e.g., image digest) to a file as JSON")
	buildCmd.Flags().Bool("no-cache", false, "Disable cache for all the vertices")
	buildCmd.Flags().String("oci-layout", "", "Allow build access to the local OCI layout")
	buildCmd.Flags().String("opt", "", "Define custom options for frontend")
	buildCmd.Flags().StringP("output", "o", "", "Define exports for build result")
	buildCmd.Flags().String("progress", "", "Set type of progress")
	buildCmd.Flags().String("ref-file", "", "Write build ref to a file")
	buildCmd.Flags().String("registry-auth-tlscontext", "", "Overwrite TLS configuration when authenticating with registries")
	buildCmd.Flags().String("secret", "", "Secret value exposed to the build")
	buildCmd.Flags().String("source-policy-file", "", "Read source policy file from a JSON file")
	buildCmd.Flags().String("ssh", "", "Allow forwarding SSH agent or a raw Unix socket to the builder")
	buildCmd.Flags().String("trace", "", "Path to trace file")
	rootCmd.AddCommand(buildCmd)

	// TODO flag completion
	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		"allow":                    carapace.ActionValues(),
		"debug-json-cache-metrics": carapace.ActionValues("stdout", "stderr"),
		"export-cache":             carapace.ActionValues(),
		"frontend":                 carapace.ActionValues(),
		"import-cache":             carapace.ActionValues(),
		"local":                    carapace.ActionValues(),
		"metadata-file":            carapace.ActionFiles(),
		"oci-layout":               carapace.ActionValues(),
		"opt":                      carapace.ActionValues(),
		"output":                   carapace.ActionValues(),
		"progress":                 carapace.ActionValues("auto", "plain", "tty", "rawjson"),
		"ref-file":                 carapace.ActionFiles(),
		"registry-auth-tlscontext": carapace.ActionValues(),
		"secret":                   carapace.ActionValues(),
		"source-policy-file":       carapace.ActionFiles(),
		"ssh":                      carapace.ActionValues(),
		"trace":                    carapace.ActionFiles(),
	})
}
