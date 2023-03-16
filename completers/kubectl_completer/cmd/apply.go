package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
	Use:     "apply",
	Short:   "Apply a configuration to a resource by file name or stdin",
	GroupID: "advanced",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applyCmd).Standalone()
	applyCmd.Flags().Bool("all", false, "Select all resources in the namespace of the specified resource types.")
	applyCmd.Flags().Bool("allow-missing-template-keys", true, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	applyCmd.Flags().String("cascade", "background", "Must be \"background\", \"orphan\", or \"foreground\". Selects the deletion cascading strategy for the dependents (e.g. Pods created by a ReplicationController). Defaults to background.")
	applyCmd.Flags().String("dry-run", "none", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	applyCmd.Flags().String("field-manager", "kubectl-client-side-apply", "Name of the manager used to track field ownership.")
	applyCmd.Flags().StringSliceP("filename", "f", []string{}, "The files that contain the configurations to apply.")
	applyCmd.Flags().Bool("force", false, "If true, immediately remove resources from API and bypass graceful deletion. Note that immediate deletion of some resources may result in inconsistency or data loss and requires confirmation.")
	applyCmd.Flags().Bool("force-conflicts", false, "If true, server-side apply will force the changes against conflicts.")
	applyCmd.Flags().Int("grace-period", -1, "Period of time in seconds given to the resource to terminate gracefully. Ignored if negative. Set to 1 for immediate shutdown. Can only be set to 0 when --force is true (force deletion).")
	applyCmd.Flags().StringP("kustomize", "k", "", "Process a kustomization directory. This flag can't be used together with -f or -R.")
	applyCmd.Flags().Bool("openapi-patch", true, "If true, use openapi to calculate diff when the openapi presents and the resource can be found in the openapi spec. Otherwise, fall back to use baked-in types.")
	applyCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	applyCmd.Flags().Bool("overwrite", true, "Automatically resolve conflicts between the modified and live configuration by using values from the modified configuration")
	applyCmd.Flags().Bool("prune", false, "Automatically delete resource objects, that do not appear in the configs and are created by either apply or create --save-config. Should be used with either -l or --all.")
	applyCmd.Flags().StringArray("prune-whitelist", []string{}, "Overwrite the default whitelist with <group/version/kind> for --prune")
	applyCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.")
	applyCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.")
	applyCmd.Flags().Bool("server-side", false, "If true, apply runs in the server instead of the client.")
	applyCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	applyCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	applyCmd.Flags().String("timeout", "0s", "The length of time to wait before giving up on a delete, zero means determine a timeout from the size of the object")
	applyCmd.Flags().String("validate", "strict", "Must be one of: strict (or true), warn, ignore (or false).")
	applyCmd.Flags().Bool("wait", false, "If true, wait for resources to be gone before returning. This waits for finalizers.")
	applyCmd.Flag("cascade").NoOptDefVal = "background"
	applyCmd.Flag("dry-run").NoOptDefVal = "unchanged"
	applyCmd.Flag("validate").NoOptDefVal = "strict"
	rootCmd.AddCommand(applyCmd)

	carapace.Gen(applyCmd).FlagCompletion(carapace.ActionMap{
		"cascade":   carapace.ActionValues("background", "orphan", "foreground"),
		"dry-run":   action.ActionDryRunModes(),
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    action.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
		"validate":  kubectl.ActionValidationModes(),
	})
}
