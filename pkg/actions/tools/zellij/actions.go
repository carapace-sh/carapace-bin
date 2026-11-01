package zellij

import "github.com/carapace-sh/carapace"

// ActionDirections completes movement directions
//
//	left
//	down
//	up
//	right
func ActionDirections() carapace.Action {
	return carapace.ActionValues("down", "left", "right", "up")
}

// ActionResizes completes resize directions
//
//	increase
//	decrease
func ActionResizes() carapace.Action {
	return carapace.ActionValues("decrease", "increase")
}

// ActionModes completes zellij modes
//
//	normal (In `Normal` mode, input is always written to the terminal, except for the shortcuts leading to other modes)
//	locked (In `Locked` mode, input is always written to the terminal and all shortcuts are disabled except the one leading back to normal mode)
func ActionModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"enter-search", "`EnterSearch` mode allows for typing in the needle for a search in the scroll buffer of a pane",
		"locked", "In `Locked` mode, input is always written to the terminal and all shortcuts are disabled except the one leading back to normal mode",
		"move", "`Move` mode allows moving the different existing panes within a tab",
		"normal", "In `Normal` mode, input is always written to the terminal, except for the shortcuts leading to other modes",
		"pane", "`Pane` mode allows creating and closing panes, as well as moving between them",
		"prompt", "`Prompt` mode allows interacting with active prompts",
		"rename-pane", "`RenamePane` mode allows assigning a new name to a pane",
		"rename-tab", "`RenameTab` mode allows assigning a new name to a tab",
		"resize", "`Resize` mode allows resizing the different existing panes",
		"scroll", "`Scroll` mode allows scrolling up and down within a pane",
		"search", "`Search` mode allows for searching a term in a pane (superset of `Scroll`)",
		"session", "`Session` mode allows detaching sessions",
		"tab", "`Tab` mode allows creating and closing tabs, as well as moving between them",
		"tmux", "`Tmux` mode allows for basic tmux keybindings functionality",
	)
}
