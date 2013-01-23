package collection

type Task struct {
	description string
	priority    int
}

type Project struct {
	name        string
	description string
	priority    int
	tasks       []Task
}

type Collection struct {
	filename string
	projects []Project
}

func New(filename string) *Collection {
	return &Collection{filename}
}
