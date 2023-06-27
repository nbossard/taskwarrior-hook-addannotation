package tools

import (
	"testing"

	"gotest.tools/assert"
)

func TestExtractNumber(t *testing.T) {
	assert.Equal(t, "123", ExtractNumber("bla bli toto ISS123", "ISS"))
	assert.Equal(t, "222", ExtractNumber("Merger JAMES MR222", "MR"))
	assert.Equal(t, "222", ExtractNumber("Merger MR222 James", "MR"))
	// multiple MR in the sentence only latest contains number
	assert.Equal(t, "531", ExtractNumber("t add regarder MR James feat: add announcement to order MR531", "MR"))
	assert.Equal(t, "", ExtractNumber("Merger JAMES", "MR"))
}
