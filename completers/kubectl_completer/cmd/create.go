package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:     "create -f FILENAME",
	Short:   "Create a resource from a file or from stdin",
	GroupID: "basic beginner",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createCmd).Standalone()

	createCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	createCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	createCmd.Flags().Bool("edit", false, "Edit the API resource before creating")
	createCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	createCmd.Flags().StringSliceP("filename", "f", nil, "Filename, directory, or URL to files to use to create the resource")
	createCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	createCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	createCmd.Flags().String("raw", "", "Raw URI to POST to the server.  Uses the transport specified by the kubeconfig file.")
	createCmd.Flags().Bool("record", false, "Record current kubectl command in the resource annotation. If set to false, do not record the command. If set to true, record the command. If not set, default to updating the existing annotation value only if one already exists.")
	createCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.")
	createCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.")
	createCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.")
	createCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	createCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	createCmd.Flags().String("validate", "", "Must be one of: strict (or true), warn, ignore (or false). \"true\" or \"strict\" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. \"warn\" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as \"ignore\" otherwise. \"false\" or \"ignore\" will not perform any schema validation, silently dropping any unknown or duplicate fields.")
	createCmd.Flags().Bool("windows-line-endings", false, "Only relevant if --edit=true. Defaults to the line ending native to your platform.")
	createCmd.Flag("dry-run").NoOptDefVal = " "
	createCmd.Flag("record").Hidden = true
	createCmd.Flag("validate").NoOptDefVal = " "
	rootCmd.AddCommand(createCmd)

	carapace.Gen(createCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":   kubectl.ActionDryRunModes(),
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    kubectl.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
		"validate":  kubectl.ActionValidationModes(),
	})
}
