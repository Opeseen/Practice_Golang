package main

import (
	"fmt"
	"path/filepath"
)

// .............................FilePath
// func main() {
// 	path := filepath.Join("Dir1", "Dir2", "Dir3", "test.txt")
// 	fmt.Println(path)
// 	fmt.Println(filepath.Dir(path))
// 	fmt.Println(filepath.Base(path))
// 	fmt.Println(filepath.IsAbs(path))
// 	fmt.Println(filepath.IsAbs("/dir/file"))
// 	fmt.Println(filepath.Ext(path))
// }

// .................................... ReadFile
// func main() {
// 	fileInfo, err1 := os.Stat("C:/Users/HomePC/Desktop/Tutorials/CodePractice/FIRST GITHUB FILES/Practice_Golang/kodekloud-go/advance-go/corePackages/file.txt")
// 	fileText, err2 := os.ReadFile("C:/Users/HomePC/Desktop/Tutorials/CodePractice/FIRST GITHUB FILES/Practice_Golang/kodekloud-go/advance-go/corePackages/file.txt")
// 	path := "./file.txt"
// 	file, err3 := os.Open(path)

// 	if err3 != nil {
// 		fmt.Println("Error occurred ", err3)
// 	}
// 	if err1 != nil {
// 		fmt.Println(err1)
// 	}
// 	if err2 != nil {
// 		fmt.Println(err2)
// 	}

// 	b := make([]byte, 4)
// 	for {
// 		n, err := file.Read(b)
// 		if err != nil {
// 			fmt.Println("Error occurred ", err)
// 			break
// 		}
// 		fmt.Println(b[0:n])
// 		fmt.Println(string(b[0:n]))
// 	}
// 	fmt.Println()
// 	fmt.Println(fileInfo.Name())
// 	fmt.Println(fileInfo.Size())
// 	fmt.Println(fileInfo.Mode())
// 	fmt.Println(fileInfo.IsDir())
// 	fmt.Println()
// 	fmt.Println(fileText)
// 	fmt.Println(string(fileText))
// }

// func main() {
// 	file, err := os.OpenFile("./file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer file.Close()
// 	_, err = file.WriteString("Hope you had a good day..? \n")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

func main() {
	fmt.Println(filepath.Join("dir1", "../dir1/dir2", "dir3", "temp.txt"))
	fmt.Println(filepath.Join("dir1", "dir2", "dir3", "temp.txt"))
	fmt.Println(filepath.Join("dir2", "../dir1/dir2", "dir3", "temp.txt"))
	fmt.Println(filepath.Join("dir2", "../dir1/dir2//", "dir3", "../dir3/temp.txt"))
}
