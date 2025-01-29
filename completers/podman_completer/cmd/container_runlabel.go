package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_runlabelCmd = &cobra.Command{
	Use:   "runlabel [options] LABEL IMAGE [ARG...]",
	Short: "Execute the command described by an image label",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_runlabelCmd).Standalone()

	container_runlabelCmd.Flags().String("authfile", "", "Path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override")
	container_runlabelCmd.Flags().String("cert-dir", "", "`Pathname` of a directory containing TLS certificates and keys")
	container_runlabelCmd.Flags().String("creds", "", "`Credentials` (USERNAME:PASSWORD) to use for authenticating to a registry")
	container_runlabelCmd.Flags().Bool("display", false, "Preview the command that the label would run")
	container_runlabelCmd.Flags().StringP("name", "n", "", "Assign a name to the container")
	container_runlabelCmd.Flags().String("opt1", "", "Optional parameter to pass for install")
	container_runlabelCmd.Flags().String("opt2", "", "Optional parameter to pass for install")
	container_runlabelCmd.Flags().String("opt3", "", "Optional parameter to pass for install")
	container_runlabelCmd.Flags().BoolP("pull", "p", false, "Pull the image if it does not exist locally prior to executing the label contents")
	container_runlabelCmd.Flags().BoolP("quiet", "q", false, "Suppress output information when installing images")
	container_runlabelCmd.Flags().Bool("replace", false, "Replace existing container with a new one from the image")
	container_runlabelCmd.Flags().String("signature-policy", "", "`Pathname` of signature policy file (not usually used)")
	container_runlabelCmd.Flags().Bool("tls-verify", false, "Require HTTPS and verify certificates when contacting registries")
	container_runlabelCmd.Flag("opt1").Hidden = true
	container_runlabelCmd.Flag("opt2").Hidden = true
	container_runlabelCmd.Flag("opt3").Hidden = true
	container_runlabelCmd.Flag("pull").Hidden = true
	container_runlabelCmd.Flag("signature-policy").Hidden = true
	containerCmd.AddCommand(container_runlabelCmd)
}
