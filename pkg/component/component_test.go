package component

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractFile(t *testing.T) {
	right := make(map[string][2]string)
	right["KEY1"] = [2]string{"VALUE1", "comment1"}
	ret, err := extractFile("./test/mysql/INPUT")
	assert.Nil(t, err)
	assert.Equal(t, right, ret)
}
