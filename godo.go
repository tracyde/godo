package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tracyde/godo/collection"
)

var noAct = errors.New("no action")

var (
	file = flag.String("file", defaultFile(".godo", "GODO"), "file to store projects and tasks")
)

func defaultFile(name, env string) string {
	if f := os.Getenv(env); f != "" {
		return f
	}
	return filepath.Join(os.Getenv("HOME"), name)
}

const usage = `Usage:
	godo
		Show top project and tasks
	godo ls
		Show all tasks
Flags:
`

func TestLoad(c *collection.Collection) {
	c.AddProject("Test0", "This is a test Project0", 2)
	c.AddProject("Test1", "This is a test Project1", 3)
	c.AddProject("Test2", "This is a test Project2", 4)
	c.AddProject("Test3", "This is a test Project3", 1)

	c.AddTask("Test0", "This task should be on Test0 project", 2)
	c.AddTask("Test0", "Task for Test0 project", 3)
	c.AddTask("Test2", "Task for Test2 project", 3)
	c.AddTask("Test3", "Test3 task ready for action", 1)
	c.AddTask("Test1", "Test1 task ready for action", 1)
	c.AddTask("Test3", "Test3 I wonder why this project is so popular", 2)
	/*
		err := c.AddTask("Test4", "complete adding func for tasks", 6)
		if err != nil {
			fmt.Fprintln(os.Stderr, "logging:", err)
		}
	*/
}

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
		flag.PrintDefaults()
	}
	flag.Parse()

	c := collection.New(*file)
	a, n := flag.Arg(0), len(flag.Args())

	err := noAct
	switch {
	case a == "ls" && n == 1:
		// list all tasks
		err = c.Read()
		c.Print()
	case n == 0:
		// no arguments; no action
		fmt.Println("No action taken")
	}
	if err == noAct {
		// no action taken
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	/*
		TestLoad(c)
		c.Print()

		fmt.Println("Saving collection to gobfile")
		err := c.Save()
		if err != nil {
			panic(err)
		}
		fmt.Printf("just saved gob with %v\n", c)

		c2 := collection.New("/tmp/test.gob")
		err = c2.Read()
		if err != nil {
			panic(err)
		}
		fmt.Printf("just read gob from file and it's showing: %v\n", c2)
		c2.Print()
	*/

	// fmt.Printf("Type: %T  ::  Value: %v\n", c, c)
}
