
## 1. 组件接口
### 1.1 查询
request
```
curl http://127.0.0.1:8080/components
```
response
```
[
  {
    "Name": "example",
    "Description": "hello",
    "Inputs": [
      {
        "Name": "CPU",
        "DefaultValue": "1",
        "CanUpdate": true,
        "Description": "cpu核数"
      },
      {
        "Name": "MEM",
        "DefaultValue": "2",
        "CanUpdate": true,
        "Description": "内存大小G"
      }
    ],
    "Outputs": [
      {
        "Name": "HOST",
        "Description": "访问host"
      },
      {
        "Name": "PORT",
  1
        "Description": "访问port"
      }
    ]
  }
]
```
## 2 产品接口  
### 2.1 添加产品
request 
```
curl -XPOST http://127.0.0.1:8080/products -d \
    '{"name":"test-product"}'
```
response
status:200, 其他为错误

### 2.2 查询产品
request
```
curl http://127.0.0.1:8080/products
```
response 
```
[{"ID":1,"CreatedAt":"2019-09-29T14:46:54.47324+08:00","UpdatedAt":"2019-09-29T14:46:54.47324+08:00","DeletedAt":null,"Name":"test-product"}]
```

## 3 设计接口
### 3.1 添加设计  
request
```
curl -XPOST http://127.0.0.1:8080/designs -d \
'{
"product_name":"test-product",
"revision":"test-revision",
"component_refers":[
{
    "role":"test-role",
    "component_name":"example",
    "input":{},
    "pre_role":[]
}
]
}'
```

### 3.2 设计列表
request
```
curl  http://127.0.0.1:8080/products/test-product/designs 
```
response 
```
[
  {
    "ProductName": "test-product",
    "Revision": "test-revision",
    "ComponentRefers": [
      {
        "Role": "test-role",
        "ComponentName": "example",
        "Input": {
          "CPU": "1",
          "MEM": "1"
        },
        "PreRole": []
      }
    ]
  }
]
```
## 4. 实例
### 4.1 创建实例
request 
```
curl -w %{http_code} http://127.0.0.1:8080/products/test-product/instances -d \
'
{
    "Name":"test-instance",
    "ProductName":"test-product",
    "DesignRevision":"test-revision",
    "KubeClusterName":"test-cluster"
}
'
```
response 
status:200
err: body

### 4.2 列表实例 
request 
```
curl http://127.0.0.1:8080/products/test-product/instances
```
response
```
[
  {
    "ID": 1,
    "CreatedAt": "2019-09-29T18:07:25.702271+08:00",
    "UpdatedAt": "2019-09-30T09:22:56.037074+08:00",
    "DeletedAt": null,
    "InstId": "test-product-test-cluster-test-instance",
    "ProductName": "test-product",
    "Name": "test-instance",
    "KubeClusterName": "test-cluster",
    "DesignRevision": "test-revision",
    "Status": "ready"
  }
]
```
