package tools

import (
	"testing"

	"gotest.tools/assert"
)

func TestExtractNumber(t *testing.T) {
	assert.Equal(t, "123", ExtractNumber("bla bla toto ISS123", "ISS"))
	assert.Equal(t, "222", ExtractNumber("Merger JAMES MR222", "MR"))
	assert.Equal(t, "222", ExtractNumber("Merger MR222 James", "MR"))
	assert.Equal(t, "", ExtractNumber("Merger JAMES", "MR"))
}
