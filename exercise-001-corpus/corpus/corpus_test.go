package corpus

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCorpus(t *testing.T) {
	result := Analyze("Are you serious? I knew you were.")
	assert.Equal(t, len(result), 5)
	assert.Equal(t, result[0].Word, "you")
	assert.Equal(t, result[0].Count, 2)
}
