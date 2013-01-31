// Package collection provides the types to manage godo projects and tasks.
package collection

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

type Task struct {
	Description string
	Priority    int
}

type Tasks []*Task // Wrapper type to allow sorting

type Project struct {
	Name        string
	Description string
	Priority    int
	Tasks       Tasks
}

type Projects []*Project // Wrapper type to allow sorting

type Collection struct {
	Filename string
	Projects Projects
}

// Wrapper types to allow different sort orders for Projects type.
type ByName struct{ Projects }
type ByPriority struct{ Projects }

// Methods required by sort.Interface for ByName type, allows
// Projects to be sorted by Name.
func (s ByName) Less(i, j int) bool {
	return s.Projects[i].Name < s.Projects[j].Name
}

// Methods required by sort.Interface for ByPriority type, allows
// Projects to be sorted by Priority.
func (s ByPriority) Less(i, j int) bool {
	return s.Projects[i].Priority < s.Projects[j].Priority
}

// Methods required by sort.Interface for Projects type, need
// to cast them as either ByName or ByPriority.
func (p Projects) Len() int {
	return len(p)
}
func (p Projects) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Methods required by sort.Interface for Tasks type, default
// sort is by Priority.
func (t Tasks) Len() int {
	return len(t)
}
func (t Tasks) Less(i, j int) bool {
	return t[i].Priority < t[j].Priority
}
func (t Tasks) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

// Method returns a pointer to an empty Collection.
func New(f string) *Collection {
	return &Collection{Filename: f, Projects: make([]*Project, 0)}
}

// Method to allow easy sorting of an entire Collection.
func (c *Collection) Sort() {
	sort.Sort(ByPriority{c.Projects})
	for _, p := range c.Projects {
		sort.Sort(p.Tasks)
	}
}

func (c *Collection) AddProject(n, d string, p int) {
	c.Projects = append(c.Projects, &Project{n, d, p, make([]*Task, 0)})
}

func (c *Collection) AddTask(pn, d string, p int) error {
	sort.Sort(ByName{c.Projects})
	i := sort.Search(len(c.Projects), func(i int) bool { return c.Projects[i].Name >= pn })
	//fmt.Printf("AddTask-> sort.Search found index: %d\n", i)
	if i < len(c.Projects) && c.Projects[i].Name == pn {
		//fmt.Printf("AddTask-> Adding new Task to project: %s\n", c.Projects[i].Name)
		c.Projects[i].Tasks = append(c.Projects[i].Tasks, &Task{d, p})
	} else {
		return errors.New("project name does not exist")
	}
	return nil
}

func (c *Collection) Print() {
	c.Sort()
	for i, p := range c.Projects {
		fmt.Printf("%d. %s - %s [%d]\n", i, p.Name, p.Description, p.Priority)
		for i, t := range p.Tasks {
			fmt.Printf("\t%d. %s [%d]\n", i, t.Description, t.Priority)
		}
		fmt.Println()
	}
}

func (c *Collection) Save() (err error) {
	b := new(bytes.Buffer)
	enc := gob.NewEncoder(b)
	enc.Encode(c.Projects)
	err = ioutil.WriteFile(c.Filename, b.Bytes(), 0600)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (c *Collection) Read() (err error) {
	b, err := ioutil.ReadFile(c.Filename)
	if err != nil {
		log.Fatal(err)
		return err
	}

	buf := bytes.NewBuffer(b)
	dec := gob.NewDecoder(buf)
	var p Projects
	fmt.Printf("project :: %T\n", p)
	err = dec.Decode(&p)
	if err != nil {
		log.Fatal(err)
		return err
	}
	c.Projects = p
	return err
}
