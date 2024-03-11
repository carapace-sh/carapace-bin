package git

import "github.com/carapace-sh/carapace"

// ActionFieldNames completes field names
//
//	author (the author header-field)
//	body (the body of the message)
func ActionFieldNames() carapace.Action {
	return carapace.ActionValuesDescribed(
		"authordate", "the date component of the author header-field",
		"authoremail", "the email component of the author header-field",
		"authorname", "the name component of the author header-field",
		"author", "the author header-field",
		"body", "the body of the message",
		"committerdate", "the date component of the committer header-field",
		"committeremail", "the email component of the committer header-field",
		"committername", "the name component of the committer header-field",
		"committer", "the committer header-field",
		"contents", "complete message",
		"creatordate", "the date component of the creator header-field",
		"creator", "the creator header-field",
		"deltabase", "object name of the delta base of the object",
		"HEAD", "* if HEAD matches ref or space otherwise",
		"numparent", "number of parent objects",
		"objectname", "object name (SHA-1)",
		"objectsize", "the size of the object",
		"object", "the object header-field",
		"objecttype", "the type of the object",
		"parent", "the parent header-field",
		"push", "name of a local ref which represents the @{push} location for the displayed ref",
		"refname", "name of the ref",
		"subject", "the subject of the message",
		"symref", "the ref which the given symbolic ref refers to",
		"taggerdate", "the date component of the tagger header-field",
		"taggeremail", "the email component of the tagger header-field",
		"taggername", "the name component of the tagger header-field",
		"tagger", "the tagger header-field",
		"tag", "the tag header-field",
		"trailers", "structured information in commit messages",
		"tree", "the tree header-field",
		"type", "the type header-field",
		"upstream", "name of a local ref which can be considered “upstream” from the displayed ref",
		"version:refname", "sort by versions",
	)
}
