/**********************************************************
 * Author        : biliucheng
 * Email         : bilc_dev@163.com
 * Last modified : 2019-09-11 17:56
 * Filename      : cluster_test.go
 * Description   :
 * Modified By   :
 * *******************************************************/
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
