# LSP

An experimental [Bash Language Server][lsp] with support for [Carapace].

![](./lsp/bash-language-server.cast)

Place the [single binary release](https://github.com/carapace-sh/bash-language-server/releases) in your [PATH] and follow the [original instructions](https://github.com/bash-lsp/bash-language-server).

Then add it to your [config](https://github.com/charmbracelet/crush#lsps).

```json
{
  "$schema": "https://charm.land/crush.json",
  "lsp": {
    "bash": {
      "command": "bash-language-server",
      "args": [
        "start"
      ]
    }
  }
}
```

> Not actually recommended at this moment.

[lsp]:https://github.com/carapace-sh/bash-language-server
[Carapace]:https://carapace.sh
[PATH]:https://en.wikipedia.org/wiki/PATH_(variable)
