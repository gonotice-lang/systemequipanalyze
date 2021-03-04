package models

// ProcessInfo - Launched process Infomations
type ProcessInfo struct {
	ProcPid  uint64
	ProcTty  string
	ProcTime string
	ProcCmd  string
}
