package main

import (
	"flag"
	"fmt"
	"github.com/karanrn/go-least-ls/helper"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

/*
	List the least used files and group them based on file type.
	List last five used in each group.
*/

// Walk the files in the path
func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
		}
		*files = append(*files, path)
		return nil
	}
}

// Check the extension of the file
// naked return : returns zero values when no match
func Find(slice []string, val string) (i int, flag bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return
}

// Flag usage
func toolUsage() {
	// Command line flags
	var helpUsage string = `go-least-ls:
A command line utility to list least recently used files in the current directory.

Usage:
go-least-ls.exe -older 30 -count 10
> Lists least 10 recently used files (older than 30 days).

-help  : Gets the help.
-older : How much older files? Accepts integer in units of days.
-count : Number of files to view.`

	fmt.Printf(helpUsage)
	os.Exit(0)
}

func main() {
	var files []string
	var lastFiles = make(map[string]time.Time)
	var days int
	var count int

	// Command line flags
	flag.Usage = toolUsage
	older := flag.String("older", "", "how older?")
	counter := flag.String("count", "", "number of files to view.")
	help := flag.String("help", "", "Get usage help.")

	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Println("Check usage- go-least-ls.exe -help")
		os.Exit(0)
	}

	// Stupid way to implement an option
	if flag.NFlag() == 1 && *help == "" {
		toolUsage()
		os.Exit(0)
	}

	if *older != "" && *counter != "" {
		var err error
		days, err = strconv.Atoi(*older)
		if err != nil {
			fmt.Println("-older value has to be integer.")
			os.Exit(-1)
		}

		count, err = strconv.Atoi(*counter)
		if err != nil {
			fmt.Println("-count value has to be integer.")
			os.Exit(-1)
		}

	} else {
		fmt.Println("Check usage- go-least-ls.exe -help")
	}

	// 30 days before current time
	archiveDate := time.Now().Local().AddDate(0, 0, -days)

	// File extension types
	//var documents = []string{".pdf", ".doc", ".docx", ".txt"}

	// Current folder/directory
	root, err := os.Getwd()
	err = filepath.Walk(root, visit(&files))
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		info, err := os.Stat(file)
		if err != nil {
			panic(err)
		}
		// _, found := Find(documents, filepath.Ext(file))

		lastAccess := info.ModTime()
		if lastAccess.Before(archiveDate) || lastAccess.Equal(archiveDate) {
			//fmt.Printf("%s : %s\n",file, info.ModTime())
			lastFiles[file] = lastAccess
		}
	}

	sortedfiles := helper.Sort(lastFiles)
	for i := 0; i < len(sortedfiles); i++ {
		if count == 0 {
			break
		}
		f := sortedfiles[i]
		fmt.Printf("%s : %s \n", f.Key, f.Value)
		count = count - 1
	}

}
