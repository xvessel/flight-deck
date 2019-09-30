/**********************************************************
 * Author        : biliucheng
 * Email         : bilc_dev@163.com
 * Last modified : 2019-09-11 17:56
 * Filename      : manager.go
 * Description   :
 * Modified By   :
 * *******************************************************/
package component

import (
	"fmt"
	"io/ioutil"
)

type Manager interface {
	Components() []Spec
	Spec(compName string) (Spec, error)
	Excute(compName string, cmd string, kubeConfig string, input map[string]string, namespace string, id string) (err error, output map[string]string)
}

type ComponentMgr struct {
	Dir     string
	compMap map[string]ComponentIf
}

func NewComponentMgr(dir string) (*ComponentMgr, error) {
	ret := &ComponentMgr{Dir: dir,
		compMap: make(map[string]ComponentIf)}

	childs, _ := ioutil.ReadDir(dir)
	for _, ch := range childs {
		if ch.IsDir() {
			comp, err := NewComponent(dir, ch.Name())
			if err != nil {
				return ret, err
			}
			ret.compMap[ch.Name()] = &comp
		}
	}
	return ret, nil
}

func (m *ComponentMgr) Components() []Spec {
	ret := make([]Spec, 0)
	for _, i := range m.compMap {
		ret = append(ret, i.GetSpec())
	}
	return ret
}

func (m *ComponentMgr) Spec(component string) (Spec, error) {
	if v, ok := m.compMap[component]; ok {
		return v.GetSpec(), nil
	} else {
		return Spec{}, fmt.Errorf("component %v not exist", component)
	}
}

func (m *ComponentMgr) Excute(component string, cmd string, kubeConfig string, input map[string]string, namespace string, id string) (err error, output map[string]string) {
	envs := make([]string, 0)
	for i, j := range input {
		envs = append(envs, i+"="+j)
	}
	envs = append(envs, "KUBECONFIG="+kubeConfig)

	if v, ok := m.compMap[component]; ok {
		return v.Run(cmd, envs, namespace, id)
	} else {
		return fmt.Errorf("commponet %s not exsit", component), nil
	}
}
