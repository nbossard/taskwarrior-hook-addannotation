package main

import (
	"addissueannotation/model"
	"addissueannotation/tools"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var task model.Task
	decoder := json.NewDecoder(os.Stdin)
	err := decoder.Decode(&task)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	nbrAddedAnnotations := 0
	// These are the rules to add annotations automatically to a task
	nbrAddedAnnotations += addAnnotation("ISS", "https://taiga.tech.orange/project/thommil-mahali-poc/issue/", &task)
	nbrAddedAnnotations += addAnnotation("MR", "https://gitlab.tech.orange/mahali/mahali-backend/-/merge_requests/", &task)
	nbrAddedAnnotations += addAnnotation("US", "https://taiga.tech.orange/project/thommil-mahali-poc/us/", &task)
	nbrAddedAnnotations += addAnnotation("TSK", "https://taiga.tech.orange/project/thommil-mahali-poc/task/", &task)

	encoder := json.NewEncoder(os.Stdout)
	err = encoder.Encode(task)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	} else {
		if nbrAddedAnnotations > 0 {
			fmt.Fprintln(os.Stderr, "Successfully added annotations to task")
		}
		os.Exit(0)
	}
}

// addAnnotation to add a new annotation to a task.
//
// param parPrefix : prefix of the annotation. E.g. "ISS" for issue, "MR" for merge request.
// param parURL : URL to include in the annotation.
func addAnnotation(parPrefix string, parURL string, parTask *model.Task) int {
	if !strings.Contains(parTask.Description, parPrefix) {
		return 0
	}
	// Extract the number after parPrefix in the task.description
	number := tools.ExtractNumber(parTask.Description, parPrefix)
	if number == "" {
		return 0
	}
	// create a new annotation
	newAnnotation := model.Annotation{
		Description: parURL + number,
		Entry:       time.Now().Format("20060102T150405Z"),
	}
	// append it to the task annotations if it does not already exist
	if tools.ContainsAnnotationDescr(parTask.Annotations, newAnnotation) {
		return 0
	}
	parTask.Annotations = append(parTask.Annotations, newAnnotation)
	return 1
}
