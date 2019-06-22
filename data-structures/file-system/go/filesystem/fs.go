package filesystem

import (
	"errors"
	"strings"
)

// Fs is the struct for the fileSystem
type Fs struct {
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
	content  []byte
}

// File contains all the methods
type File interface {
	changeDir(path string) (*file, error)
	createFile(path string, content string) error
	createDir(name string) error
	create(file file) error
	read(path string)
}

// file implementations

func (f *file) changeDir(path string) (*file, error) {
	if !f.isDir {
		return nil, errors.New("fs: can't change to a file")
	}

	// walk up the tree to the root
	if strings.HasPrefix(path, "/") && f.parent != nil {
		return f.parent.changeDir(path)
	}

	// if our target is the root and we have walked up to it, just return it
	if path == "/" {
		return f, nil
	}

	// get all the files in the path
	files := strings.Split(strings.Trim(path, "/"), "/")
	file, rest := files[0], files[1:]

	// going up a dir
	if file == ".." {
		// no parent means we are at root
		if f.parent == nil {
			// we can just ignore the ../
			return f.changeDir(strings.Join(rest, "/"))
		}

		// if there are no files left, we just return the parent
		if len(files) == 1 {
			return f.parent, nil
		}

		// we have to keep going if there are files left
		return f.parent.changeDir(strings.Join(rest, "/"))
	}

	cf, ok := f.children[file]

	if !ok {
		return nil, errors.New("fs: can't change to a directory that doesn't exist")
	}

	// we have reached the end of the path
	if len(files) == 1 {
		return cf, nil
	}

	// recursively keep changing
	return cf.changeDir(strings.Join(rest, "/"))
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

// ChangeDir changes to a directory
func (f *Fs) ChangeDir(path string) error {
	cf, err := f.currentDir.changeDir(path)

	if err != nil {
		return err
	}

	f.currentDir = cf

	return nil
}

// CreateDir creates a new directory in the current directory
func (f *Fs) CreateDir(name string) error {
	err := f.currentDir.createDir(name)

	if err != nil {
		return err
	}

	return nil
}

// New creates a new fileSystem
func New() Fs {
	root := &file{
		name:     "/",
		path:     "/",
		isDir:    true,
		children: make(children),
		parent:   nil,
	}

	return Fs{
		root:       root,
		currentDir: root,
	}
}

// PrintWorkingDirectory returns the current directory's path
func (f *Fs) PrintWorkingDirectory() string {
	return f.currentDir.path
}
