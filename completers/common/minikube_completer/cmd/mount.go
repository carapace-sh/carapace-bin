package cmd

import (
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var mountCmd = &cobra.Command{
	Use:     "mount [flags] <source directory>:<target directory>",
	Short:   "Mounts the specified directory into minikube",
	GroupID: "advanced",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mountCmd).Standalone()

	mountCmd.Flags().String("9p-version", "", "Specify the 9p version that the mount should use")
	mountCmd.Flags().String("gid", "", "Default group id used for the mount")
	mountCmd.Flags().String("ip", "", "Specify the ip that the mount should be setup on")
	mountCmd.Flags().Bool("kill", false, "Kill the mount process spawned by minikube start")
	mountCmd.Flags().String("msize", "", "The number of bytes to use for 9p packet payload")
	mountCmd.Flags().StringSlice("options", nil, "Additional mount options, such as cache=fscache")
	mountCmd.Flags().String("port", "", "Specify the port that the mount should be setup on, where 0 means any free port.")
	mountCmd.Flags().String("type", "", "Specify the mount filesystem type (supported types: 9p)")
	mountCmd.Flags().String("uid", "", "Default user id used for the mount")
	rootCmd.AddCommand(mountCmd)

	carapace.Gen(mountCmd).FlagCompletion(carapace.ActionMap{
		"gid":  os.ActionGroups(),
		"port": net.ActionPorts(),
		"type": carapace.ActionValues("9p"),
		"uid":  os.ActionUsers(),
	})

	carapace.Gen(mountCmd).PositionalCompletion(
		carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionDirectories().NoSpace()
			case 1:
				path := filepath.Dir(c.Value)
				if path == "" {
					path = "/"
				}
				return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
					return carapace.ActionExecCommand("minikube", "ssh", "ls", `\-1`, `\-p`, path)(func(output []byte) carapace.Action {
						lines := strings.Split(string(output), "\n")
						return carapace.ActionValues(lines[:len(lines)-1]...)
					})
				})
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
