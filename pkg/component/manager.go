package component

import (
	"fmt"
	"io/ioutil"
)

type Manager interface {
	Components() []string
	Input(component string) (ret map[string][2]string, err error)
	Output(component string) (ret map[string][2]string, err error)
	Excute(component string, cmd string, kubeConfig string, input map[string]string, namespace string, id string) (err error, output []string)
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

func (m *ComponentMgr) Components() []string {
	ret := make([]string, 0)
	for i, _ := range m.compMap {
		ret = append(ret, i)
	}
	return ret
}

func (m *ComponentMgr) Input(component string) (ret map[string][2]string, err error) {
	if v, ok := m.compMap[component]; ok {
		return v.Input()
	} else {
		return nil, fmt.Errorf("commponet %s not exsit", component)
	}
}

func (m *ComponentMgr) Output(component string) (ret map[string][2]string, err error) {
	if v, ok := m.compMap[component]; ok {
		return v.Output()
	} else {
		return nil, fmt.Errorf("commponet %s not exsit", component)
	}
}

func (m *ComponentMgr) Excute(component string, cmd string, kubeConfig string, input map[string]string, namespace string, id string) (err error, output []string) {
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
