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
			fs.CreateDir(strings.Join(args, " "))
		case "cd":
			fs.ChangeDir(strings.Join(args, " "))
		case "touch":
			fs.CreateFile(args[0], []byte(strings.Join(args[1:], " ")))
		case "cat":
			content, err := fs.ReadFile(strings.Join(args, " "))

			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(string(content))

		case "exit":
			return
		default:
			fmt.Println(cmd + " doesn't exist")
		}

		fmt.Printf("[%s@%s %s] ~ ", user.Name, hostname, fs.PrintWorkingDirectory())
	}
}
