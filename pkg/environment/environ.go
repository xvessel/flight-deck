package envrionment

import (
	"sort"
)

type EnvrionmentMgr struct {
	KubeConfigs map[string]string
}

func NewEnvrionmentMgr(m map[string]string) *EnvrionmentMgr {
	return &EnvrionmentMgr{m}
}

func (m *EnvrionmentMgr) List() []string {
	ret := make([]string, 0)
	for i, _ := range m.KubeConfigs {
		ret = append(ret, i)
	}
	sort.Strings(ret)
	return ret
}

func (m *EnvrionmentMgr) GetKubeConfig(environ string) (string, bool) {
	v, err := m.KubeConfigs[environ]
	return v, err
}
