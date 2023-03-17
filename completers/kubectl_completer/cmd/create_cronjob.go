package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var create_cronjobCmd = &cobra.Command{
	Use:     "cronjob NAME --image=image --schedule='0/5 * * * ?' -- [COMMAND] [args...]",
	Short:   "Create a cron job with the specified name",
	Aliases: []string{"cj"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_cronjobCmd).Standalone()

	create_cronjobCmd.Flags().Bool("allow-missing-template-keys", true, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	create_cronjobCmd.Flags().String("dry-run", "none", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	create_cronjobCmd.Flags().String("field-manager", "kubectl-create", "Name of the manager used to track field ownership.")
	create_cronjobCmd.Flags().String("image", "", "Image name to run.")
	create_cronjobCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	create_cronjobCmd.Flags().String("restart", "", "job's restart policy. supported values: OnFailure, Never")
	create_cronjobCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.")
	create_cronjobCmd.Flags().String("schedule", "", "A schedule in the Cron format the job should be run with.")
	create_cronjobCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	create_cronjobCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	create_cronjobCmd.Flags().String("validate", "strict", "Must be one of: strict (or true), warn, ignore (or false).")
	create_cronjobCmd.Flag("dry-run").NoOptDefVal = "unchanged"
	create_cronjobCmd.Flag("validate").NoOptDefVal = "strict"
	createCmd.AddCommand(create_cronjobCmd)

	carapace.Gen(create_cronjobCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":  kubectl.ActionDryRunModes(),
		"output":   kubectl.ActionOutputFormats(),
		"restart":  carapace.ActionValues("OnFailure", "Never"),
		"template": carapace.ActionFiles(),
		"validate": kubectl.ActionValidationModes(),
	})
}
