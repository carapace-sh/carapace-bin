package gh

import (
	"strconv"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
)

type discussion struct {
	Number int
	Title  string
	Closed bool
}

type discussionQuery struct {
	Data struct {
		Repository struct {
			Discussions struct {
				Nodes []discussion
			}
		}
	}
}

// ActionDiscussions completes discussions
//
//	11 (discussion title)
//	12 (discussion title)
func ActionDiscussions(opts RepoOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult discussionQuery
		return graphQlAction(opts, `repository(owner: $owner, name: $repo){ discussions(first: 100, orderBy: {direction: DESC, field: UPDATED_AT}) { nodes { number, title, closed } } }`, &queryResult, func() carapace.Action {
			discussions := queryResult.Data.Repository.Discussions.Nodes
			vals := make([]string, 0)
			for _, d := range discussions {
				s := styles.Gh.StateOpen
				if d.Closed {
					s = styles.Gh.StateClosed
				}
				vals = append(vals, strconv.Itoa(d.Number), d.Title, s)
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	}).Tag("discussions")
}

type discussionCategory struct {
	Name        string
	Description string
}

type discussionCategoriesQuery struct {
	Data struct {
		Repository struct {
			DiscussionCategories struct {
				Nodes []discussionCategory
			}
		}
	}
}

// ActionDiscussionCategories completes discussion categories
//
//	Announcements (General announcements)
//	Ideas (Ideas and suggestions)
func ActionDiscussionCategories(opts RepoOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult discussionCategoriesQuery
		return graphQlAction(opts, `repository(owner: $owner, name: $repo){ discussionCategories( first: 100) { nodes { name, description } } }`, &queryResult, func() carapace.Action {
			categories := queryResult.Data.Repository.DiscussionCategories.Nodes
			vals := make([]string, len(categories)*2)
			for index, category := range categories {
				vals[index*2] = category.Name
				vals[index*2+1] = category.Description
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	}).Tag("discussion categories")
}

// ActionDiscussionListFields completes discussion list fields
//
//	answerChosenAt
//	answerChosenBy
func ActionDiscussionListFields() carapace.Action {
	return carapace.ActionValues(
		"answerChosenAt",
		"answerChosenBy",
		"answered",
		"author",
		"body",
		"category",
		"closed",
		"closedAt",
		"createdAt",
		"id",
		"labels",
		"locked",
		"number",
		"stateReason",
		"title",
		"updatedAt",
		"url",
	)
}

// ActionDiscussionViewFields completes discussion view fields
//
//	answerChosenAt
//	answerChosenBy
func ActionDiscussionViewFields() carapace.Action {
	return carapace.ActionValues(
		"answerChosenAt",
		"answerChosenBy",
		"answered",
		"author",
		"body",
		"category",
		"closed",
		"closedAt",
		"comments",
		"createdAt",
		"id",
		"labels",
		"locked",
		"number",
		"reactionGroups",
		"state",
		"stateReason",
		"title",
		"updatedAt",
		"url",
	)
}
