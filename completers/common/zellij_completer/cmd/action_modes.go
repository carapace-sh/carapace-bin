package cmd

import "github.com/carapace-sh/carapace"

func actionModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"normal", "In `Normal` mode, input is always written to the terminal, except for the shortcuts leading to other modes",
		"locked", "In `Locked` mode, input is always written to the terminal and all shortcuts are disabled except the one leading back to normal mode",
		"resize", "`Resize` mode allows resizing the different existing panes",
		"pane", "`Pane` mode allows creating and closing panes, as well as moving between them",
		"tab", "`Tab` mode allows creating and closing tabs, as well as moving between them",
		"scroll", "`Scroll` mode allows scrolling up and down within a pane",
		"enter-search", "`EnterSearch` mode allows for typing in the needle for a search in the scroll buffer of a pane",
		"search", "`Search` mode allows for searching a term in a pane (superset of `Scroll`)",
		"rename-tab", "`RenameTab` mode allows assigning a new name to a tab",
		"rename-pane", "`RenamePane` mode allows assigning a new name to a pane",
		"session", "`Session` mode allows detaching sessions",
		"move", "`Move` mode allows moving the different existing panes within a tab",
		"prompt", "`Prompt` mode allows interacting with active prompts",
		"tmux", "`Tmux` mode allows for basic tmux keybindings functionality",
	)
}
