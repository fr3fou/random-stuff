package main

import (
	"fmt"
)

func main() {
	root := &file{
		name:     "/",
		path:     "/",
		isDir:    true,
		children: make(children),
		parent:   nil,
	}

	fs := fs{
		root:       root,
		currentDir: root,
	}

	fs.createDir("usr")
	fs.changeDir("usr")

	fmt.Println(fs.currentDir.path)

	fs.createDir("share")
	fs.changeDir("share")

	fmt.Println(fs.currentDir.path)

	fs.changeDir("../..")

	fmt.Println(fs.currentDir.path)

	fs.changeDir("usr/share/")

	fmt.Println(fs.currentDir.path)

	fs.changeDir("/usr/share")

	fmt.Println(fs.currentDir.path)
}
