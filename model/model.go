package model

// Annotation is a struct to represent an annotation
// a sub-struct of task.
type Annotation struct {
	Description string `json:"description"`
	// Entry is the date and time the annotation was added.
	// e.g. : "20230227T100032Z"
	Entry string `json:"entry"`
}

// Task is A taskwarrior task.
type Task struct {
	Description string       `json:"description"`
	Annotations []Annotation `json:"annotations,omitempty"`
	Status      string       `json:"status"`
	Entry       string       `json:"entry,omitempty"`
	UUID        string       `json:"uuid"`
	Modified    string       `json:"modified,omitempty"`
	Tags        []string     `json:"tags,omitempty"`
	Due         string       `json:"due,omitempty"`
	Urgency     float64      `json:"urgency,omitempty"`
	End         string       `json:"end,omitempty"`
	Project     string       `json:"project,omitempty"`
}

// Rule is a struct to represent a rule found in .taskrc.
type Rule struct {
	Prefix string `json:"prefix"`
	URL    string `json:"url"`
}
