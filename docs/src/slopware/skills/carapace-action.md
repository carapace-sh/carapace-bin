# carapace-action

The [`carapace-action`] describes in detail how to create [Custom Actions](https://carapace-sh.github.io/carapace/carapace/customActions.html).

![](./carapace-action/passwd.cast)

## Instructions

1. start [crush] in an cobra project
2. `ctrl+p`, `TAB`, then select the **skill**
3. run with `create (public/private/anoymous) action ...`

### Public Action

  > for the repo completion query the api for repos where the current user has permissions to close pull requests.\
  > the repo flag should just accept OWNER/REPO.\
  > create a custom action in pkg/actions.\
  > also, cache it for a day as these shouldn't change often

  ![](./carapace-action/public.cast)

[Carapace]:https://carapace.sh
[`carapace-action`]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/carapace-action/SKILL.md
[crush]:https://github.com/charmbracelet/crush
