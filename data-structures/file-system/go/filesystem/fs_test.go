package filesystem

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	fs := New()

	root := &file{
		name:     "/",
		path:     "/",
		isDir:    true,
		children: make(Children),
		parent:   nil,
	}

	expected := Fs{
		root:       root,
		currentDir: root,
	}

	if !reflect.DeepEqual(fs, expected) {
		t.Errorf("filesystem.New() = %+v, want %+v", fs, expected)
	}
}

func TestPrintWorkingDirectory(t *testing.T) {
	fs := New()

	pwd := fs.PrintWorkingDirectory()

	if pwd != fs.currentDir.path {
		t.Errorf("filesystem.PrintWorkingDirectory() = %+v, want %+v", pwd, fs.currentDir.path)
	}
}

func TestCreateDir(t *testing.T) {
	fs := New()

	createDir := func(name string) func(t *testing.T) {
		return func(t *testing.T) {
			err := fs.CreateDir(name)

			if err != nil {
				t.Errorf("filesystem.CreateDir(%s) returned an error %s", name, err)
			}

			err = fs.ChangeDir(name)

			if err != nil {
				t.Errorf("filesystem.CreateDir(%s) did not create a dir %s; error: %s", name, name, err)
			}

			fs.ChangeDir("/")
		}
	}

	t.Run("usr", createDir("usr"))

	t.Run("etc/", createDir("etc/"))

	t.Run("/opt/", createDir("/opt/"))

	t.Run("/home", createDir("/home"))
}
