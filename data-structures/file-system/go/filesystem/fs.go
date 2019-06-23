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

// Children is the underlying map which for items
type Children map[string]*file

type file struct {
	name     string
	path     string
	children Children
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

	// if our target is the root and we have walked up to it, or if the path is empty just return it
	if path == "/" || path == "" {
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

// walkToParent walks up to the given path and returns the parent of the last file.
func (f *file) walkToParent(path string, fn func(cf *file, name string) (interface{}, error)) (interface{}, error) {
	// remove the trailing slash at the end
	path = strings.TrimRight(path, "/")

	// get the path up until the last element
	lastItem := strings.LastIndex(path, "/")

	var (
		name string
		cf   *file
		err  error
	)

	// if the only slash is at the beginning, it's an absolute file
	if lastItem == 0 {
		// for cases like /usr
		// walk up until the last item
		cf, err = f.walk(path)
		// the name is going to be our item but without the /
		name = path[lastItem+1:]
	} else if lastItem > -1 {
		// for cases like /usr/share/local
		// if we are trying to make a nested file, we should check if all the directories preceding it actually exist
		// walk up until the last item
		cf, err = f.walk(path[:lastItem])
		// the name is going to be our last item
		name = path[lastItem+1:]
	} else {
		// for cases like usr
		// if it's not nested, we can assume it's in the current directory
		cf = f
		err = nil
		name = path
	}

	if err != nil {
		return nil, err
	}

	return fn(cf, name)
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
	_, err := f.currentDir.walkToParent(path, func(cf *file, name string) (interface{}, error) {
		if _, ok := cf.children[name]; ok {
			return nil, errors.New("fs: can't create a directory that already exists")
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
			children: make(Children),
			path:     path,
		}

		return nil, nil
	})

	return err
}

// ListDirectoryContents lists all of the items inside a directory
func (f *Fs) ListDirectoryContents(path string) (Children, error) {
	cf, err := f.currentDir.walk(path)

	if err != nil {
		return nil, err
	}

	if cf.children == nil {
		return nil, errors.New("fs: can't list items inside a directory that doesn't exist")
	}

	if !cf.isDir {
		return nil, errors.New("fs: can't list items inside a file")
	}

	return cf.children, nil
}

// DeleteDirectory deletes the directory at a given path
func (f *Fs) DeleteDirectory(path string) error {
	_, err := f.currentDir.walkToParent(path, func(cf *file, name string) (interface{}, error) {
		if _, ok := cf.children[name]; !ok {
			return nil, errors.New("fs: can't delete a directory that doesn't exist")
		}

		delete(cf.children, name)

		return nil, nil

	})

	return err
}

// CreateFile creates a new file in the current directory
func (f *Fs) CreateFile(path string, content []byte) error {
	_, err := f.currentDir.walkToParent(path, func(cf *file, name string) (interface{}, error) {
		if _, ok := cf.children[name]; ok {
			return nil, errors.New("fs: can't create a file that already exists")
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

		return nil, nil
	})

	return err
}

// DeleteFile deletes the file at a given path
func (f *Fs) DeleteFile(path string) error {
	_, err := f.currentDir.walkToParent(path, func(cf *file, name string) (interface{}, error) {

		if _, ok := cf.children[name]; !ok {
			return nil, errors.New("fs: can't delete a file that doesn't exist")
		}

		delete(cf.children, name)

		return nil, nil

	})

	return err
}

// ReadFile returns the content of a file at a given path
func (f *Fs) ReadFile(path string) ([]byte, error) {
	cf, err := f.currentDir.walk(path)

	if err != nil {
		return nil, err
	}

	if cf.isDir {
		return nil, errors.New("fs: can't read content of a directory")
	}

	return cf.content, nil
}

// EditFile edits a file in the current directory
func (f *Fs) EditFile(path string, content []byte) error {
	_, err := f.currentDir.walkToParent(path, func(cf *file, name string) (interface{}, error) {
		if _, ok := cf.children[name]; !ok {
			return nil, errors.New("fs: can't edit a file that doesn't exists")
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

		return nil, nil
	})

	return err
}

// PrintWorkingDirectory returns the current directory's path
func (f *Fs) PrintWorkingDirectory() string {
	return f.currentDir.path
}

// New creates a new fileSystem
func New() Fs {
	root := &file{
		name:     "/",
		path:     "/",
		isDir:    true,
		children: make(Children),
		parent:   nil,
	}

	return Fs{
		root:       root,
		currentDir: root,
	}
}

