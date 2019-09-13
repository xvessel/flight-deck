/**********************************************************
 * Author        : biliucheng
 * Email         : bilc_dev@163.com
 * Last modified : 2019-09-11 17:55
 * Filename      : manager_component.go
 * Description   :
 * Modified By   :
 * *******************************************************/
package product

import (
	"bytes"
	"encoding/json"
	"fmt"
	"text/template"
)

func (m *Manager) GetCompnentObj(prodInstId string, role string) (ComponentObj, error) {
	var model ComponentObjModel
	err := m.db.Where("product_inst_id = ? and role = ?", prodInstId, role).First(&model).Error
	if err != nil {
		return ComponentObj{}, err
	}
	obj := ComponentObj{
		ProductInstId:   model.ProductInstId,
		ComponentName:   model.ComponentName,
		KubeClusterName: model.KubeClusterName,
		Role:            model.Role,
		Input:           make(map[string]string),
		Output:          make(map[string]string),
	}
	json.Unmarshal([]byte(model.InputJson), &obj.Input)
	json.Unmarshal([]byte(model.OutputJson), &obj.Output)
	return obj, nil
}

func (m *Manager) RenderInput(instId string, refer ComponentRefer) map[string]string {
	args := make(map[string]interface{})
	//渲染参数
	for _, r := range refer.PreRole {
		component, err := m.GetCompnentObj(instId, r)
		if err != nil {
			fmt.Println(err)
		}
		args[r] = component.Output
	}
	for inK, inV := range refer.Input {
		buf := new(bytes.Buffer)
		tpl, _ := template.New("letter").Parse(inV)
		tpl.Execute(buf, args)
		refer.Input[inK] = buf.String()
	}
	fmt.Println("render ", refer.Input)
	return refer.Input
}

func (m *Manager) CreateComponentObj(prodInst ProductInstModel, refer ComponentRefer) error {
	var count int
	err := m.db.Model(&ComponentObjModel{}).Where("product_inst_id=? and role=?", prodInst.InstId, refer.Role).Count(&count).Error
	if err != nil {
		return err
	}
	if count == 1 {
		fmt.Println("count == 1")
		return nil
	}

	input := m.RenderInput(prodInst.InstId, refer)
	//创建组件
	conf, ok := m.clusterMgr.GetKubeConfig(prodInst.KubeClusterName)
	if !ok {
		return fmt.Errorf("kubeclustername %s not exist", prodInst.KubeClusterName)
	}
	err, out := m.componentMgr.Excute(refer.ComponentName, "CREATE", conf, input, prodInst.InstId, refer.Role)
	if err != nil {
		return err
	}
	//写库
	obj := ComponentObj{
		ProductInstId:   prodInst.InstId,
		ComponentName:   refer.ComponentName,
		KubeClusterName: prodInst.KubeClusterName,
		Role:            refer.Role,
		DesignRevision:  prodInst.DesignRevision,
		Input:           input,
		Output:          out,
	}
	fmt.Println("CreateComponentObj ", obj)
	objModel := ComponentObj2Model(obj)
	return m.db.Create(&objModel).Error
}

func (m *Manager) UpdateComponentObj(prodInst ProductInstModel, refer ComponentRefer, revision string) error {
	obj, _ := m.GetCompnentObj(prodInst.InstId, refer.Role)
	if obj.DesignRevision == revision {
		return nil
	}
	input := m.RenderInput(prodInst.InstId, refer)

	conf, ok := m.clusterMgr.GetKubeConfig(prodInst.KubeClusterName)
	if !ok {
		return fmt.Errorf("kubeclustername %s not exist", prodInst.KubeClusterName)
	}
	err, out := m.componentMgr.Excute(obj.ComponentName, "UPDATE", conf, input, prodInst.InstId, refer.Role)
	if err == nil {
		return err
	}
	inByte, _ := json.Marshal(input)
	outByte, _ := json.Marshal(out)
	return m.db.Model(&ComponentObjModel{}).
		Where("product_inst_id = ? and role = ?", obj.ProductInstId, obj.Role).
		Update(map[string]interface{}{
			"input_json":  string(inByte),
			"output_json": string(outByte),
			"revision":    revision,
			"status":      StatusUpdate}).Error
}

func (m *Manager) ComponentObjReady(prodInst ProductInstModel, role string, component string) error {
	conf, ok := m.clusterMgr.GetKubeConfig(prodInst.KubeClusterName)
	if !ok {
		return fmt.Errorf("kubeclustername %s not exist", prodInst.KubeClusterName)
	}
	err, _ := m.componentMgr.Excute(component, "READY", conf, nil, prodInst.InstId, role)
	if err != nil {
		return err
	}
	return m.db.Model(&ComponentObjModel{}).
		Where("product_inst_id = ? and role = ?", prodInst.InstId, role).
		Updates(map[string]interface{}{
			"status": StatusReady}).Error

}
