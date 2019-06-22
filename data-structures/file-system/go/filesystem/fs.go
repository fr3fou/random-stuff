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
	walk(path string) (*file, error)
}

// file implementations

// walk takes a path and returns the file at that path
func (f *file) walk(path string) (*file, error) {
	// walk up the tree to the root
	if strings.HasPrefix(path, "/") && f.parent != nil {
		return f.parent.walk(path)
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
			return f.walk(strings.Join(rest, "/"))
		}

		// if there are no files left, we just return the parent
		if len(files) == 1 {
			return f.parent, nil
		}

		// we have to keep going if there are files left
		return f.parent.walk(strings.Join(rest, "/"))
	}

	cf, ok := f.children[file]

	if !ok {
		return nil, errors.New("fs: can't walk to a file that doesn't exist")
	}

	// we have reached the end of the path
	if len(files) == 1 {
		return cf, nil
	}

	// recursively keep walking
	return cf.walk(strings.Join(rest, "/"))
}

// fs implementations

// ChangeDir changes to a directory
func (f *Fs) ChangeDir(path string) error {
	cf, err := f.currentDir.walk(path)

	if err != nil {
		return err
	}

	if !cf.isDir {
		return errors.New("fs: can't cd to a file")
	}

	f.currentDir = cf

	return nil
}

// CreateDir creates a new directory in the current directory
func (f *Fs) CreateDir(path string) error {
	// get the path up until the last element
	lastItem := strings.LastIndex(path, "/")

	var (
		name string
		cf   *file
		err  error
	)

	// if we are trying to make a nested directory, we should check if all the directorie preceding it actually exist
	if lastItem > -1 {
		// walk up until the last item
		cf, err = f.currentDir.walk(path[:lastItem])
		// the name is going to be our last item
		name = path[lastItem+1:]
	} else {
		cf = f.currentDir
		err = nil
		name = path
	}

	if err != nil {
		return err
	}

	if _, ok := cf.children[name]; ok {
		return errors.New("fs: can't create a directory that already exists")
	}

	// no parent means path is at root
	if cf.parent == nil {
		path = cf.path + name
	} else {
		path = cf.path + "/" + name
	}

	cf.children[name] = &file{
		isDir:    true,
		parent:   cf,
		name:     strings.Trim(name, "/"),
		children: make(children),
		path:     path,
	}

	return nil
}

// CreateFile creates a new file in the current directory
func (f *Fs) CreateFile(path string, content []byte) error {
	// get the path up until the last element
	lastItem := strings.LastIndex(path, "/")

	var (
		name string
		cf   *file
		err  error
	)

	// if we are trying to make a nested directory, we should check if all the directorie preceding it actually exist
	if lastItem > -1 {
		// walk up until the last item
		cf, err = f.currentDir.walk(path[:lastItem])
		// the name is going to be our last item
		name = path[lastItem+1:]
	} else {
		cf = f.currentDir
		err = nil
		name = path
	}

	if err != nil {
		return err
	}

	if _, ok := cf.children[name]; ok {
		return errors.New("fs: can't create a file that already exists")
	}

	// no parent means path is at root
	if cf.parent == nil {
		path = cf.path + name
	} else {
		path = cf.path + "/" + name
	}

	cf.children[name] = &file{
		isDir:   false,
		parent:  cf,
		name:    strings.Trim(name, "/"),
		content: content,
		path:    path,
	}

	return nil
}

// ReadFile returns the content of a file at a given path
func (f *Fs) ReadFile(path string) ([]byte, error) {
	cf, err := f.currentDir.walk(path)

	if err != nil {
		return nil, err
	}

	if cf.isDir {
		return nil, errors.New("fs: can't read content of a file")
	}

	return cf.content, nil
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

func parsePath(str string) []string {
	s := strings.FieldsFunc(str, func(c rune) bool {
		return c == '/'
	})

	return s
}
