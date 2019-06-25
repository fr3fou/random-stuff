package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/fr3fou/go-fs/filesystem"
)

func main() {
	fs := filesystem.New()

	scanner := bufio.NewScanner(os.Stdin)
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	hostname, err := os.Hostname()

	if err != nil {
		panic(err)
	}

	fmt.Printf("[%s@%s %s] ~ ", user.Name, hostname, fs.PrintWorkingDirectory())
	for scanner.Scan() {
		txt := strings.Split(scanner.Text(), " ")

		cmd := txt[0]
		args := txt[1:]

		switch cmd {
		case "pwd":
			fmt.Println(fs.PrintWorkingDirectory())
		case "ls":
			dirs, err := fs.ListDirectoryContents(strings.Join(args, " "))

			if err != nil {
				fmt.Println(err)
			}

			for key := range dirs {
				fmt.Print(key + " ")
			}
			fmt.Println()
		case "mkdir":
			if len(args) >= 1 {
				err := fs.CreateDir(strings.Join(args, " "))

				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("mkdir requires at least 1 arg")
			}
		case "cd":
			if len(args) >= 1 {
				err := fs.ChangeDir(strings.Join(args, " "))

				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("cd requires at least 1 arg")
			}
		case "touch":
			if len(args) >= 2 {
				err := fs.CreateFile(args[0], []byte(strings.Join(args[1:], " ")))

				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("touch needs at least 2 args")
			}
		case "rm":
			ln := len(args)
			if ln >= 1 {
				if args[0] == "-r" {
					if ln >= 2 {
						err := fs.DeleteDirectory(strings.Join(args[1:], " "))
						if err != nil {
							fmt.Println(err)
						}
					} else {
						fmt.Println("rm -r requires at least 2 args")
					}
				} else {
					err := fs.DeleteFile(strings.Join(args, " "))
					if err != nil {
						fmt.Println(err)
					}
				}
			} else {
				fmt.Println("rm requires at least 1 arg")
			}

		case "cat":
			if len(args) >= 1 {
				content, err := fs.ReadFile(strings.Join(args, " "))

				if err != nil {
					fmt.Println(err)
				}

				fmt.Println(string(content))
			} else {
				fmt.Println("cat requires at least 1 arg")
			}

		case "exit":
			return
		default:
			fmt.Println(cmd + " doesn't exist")
		}

		fmt.Printf("[%s@%s %s] ~ ", user.Name, hostname, fs.PrintWorkingDirectory())
	}
}
