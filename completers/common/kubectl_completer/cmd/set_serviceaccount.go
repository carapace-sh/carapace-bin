package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var set_serviceaccountCmd = &cobra.Command{
	Use:     "serviceaccount (-f FILENAME | TYPE NAME) SERVICE_ACCOUNT",
	Short:   "Update the service account of a resource",
	Aliases: []string{"sa"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(set_serviceaccountCmd).Standalone()

	set_serviceaccountCmd.Flags().Bool("all", false, "Select all resources, in the namespace of the specified resource types")
	set_serviceaccountCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	set_serviceaccountCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	set_serviceaccountCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	set_serviceaccountCmd.Flags().StringSliceP("filename", "f", nil, "Filename, directory, or URL to files identifying the resource to get from a server.")
	set_serviceaccountCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	set_serviceaccountCmd.Flags().Bool("local", false, "If true, set serviceaccount will NOT contact api-server but run locally.")
	set_serviceaccountCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, kyaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	set_serviceaccountCmd.Flags().Bool("record", false, "Record current kubectl command in the resource annotation. If set to false, do not record the command. If set to true, record the command. If not set, default to updating the existing annotation value only if one already exists.")
	set_serviceaccountCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.")
	set_serviceaccountCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	set_serviceaccountCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	set_serviceaccountCmd.Flag("dry-run").NoOptDefVal = " "
	set_serviceaccountCmd.Flag("record").Hidden = true
	setCmd.AddCommand(set_serviceaccountCmd)

	carapace.Gen(set_serviceaccountCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":   kubectl.ActionDryRunModes(),
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    kubectl.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
	})

	carapace.Gen(set_serviceaccountCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if set_serviceaccountCmd.Flag("filename").Changed {
				return kubectl.ActionResources(kubectl.ResourceOpts{
					Context:   rootCmd.Flag("context").Value.String(),
					Namespace: rootCmd.Flag("namespace").Value.String(),
					Types:     "serviceaccounts",
				})
			} else {
				return kubectl.ActionApiResources()
			}
		}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if set_serviceaccountCmd.Flag("filename").Changed {
				return carapace.ActionValues()
			} else {
				return kubectl.ActionResources(kubectl.ResourceOpts{
					Context:   rootCmd.Flag("context").Value.String(),
					Namespace: rootCmd.Flag("namespace").Value.String(),
					Types:     c.Args[0],
				})
			}
		}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if set_serviceaccountCmd.Flag("filename").Changed {
				return carapace.ActionValues()
			} else {
				return kubectl.ActionResources(kubectl.ResourceOpts{
					Context:   rootCmd.Flag("context").Value.String(),
					Namespace: rootCmd.Flag("namespace").Value.String(),
					Types:     "serviceaccounts",
				})
			}
		}),
	)
}
