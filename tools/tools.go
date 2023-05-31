package tools

import (
	"bufio"
	"encoding/json"
	"hookaddannotation/model"
	"os"
	"strings"
)

// ExtractNumber to extract a number from a string after a marker.
// returns "" if no number is found.
func ExtractNumber(parString string, parMarker string) string {
	if strings.Contains(parString, parMarker) {
		// Extract the number after parPrefix in the task.description
		startIndex := strings.Index(parString, parMarker) + len(parMarker)
		endIndex := startIndex
		for endIndex < len(parString) && parString[endIndex] >= '0' && parString[endIndex] <= '9' {
			endIndex++
		}
		if startIndex != endIndex {
			return parString[startIndex:endIndex]
		}
	}
	return ""
}

// ContainsAnnotationDescr to check if a slice of annotations contains an annotation with the same description.
func ContainsAnnotationDescr(parAnnotations []model.Annotation, parAnnotation model.Annotation) bool {
	for _, annotation := range parAnnotations {
		if annotation.Description == parAnnotation.Description {
			return true
		}
	}
	return false
}

// ContainsAnnotationPrefix to check if a slice of annotations contains an annotation with the provided prefix.
func ContainsAnnotationPrefix(parAnnotations []model.Annotation, parPrefix string) bool {
	for _, annotation := range parAnnotations {
		if strings.Contains(annotation.Description, parPrefix) {
			return true
		}
	}
	return false
}

// LoadConfig to load and parse a taskwarrior config file
// config files contains lines with the following format:
// titi.tata.toto = "tutu"
func LoadConfig(parConfigPath string) ([]model.Rule, error) {
	var rules []model.Rule
	// load file at path as text

	f, err := os.Open(parConfigPath)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		curLine := scanner.Text()
		if strings.HasPrefix(curLine, "hookaddannotation.rule") {
			// get the value part, after the equal sign till the end of the line
			startIndex := strings.Index(curLine, "=") + 1
			endIndex := len(curLine)
			ruleString := curLine[startIndex:endIndex]
			// parse it as json to get a rule
			var rule model.Rule
			err := json.Unmarshal([]byte(ruleString), &rule)
			if err != nil {
				return nil, err
			}
			rules = append(rules, rule)
		}
	}
	return rules, nil
}
