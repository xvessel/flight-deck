/**********************************************************
 * Author        : biliucheng
 * Email         : bilc_dev@163.com
 * Last modified : 2019-09-13 16:04
 * Filename      : handler.go
 * Description   :
 * Modified By   :
 * *******************************************************/
package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"product"
)

func listComponentHandler(w http.ResponseWriter, r *http.Request) {
	ret := componentMgr.Components()
	b, _ := json.Marshal(ret)
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func detailComponentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	c, err := componentMgr.Component(vars["componentName"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		b, _ := json.Marshal(c)
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func newProdHandler(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	var prod product.Product
	err := json.Unmarshal(b, &prod)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	err = productMgr.NewProduct(prod.Name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func listProdHandler(w http.ResponseWriter, r *http.Request) {
	ret, err := productMgr.Products()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	b, _ := json.Marshal(ret)
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func newDesignHandler(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	var design product.Design
	err := json.Unmarshal(b, &design)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	err = productMgr.NewDesign(&design)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}
func listDesignHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ret, err := productMgr.GetDesigns(vars["prodName"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	b, _ := json.Marshal(ret)
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func getDesignHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ret, err := productMgr.GetDesign(vars["prodName"], vars["revision"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	b, _ := json.Marshal(ret)
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func listInstanceHandler(w http.ResponseWriter, r *http.Request) {
}

func getInstanceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ret, err := productMgr.GetProductInst1(vars["prodName"], vars["instanceName"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	b, _ := json.Marshal(ret)
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func newInstanceHandler(w http.ResponseWriter, r *http.Request) {
	var param product.ProductInst

	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, &param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	err = productMgr.NewProductInst(&param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func updateInstanceHandler(w http.ResponseWriter, r *http.Request) {
}
