package main

import (
	"addissueannotation/tools"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type annotation struct {
	Description string `json:"description"`
	// Entry is the date and time the annotation was added.
	// e.g. : "20230227T100032Z"
	Entry string `json:"entry"`
}
type task struct {
	Description string       `json:"description"`
	Annotations []annotation `json:"annotations,omitempty"`
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

func main() {
	var task task
	decoder := json.NewDecoder(os.Stdin)
	err := decoder.Decode(&task)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	addAnnotation("ISS", "https://taiga.tech.orange/project/thommil-mahali-poc/issue/", &task)
	addAnnotation("MR", "https://gitlab.tech.orange/mahali/mahali-backend/-/merge_requests/", &task)
	addAnnotation("US", "https://taiga.tech.orange/project/thommil-mahali-poc/us/", &task)
	addAnnotation("TSK", "https://taiga.tech.orange/project/thommil-mahali-poc/task/", &task)

	encoder := json.NewEncoder(os.Stdout)
	err = encoder.Encode(task)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

// addAnnotation to add a new annotation to a task.
// param parPrefix : prefix of the annotation. E.g. "ISS" for issue, "MR" for merge request
// param parURL : URL to include in the annotation
func addAnnotation(parPrefix string, parURL string, parTask *task) {
	if strings.Contains(parTask.Description, parPrefix) {
		// Extract the number after parPrefix in the task.description
		number := tools.ExtractNumber(parTask.Description, parPrefix)
		if number != "" {
			// append a new annotation
			newAnnotation := annotation{
				Description: parURL + number,
				Entry:       time.Now().Format("20060102T150405Z"),
			}
			parTask.Annotations = append(parTask.Annotations, newAnnotation)
		}
	}
}
