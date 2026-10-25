package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var oci_runCmd = &cobra.Command{
	Use:   "run",
	Short: "[experimental] Build the project's image and run a command in it",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(oci_runCmd).Standalone()

	oci_runCmd.Flags().String("engine", "auto", "Container engine to use; `auto` prefers podman and falls back to docker")
	oci_runCmd.Flags().StringSliceP("env", "e", nil, "Set environment variable in the container (repeatable, `KEY=VAL`)")
	oci_runCmd.Flags().String("from", "", "Base image for the build")
	oci_runCmd.Flags().BoolP("help", "h", false, "Print help")
	oci_runCmd.Flags().String("image-dir", "", "Use an existing OCI image layout instead of building one")
	oci_runCmd.Flags().Bool("include-global", false, "Also package tools from the global and system configs")
	oci_runCmd.Flags().BoolP("interactive", "i", false, "Run interactively (pass `-i` to the engine)")
	oci_runCmd.Flags().Bool("keep", false, "Keep the loaded image in the engine's storage after the run")
	oci_runCmd.Flags().StringSlice("mount", nil, "Bind-mount a host path (repeatable, `HOST:CONTAINER[:MODE]`)")
	oci_runCmd.Flags().String("mount-point", "", "Where tools install inside the image")
	oci_runCmd.Flags().Bool("no-mise", false, "Do not embed the running mise binary at /usr/local/bin/mise")
	oci_runCmd.Flags().String("owner", "", "UID[:GID] to assign to every tar entry when building")
	oci_runCmd.Flags().BoolP("tty", "t", false, "Allocate a TTY (pass `-t` to the engine)")
	oci_runCmd.Flags().StringSlice("volume", nil, "Bind-mount a host path (repeatable, `HOST:CONTAINER[:MODE]`)")
	oci_runCmd.Flags().StringP("workdir", "w", "", "Working directory inside the container")
	oci_runCmd.Flag("mount").Hidden = true
	ociCmd.AddCommand(oci_runCmd)

	carapace.Gen(oci_runCmd).FlagCompletion(carapace.ActionMap{
		"engine":    carapace.ActionValues("auto", "docker", "podman"),
		"image-dir": carapace.ActionDirectories(),
	})
}
