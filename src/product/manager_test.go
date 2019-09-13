/**********************************************************
 * Author        : biliucheng
 * Email         : bilc_dev@163.com
 * Last modified : 2019-09-13 16:18
 * Filename      : manager_test.go
 * Description   :
 * Modified By   :
 * *******************************************************/
package product

import (
	"os"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"product/mock_cluster"
	"product/mock_component"
)

func TestDesign(t *testing.T) {
	os.Remove("test.db")
	refer1 := ComponentRefer{Role: "role1", ComponentName: "compnent1"}
	refer2 := ComponentRefer{Role: "role2",
		ComponentName: "compnent2",
		Input:         map[string]string{"ROLE2_HOST": "{{.role1.HOST}}+{{.role1.PORT}}"},
		PreRole:       []string{"role1"}}
	design1 := Design{ProductName: "product1", Revision: "revision1",
		ComponentRefers: []ComponentRefer{refer1, refer2}}

	m := NewManager(nil, nil, "test")
	m.NewProduct("product1")
	pm, err := m.GetProduct("product1")
	assert.Nil(t, err)
	assert.Equal(t, "product1", pm.Name)

	err = m.NewDesign(&design1)
	assert.Nil(t, err)
	d, err := m.GetDesign("product1", "revision1")
	assert.Nil(t, err)
	assert.Equal(t, design1, d)

	inst := ProductInst{
		Name:            "inst1",
		ProductName:     "product1",
		DesignRevision:  "revision1",
		KubeClusterName: "cluster1",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock_clust := mock_cluster.NewMockManager(ctrl)
	mock_comp := mock_component.NewMockManager(ctrl)

	m.clusterMgr = mock_clust
	m.componentMgr = mock_comp

	mock_clust.EXPECT().GetKubeConfig(gomock.Eq("cluster1")).Return("ok", true).AnyTimes()
	mock_comp.EXPECT().Excute(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil, map[string]string{"HOST": "127.0.0.1", "PORT": "8080"}).AnyTimes()
	model := ProductInst2Model(&inst)
	err = m.CreateComponentObj(model, refer1)
	assert.Nil(t, err)
	obj, err := m.GetCompnentObj(model.InstId, "role1")
	assert.Nil(t, err)
	assert.Equal(t, model.InstId, obj.ProductInstId)
	assert.Equal(t, "role1", obj.Role)
	assert.Equal(t, map[string]string{"HOST": "127.0.0.1", "PORT": "8080"}, obj.Output)

	err = m.CreateComponentObj(model, refer2)
	assert.Nil(t, err)
	err = m.ComponentObjReady(model, refer2.Role, refer2.ComponentName)
	assert.Nil(t, err)
	obj2, err := m.GetCompnentObj(model.InstId, "role2")
	assert.Nil(t, err)
	assert.Equal(t, map[string]string{"ROLE2_HOST": "127.0.0.1+8080"}, obj2.Input)

	refer3 := ComponentRefer{Role: "role3", ComponentName: "compnent3", PreRole: []string{"role4"}}
	design2 := Design{ProductName: "product1", Revision: "revision2",
		ComponentRefers: []ComponentRefer{refer1, refer2, refer3}}

	err = m.NewDesign(&design2)
	assert.NotNil(t, err)
}

func TestProduct(t *testing.T) {
	os.Remove("test1.db")
	refer1 := ComponentRefer{Role: "role1", ComponentName: "compnent1"}
	refer2 := ComponentRefer{Role: "role2", ComponentName: "compnent2", PreRole: []string{"role1"}}
	design1 := Design{ProductName: "product1", Revision: "revision1",
		ComponentRefers: []ComponentRefer{refer1, refer2}}

	m := NewManager(nil, nil, "test1")
	err := m.NewProduct("product1")
	assert.Nil(t, err)
	err = m.NewDesign(&design1)
	assert.Nil(t, err)

	inst := ProductInst{
		Name:            "inst1",
		ProductName:     "product1",
		DesignRevision:  "revision1",
		KubeClusterName: "cluster1",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock_clust := mock_cluster.NewMockManager(ctrl)
	mock_comp := mock_component.NewMockManager(ctrl)

	m.clusterMgr = mock_clust
	m.componentMgr = mock_comp

	mock_clust.EXPECT().GetKubeConfig(gomock.Eq("cluster1")).Return("ok", true).AnyTimes()
	mock_comp.EXPECT().Excute(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil, map[string]string{"HOST": "127.0.0.1", "PORT": "8080"}).AnyTimes()

	err = m.NewProductInst(&inst)
	assert.Nil(t, err)
	time.Sleep(time.Second * 2)
	instModel, err := m.GetProductInst1(inst.ProductName, inst.Name)
	assert.Nil(t, err)
	assert.Equal(t, StatusReady, instModel.Status)
}
