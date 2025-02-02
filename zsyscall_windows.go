// Code generated by 'go generate'; DO NOT EDIT.

package serial

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL     error = syscall.EINVAL
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modadvapi32 = windows.NewLazySystemDLL("advapi32.dll")
	modkernel32 = windows.NewLazySystemDLL("kernel32.dll")

	procRegEnumValueW       = modadvapi32.NewProc("RegEnumValueW")
	procClearCommBreak      = modkernel32.NewProc("ClearCommBreak")
	procCreateEventW        = modkernel32.NewProc("CreateEventW")
	procEscapeCommFunction  = modkernel32.NewProc("EscapeCommFunction")
	procGetCommModemStatus  = modkernel32.NewProc("GetCommModemStatus")
	procGetCommState        = modkernel32.NewProc("GetCommState")
	procGetOverlappedResult = modkernel32.NewProc("GetOverlappedResult")
	procPurgeComm           = modkernel32.NewProc("PurgeComm")
	procResetEvent          = modkernel32.NewProc("ResetEvent")
	procSetCommBreak        = modkernel32.NewProc("SetCommBreak")
	procSetCommState        = modkernel32.NewProc("SetCommState")
	procSetCommTimeouts     = modkernel32.NewProc("SetCommTimeouts")
	procSetupComm           = modkernel32.NewProc("SetupComm")
)

func regEnumValue(key syscall.Handle, index uint32, name *uint16, nameLen *uint32, reserved *uint32, class *uint16, value *uint16, valueLen *uint32) (regerrno error) {
	r0, _, _ := syscall.Syscall9(procRegEnumValueW.Addr(), 8, uintptr(key), uintptr(index), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(nameLen)), uintptr(unsafe.Pointer(reserved)), uintptr(unsafe.Pointer(class)), uintptr(unsafe.Pointer(value)), uintptr(unsafe.Pointer(valueLen)), 0)
	if r0 != 0 {
		regerrno = syscall.Errno(r0)
	}
	return
}

func clearCommBreak(handle syscall.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procClearCommBreak.Addr(), 1, uintptr(handle), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func createEvent(eventAttributes *uint32, manualReset bool, initialState bool, name *uint16) (handle syscall.Handle, err error) {
	var _p0 uint32
	if manualReset {
		_p0 = 1
	}
	var _p1 uint32
	if initialState {
		_p1 = 1
	}
	r0, _, e1 := syscall.Syscall6(procCreateEventW.Addr(), 4, uintptr(unsafe.Pointer(eventAttributes)), uintptr(_p0), uintptr(_p1), uintptr(unsafe.Pointer(name)), 0, 0)
	handle = syscall.Handle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

func escapeCommFunction(handle syscall.Handle, function uint32) (res bool) {
	r0, _, _ := syscall.Syscall(procEscapeCommFunction.Addr(), 2, uintptr(handle), uintptr(function), 0)
	res = r0 != 0
	return
}

func getCommModemStatus(handle syscall.Handle, bits *uint32) (res bool) {
	r0, _, _ := syscall.Syscall(procGetCommModemStatus.Addr(), 2, uintptr(handle), uintptr(unsafe.Pointer(bits)), 0)
	res = r0 != 0
	return
}

func getCommState(handle syscall.Handle, dcb *dcb) (err error) {
	r1, _, e1 := syscall.Syscall(procGetCommState.Addr(), 2, uintptr(handle), uintptr(unsafe.Pointer(dcb)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func getOverlappedResult(handle syscall.Handle, overlapEvent *syscall.Overlapped, n *uint32, wait bool) (err error) {
	var _p0 uint32
	if wait {
		_p0 = 1
	}
	r1, _, e1 := syscall.Syscall6(procGetOverlappedResult.Addr(), 4, uintptr(handle), uintptr(unsafe.Pointer(overlapEvent)), uintptr(unsafe.Pointer(n)), uintptr(_p0), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func purgeComm(handle syscall.Handle, flags uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procPurgeComm.Addr(), 2, uintptr(handle), uintptr(flags), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func resetEvent(handle syscall.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procResetEvent.Addr(), 1, uintptr(handle), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func setCommBreak(handle syscall.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procSetCommBreak.Addr(), 1, uintptr(handle), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func setCommState(handle syscall.Handle, dcb *dcb) (err error) {
	r1, _, e1 := syscall.Syscall(procSetCommState.Addr(), 2, uintptr(handle), uintptr(unsafe.Pointer(dcb)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func setCommTimeouts(handle syscall.Handle, timeouts *commTimeouts) (err error) {
	r1, _, e1 := syscall.Syscall(procSetCommTimeouts.Addr(), 2, uintptr(handle), uintptr(unsafe.Pointer(timeouts)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func setupComm(handle syscall.Handle, rxSize uint32, txSize uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupComm.Addr(), 3, uintptr(handle), uintptr(rxSize), uintptr(txSize))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}
