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

	fmt.Println("fs: changed dir to: " + fs.PrintWorkingDirectory())

	fs.CreateFile("nested", []byte("some content"))
	fs.ChangeDir("../")

	fmt.Println("fs: changed dir to: " + fs.PrintWorkingDirectory())

	content, _ = fs.ReadFile("usr/nested")
	fmt.Println(string(content))

	fs.CreateFile("/usr/share/testing", []byte("testing root files"))

	content, _ = fs.ReadFile("/usr/share/testing")
	fmt.Println(string(content))

	fs.CreateDir("/usr/share/fog")
	fs.ChangeDir("/usr/share/fog")

	fmt.Println("fs: changed dir to: " + fs.PrintWorkingDirectory())

	fs.CreateDir("/usr/share/fog/this/shouldnt-work")
	fs.ChangeDir("/usr/share/fog/this/shouldnt-work")
}
