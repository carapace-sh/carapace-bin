package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

func ActionOwnerRepositories(cmd *cobra.Command) carapace.Action {
	return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
		// TODO hack to enable completion outside git repo - this needs to be fixed in GraphQlAction/repooverride though
		if cmd.Flag("repo") == nil {
			fakeRepoFlag(cmd, "", "")
		}

		switch len(c.Parts) {
		case 0:
			return gh.ActionOwners(gh.HostOpts{}).Invoke(c).Suffix("/").ToA() // TODO host
		case 1:
			return gh.ActionRepositories(gh.OwnerOpts{Owner: c.Parts[0]})
		default:
			return carapace.ActionValues()
		}
	})
}

func ActionRepositoryFields() carapace.Action {
	return carapace.ActionValues(
		"assignableUsers",
		"codeOfConduct",
		"contactLinks",
		"createdAt",
		"defaultBranchRef",
		"deleteBranchOnMerge",
		"description",
		"diskUsage",
		"forkCount",
		"fundingLinks",
		"hasIssuesEnabled",
		"hasProjectsEnabled",
		"hasWikiEnabled",
		"homepageUrl",

		"id",
		"isArchived",
		"isBlankIssuesEnabled",

		"isEmpty",
		"isFork",
		"isInOrganization",
		"isMirror",
		"isPrivate",
		"isSecurityPolicyEnabled",
		"isTemplate",
		"isUserConfigurationRepository",

		"issueTemplates",
		"issues",
		"labels",
		"languages",
		"latestRelease",

		"licenseInfo",
		"mentionableUsers",
		"mergeCommitAllowed",
		"milestones",
		"mirrorUrl",
		"name",
		"nameWithOwner",
		"openGraphImageUrl",
		"owner",
		"parent",
		"primaryLanguage",
		"projects",
		"pullRequestTemplates",
		"pullRequests",
		"pushedAt",
		"rebaseMergeAllowed",
		"repositoryTopics",
		"securityPolicyUrl",
		"squashMergeAllowed",
		"sshUrl",
		"stargazerCount",
		"templateRepository",

		"updatedAt",
		"url",
		"usesCustomOpenGraphImage",
		"viewerCanAdminister",
		"viewerDefaultCommitEmail",
		"viewerDefaultMergeMethod",
		"viewerHasStarred",
		"viewerPermission",

		"viewerPossibleCommitEmails",
		"viewerSubscription",
		"watchers",
	)
}

func ActionRepositorySearchFields() carapace.Action {
	return carapace.ActionValues(
		"createdAt",
		"defaultBranch",
		"description",
		"forksCount",
		"fullName",
		"hasDownloads",
		"hasIssues",
		"hasPages",
		"hasProjects",
		"hasWiki",
		"homepage",
		"id",
		"isArchived",
		"isDisabled",
		"isFork",
		"isPrivate",
		"language",
		"license",
		"name",
		"openIssuesCount",
		"owner",
		"pushedAt",
		"size",
		"stargazersCount",
		"updatedAt",
		"visibility",
		"watchersCount",
	)
}
