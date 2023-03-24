package tools

import (
	"hookaddannotation/model"
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
