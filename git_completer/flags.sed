#!/bin/sed -rf
# escape "
s/"/\\"/g
# shorthand and name
s/^ +-([^,-]), --([^[ ]+)/\1;\2;/
# only name
s/^ +--([^ [<=]+)[ [<=]?/;\1;/
# only shorthand (git grep has a shorthand with more than one char `-NUM`)
s/^ +-([^, [<=-]+)[ [<=]/\1;;/
# any kind of value
s/(.*;.*;)[^ ]+ */\1string;/
# else bool
s/^([^;]*;[^;]*;)( +|$)/\1bool;/
#----
# bool flag name only
s/^;([^;]+);bool;(.*)/cmd.Flags().Bool("\1", false, "\2")/
# bool flag shorthand only
s/^([^;]+);;bool;(.*)/cmd.Flags().BoolP("\1", "\1", false, "\2")/
# bool flag name and shorthand
s/^(.);([^;]+);bool;(.*)/cmd.Flags().BoolP("\2", "\1", false, "\3")/
# ---
# string flag name only
s/^;([^;]+);string;(.*)/cmd.Flags().String("\1", "", "\2")/
# string flag shorthand only
s/^([^;]);;string;(.*)/cmd.Flags().StringP("\1", "\1", "", "\2")/
# string flag name and shorthand
s/^(.);([^;]+);string;(.*)/cmd.Flags().StringP("\2", "\1", "", "\3")/
