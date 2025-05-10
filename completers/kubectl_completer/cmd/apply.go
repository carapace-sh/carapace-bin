package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
	Use:     "apply (-f FILENAME | -k DIRECTORY)",
	Short:   "Apply a configuration to a resource by file name or stdin",
	GroupID: "advanced",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applyCmd).Standalone()

	applyCmd.Flags().Bool("all", false, "Select all resources in the namespace of the specified resource types.")
	applyCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	applyCmd.Flags().String("cascade", "", "Must be \"background\", \"orphan\", or \"foreground\". Selects the deletion cascading strategy for the dependents (e.g. Pods created by a ReplicationController). Defaults to background.")
	applyCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	applyCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	applyCmd.Flags().StringSliceP("filename", "f", nil, "The files that contain the configurations to apply.")
	applyCmd.Flags().Bool("force", false, "If true, immediately remove resources from API and bypass graceful deletion. Note that immediate deletion of some resources may result in inconsistency or data loss and requires confirmation.")
	applyCmd.Flags().Bool("force-conflicts", false, "If true, server-side apply will force the changes against conflicts.")
	applyCmd.Flags().String("grace-period", "", "Period of time in seconds given to the resource to terminate gracefully. Ignored if negative. Set to 1 for immediate shutdown. Can only be set to 0 when --force is true (force deletion).")
	applyCmd.Flags().StringP("kustomize", "k", "", "Process a kustomization directory. This flag can't be used together with -f or -R.")
	applyCmd.Flags().Bool("openapi-patch", false, "If true, use openapi to calculate diff when the openapi presents and the resource can be found in the openapi spec. Otherwise, fall back to use baked-in types.")
	applyCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	applyCmd.Flags().Bool("overwrite", false, "Automatically resolve conflicts between the modified and live configuration by using values from the modified configuration")
	applyCmd.Flags().Bool("prune", false, "Automatically delete resource objects, that do not appear in the configs and are created by either apply or create --save-config. Should be used with either -l or --all.")
	applyCmd.Flags().StringSlice("prune-allowlist", nil, "Overwrite the default allowlist with <group/version/kind> for --prune")
	applyCmd.Flags().Bool("record", false, "Record current kubectl command in the resource annotation. If set to false, do not record the command. If set to true, record the command. If not set, default to updating the existing annotation value only if one already exists.")
	applyCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.")
	applyCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.")
	applyCmd.Flags().Bool("server-side", false, "If true, apply runs in the server instead of the client.")
	applyCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	applyCmd.Flags().String("subresource", "", "If specified, apply will operate on the subresource of the requested object.  Only allowed when using --server-side.")
	applyCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	applyCmd.Flags().String("timeout", "", "The length of time to wait before giving up on a delete, zero means determine a timeout from the size of the object")
	applyCmd.Flags().String("validate", "", "Must be one of: strict (or true), warn, ignore (or false). \"true\" or \"strict\" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. \"warn\" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as \"ignore\" otherwise. \"false\" or \"ignore\" will not perform any schema validation, silently dropping any unknown or duplicate fields.")
	applyCmd.Flags().Bool("wait", false, "If true, wait for resources to be gone before returning. This waits for finalizers.")
	applyCmd.Flag("cascade").NoOptDefVal = " "
	applyCmd.Flag("dry-run").NoOptDefVal = " "
	applyCmd.Flag("record").Hidden = true
	applyCmd.Flag("validate").NoOptDefVal = " "
	rootCmd.AddCommand(applyCmd)

	carapace.Gen(applyCmd).FlagCompletion(carapace.ActionMap{
		"cascade":   carapace.ActionValues("background", "orphan", "foreground"),
		"dry-run":   kubectl.ActionDryRunModes(),
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    kubectl.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
		"validate":  kubectl.ActionValidationModes(),
	})
}
