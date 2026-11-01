# carapace

Describes how to use [carapace](https://carapace.sh).

| reference                        | description                                          |
|----------------------------------|------------------------------------------------------|
| [action][reference-action]       | create and modify shell completion actions           |
| [choice][reference-choice]       | how carapace resolves which completer to use         |
| [convert][reference-convert]     | convert a user spec into a native Go completer       |
| [env][reference-env]             | define and customize environment variable completion |
| [integrate][reference-integrate] | integrate carapace into a cobra application          |
| [lexer][reference-lexer]         | build a parser/lexer for a DSL                       |
| [macro][reference-macro]         | define and customize macros                          |
| [man][reference-man]             | provide inline documentation                         |
| [mcp][reference-mcp]             | work with the MCP server                             |
| [scrape][reference-scrape]       | scrape command into a spec                           |
| [setup][reference-setup]         | install and configure carapace                       |
| [spec][reference-spec]           | define and customize specs                           |
| [update][reference-update]       | update an existing completer                         |


[reference-action]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/carapace/references/action.md 
[reference-choice]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/carapace/references/choice.md 
[reference-convert]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/carapace/references/convert.md 
[reference-env]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/carapace/references/env.md 
[reference-integrate]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/carapace/references/integrate.md 
[reference-lexer]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/carapace/references/lexer.md 
[reference-macro]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/carapace/references/macro.md 
[reference-man]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/carapace/references/man.md 
[reference-mcp]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/carapace/references/mcp.md 
[reference-scrape]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/carapace/references/scrape.md 
[reference-setup]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/carapace/references/setup.md 
[reference-spec]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/carapace/references/spec.md 
[reference-update]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/carapace/references/update.md 

## Examples

> Some asciicasts show an older version of the skills before these were combined into one.

### Action

> ![](./carapace/action-passwd.cast)
>
> create and show me an action that completes users from /etc/passwd .\
> style them in red (root), yellow  (system), default (regular).\
> examples: dbus (System Message Bus) dnsmasq (dnsmasq daemon)

> ![](./carapace/action-repo.cast)
>
> for the repo completion query the api for repos where the current user has permissions to close pull requests.\
> the repo flag should just accept OWNER/REPO.\
> create a custom action in pkg/actions.\
> also, cache it for a day as these shouldn't change often


### Env

> ![](./carapace/env-proxy.cast)
>
> create custom variable completion for https proxy with default ports\
> i've got one for development at localhost and a production one at proxy.example\
> style them accordingly

> ![](./carapace/env-git.cast)
>
> look at the manpage of `git` and add missing environment variable completions for the ones with `GIT_` prefix

> ![](./carapace/env-bat.cast)
>
> look at the output of `bat --help` and create environment variable completion as user specs

### Integrate

> ![](./carapace/integrate-crush.cast)
>
> integrate carapace

> ![](./carapace/integrate-tail.cast)
>
> create a standalone completer for the shell command `tail`

  
> ![](./carapace/integrate-minigit.cast)
>
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

### Scrape

> ![](./carapace/scrape-pgxcli.cast)
>
> scrape https://github.com/balajz/pgxcli

### Spec

> ![](./carapace/spec-summarize.cast)
>
> create a runnable spec named `summarize` that takes two git refs as positional argument and invokes
> `crush` to summarize the difference between them
