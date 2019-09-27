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
	Components() []string
	Spec(compName string) (Spec, error)
	Excute(compName string, cmd string, kubeConfig string, input map[string]string, namespace string, id string) (err error, output map[string]string)
}

var CMD_CREATE string = "CREATE"
var CMD_READY string = "READY"
var CMD_UPDATE_CHECK string = "UPDATE_CHECK"
var CMD_UPDATE string = "UPDATE"
var CMD_DELETE string = "DELETE"

type ComponentMgr struct {
	Dir     string
	compMap map[string]Component
}

func NewComponentMgr(dir string) *ComponentMgr {
	ret := &ComponentMgr{Dir: dir,
		compMap: make(map[string]Component)}

	childs, _ := ioutil.ReadDir(dir)
	for _, c := range childs {
		if c.IsDir() {
			ret.compMap[c.Name()] = Component{Dir: dir + "/" + c.Name(), Name: c.Name()}
		}
	}
	return ret
}

//func (m *ComponentMgr) Component(name string) (Component, error) {
//	input, err := m.Input(name)
//	if err != nil {
//		return Component{}, err
//	}
//	output, _ := m.Output(name)
//	return Component{Name: name, InputSpec: input, OutputSpec: output}, nil
//}

func (m *ComponentMgr) Components() []string {
	ret := make([]string, 0)
	for i, _ := range m.compMap {
		ret = append(ret, i)
	}
	return ret
}

func (m *ComponentMgr) Spec(component string) (Spec, error) {
	if v, ok := m.compMap[component]; ok {
		return v.spec, nil
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
