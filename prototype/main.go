package main

import "fmt"

func main() {
	file1 := &file{name: "file1"}
	file2 := &file{name: "file2"}
	file3 := &file{name: "file3"}
	folder1 := &folder{
		childrens: []inode{file1},
		name:      "folder1",
	}
	folder2 := &folder{
		childrens: []inode{folder1, file2, file3},
		name:      "folder2",
	}
	fmt.Println("printing hierarchy for folder2")
	folder2.print(" ")
	cloneFolder := folder2.clone()
	fmt.Println("print hierarchy for clone folder")
	cloneFolder.print(" ")
}
