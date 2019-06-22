package main

import (
	"fmt"
	"github.com/fr3fou/go-fs/filesystem"
)

func main() {
	fs := filesystem.New()

	fs.CreateDir("usr")
	fs.ChangeDir("usr")

	fmt.Println("fs: changed dir to: " + fs.PrintWorkingDirectory())

	fs.CreateDir("share")
	fs.ChangeDir("share")

	fmt.Println("fs: changed dir to: " + fs.PrintWorkingDirectory())

	fs.ChangeDir("../..")

	fmt.Println("fs: changed dir to: " + fs.PrintWorkingDirectory())

	fs.ChangeDir("usr/share/")

	fmt.Println("fs: changed dir to: " + fs.PrintWorkingDirectory())

	fs.ChangeDir("/usr/share")

	fmt.Println("fs: changed dir to: " + fs.PrintWorkingDirectory())

	fs.ChangeDir("..")

	fmt.Println("fs: changed dir to: " + fs.PrintWorkingDirectory())

	fs.ChangeDir("/")

	fmt.Println("fs: changed dir to: " + fs.PrintWorkingDirectory())

	fs.CreateFile("kernel", []byte("hello"))

	content, _ := fs.ReadFile("kernel")
	fmt.Println(string(content))

	fs.ChangeDir("usr")

	fs.CreateFile("nested", []byte("some content"))
	fs.ChangeDir("../")

	content, _ = fs.ReadFile("usr/nested")
	fmt.Println(string(content))
}
