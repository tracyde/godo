// Package collection provides the types to manage godo projects and tasks.
package collection

import (
	"errors"
	"fmt"
	"sort"
)

type Task struct {
	description string
	priority    int
}

type Tasks []*Task // Wrapper type to allow sorting

type Project struct {
	name        string
	description string
	priority    int
	tasks       Tasks
}

type Projects []*Project // Wrapper type to allow sorting

type Collection struct {
	Filename string
	projects Projects
}

// Wrapper types to allow different sort orders for Projects type.
type ByName struct{ Projects }
type ByPriority struct{ Projects }

// Methods required by sort.Interface for ByName type, allows
// Projects to be sorted by name.
func (s ByName) Less(i, j int) bool {
	return s.Projects[i].name < s.Projects[j].name
}

// Methods required by sort.Interface for ByPriority type, allows
// Projects to be sorted by priority.
func (s ByPriority) Less(i, j int) bool {
	return s.Projects[i].priority < s.Projects[j].priority
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
// sort is by priority.
func (t Tasks) Len() int {
	return len(t)
}
func (t Tasks) Less(i, j int) bool {
	return t[i].priority < t[j].priority
}
func (t Tasks) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

// Method returns a pointer to an empty Collection.
func New(f string) *Collection {
	return &Collection{Filename: f, projects: make([]*Project, 0)}
}

// Method to allow easy sorting of an entire Collection.
func (c *Collection) Sort() {
	sort.Sort(ByPriority{c.projects})
	for _, p := range c.projects {
		sort.Sort(p.tasks)
	}
}

func (c *Collection) AddProject(n, d string, p int) {
	c.projects = append(c.projects, &Project{n, d, p, make([]*Task, 0)})
}

func (c *Collection) AddTask(pn, d string, p int) error {
	sort.Sort(ByName{c.projects})
	i := sort.Search(len(c.projects), func(i int) bool { return c.projects[i].name >= pn })
	//fmt.Printf("AddTask-> sort.Search found index: %d\n", i)
	if i < len(c.projects) && c.projects[i].name == pn {
		//fmt.Printf("AddTask-> Adding new Task to project: %s\n", c.projects[i].name)
		c.projects[i].tasks = append(c.projects[i].tasks, &Task{d, p})
	} else {
		return errors.New("project name does not exist")
	}
	return nil
}

func (c *Collection) Print() {
	c.Sort()
	for i, p := range c.projects {
		fmt.Printf("%d. %s - %s [%d]\n", i, p.name, p.description, p.priority)
		for i, t := range p.tasks {
			fmt.Printf("\t%d. %s [%d]\n", i, t.description, t.priority)
		}
		fmt.Println()
	}
}
