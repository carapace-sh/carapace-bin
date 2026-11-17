$env.PATH = ($env.PATH | split row (char esep) | where { $in != "/tmp/carapace-snippet-fixtures/carapace/bin" } | prepend "/tmp/carapace-snippet-fixtures/carapace/bin")

let carapace_completer = {|place|
  # backwards compatible workaround for positional completer input, see nushell #18791
  let spans = (if ($place | describe) =~ "record" { $place.command } else { $place })

  # if the current command is an alias, get it's expansion
  let expanded_alias = (scope aliases | where name == $spans.0 | $in.0?.expansion?)

  # overwrite
  let spans = (if $expanded_alias != null  {
    # put the first word of the expanded alias first in the span
    $spans | skip 1 | prepend ($expanded_alias | split row " " | take 1)
  } else {
    $spans | skip 1 | prepend ($spans.0)
  })

  if ($spans | length) == 1 {
    null
  } else {
    let carapace_env = { CARAPACE_SHELL: 'nushell' }
    with-env $carapace_env {
      carapace $spans.0 nushell ...$spans | from json
    }
  }
}

mut current = (($env | default {} config).config | default {} completions)
$current.completions = ($current.completions | default {} external)
$current.completions.external = ($current.completions.external
| default true enable
# backwards compatible workaround for default, see nushell #15654
| upsert completer { if $in == null { $carapace_completer } else { $in } })

$env.config = $current
    
