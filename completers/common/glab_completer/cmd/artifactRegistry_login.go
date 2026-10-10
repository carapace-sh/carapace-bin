package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var artifactRegistry_loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate a package manager against the GitLab Artifact Registry. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(artifactRegistry_loginCmd).Standalone()

	artifactRegistry_loginCmd.Flags().Bool("docker", false, "Configure Docker to authenticate against the registry. Writes to $DOCKER_CONFIG, or ~/.docker when it is unset.")
	artifactRegistry_loginCmd.Flags().String("duration", "15m0s", "How long the exchanged token should remain valid. Ignored for --docker.")
	artifactRegistry_loginCmd.Flags().Bool("gradle", false, "Configure Gradle to authenticate against the registry. Writes to ~/.gradle/gradle.properties.")
	artifactRegistry_loginCmd.Flags().String("hostname", "", "GitLab hostname to request the token from. Defaults to the configured GitLab instance.")
	artifactRegistry_loginCmd.Flags().Bool("maven", false, "Configure Maven to authenticate against the registry. Writes to ~/.m2/settings.xml.")
	artifactRegistry_loginCmd.Flags().Bool("npm", false, "Configure npm to authenticate against the registry. Writes to ~/.npmrc.")
	artifactRegistry_loginCmd.Flags().String("registry", "", "Registry to authenticate against. For --docker, a bare hostname; for others, typically a URL.")
	artifactRegistry_loginCmd.Flags().String("registry-alias", "", "Alias/ID to register the registry under (Maven/Gradle only). Defaults to a name derived from --registry.")
	artifactRegistry_loginCmd.Flags().Bool("sbt", false, "Configure sbt to authenticate against the registry. Writes to ~/.sbt/1.0/credentials.sbt.")
	artifactRegistry_loginCmd.MarkFlagRequired("registry")
	artifactRegistryCmd.AddCommand(artifactRegistry_loginCmd)

	carapace.Gen(artifactRegistry_loginCmd).FlagCompletion(carapace.ActionMap{
		"hostname": action.ActionConfigHosts(),
	})
}
