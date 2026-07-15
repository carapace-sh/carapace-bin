---
name: ruby
description: >
  Use when working with Ruby — the dynamic, object-oriented language and its
  runtime (MRI/CRuby). Covers the object model, metaprogramming, core types,
  blocks/procs/lambdas, control flow and exceptions, I/O and encoding, the
  concurrency model (Threads, Fibers, Ractors, GVL), garbage collection and
  ObjectSpace, the RubyGems/Bundler package system, the interpreter CLI and
  tooling (irb, erb, rake, rdoc/ri), environment variables, and the standard
  library. Triggers on: "ruby", "Ruby", "MRI", "CRuby", "YARV", "Ripper",
  "object model", "method lookup", "eigenclass", "singleton class",
  "metaprogramming", "define_method", "method_missing", "class_eval",
  "instance_eval", "Refinements", "Binding", "Proc", "lambda", "block",
  "yield", "rescue", "ensure", "raise", "begin", "end", "Encoding",
  "String#encode", "Thread", "Fiber", "Ractor", "GVL", "GIL", "Mutex",
  "ObjectSpace", "GC", "WeakRef", "gem", "Gemfile", "Bundler", "gemspec",
  "bundle exec", "irb", "erb", "rake", "Rakefile", "rdoc", "ri", "RUBYOPT",
  "RUBYLIB", "GEM_HOME", "GEM_PATH", "BUNDLE_GEMFILE",
  "frozen string literal", "RUBY_VERSION", "RUBY_ENGINE".
user-invocable: true
---

# Ruby In-Depth Reference

Comprehensive reference for the Ruby programming language and its reference implementation (MRI/CRuby) — covering the object model, metaprogramming, core types, control flow, I/O and encoding, concurrency, garbage collection, the RubyGems/Bundler package system, the interpreter CLI and tooling, environment variables, and the standard library.

## Data Flow

```
Source file (.rb)
  → Ripper (parse to AST)
    → YARV compiler (compile AST to bytecode instructions)
      → YARV VM (per-instruction dispatch)
        → runtime services (ObjectSpace, GC, Thread scheduler, load path)
          → I/O, exit
```

## Sub-Resources

Load the reference that matches your task. When in doubt, load multiple references.

| Keywords | Reference |
|----------|----------|
| object model, class, module, BasicObject, Object, Kernel, ancestor, method lookup, eigenclass, singleton class, superclass, super, included, prepended, inherited, ancestors, const_get, constants, autoload, refine, using, Module vs Class, instance variable, class variable, @@, constant resolution, toplevel, ::, ObjectSpace.each_object, method_missing | [references/object-model.md](references/object-model.md) |
| metaprogramming, define_method, class_eval, instance_eval, module_eval, eval, Binding, send, __send__, public_send, respond_to?, method, method_missing, method_added, Refinements, refine, using, Module.new, Class.new, remove_method, undef_method, alias_method, attr_accessor, attr_reader, attr_writer, define_singleton_method, const_set, trace_point, TracePoint, set_trace_func, caller, caller_locations | [references/metaprogramming.md](references/metaprogramming.md) |
| String, Symbol, Integer, Float, Numeric, Rational, Complex, BigDecimal, Array, Hash, Range, Regexp, MatchData, Proc, Method, nil, NilClass, true, false, TrueClass, FalseClass, Boolean, truthiness, eql?, equal?, ==, ===, <=>, Comparable, freeze, frozen?, immutability, dup, clone, deep_copy, hash, inspect, to_s, to_str, coerce | [references/types.md](references/types.md) |
| block, yield, proc, lambda, Proc, Method, method(:foo), to_proc, &block, ampersand, &, proc call, lambda call, return in block, return in lambda, argument strictness, arity, curry, splat, double splat, **, block_given?, &method(:foo), proc vs lambda | [references/blocks-procs-lambdas.md](references/blocks-procs-lambdas.md) |
| if, unless, case, when, while, until, for, loop, break, next, redo, retry, throw, catch, raise, begin, rescue, ensure, else, Exception, StandardError, RuntimeError, custom exception, rescue modifier, rescue chain, $!, $@, $ERROR_INFO, retry, else clause, fail | [references/control-flow.md](references/control-flow.md) |
| IO, File, StringIO, File::open, File.open, IO.foreach, IO.read, IO.write, Encoding, String#encode, Encoding.default_external, Encoding.default_internal, magic comment, # encoding, # coding, # frozen_string_literal, external encoding, internal encoding, transcoding, binary mode, text mode, STDIN, STDOUT, STDERR, ARGF, ARGV, open, File.read, File.write, flock, IO.select, pipe, IO.pipe, io/nonblock | [references/io-encoding.md](references/io-encoding.md) |
| Thread, Fiber, Ractor, GVL, GIL, Mutex, ConditionVariable, Queue, SizedQueue, Thread::SizedQueue, Thread::Mutex, Thread::Queue, Thread.group, Thread.list, Thread.current, Thread.main, Thread.pass, Fiber.new, Fiber.resume, Fiber.yield, Fiber.schedule, scheduler, non-blocking, concurrent-ruby, async, Async, actor, Ractor.new, Ractor.receive, Ractor.send, Ractor.yield, Thread.pool, Thread.abort_on_exception, thread variable, Thread.current[:name], GVL release | [references/concurrency.md](references/concurrency.md) |
| ObjectSpace, ObjectSpace.each_object, ObjectSpace.define_finalizer, ObjectSpace._id2ref, GC, GC.start, GC.stat, GC.disable, GC.enable, GC.stress, generational GC, RGENGC, write barrier, weak reference, WeakRef, WeakMap, ObjectSpace::WeakMap, finalizer, _id2ref, VALUE, immediate value, tagged pointer, RVALUE, RSTRING, memory layout, object shape, shape pool, frozen heap, malloc/free | [references/gc-memory.md](references/gc-memory.md) |
| RubyGems, gem, gem install, gem list, gem uninstall, gem update, gem env, gem which, gemspec, Gem::Specification, Gemfile, Gemfile.lock, Bundler, bundle, bundle install, bundle update, bundle exec, bundle config, bundle add, bundle outdated, bundle lock, bundle check, bundle clean, bundle open, bundle binstubs, groups, :development, :test, :production, source, rubygems.org, git source, path source, BUNDLE_GEMFILE, Gemfile.lock, vendor/cache, Gem::Version, Gem::Requirement | [references/gems-bundler.md](references/gems-bundler.md) |
| ruby, ruby interpreter, ruby flag, -e, -I, -r, -n, -p, -a, -F, -l, -0, -K, -E, -w, -W, -d, --enable-gems, --disable-gems, --enable-frozen-string-literal, --disable-frozen-string-literal, irb, .irbrc, binding.irb, erb, ERB, Rake, Rakefile, rake task, rake file task, rake namespace, rake multitask, rake rule, rdoc, ri, RDoc, RDoc::Markup, ripper, ruby-prof, stackprof, memory_profiler, debug, debug.rb, byebug, rdbg, rbs, typeprof, rubocop, standard, ruby version manager, rvm, rbenv, chruby, asdf | [references/cli.md](references/cli.md) |
| environment variable, RUBYOPT, RUBYLIB, RUBY_ENGINE, RUBY_VERSION, RUBY_PLATFORM, RUBY_PATCHLEVEL, RUBY_RELEASE_DATE, RUBY_DESCRIPTION, GEM_HOME, GEM_PATH, GEM_ROOT, GEM_RC, BUNDLE_GEMFILE, BUNDLE_PATH, BUNDLE_FROZEN, BUNDLE_JOBS, BUNDLE_WITHOUT, BUNDLE_BIN, BUNDLE_IGNORE_CONFIG, RUBY_CONFIGURE_OPTS, BUNDLER_VERSION, BUNDLE_CACHE_PATH, PATH modifications, rubyopt parsing, frozen-string-literal default | [references/environment.md](references/environment.md) |
| standard library, stdlib, set, monitor, singleton, forwardable, delegate, observer, logger, optparse, fileutils, pathname, tempfile, tmpdir, json, yaml, psych, csv, erb, digest, securerandom, time, date, matrix, tsort, pp, prettyprint, benchmark, debug_inspector, objspace, io/console, io/nonblock, socket, openssl, net/http, net/smtp, net/pop, net/imap, uri, cgi | [references/stdlib.md](references/stdlib.md) |

## Quick Guide

- **How does method lookup work in Ruby?** → [references/object-model.md](references/object-model.md)
- **What's the difference between include and prepend?** → [references/object-model.md](references/object-model.md)
- **How do I dynamically define a method?** → [references/metaprogramming.md](references/metaprogramming.md)
- **How does method_missing work?** → [references/metaprogramming.md](references/metaprogramming.md)
- **What are Refinements?** → [references/metaprogramming.md](references/metaprogramming.md)
- **What's the difference between ==, eql?, and equal??** → [references/types.md](references/types.md)
- **Are Symbols just frozen Strings?** → [references/types.md](references/types.md)
- **What's the difference between a Proc and a lambda?** → [references/blocks-procs-lambdas.md](references/blocks-procs-lambdas.md)
- **How do I capture a block as a variable?** → [references/blocks-procs-lambdas.md](references/blocks-procs-lambdas.md)
- **How do exceptions and rescue work?** → [references/control-flow.md](references/control-flow.md)
- **What does rescue without a class catch?** → [references/control-flow.md](references/control-flow.md)
- **How do I read a file in Ruby?** → [references/io-encoding.md](references/io-encoding.md)
- **How does string encoding work?** → [references/io-encoding.md](references/io-encoding.md)
- **How do I freeze string literals by default?** → [references/io-encoding.md](references/io-encoding.md)
- **Does Ruby have a GIL?** → [references/concurrency.md](references/concurrency.md)
- **What's the difference between Thread and Fiber?** → [references/concurrency.md](references/concurrency.md)
- **What is a Ractor?** → [references/concurrency.md](references/concurrency.md)
- **How does the Ruby GC work?** → [references/gc-memory.md](references/gc-memory.md)
- **How do I find all instances of a class?** → [references/gc-memory.md](references/gc-memory.md)
- **How do I add a gem dependency?** → [references/gems-bundler.md](references/gems-bundler.md)
- **What's the difference between Gemfile and gemspec?** → [references/gems-bundler.md](references/gems-bundler.md)
- **How does `bundle exec` work?** → [references/gems-bundler.md](references/gems-bundler.md)
- **What CLI flags does the ruby interpreter have?** → [references/cli.md](references/cli.md)
- **How do I start an interactive Ruby session?** → [references/cli.md](references/cli.md)
- **What environment variables affect Ruby?** → [references/environment.md](references/environment.md)
- **What's in the standard library?** → [references/stdlib.md](references/stdlib.md)

## Cross-Project References

- For **Rails** CLI development (commands, generators, database tasks), use the **rails** skill (in this repo).
- For carapace library development (Action API, traverse engine, shell formatters), use the **carapace-dev** skill (in this repo).
- For carapace user-facing integration (specs, macros, setup), use the **carapace** skill (in this repo).
- For the official Ruby language reference and C-API, see <https://ruby-doc.org/> — this skill covers the language and runtime model, not the C extension API.
- For the RubySpec (executable spec suite that defines conformance), see <https://github.com/ruby/spec>.
