# Environment

Complex environment variable completion is provided with `get-env`, `set-env` and `unset-env`.

In `elvish` the completion is simply overridden.
For other shells custom functions are added.

> Setting the environment variable `CARAPACE_ENV` to `0`
> before sourcing `carapace _carapace` disables this behaviour.

![](./environment.cast)

## Custom definitions

*TODO:* define using `~/.config/carapace/env.yaml`