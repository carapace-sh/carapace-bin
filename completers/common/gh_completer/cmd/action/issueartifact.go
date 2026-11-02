package action

import (
	"fmt"
	"strconv"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

type issueArtifact struct {
	Number int
	Type   string
	Name   string
}

type issueArtifactVersion struct {
	Version int
	Name    string
}

type issueArtifactVersionDetails struct {
	Versions []issueArtifactVersion
}

func ActionIssueArtifacts(cmd *cobra.Command, issueNumber string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		number, err := strconv.Atoi(issueNumber)
		if err != nil {
			return carapace.ActionValues()
		}
		repo, err := repoOverride(cmd, c)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		var issueArtifacts []issueArtifact
		return ApiV3Action(cmd, fmt.Sprintf("repos/%v/%v/issues/%v/artifacts", repo.RepoOwner(), repo.RepoName(), number), &issueArtifacts, func() carapace.Action {
			vals := make([]string, 0)
			for _, artifact := range issueArtifacts {
				vals = append(vals, strconv.Itoa(artifact.Number), fmt.Sprintf("%v: %v", artifact.Type, artifact.Name))
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	}).Tag("issue artifacts")
}

func ActionIssueArtifactVersions(cmd *cobra.Command, issueNumber, artifactNumber string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		number, err := strconv.Atoi(issueNumber)
		if err != nil {
			return carapace.ActionValues()
		}
		artifact, err := strconv.Atoi(artifactNumber)
		if err != nil {
			return carapace.ActionValues()
		}
		repo, err := repoOverride(cmd, c)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		var details issueArtifactVersionDetails
		return ApiV3Action(cmd, fmt.Sprintf("repos/%v/%v/issues/%v/artifacts/%v", repo.RepoOwner(), repo.RepoName(), number, artifact), &details, func() carapace.Action {
			vals := make([]string, 0)
			for _, version := range details.Versions {
				name := version.Name
				if name == "" {
					name = "(current content)"
				}
				vals = append(vals, strconv.Itoa(version.Version), name)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	}).Tag("issue artifact versions")
}
