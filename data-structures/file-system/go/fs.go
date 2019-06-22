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
	changeDir(path string) (*file, error)
	createFile(path string, content string) error
	createDir(name string)
	read(path string)
}

// file implementations

func (f *file) changeDir(path string) (*file, error) {
	if !f.isDir {
		return nil, errors.New("fs: can't change to a file")
	}

	// get all the files in the path
	files := strings.Split(path, "/")

	file, rest := files[0], files[1:]

	// up a dir
	if file == ".." {
		// no parent means we are at root
		if f.parent == nil {
			// we can just ignore the ../
			cf, err := f.changeDir(strings.Join(rest, "/"))

			if err != nil {
				return nil, err
			}

			return cf, nil
		}

		// if there are no files left, we just return the parent
		if len(files) == 1 {
			return f.parent, nil
		}

		// we have to keep going if there are files left
		cf, err := f.parent.changeDir(strings.Join(rest, "/"))

		if err != nil {
			return nil, err
		}

		return cf, nil

	}

	cf, ok := f.children[file]

	if !ok {
		return nil, errors.New("fs: can't change to a directory that doesn't exist")
	}

	// we have reached the end of the path
	if len(files) == 1 {
		return cf, nil
	}

	cf, err := cf.changeDir(strings.Join(rest, "/"))

	if err != nil {
		return nil, err
	}

	return cf, nil
}

func (f *file) createDir(name string) error {
	if _, ok := f.children[name]; ok {
		return errors.New("fs: can't create a directory that already exists")
	}

	var path string

	// no parent means we are at root
	if f.parent == nil {
		path = f.path + name
	} else {
		path = f.path + "/" + name
	}

	f.children[name] = &file{
		isDir:    true,
		parent:   f,
		name:     name,
		children: make(children),
		path:     path,
	}

	return nil
}

// fs implementations

func (f *fs) changeDir(path string) (*file, error) {
	cf, err := f.currentDir.changeDir(path)

	if err != nil {
		return nil, err
	}

	f.currentDir = cf

	return f.currentDir, nil
}

func (f *fs) createDir(name string) error {
	return f.currentDir.createDir(name)
}
