package main

import (
	"encoding/json"
	"fmt"
	"hookaddannotation/model"
	"hookaddannotation/tools"
	"os"
	"strings"
	"time"
)

const (
	appName       = "🪄HookAddAnnotation"
	appDispPrefix = appName + ": "
)

func main() {
	// parse args
	// check version of hook API, one of args should be "api:" followed by number 2
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "api:") {
			apiVersion := strings.TrimPrefix(arg, "api:")
			if apiVersion != "2" {
				fmt.Fprintln(os.Stderr, "Error: unsupported API version:", apiVersion)
				os.Exit(1)
			}
		}
	}

	// get arg path of config file
	var configPath string
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "rc:") {
			configPath = strings.TrimPrefix(arg, "rc:")
		}
	}

	// parse input
	var task model.Task
	decoder := json.NewDecoder(os.Stdin)
	err := decoder.Decode(&task)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	// find rules in config file
	// load and parse config file
	rules, err := tools.LoadConfig(configPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error loading config file:", err)
		os.Exit(1)
	}

	resAddAnnotations := ""

	// execute rules from config file
	for _, rule := range rules {
		resAddAnnotations += addAnnotation(rule.Prefix, rule.URL, &task)
	}

	// write output
	encoder := json.NewEncoder(os.Stdout)
	err = encoder.Encode(task)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	} else {
		if resAddAnnotations != "" {
			fmt.Fprintln(os.Stderr, resAddAnnotations)
		}
		os.Exit(0)
	}
}

// addAnnotation to add a new annotation to a task if prefix found and not already existing.
//
// param parPrefix : prefix of the annotation to be searched. E.g. "ISS" for issue, "MR" for merge request.
// If a number is found after this prefix in the task description, this number will be added at the end of the URL
// param parURL : URL to include in the annotation.
func addAnnotation(parPrefix string, parURL string, parTask *model.Task) string {
	res := ""

	// quick exit if prefix is not found in task description nor in task annotations
	if !strings.Contains(parTask.Description, parPrefix) &&
		!tools.ContainsAnnotationPrefix(parTask.Annotations, parPrefix) {
		return ""
	}
	res += appDispPrefix + "Found prefix \"" + parPrefix + "\"\n"

	// Extract the number after parPrefix in the task.description
	number := tools.ExtractNumber(parTask.Description, parPrefix)
	if number == "" {
		// fallback, search in annotations
		for _, annotation := range parTask.Annotations {
			number = tools.ExtractNumber(annotation.Description, parPrefix)
			if number != "" {
				break
			}
		}
	}
	if number == "" {
		return appDispPrefix + "❌ Did not found number, cancelling adding annotation\n"
	}

	// create a new annotation
	newAnnotation := model.Annotation{
		Description: parURL + number,
		Entry:       time.Now().Format("20060102T150405Z"),
	}

	// append it to the task annotations if it does not already exist
	if tools.ContainsAnnotationDescr(parTask.Annotations, newAnnotation) {
		return appDispPrefix + "❌ already contains annotation\n"
	}

	parTask.Annotations = append(parTask.Annotations, newAnnotation)
	res += appDispPrefix + "✅ Added annotation \"" + newAnnotation.Description + "\"\n"

	return res
}
