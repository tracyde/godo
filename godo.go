package main

import (
	"fmt"
	//	"os"
	"bytes"
	"encoding/gob"
	"github.com/tracyde/godo/collection"
	"io/ioutil"
)

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
	c := collection.New("/tmp/test.gob")
	TestLoad(c)
	c.Print()

	fmt.Println("Saving collection to gobfile")
	b := new(bytes.Buffer)
	enc := gob.NewEncoder(b)
	enc.Encode(c)
	err := ioutil.WriteFile(c.Filename, b.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
	fmt.Printf("just saved gob with %v\n", c)

	n, err := ioutil.ReadFile(c.Filename)
	if err != nil {
		panic(err)
	}
	p := bytes.NewBuffer(n)
	dec := gob.NewDecoder(p)
	e := collection.New("/tmp/test.gob")
	err = dec.Decode(&e)
	if err != nil {
		panic(err)
	}
	fmt.Printf("just read gob from file and it's showing: %v\n", e)
	e.Print()
	// fmt.Printf("Type: %T  ::  Value: %v\n", c, c)
}
