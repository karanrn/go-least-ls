package main

import(
	"fmt"
	"os"
	"path/filepath"
	"time"
	"github.com/karanrn/go-least-ls/helper"
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

func main(){
	var files []string
	var lastFiles = make(map[string]time.Time)

	// 30 days before current time
	archiveDate := time.Now().Local().AddDate(0, 0, -30)

	// File extension types
	//var documents = []string{".pdf", ".doc", ".docx", ".txt"}

	// Current folder/directory
	root, err := os.Getwd()
	err = filepath.Walk(root, visit(&files))
	if err != nil{
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
	count := 5
	for i := 0; i < len(sortedfiles) ; i++ {
		if count == 0 {
			break
		}
		f := sortedfiles[i]
		fmt.Printf("%s : %s \n", f.Key, f.Value)
		count = count - 1
	}
	
}