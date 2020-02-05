package tests

import (
	"testing"
	"github.com/karanrn/go-least-ls/helper"
)

func TestContains(t *testing.T){
	slist := []string {"file1", "file2", "file3"}
	strue := "file1"
	//sfalse := "file"

	res := helper.Contains(slist, strue)
	if !res {
		t.Errorf("%s does contain in the list", strue)
	}
}

func TestIsHidden(t *testing.T){
	fileName := ".hello.txt"
	directory := ".vscode/file1.exe"

	// Check for hidden file
	res := helper.IsHidden(fileName)
	if !res {
		t.Error("File is actually hidden")
	}

	// Check for hidden directory
	res = helper.IsHidden(directory)
	if !res {
		t.Error("Directory is actually hidden")
	}
}
