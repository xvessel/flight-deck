/**********************************************************
 * Author        : biliucheng
 * Email         : bilc_dev@163.com
 * Last modified : 2019-09-13 16:03
 * Filename      : main.go
 * Description   :
 * Modified By   :
 * *******************************************************/
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	yaml "gopkg.in/yaml.v2"

	"cluster"
	"component"
	"product"
)

var clusterMgr cluster.Manager
var componentMgr component.Manager
var productMgr *product.Manager

type Config struct {
	Listen       string            `yaml:"Listen"`
	ComponentDir string            `yaml:"ComponentDir"`
	KubeConfig   map[string]string `yaml:"KubeConfig"`
}

func main() {
	configFile := flag.String("conf", "", "config file ")
	flag.Parse()

	b, err := ioutil.ReadFile(*configFile)
	if err != nil {
		return
	}
	var conf Config
	err = yaml.Unmarshal(b, &conf)
	if err != nil {
		return
	}
	fmt.Println(string(b), conf)

	clusterMgr = cluster.NewClusterMgr(conf.KubeConfig)
	componentMgr, err = component.NewComponentMgr(conf.ComponentDir)
	if err != nil {
		fmt.Println("componentMgr err", err)
	}
	productMgr = product.NewManager(componentMgr, clusterMgr, "sqlite")

	r := mux.NewRouter()
	//product
	r.HandleFunc("/products", newProdHandler).Methods(http.MethodPost)
	r.HandleFunc("/products", listProdHandler).Methods(http.MethodGet)
	//components
	r.HandleFunc("/components", listComponentHandler).Methods(http.MethodGet)
	r.HandleFunc("/components/{componentName}", detailComponentHandler).Methods(http.MethodGet)
	//design
	r.HandleFunc("/designs", newDesignHandler).Methods(http.MethodPost)
	r.HandleFunc("/products/{prodName}/designs", listDesignHandler).Methods(http.MethodGet)
	r.HandleFunc("/products/{prodName}/designs/{revision}", getDesignHandler).Methods(http.MethodGet)

	//product instance
	r.HandleFunc("/products/{prodName}/instances", newInstanceHandler).Methods(http.MethodPost)
	r.HandleFunc("/products/{prodName}/instances", listInstanceHandler).Methods(http.MethodGet)
	r.HandleFunc("/products/{prodName}/instances/{instanceName}", getInstanceHandler).Methods(http.MethodGet)
	r.HandleFunc("/products/{prodName}/instances", updateInstanceHandler).Methods(http.MethodPut)

	fmt.Println(http.ListenAndServe(conf.Listen, r))
}
