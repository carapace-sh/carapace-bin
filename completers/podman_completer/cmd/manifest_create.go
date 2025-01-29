package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var manifest_createCmd = &cobra.Command{
	Use:   "create [options] LIST [IMAGE...]",
	Short: "Create manifest list or image index",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(manifest_createCmd).Standalone()

	manifest_createCmd.Flags().Bool("all", false, "add all of the lists' images if the images to add are lists")
	manifest_createCmd.Flags().BoolP("amend", "a", false, "modify an existing list if one with the desired name already exists")
	manifest_createCmd.Flags().StringSlice("annotation", []string{}, "set annotations on the new list")
	manifest_createCmd.Flags().Bool("insecure", false, "neither require HTTPS nor verify certificates when accessing the registry")
	manifest_createCmd.Flags().Bool("tls-verify", false, "require HTTPS and verify certificates when accessing the registry")
	manifest_createCmd.Flag("insecure").Hidden = true
	manifestCmd.AddCommand(manifest_createCmd)
}
