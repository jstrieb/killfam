// killfam provides a pure-Go, cross-platform interface for creating processes
// whose entire process tree can be killed.
//
// In other words, processes created with killfam are augmented with a method
// that allows them to be killed with an OS-agnostic function such that their
// child processes are also killed.
package killfam
