package main

import "fmt"

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

	fs.createDir("etc")

	fmt.Printf("%v", fs.currentDir.children)
}
