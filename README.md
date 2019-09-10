这是一个基于k8s的构建私有云的支撑系统。  
目的是利用k8s的helm和operator的能力提供中间件，在此上构建应用系统。  

概念:
```
Component:基础组件：各种中间件，常用服务  
Design:产品设计
Instance:产品实例
Cluster:底层k8s环境
```

逻辑关系： 
```
Product:
  ProductName: product
  Designs:
    - Revision: xxxx
      Compnents: 
        - Role: aa
          Component: mysql
          Input: asdf
          PreCondition: 
  Instances:
    - Name: test1
      Environment: test
      DesignRevision: xxxx
      CompnentRefers: 
        - ID: product-test-test1-db
	  Role: db
          Component: mysql
          Input: asdf
          PreRole: 
```
[component定义](./components/README.md)  

命名空间
```

productName---envrionmentName---instanceName---role
              |-k8s cluster-|               
  |-------------k8s namespace-------------|
  |-------------ComponentObject------------------|
  
ComponentObject: 可能是helm的release，或者是operator的资源id，或者是个deploymentid  
```
