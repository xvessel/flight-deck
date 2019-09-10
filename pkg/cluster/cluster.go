package cluster

import (
	"sort"
)

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
