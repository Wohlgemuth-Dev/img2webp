package sysutil

import (
	"syscall"
	"unsafe"
)

// IsLaunchedFromExplorer uses the Windows API to check how many processes are attached to
// the current console. If it's only 1 (this program), it means it created its own console
// window because it was double-clicked in Explorer. If >1, it was launched from an existing terminal.
func IsLaunchedFromExplorer() bool {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	procGetConsoleProcessList := kernel32.NewProc("GetConsoleProcessList")

	var pids [2]uint32
	ret, _, _ := procGetConsoleProcessList.Call(
		uintptr(unsafe.Pointer(&pids[0])),
		uintptr(len(pids)),
	)
	return ret == 1
}
