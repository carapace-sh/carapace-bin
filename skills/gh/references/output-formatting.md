# Output Formatting — JSON, jq, and Go Templates

How to format `gh` command output using `--json`, `--jq`, and `--template` flags.

## Overview

By default, `gh` commands output line-based plain text. Commands that support `--json` can output structured JSON, which can then be further formatted with `--jq` (jq query syntax) or `--template` (Go template syntax).

These flags are available on many list/view commands: `pr list`, `pr view`, `issue list`, `issue view`, `repo list`, `repo view`, `release list`, `release view`, `run list`, `run view`, `workflow list`, `cache list`, `label list`, `secret list`, `variable get`, `variable list`, `search *`, `extension search`, `auth status`, etc.

## `--json`

Output JSON with the specified fields. Pass a comma-separated list of field names.

```sh
$ gh pr list --json number,title,author
```

```json
[
  {
    "author": { "login": "monalisa" },
    "number": 123,
    "title": "A helpful contribution"
  }
]
```

To discover available JSON field names for a command, omit the field argument:

```sh
$ gh pr list --json
```

When connected to a terminal, JSON output is automatically pretty-printed.

## `--jq`

Filter and select values using jq syntax. The `jq` utility does **not** need to be installed. Requires `--json` to be passed with field names.

```sh
# Select fields from an array
$ gh pr list --json author --jq '.[].author.login'
monalisa
codercat

# Complex filtering
$ gh issue list --json number,title,labels --jq \
  'map(select((.labels | length) > 0))    # must have labels
  | map(.labels = (.labels | map(.name))) # show only label names
  | .[:3]'                                # select first 3
```

When connected to a terminal, jq output is automatically pretty-printed.

For jq query syntax, see <https://jqlang.github.io/jq/manual/>.

## `--template`

Format output using Go template syntax. Requires `--json` with field names.

```sh
# Simple template
$ gh pr list --json number,title,headRefName,updatedAt --template \
  '{{range .}}{{tablerow (printf "#%v" .number | autocolor "green") .title .headRefName (timeago .updatedAt)}}{{end}}'

# Complex template with tables
$ gh pr view 3519 --json number,title,body,reviews,assignees --template \
'{{printf "#%v" .number}} {{.title}}

{{.body}}

{{tablerow "ASSIGNEE" "NAME"}}{{range .assignees}}{{tablerow .login .name}}{{end}}{{tablerender}}
{{tablerow "REVIEWER" "STATE" "COMMENT"}}{{range .reviews}}{{tablerow .author.login .state .body}}{{end}}
'

# Hyperlinks
$ gh issue list --json title,url --template '{{range .}}{{hyperlink .url .title}}{{"\n"}}{{end}}'
```

For Go template syntax, see <https://golang.org/pkg/text/template/>.

## Custom Template Functions

In addition to Go template standard library functions, gh provides:

| Function | Description |
|----------|-------------|
| `autocolor <style> <input>` | Like `color`, but only emits color to terminals |
| `color <style> <input>` | Colorize input using [mgutz/ansi](https://github.com/mgutz/ansi) |
| `join <sep> <list>` | Join values in a list using a separator |
| `pluck <field> <list>` | Collect values of a field from all items in the input |
| `tablerow <fields>...` | Align fields in output vertically as a table |
| `tablerender` | Render fields added by `tablerow` in place |
| `timeago <time>` | Render a timestamp as relative to now |
| `timefmt <format> <time>` | Format a timestamp using Go's `Time.Format` |
| `truncate <length> <input>` | Ensure input fits within length |
| `hyperlink <url> <text>` | Render a terminal hyperlink |

### Sprig Template Functions

The following [Sprig](https://masterminds.github.io/sprig/) library functions are also available:

| Function | Description |
|----------|-------------|
| `contains <arg> <string>` | Check if `string` contains `arg` |
| `hasPrefix <prefix> <string>` | Check if `string` starts with `prefix` |
| `hasSuffix <suffix> <string>` | Check if `string` ends with `suffix` |
| `regexMatch <regex> <string>` | Check if `string` has matches for `regex` |

## `gh api` Formatting

The `gh api` command also supports `--jq` and `--template` for formatting API responses. Unlike other commands, `gh api` does not require `--json` — the response is already JSON.

```sh
$ gh api repos/{owner}/{repo}/issues --jq '.[].title'
$ gh api repos/{owner}/{repo}/issues --template \
  '{{range .}}{{.title}} ({{.labels | pluck "name" | join ", " | color "yellow"}}){{"\n"}}{{end}}'
```

See [api.md](api.md) for `gh api` details.

## Common Patterns

### Extract a Single Field

```sh
$ gh pr list --json number --jq '.[].number'
123
124
```

### Filter and Transform

```sh
# PRs with more than 0 labels, showing only label names
$ gh issue list --json number,title,labels --jq \
  'map(select((.labels | length) > 0)) | map(.labels = (.labels | map(.name)))'
```

### Table Output

```sh
$ gh pr list --json number,title,headRefName,updatedAt --template \
  '{{range .}}{{tablerow (printf "#%v" .number | autocolor "green") .title .headRefName (timeago .updatedAt)}}{{end}}'
```

### Hyperlinks

```sh
$ gh issue list --json title,url --template \
  '{{range .}}{{hyperlink .url .title}}{{"\n"}}{{end}}'
```

## Related

- For the `gh api` command, see [api.md](api.md).
- For environment variables affecting output (`NO_COLOR`, `CLICOLOR`, `GH_FORCE_TTY`), see [config.md](config.md).
