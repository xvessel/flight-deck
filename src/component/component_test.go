/**********************************************************
 * Author        : biliucheng
 * Email         : bilc_dev@163.com
 * Last modified : 2019-09-11 17:56
 * Filename      : component_test.go
 * Description   :
 * Modified By   :
 * *******************************************************/
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
