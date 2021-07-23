# `killfam`

`killfam` is a Go package to provide cross-platform process tree killing
capability.

Doing `cmd.Process.Kill()` kills a process, but not its children, which
continue to run. Whereas `killfam.KillTree(cmd)` kills the whole family,
parents and children alike. Hence the (dark) name.

It is particularly useful for terminating shell scripts that execute other
processes as children. 
