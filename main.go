package main

import (
	"fmt"
	"os"
	"strings"
)

func isImage(extension string) bool {
	switch extension {
	case "jpeg", "jpg", "png", "gif", "ai", "indd", "raw", "tiff", "eps":
		return true
	default:
		return false
	}

}

func isVideo(extension string) bool {
	switch extension {
	case "mp4", "webm", "mkv", "flv", "vob", "mov", "qt", "avi", "wmv", "rmvb", "amv", "rm", "yuv", "MTS", "TS", "M2TS", "mng":
		return true
	default:
		return false
	}

}

func isPdf(extension string) bool {
	return extension == "pdf"
}

func isExcel(extension string) bool {
	return extension == "xlsx" || extension == "csv"
}

func isText(extension string) bool {
	return extension == "txt" || extension == "odt" || extension == "doc"
}

func isAnOrganizingDir(dirName string) bool {
	return dirName == "Images" || dirName == "Videos" || dirName == "PDF" || dirName == "Excels" || dirName == "Docs" || dirName == "Others"
}

func correspondingDir(extension string) string {
	var directory string
	if isPdf(extension) {
		directory = "PDF"
	} else if isText(extension) {
		directory = "Docs"
	} else if isImage(extension) {
		directory = "Images"
	} else if isVideo(extension) {
		directory = "Videos"
	} else if isExcel(extension) {
		directory = "Excels"
	} else {
		directory = "Others"
	}

	return directory
}

func moveEntry(extension, entryName string) {
	newDirectory := correspondingDir(extension)

	err := os.Rename("/home/hazem/Downloads/"+entryName, "/home/hazem/Downloads/"+newDirectory+"/"+entryName)
	if err != nil {
		fmt.Println("Error while moving entry : ", err.Error())
		return
	}

	fmt.Println(entryName, " moved successfully to : ", newDirectory)

}

func main() {
	entries, err := os.ReadDir("/home/hazem/Downloads")
	if err != nil {
		fmt.Println("error reading the directory: ", err.Error())
	}

	for _, entry := range entries {
		if entry.IsDir() && isAnOrganizingDir(entry.Name()) {
			continue
		}

		splittedEntryName := strings.Split(entry.Name(), ".")
		var extension string
		if len(splittedEntryName) >= 2 {
			extension = splittedEntryName[1]
		} else {
			extension = ""
		}

		moveEntry(extension, entry.Name())
	}

}
