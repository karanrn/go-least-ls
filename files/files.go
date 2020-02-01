package main

import (
	"flag"
	"fmt"
	"github.com/karanrn/go-least-ls/helper"
	"os"
	"path/filepath"
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
go-least-ls.exe -older 30 -count 10 -filetype .txt [-all]
> Lists least 10 recently used files (older than 30 days).

-help     : Gets the help.
-older    : How much older files? Accepts integer in units of days. Default 30
-filetype : Extension of files to be searched. Ex: .exe, .doc, .txt
-count    : Number of files to view. Default is 5
-all      : Optional flag to list hidden files in the output
`

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
	older := flag.Int("older", 30, "How older?")
	counter := flag.Int("count", 5, "Number of files to view.")
	fileType := flag.String("filetype", "", "Enter the file type to search.")
	hidden := flag.Bool("all", false, "Include hidden files.")
	help := flag.Bool("help", false, "Get usage help.")

	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Println("Check usage- go-least-ls.exe -help")
		os.Exit(0)
	}

	// Stupid way to implement an option
	if *help {
		toolUsage()
		os.Exit(0)
	}

	// Check for file type
	if *fileType == ""{
		fmt.Println("File type is not given")
		fmt.Println("For more help: go-least-ls -help")
		os.Exit(0)
	}

	extType := *fileType
	days = *older
	count = *counter
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
		if extType != filepath.Ext(file){
			continue
		}
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

	if len(lastFiles) != 0{
		sortedfiles := helper.Sort(lastFiles)
		for i := 0; i < len(sortedfiles); i++ {
			if count == 0 {
				break
			}
			f := sortedfiles[i]
			rel, _ := filepath.Rel(root, f.Key)
			
			// Default show all files
			if *hidden{
				fmt.Printf("%s : %s \n", rel, f.Value)
			}else {
				if helper.IsHidden(rel){
					count = count - 1
					continue
				}else {
					fmt.Printf("%s : %s \n", rel, f.Value)
				}
			}
			count = count - 1
		}
	}else{
		fmt.Printf("No files found for %s extension", *fileType)
		fmt.Println("Extension is either invalid or no files with such extension exist.")
	}

}
