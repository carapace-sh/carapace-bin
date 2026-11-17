# fish_add_path "/tmp/carapace-snippet-fixtures/carapace/bin"

function get-env -d "get environment variable"; echo $$argv[1]; end
function set-env -d "set environment variable"; set -g -x $argv[1] $argv[2]; end
function unset-env -d "unset environment variable"; set -e $argv[1]; end

function _carapace_completer
  set --local --export CARAPACE_SHELL fish
  set --local --export CARAPACE_SHELL_ALIASES (alias | string replace -r 'alias ' '' | string replace -r '=.*' '')
  set --local --export CARAPACE_SHELL_BUILTINS (builtin --names)
  set --local --export CARAPACE_SHELL_FUNCTIONS (functions --all)
  set --local --export CARAPACE_SHELL_JOBS (jobs 2>/dev/null | string match -r '^\d+\s+\d+\s+(running|stopped|suspended|continued)' | string replace -r '^(\d+).*' '%%$1')
  set --local --export CARAPACE_SHELL_VARIABLES (set --names)
  set --local data
  IFS='' set data (echo (commandline -cp)'' | sed "s/ \$/ ''/" | xargs carapace $argv[1] fish 2>/dev/null)
  if [ $status -eq 1 ]
    IFS='' set data (echo (commandline -cp)"'" | sed "s/ \$/ ''/" | xargs carapace $argv[1] fish 2>/dev/null)
    if [ $status -eq 1 ]
      IFS='' set data (echo (commandline -cp)'"' | sed "s/ \$/ ''/" | xargs carapace $argv[1] fish 2>/dev/null)
    end
  end
  echo $data
end

complete -e example
complete -c example -f -a '(_carapace_completer example)'

