package rustdoc

import "github.com/carapace-sh/carapace"

// ActionLints completes lints
//
//	private_intra_doc_links (Detects when intra-doc links from public to private items)
//	missing_docs (Detects items missing documentation)
func ActionLints() carapace.Action {
	return carapace.ActionValuesDescribed(
		"broken_intra_doc_links", "Detects when an intra-doc link fails to be resolved",
		"private_intra_doc_links", "Detects when intra-doc links from public to private items",
		"missing_docs", "Detects items missing documentation",
		"missing_crate_level_docs", "Detects if there is no documentation at the crate root",
		"missing_doc_code_examples", "Detects when a documentation block is missing a code example",
		"private_doc_tests", "Detects documentation tests when they are on a private item",
		"invalid_codeblock_attributes", "Detects code block attributes in documentation examples that have potentially mis-typed values",
		"invalid_html_tags", "Detects unclosed or invalid HTML tags",
		"invalid_rust_codeblocks", "Detects Rust code blocks in documentation examples that are invalid",
		"bare_urls", "Detects URLs which are not links",
		"unescaped_backticks", "Detects backticks (`) that are not escaped",
	)
}
