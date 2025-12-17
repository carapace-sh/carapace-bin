package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var autoscaleCmd = &cobra.Command{
	Use:     "autoscale (-f FILENAME | TYPE NAME | TYPE/NAME) [--min=MINPODS] --max=MAXPODS [--cpu=CPU] [--memory=MEMORY]",
	Short:   "Auto-scale a deployment, replica set, stateful set, or replication controller",
	GroupID: "deploy",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaleCmd).Standalone()

	autoscaleCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	autoscaleCmd.Flags().String("cpu", "", "Target CPU utilization over all the pods. When specified as a percentage (e.g.\"70%\" for 70% of requested CPU) it will target average utilization. When specified as quantity (e.g.\"500m\" for 500 milliCPU) it will target average value. Value without units is treated as a quantity with miliCPU being the unit (e.g.\"500\" is \"500m\").")
	autoscaleCmd.Flags().String("cpu-percent", "", "The target average CPU utilization (represented as a percent of requested CPU) over all the pods. If it's not specified or negative, a default autoscaling policy will be used.")
	autoscaleCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	autoscaleCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	autoscaleCmd.Flags().StringSliceP("filename", "f", nil, "Filename, directory, or URL to files identifying the resource to autoscale.")
	autoscaleCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	autoscaleCmd.Flags().String("max", "", "The upper limit for the number of pods that can be set by the autoscaler. Required.")
	autoscaleCmd.Flags().String("memory", "", "Target memory utilization over all the pods. When specified  as a percentage (e.g.\"60%\" for 60% of requested memory) it will target average utilization. When specified as quantity (e.g.\"200Mi\" for 200 MiB, \"1Gi\" for 1 GiB) it will target average value. Value without units is treated as a quantity with mebibytes being the unit (e.g.\"200\" is \"200Mi\").")
	autoscaleCmd.Flags().String("min", "", "The lower limit for the number of pods that can be set by the autoscaler. If it's not specified or negative, the server will apply a default value.")
	autoscaleCmd.Flags().String("name", "", "The name for the newly created object. If not specified, the name of the input resource will be used.")
	autoscaleCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, kyaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	autoscaleCmd.Flags().Bool("record", false, "Record current kubectl command in the resource annotation. If set to false, do not record the command. If set to true, record the command. If not set, default to updating the existing annotation value only if one already exists.")
	autoscaleCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.")
	autoscaleCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.")
	autoscaleCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	autoscaleCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	autoscaleCmd.Flag("cpu-percent").Hidden = true
	autoscaleCmd.Flag("dry-run").NoOptDefVal = " "
	autoscaleCmd.MarkFlagRequired("max")
	autoscaleCmd.Flag("record").Hidden = true
	rootCmd.AddCommand(autoscaleCmd)

	// TODO field-manager
	carapace.Gen(autoscaleCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":   kubectl.ActionDryRunModes(),
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    kubectl.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
	})

	carapace.Gen(autoscaleCmd).PositionalCompletion(
		carapace.ActionValues("deployments", "replicasets", "replicationcontrollers"),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return kubectl.ActionResources(kubectl.ResourceOpts{
				Context:   rootCmd.Flag("context").Value.String(),
				Namespace: rootCmd.Flag("namespace").Value.String(),
				Types:     c.Args[0],
			})
		}),
	)
}
