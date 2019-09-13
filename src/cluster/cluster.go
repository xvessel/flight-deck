/**********************************************************
 * Author        : biliucheng
 * Email         : bilc_dev@163.com
 * Last modified : 2019-09-11 17:56
 * Filename      : cluster.go
 * Description   :
 * Modified By   :
 * *******************************************************/
package cluster

import (
	"sort"
)

type Manager interface {
	List() []string
	GetKubeConfig(environ string) (string, bool)
}

type ClusterMgr struct {
	KubeConfigs map[string]string
}

func NewClusterMgr(m map[string]string) *ClusterMgr {
	return &ClusterMgr{m}
}

func (m *ClusterMgr) List() []string {
	ret := make([]string, 0)
	for i, _ := range m.KubeConfigs {
		ret = append(ret, i)
	}
	sort.Strings(ret)
	return ret
}

func (m *ClusterMgr) GetKubeConfig(environ string) (string, bool) {
	v, ok := m.KubeConfigs[environ]
	return v, ok
}
