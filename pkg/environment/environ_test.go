package envrionment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnv(t *testing.T) {
	m := make(map[string]string)
	config := "testconfig"
	env := "prod"
	m[env] = config
	mgr := NewEnvrionmentMgr(m)

	ret, err := mgr.GetKubeConfig(env)

	assert.Equal(t, config, ret)
	assert.True(t, err)
}
