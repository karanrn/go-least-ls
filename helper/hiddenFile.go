package helper

import (
	"runtime"
	"fmt"
 	_ "syscall"
)

// IsHidden method checks if the file/directory is hidden.
// Returns true if hidden, checks for both windows and unix/linux systems.
func IsHidden(filename string) (bool){
	
	// Check the OS type
	if runtime.GOOS != "windows"{

		// Unix/linux file/directory starting with . is hidden
		if filename[0:1] == "." {
			return true
		}else {
			return false
		}
	}else {
		fmt.Println("Unable to check if file is hidden under this OS!")
	}
	return false
}