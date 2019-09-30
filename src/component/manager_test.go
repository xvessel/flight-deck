/**********************************************************
 * Author        : biliucheng
 * Email         : bilc_dev@163.com
 * Last modified : 2019-09-11 17:56
 * Filename      : manager_test.go
 * Description   :
 * Modified By   :
 * *******************************************************/
package component

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManger(t *testing.T) {

	mgr, err := NewComponentMgr("./test")
	assert.Nil(t, err)
	comps := mgr.Components()
	assert.Equal(t, 1, len(comps))
	assert.Equal(t, "mysql", comps[0].Name)
	env := make(map[string]string)
	env["A"] = "a"
	err, output := mgr.Excute("mysql", "CREATE", "kubetest", env, "testns", "testid")
	assert.Nil(t, err)
	assert.Equal(t, map[string]string{"A": "a"}, output)
}
