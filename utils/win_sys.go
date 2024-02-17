package utils

import (
	"golang.org/x/sys/windows"
)

type (
	BOOL   int32
	HANDLE uintptr
)

type SECURITY_ATTRIBUTES struct {
	Length             uint32
	SecurityDescriptor uintptr
	InheritHandle      BOOL
}

type STARTUPINFOW struct {
	cb            uint32
	_             *uint16
	Desktop       *uint16
	Title         *uint16
	X             uint32
	Y             uint32
	XSize         uint32
	YSize         uint32
	XCountChars   uint32
	YCountChars   uint32
	FillAttribute uint32
	Flags         uint32
	ShowWindow    uint16
	_             uint16
	_             *uint8
	StdInput      HANDLE
	StdOutput     HANDLE
	StdError      HANDLE
}

type PROCESS_INFORMATION struct {
	Process   HANDLE
	Thread    HANDLE
	ProcessId uint32
	ThreadId  uint32
}

// StartupInfoEx Used for outdated CreateProcessA and not CreateProcessW
type StartupInfoEx struct {
	StartupInfo   windows.StartupInfo
	AttributeList *PROC_THREAD_ATTRIBUTE_LIST
}

type PROC_THREAD_ATTRIBUTE_LIST struct {
	dwFlags  uint32
	size     uint64
	count    uint64
	reserved uint64
	unknown  *uint64
	entries  []*PROC_THREAD_ATTRIBUTE_ENTRY
}

type PROC_THREAD_ATTRIBUTE_ENTRY struct {
	attribute *uint32
	cbSize    uintptr
	lpValue   uintptr
}

//var (
//	modKernel32        = windows.NewLazySystemDLL("kernel32.dll")
//	procCreateProcessW = modKernel32.NewProc("CreateProcessW")
//	procVirtualAllocEx = modKernel32.NewProc("VirtualAllocEx")
//)
//
//
//func CreateProcessW(
//	lpApplicationName, lpCommandLine string,
//	lpProcessAttributes, lpThreadAttributes *SECURITY_ATTRIBUTES,
//	bInheritHandles BOOL,
//	dwCreationFlags uint32,
//	lpEnvironment unsafe.Pointer,
//	lpCurrentDirectory string,
//	lpStartupInfo *STARTUPINFOW,
//	lpProcessInformation *PROCESS_INFORMATION,
//) (e error) {
//	var lpAN, lpCL, lpCD *uint16
//	if len(lpApplicationName) > 0 {
//		lpAN, e = windows.UTF16PtrFromString(lpApplicationName)
//		if e != nil {
//			return
//		}
//	}
//	if len(lpCommandLine) > 0 {
//		lpCL, e = windows.UTF16PtrFromString(lpCommandLine)
//		if e != nil {
//			return
//		}
//	}
//	if len(lpCurrentDirectory) > 0 {
//		lpCD, e = windows.UTF16PtrFromString(lpCurrentDirectory)
//		if e != nil {
//			return
//		}
//	}
//	ret, _, lastErr := procCreateProcessW.Call(
//		uintptr(unsafe.Pointer(lpAN)),
//		uintptr(unsafe.Pointer(lpCL)),
//		uintptr(unsafe.Pointer(lpProcessAttributes)),
//		uintptr(unsafe.Pointer(lpProcessInformation)),
//		uintptr(bInheritHandles),
//		uintptr(dwCreationFlags),
//		uintptr(lpEnvironment),
//		uintptr(unsafe.Pointer(lpCD)),
//		uintptr(unsafe.Pointer(lpStartupInfo)),
//		uintptr(unsafe.Pointer(lpProcessInformation)),
//	)
//	if ret == 0 {
//		e = lastErr
//	}
//	return
//}
