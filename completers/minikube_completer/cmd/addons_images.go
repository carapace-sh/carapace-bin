package cmd

import (
	"github.com/spf13/cobra"
)

var addons_imagesCmd = &cobra.Command{
	Use:   "images",
	Short: "List image names the addon w/ADDON_NAME used. For a list of available addons use: minikube addons list",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	addonsCmd.AddCommand(addons_imagesCmd)
}
