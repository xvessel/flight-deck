package main

import (
	"component"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"os"
)

// TODO: component add command
func NewComponentCommand() *cobra.Command {
	cmd := &cobra.Command{Use: "component [sub command]"}
	cmd.AddCommand(&cobra.Command{Use: "init [name]", Run: ComponentInit, Args:cobra.ExactArgs(1)})
	cmd.AddCommand(&cobra.Command{Use: "list", Run: ComponentList})
	cmd.AddCommand(&cobra.Command{Use: "get [name]", Run: ComponentGet})
	return cmd
}

// component template
const CREATE = `#!/bin/bash
# create instance
namespace=$1
instance-id=$2

## do some thing to create instance in namespace
## example
#sed xxx xxx xxx
#kubectl apply -f xxx.yml -n $namespace

## success example:
#echo SERVICE_HOST="127.0.0.1"
#echo SERVICE_PORT="8080"
#exit 0

## failed example:
#echo "some error message" > &2 
#exit 1
`
const DELETE = `#!/bin/bash
# delete instance
namespace=$1
instance-id=$2

## do some thing to delete instance in namespace
## example
#kubectl delete xxx -n $namespace
#exit 0
#exit 1
`
const PRE = `#!/bin/bash
# component init script
exit 0
`

const READY = `#!/bin/bash
# check instance status 
namespace=$1
instance-id=$2
## if instance is ready
#    exit 0
## else instance is ready
#    exit 1
`
const UPDATE = `#!/bin/bash
# update instance 
namespace=$1
instance-id=$2
# exit 0 # success
# exit 1 # failed
`

const UPDATECHECK = `#!/bin/bash
# update check instance 
namespace=$1
instance-id=$2
# exit 0 # success
# exit 1 # failed
`

const SPEC = `
description: hello
inputs:
- name: CPU
default_value: 1
can_update: true
description: "cpu核数"
- name: MEM
default_value: 2
can_update: true
description: "内存大小G"
outputs:
- name: HOST
description: "访问host"
- name: PORT
description: "访问port"
`

func createTemplateFile(filename, content string){
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	ExitIfError(err)
	_, err = f.WriteString(content)
	ExitIfError(err)
	f.Close()
}

func ComponentInit(cmd *cobra.Command, args []string) {
	name := args[0]
	// TODO: check -- the dir must be clean
	// init dir name
	err := os.MkdirAll(name, 0755)
	ExitIfError(err)

	// template file
	createTemplateFile(name+"/CREATE", CREATE)
	createTemplateFile(name+"/DELETE", DELETE)
	createTemplateFile(name+"/PRE", PRE)
	createTemplateFile(name+"/READY", READY)
	createTemplateFile(name+"/UPDATE", UPDATE)
	createTemplateFile(name+"/UPDATE_CHECK", UPDATECHECK)
	createTemplateFile(name+"/SPEC.yml", SPEC)
}

func ComponentGet(cmd *cobra.Command, args []string) {
	name := args[0]
	req, err := http.NewRequest(http.MethodGet, DeckAddr + "/components/" + name, nil)
	ExitIfError(err)
	resp, err := http.DefaultClient.Do(req)
	ExitIfError(err)
	if resp.StatusCode != http.StatusOK{
		ExitIfError(fmt.Errorf("http responde code is not 200. %s", resp.Status))
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	ExitIfError(err)
	c := &component.Spec{}
	err = json.Unmarshal(data, c)
	ExitIfError(err)
	pretty, err := json.MarshalIndent(c, "    ", "    ")
	ExitIfError(err)
	fmt.Println(string(pretty))

}

func ComponentList(cmd *cobra.Command, args []string) {
	req, err := http.NewRequest(http.MethodGet, DeckAddr + "/components", nil)
	ExitIfError(err)
	resp, err := http.DefaultClient.Do(req)
	ExitIfError(err)
	if resp.StatusCode != http.StatusOK{
		ExitIfError(fmt.Errorf("http responde code is not 200. %s", resp.Status))
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	ExitIfError(err)
	cs := make([]component.Spec,0)
	err = json.Unmarshal(data, &cs)
	ExitIfError(err)
	for _, c := range cs{
		fmt.Println(c.Name, c.Description)
	}
}

// TODO: add package and post component CMD
