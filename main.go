package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

//рекурсивный проход директории
func recursiveWatcher(output io.Writer, path string, printFiles bool, indent string) error {
	files, err := getEntriesFromFolder(path)
	var elementsCount int = getElementsCount(files, printFiles)
	var fileSize string
	var indentString string
	var localIndent string
	var isLastElement bool
	var counter int

	if err != nil {
		return fmt.Errorf("Something wrong")
	}

	for _, file := range files {
		isDir := file.IsDir()

		if !printFiles && !isDir {
			continue
		} else {
			counter++
		}

		if !isDir {
			if file.Size() > 0 {
				fileSize = " (" + strconv.Itoa(int(file.Size())) + "b)"
			} else {
				fileSize = " (empty)"
			}
		} else {
			fileSize = ""
		}

		if counter == elementsCount {
			isLastElement = true
			indentString = "└───"
		} else {
			isLastElement = false
			indentString = "├───"
		}

		fmt.Fprintf(output, indent+indentString+file.Name()+fileSize+"\n")

		if isDir {
			if isLastElement {
				localIndent = indent + "\t"
			} else {
				localIndent = indent + "│" + "\t"
			}
			recursiveWatcher(output, path+string(os.PathSeparator)+file.Name(), printFiles, localIndent)
		}
	}

	return nil
}

//количество елементов (учитывая/неучитывая) файлы
func getElementsCount(elements []os.FileInfo, withFiles bool) int {
	if withFiles {
		return len(elements)
	} else {
		var count = 0
		for _, element := range elements {
			if element.IsDir() {
				count++
			}
		}
		return count
	}
}

//получение списка файлов и директорий
func getEntriesFromFolder(path string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		return nil, fmt.Errorf("Something wrong")
	}

	return files, nil
}

func dirTree(output io.Writer, path string, printFiles bool) error {

	err := recursiveWatcher(output, path, printFiles, "")

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func main() {
	out := os.Stdout

	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		fmt.Println("No folder in input. Run program like 'go run main.go . [-f]'")
		os.Exit(2)
	}

	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"

	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
