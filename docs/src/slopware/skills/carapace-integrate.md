# carapace-integrate

The [`carapace-integrate`] skill integrates [Carapace] into a cobra-based CLI application.

![](./carapace-integrate/integrate.cast)

## Instructions

1. start [crush] in an cobra project
2. `ctrl+p`, `TAB`, then select the **skill**
3. run with `integrate carapace`

### Create
  > create a standalone completer for the shell command `tail`

  ![](./carapace-integrate/create.cast)

### Modify

  > i'm building my own git command.  create completions for it.
  > - positional arguments are git refs
  > - default behaviour is to show a prettier git log for them (oneline, graph, make it fancy)
  > - add ref completion with a carapace macro (loosely coupled)
  > - dynamically embed git plugin commands at root level and bridge completions (executables with `git-` prefix)
  > - add dash completion since refs could clash with subcommands
  >
  > ```sh
  > minigit HEAD~1 master # show log for refs
  > minigit clang-format  # execute git-clang-format
  > ```
  
  ![](./carapace-integrate/modify.cast)


[Carapace]:https://carapace.sh
[`carapace-integrate`]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/carapace-integrate/SKILL.md
[crush]:https://github.com/charmbracelet/crush
