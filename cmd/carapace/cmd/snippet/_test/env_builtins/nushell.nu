$env.PATH = ($env.PATH | split row (char esep) | where { $in != "/tmp/carapace-snippet-fixtures/carapace/bin" } | prepend "/tmp/carapace-snippet-fixtures/carapace/bin")

# nushell passes a place record since 0.116, the element text since 0.115 and only the token before that
def _carapace-env-span [command: string, place: any] {
  let described = ($place | describe)
  let parts = (if $described =~ "record" {
      ($place.command? | default [])
    } else if $described =~ "list" {
      $place
    } else {
      ($place | split row " ")
    })

  if ($parts | is-empty) {
    [$command, ""]
  } else if (($parts | length) > 1) and (($parts | first) == $command) {
    $parts
  } else {
    [$command, ...$parts]
  }
}

def _carapace-env-complete [command: string, place: any] {
  let spans = (_carapace-env-span $command $place)

  with-env { CARAPACE_SHELL: "nushell" } {
    carapace $spans.0 nushell ...$spans | from json
  }
}

def _carapace-get-env [place?] { _carapace-env-complete get-env $place }
def _carapace-set-env [place?] { _carapace-env-complete set-env $place }
def _carapace-unset-env [place?] { _carapace-env-complete unset-env $place }

def --env get-env [name: any@_carapace-get-env] { $env | get $name }
def --env set-env [name: any@_carapace-set-env, value: any@_carapace-set-env] { load-env { $name: $value } }
def --env unset-env [name: any@_carapace-unset-env] { hide-env $name }

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
    let carapace_state = (if (($env.XDG_RUNTIME_DIR? | default "") != "") and (($env.XDG_RUNTIME_DIR | path type) == "dir") {
      let dir = ($env.XDG_RUNTIME_DIR | path join "carapace" "builtins" (id -u | into string) ($nu.pid | into string))
      mkdir $dir
      $dir
    } else {
      mktemp -d
    })
    (scope aliases | get name | uniq | str join "\n") | save -f ($carapace_state | path join "aliases")
    (help commands | where category != "" | get name | each { split row " " | first } | uniq  | str join "\n") | save -f ($carapace_state | path join "builtins")
    (help commands | where category == "" | get name | each { split row " " | first } | uniq  | str join "\n") | save -f ($carapace_state | path join "functions")
    (scope variables | get name | uniq | str join "\n") | save -f ($carapace_state | path join "variables")
    let carapace_env = { CARAPACE_SHELL: 'nushell', CARAPACE_SHELL_STATE: $carapace_state }
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
    
