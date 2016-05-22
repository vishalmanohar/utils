package main

import (
  "fmt"
  "os"
  "strconv"
  "io/ioutil"
  "time"
)

func main() {
  startTime := time.Now()
  directory := os.Args[1]
  updateFileSizes(directory, directory)

  files := getFiles(directory)
  for _, fileInfo := range files {
    fileName := fileInfo.Name()
    subDirName := subDirName(directory, fileName)
    updateFileSizes(subDirName, subDirName)
  }

  fmt.Printf("File size: %s - %s\n", directory, formatFileSize(fileSizes[directory]))

  fmt.Println("|--")
  for dirName, fileSize := range fileSizes {
    fmt.Printf(" -- %s - %s\n", dirName, formatFileSize(fileSize))
  }

  fmt.Printf("Completed in %f secs\n", time.Since(startTime).Seconds())
}

var fileSizes = make(map[string] int64)

func formatFileSize(fileSize int64) string {
  return strconv.FormatInt(fileSize / 1000000, 10) + " Mb"
}

func getFiles(directory string) []os.FileInfo {
  fileInfos, _ := ioutil.ReadDir(directory)
  return fileInfos
}

func subDirName(dirPath string, childDirName string) string {
  return dirPath + "/" + childDirName
}

func updateFileSizes(directory string, currentDirectory string) {
  fileInfos, _ := ioutil.ReadDir(currentDirectory)
  for _, fileInfo := range fileInfos {
    if !fileInfo.IsDir() {
      fileSizes[directory] += fileInfo.Size()
    } else {
      updateFileSizes(directory, subDirName(currentDirectory, fileInfo.Name()))
    }
  }
}
