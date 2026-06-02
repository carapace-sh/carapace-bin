# carapace-env

The [`carapace-env`] describes how to define environment variable completion.

![](./carapace-env/proxy-env.cast)

## Instructions

1. start [crush] in an cobra project
2. `ctrl+p`, `TAB`, then select the **skill**
3. run with `create variable completion ...`

### Go

  > look at the manpage of `git` and add missing environment variable completions for the ones with `GIT_` prefix

  ![](./carapace-env/git-env.cast)

### Spec

  > look at the output of `bat --help` and create environment variable completion as user specs

  ![](./carapace-env/bat-env.cast)

[Carapace]:https://carapace.sh
[`carapace-env`]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/carapace-env/SKILL.md
[crush]:https://github.com/charmbracelet/crush
