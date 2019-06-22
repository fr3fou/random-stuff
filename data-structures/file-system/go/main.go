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

}
