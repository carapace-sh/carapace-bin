# Ruby CLI and Tooling

The `ruby` interpreter command-line interface, `irb`, `erb`, `rake`, `rdoc`/`ri`, profiling and debugging tools, and version managers.

> **Source of truth**: `ruby --help`, `man ruby`, <https://github.com/ruby/ruby>. For **environment variables**, see [environment.md](environment.md). For **RubyGems/Bundler**, see [gems-bundler.md](gems-bundler.md).

## ruby (Interpreter)

```
ruby [options] [script.rb] [arguments...]
ruby [options] -e "code"
```

### Common Flags

| Flag | Description |
|------|-------------|
| `-e 'code'` | Execute the given string |
| `-I path` | Prepend to `$LOAD_PATH` (can repeat) |
| `-r lib` | Require the library before execution |
| `-n` | Wrap script in `while gets ... end` (read line-by-line) |
| `-p` | Same as `-n` but prints `$_` at end of each loop |
| `-a` | Auto-split `$_` into `$F` (use with `-n` or `-p`) |
| `-F pattern` | Set `$/`-split pattern for `-a` (default: whitespace) |
| `-l` | Chomp `$/` from input, append `$/` to output |
| `-0 octal` | Set `$/` (input record separator) |
| `-K code` | Set source encoding (legacy; use magic comment) |
| `-E ex[:in]` | Set default external[:internal] encoding |
| `-w` | Enable warnings |
| `-W level` | Set warning level: 0=silence, 1=medium, 2=verbose |
| `-d` | Debug mode (set `$DEBUG = true`) |
| `-v` / `--verbose` | Print version and enable warnings |
| `--version` | Print version and exit |
| `--copyright` | Print copyright and exit |
| `--enable-gems` | Enable RubyGems (default) |
| `--disable-gems` | Disable RubyGems |
| `--enable-frozen-string-literal` | Freeze all string literals |
| `--disable-frozen-string-literal` | Don't freeze (default) |
| `--yjit` | Enable YJIT (in-process JIT, since Ruby 3.1) |
| `--disable-yjit` | Disable YJIT |
| `--yjit-exec-mem-size=N` | YJIT executable memory size |
| `-C dir` | Change to directory before execution |
| `-S` | Search `$PATH` for the script |
| `-x` | Strip text before `#!ruby` line (for embedded scripts) |

### -n / -p / -a patterns

```bash
# Print lines containing "error"
ruby -ne 'puts $_ if /error/' log.txt

# Awk-style: print first field of each line
ruby -nae 'puts $F.first' data.txt

# In-place edit (like sed -i)
ruby -i -pe 'gsub(/foo/, "bar")' *.txt
```

### Shebang

```ruby
#!/usr/bin/env ruby
# or
#!/usr/bin/env -S ruby -w
```

### STDIN reading

```ruby
# Without -n: read all of STDIN
input = STDIN.read

# With -n: each line in $_
while gets
  process($_)
end
```

## irb (Interactive Ruby)

```bash
irb                    # start
irb -r foo             # require foo on startup
irb --simple-prompt    # minimal prompt
```

Inside irb:

| Command | Description |
|---------|-------------|
| `help` | Start ri for a method/class |
| `show_source` | Show source of a method |
| `ls` | List methods (like `obj.methods`) |
| `inspect` | Inspect an object |
| `_` | Last result |
| `binding.irb` | Drop into irb at current binding (debugging) |

### .irbrc

`~/.irbrc` is loaded on startup:

```ruby
# ~/.irbrc
require 'irb/completion'  # tab completion
IRB.conf[:PROMPT_MODE] = :SIMPLE
IRB.conf[:AUTO_INDENT] = true

# Custom helpers
def m(obj)
  obj.methods - obj.class.superclass.instance_methods
end
```

### binding.irb

```ruby
def debug_method(x)
  binding.irb  # drops into an irb session here
  # x and local variables are accessible
end
```

`binding.irb` starts an irb session with the current binding — the local variables, `self`, and method dispatch context are all preserved. Exit with `exit` or `Ctrl+D`.

## erb (Embedded Ruby)

```erb
<% code %>      # execute Ruby
<%= expr %>     # execute and insert result
<%# comment %>  # comment
```

```ruby
require 'erb'
template = ERB.new("Hello, <%= name %>")
template.result_with_hash(name: "World")  # => "Hello, World"

# With binding
name = "Alice"
ERB.new("Hi <%= name %>").result(binding)  # => "Hi Alice"
```

CLI:

```bash
erb template.erb > output.html
erb -T- template.erb  # no trim mode
```

## rake (Ruby Make)

`Rakefile` or `rakefile.rb`:

```ruby
task :default => [:test]

desc "Run tests"
task :test do
  sh "rspec"
end

desc "Build gem"
task :build do
  sh "gem build mygem.gemspec"
end

# File task (runs if file is missing or older than prerequisites)
file "lib/version.rb" => ["VERSION"] do |t|
  cp t.prerequisites.first, t.name
end

# Namespace
namespace :db do
  task :migrate do
    puts "migrating"
  end
end

# Multitask (parallel prerequisites)
multitask :deploy => [:build, :test]

# Rule (pattern-based)
rule '.o' => '.c' do |t|
  sh "cc -c #{t.source} -o #{t.name}"
end
```

Common commands:

| Command | Description |
|---------|-------------|
| `rake` | Run default task |
| `rake -T` | List tasks with descriptions |
| `rake -P` | Show task dependencies |
| `rake -t` | Trace (full backtrace on error) |
| `rake -n` | Dry run (print commands without executing) |
| `rake <task>` | Run a specific task |
| `rake <namespace>:<task>` | Run a namespaced task |

## rdoc and ri

`rdoc` generates HTML documentation from source comments:

```bash
rdoc lib/              # generate docs for lib/
rdoc --main README.md  # set the main page
rdoc --op doc/         # output directory
```

`ri` is the command-line doc viewer:

```bash
ri String
ri String#gsub
ri Array#map
ri File.open
```

Inside `irb`, `help` runs `ri`.

## Profiling and Debugging

### stackprof

Sampling profiler (gem):

```ruby
require 'stackprof'
StackProf.run(mode: :wall, out: "prof.dump") do
  my_hot_code
end
```

```bash
stackprof prof.dump --text
stackprof prof.dump --method Foo#bar
```

### ruby-prof

Call-graph profiler (gem):

```ruby
require 'ruby-prof'
result = RubyProf.profile do
  my_code
end
printer = RubyProf::GraphPrinter.new(result)
printer.print(STDOUT, {})
```

### debug.rb / rdbg

Modern debugger (`debug` gem, bundled since Ruby 3.1):

```ruby
require 'debug'
# or use binding.break (alias: debugger)
def foo
  binding.break  # drop into the debugger
end
```

```bash
rdbg -c script.rb       # start with debugger
rdbg -A                 # attach to a running process
```

Commands: `step` (s), `next` (n), `continue` (c), `break` (b), `info` (i), `p`, `where` (w), `finish` (fin), `catch`, `watch`.

### tracepoint

```ruby
TracePoint.new(:call, :return) do |tp|
  puts "#{tp.event} #{tp.defined_class}##{tp.method_id}"
end.enable do
  my_method
end
```

See [metaprogramming.md](metaprogramming.md) for full details.

## Type Tools

### rbs

Type signature files (`.rbs`):

```ruby
# sig/foo.rbs
class Foo
  def greet: (String name) -> String
end
```

```bash
rbs validate
rbs annotate lib/foo.rb  # generate prototype
```

### typeprof

Type profiler (infers types from code):

```bash
typeprof lib/foo.rb
```

## Linters and Formatters

| Tool | Description |
|------|-------------|
| `rubocop` | Style and lint checker (configurable) |
| `standard` | Standard Ruby formatter (opinionated, no config) |
| `ruby -w` | Built-in warnings |

```bash
rubocop                    # check
rubocop -a                 # auto-correct
rubocop --only Style/StringLiterals  # specific cops
standardrb                 # Standard Ruby
```

## Version Managers

| Tool | Description |
|------|-------------|
| `rbenv` | shim-based, `rbenv install 3.2.0`, `rbenv global 3.2.0` |
| `rvm` | shell-function-based, `rvm install 3.2.0`, `rvm use 3.2.0` |
| `chruby` | lightweight, `chruby 3.2.0` |
| `asdf` | multi-language, `asdf install ruby 3.2.0` |

All manipulate `PATH` and `GEM_HOME` to switch Ruby versions.

## Edge Cases and Known Issues

- **`-n`/`-p` modify `$_`** — these flags set `$_` to the current line; `$_` is thread-local and method-local
- **`-i` in-place editing** — `ruby -i -pe 'gsub(/x/, "y")' *.txt` modifies files in place; use `-i.bak` to keep backups
- **`irb` history** — stored in `~/.irb_history` by default; can be customized via `IRB.conf[:SAVE_HISTORY]`
- **`binding.irb` vs `binding.break`** — the latter uses the `debug` gem; `binding.irb` starts irb
- **`rake` task with no block** — `task :foo` is a no-op; useful as a dependency anchor
- **`rdoc` on C extensions** — uses `.c` source comments; needs `Init_foo` function to be found
- **YJIT** — enabled by default since Ruby 3.3; disable with `--disable-yjit` if it causes issues
- **`ruby -e` and ARGV** — `ruby -e 'puts ARGV.inspect' a b c` → `["a", "b", "c"]`; `-e` doesn't consume args

## Related

- [environment.md](environment.md) — `RUBYOPT`, `RUBYLIB`, `GEM_HOME`
- [gems-bundler.md](gems-bundler.md) — `gem`, `bundle exec`, binstubs
- [metaprogramming.md](metaprogramming.md) — `TracePoint`, `binding`
