package component

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManger(t *testing.T) {

	mgr := NewComponentMgr("test")

	list := []string{"mysql"}
	assert.Equal(t, list, mgr.Components())
	err, output := mgr.Excute("mysql", "CREATE", "kubetest", nil, "testns", "testid")
	assert.Nil(t, err)
	assert.Equal(t, []string{"testns=testid"}, output)
}
