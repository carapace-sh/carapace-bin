# fish_add_path "/tmp/carapace-snippet-fixtures/carapace/bin"

function _carapace_completer
  set --local --export CARAPACE_SHELL fish
  set --local data
  IFS='' set data (echo (commandline -cp)'' | sed "s/ \$/ ''/" | xargs carapace $argv[1] fish 2>/dev/null)
  set --local _carapace_rc $status
  if test $_carapace_rc -eq 1
    IFS='' set data (echo (commandline -cp)"'" | sed "s/ \$/ ''/" | xargs carapace $argv[1] fish 2>/dev/null)
    set _carapace_rc $status
    if test $_carapace_rc -eq 1
      IFS='' set data (echo (commandline -cp)'"' | sed "s/ \$/ ''/" | xargs carapace $argv[1] fish 2>/dev/null)
      set _carapace_rc $status
    end
  end
  if test $_carapace_rc -ne 0; and set -q CARAPACE_SHELL_STATE
    set --local _carapace_old_state "$CARAPACE_SHELL_STATE"
    set -e CARAPACE_SHELL_STATE
    IFS='' set data (echo (commandline -cp)'' | sed "s/ \$/ ''/" | xargs carapace $argv[1] fish 2>/dev/null)
    set _carapace_rc $status
    if test $_carapace_rc -eq 1
      IFS='' set data (echo (commandline -cp)"'" | sed "s/ \$/ ''/" | xargs carapace $argv[1] fish 2>/dev/null)
      set _carapace_rc $status
      if test $_carapace_rc -eq 1
        IFS='' set data (echo (commandline -cp)'"' | sed "s/ \$/ ''/" | xargs carapace $argv[1] fish 2>/dev/null)
      end
    end
    set --export CARAPACE_SHELL_STATE "$_carapace_old_state"
  end
  if test -n "$_carapace_state_cleanup"
    rm -rf "$_carapace_state_cleanup"
  end
  echo $data
end

complete -e example
complete -c example -f -a '(_carapace_completer example)'

