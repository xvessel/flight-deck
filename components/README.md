
服务组件的定义规范  

### 执行文件需要的环境变量说明  
- 格式：KEY=VALUE  #KEY的相关说明，帮助用户明确设置;
- VALUE为默认值  
```
INPUT
OUTPUT
```

### 初始化执行的操作
```
PRE
```

### 可执行文件
- 执行脚本参数为  namespace id
- 其他需要变量会以环境变量的方式使用  
- 退出码表示是否正确执行  
- stdout为用户访问的变量，同OUTPUT格式
- stderr, 错误具体信息
- KUBECONFIG,k8s的配置会以环境变量的方式提供

```
CREATE namespace id
READY namespace id
UPDATE_CHECK
UPDATE
DELETE 
```

