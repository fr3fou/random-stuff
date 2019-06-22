package main

import (
	"errors"
	"strings"
)

type fs struct {
	root       *file
	currentDir *file
}

type children map[string]*file

type file struct {
	name     string
	path     string
	children children
	isDir    bool
	parent   *file
}

// File contains all the methods
type File interface {
	changeDir(path string) *file
	createFile(path string, content string) error
	createDir(name string)
	read(path string)
}

func (f *file) changeDir(path string) *file {
	if !f.isDir {
		return f
	}

	files := strings.Split(path, "/")
	file := files[0]

	if _, ok := f.children[file]; ok {
		rest := files[1:]
		f.path = f.children[file].changeDir(strings.Join(rest, "/")).path
	}

	return f
}

func (f *file) createDir(name string) error {

	if _, ok := f.children[name]; ok {
		return errors.New("fs: can't create a directory that already exists")
	}

	f.children[name] = &file{

		isDir:    true,
		parent:   f,
		name:     name,
		children: make(children),
		path:     f.path + "/" + name,
	}

	return nil
}

func (f *fs) changeDir(path string) *file {
	f.currentDir = f.currentDir.changeDir(path)

	return f.currentDir
}

func (f *fs) createDir(name string) error {
	return f.currentDir.createDir(name)
}
