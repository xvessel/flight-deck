package cluster

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCluster(t *testing.T) {
	m := make(map[string]string)
	config := "testconfig"
	env := "prod"
	m[env] = config
	mgr := NewClusterMgr(m)

	ret, err := mgr.GetKubeConfig(env)

	assert.Equal(t, config, ret)
	assert.True(t, err)
}
