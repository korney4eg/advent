package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var input = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

type command struct {
	command  string
	argument string
	outputs  []string
}

type node struct {
	parent   *node
	name     string
	nodeType string
	children []*node
	size     int
}

var rootFolder = &node{name: "/", nodeType: "dir"}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)

	}
	commands := splitToCommands(string(input))
	cwd := rootFolder
	for _, cmd := range commands {
		cwd = runCommand(cwd, cmd)
	}
	_ = rootFolder.setSize()
	rootFolder.draw("")
	fmt.Println(rootFolder.sumFoldersSmallerThen(100000))
}

func runCommand(cwd *node, cmd command) *node {
	fmt.Println("command: ", cmd.command)
	if cmd.command == "ls" {
		cwd.children = runLsCommand(cwd, cmd.outputs)
		return cwd
	}
	if cmd.command == "cd" {
		return cwd.find(cmd.argument)
	}
	return nil

}
func (n *node) setSize() int {
	if n.nodeType == "dir" {
		totalSize := 0
		for _, child := range n.children {
			totalSize += child.setSize()
		}
		n.size = totalSize
		return totalSize
	}
	return n.size

}

func (n *node) draw(indent string) {
	fmt.Printf("%s- %s (%s, size=%d)\n", indent, n.name, n.nodeType, n.size)
	if n.nodeType == "dir" {
		for _, child := range n.children {
			child.draw(indent + "  ")
		}
	}
}

func (n *node) sumFoldersSmallerThen(limit int) int {
	sum := 0
	if n.nodeType == "dir" && n.size <= limit {
		sum += n.size
	}
	if n.nodeType == "dir" {
		for _, child := range n.children {
			sum += child.sumFoldersSmallerThen(limit)
		}
	}
	return sum
}

func runLsCommand(cwd *node, outputs []string) (nodes []*node) {
	for _, output := range outputs {
		if output == "" {
			continue
		}
		fmt.Println(output)
		typeOrSize := strings.Split(output, " ")[0]
		nodeName := strings.Split(output, " ")[1]
		newNode := &node{name: nodeName, parent: cwd}
		if typeOrSize == "dir" {
			newNode.nodeType = "dir"
		} else {
			newNode.nodeType = "file"
			nodeSize, _ := strconv.Atoi(typeOrSize)
			newNode.size = nodeSize
		}
		fmt.Printf("new node: '%+v'\n", newNode)
		nodes = append(nodes, newNode)
	}

	return nodes
}

func (n *node) find(name string) *node {
	if name == "/" {
		return rootFolder
	}
	if name == ".." {
		return n.parent
	}
	for _, child := range n.children {
		if child.name == name {
			return child
		}
	}
	return nil
}

func splitToCommands(input string) []command {
	commands := []command{}
	for _, fullCommand := range strings.Split(input, "$ ") {
		if fullCommand == "" {
			continue
		}
		cmd := command{}
		commandLine := strings.Split(fullCommand, "\n")[0]
		outputs := strings.Split(fullCommand, "\n")[1:]
		commandLineArray := strings.Split(commandLine, " ")
		cmd.command = commandLineArray[0]
		if len(commandLineArray) > 1 {
			cmd.argument = commandLineArray[1]
		}
		cmd.outputs = append(cmd.outputs, outputs...)

		commands = append(commands, cmd)

	}

	return commands
}
