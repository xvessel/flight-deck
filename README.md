flight-deck是将组件经过编排组装成产品(这里指一整套系统)的服务.

### 流程
```
           组件-------------->产品设计------------->产品实例
涉及人员: 中间件开发            业务ops               业务ops

```


### 概念
```
Component:基础组件：各种中间件，常用服务  
Design:产品设计
Instance:产品实例
Cluster:创建实例底层k8s集群
```

### 组织结构 
```
Product:
  ProductName: product
  Designs:
    - Revision: xxx           --------v
      CompnentRefers: 
        - Role: db            --------------|v|
          ComponentName: mysql
          Input: {USER:root,PASS:root}
          PreRoles: 
  Instances:
    - Name: test1
      KubeClusterName: test
      DesignRevision: xxxx    --------^
      CompnentObjs: 
        - Role: db            --------------|^|
          ComponentName: mysql
          Input: {USER:root,PASS:root}
          Output: 
```

### 中间件开发人员如何接入： 
[component定义](./components/README.md)  

### kubernetes使用  
```
productName---envrionmentName---instanceName---role
              |-k8s cluster-|               
  |-------------k8s namespace-------------|
  |-------------ComponentObject------------------|
```
