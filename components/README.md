
服务组件的定义规范  
- 在components目录下创建子目录，目录名字为组件名 
- 目录中实现以下文件INPUT, OUTPUT, PRE, CREATE, READY, UPDATE_CHECK, UPDATE, DELETE

### PRE 初始化执行的操作
创建组件实例前要做的准备，比如安装operator  

### INPUT OUTPUT 环境变量说明文件  
- 文件格式：KEY=VALUE  #KEY的相关说明，帮助用户明确设置;
- VALUE为默认值或者示意，运行时会由用户输入  

### 对实例操作的可执行文件
- 执行脚本参数为xxx  namespace id
- INPUT中的变量，可执行文件运行时，可以直接用环境变量引用${KEY}
- 退出码表示是否成功,若失败，stderr可以用来返回原因  
- 建议用shell实现，可以直接使用helm，kubectl等命令，KUBECONFIG会设置好 
- CREATE执行成功，需要在stdout输出用户使用该组件的信息，遵守OUTPUT格式

```
CREATE namespace id
READY  namespace id
UPDATE_CHECK namespace id
UPDATE namespace id
DELETE namespace id
```

### 测试
test.sh是测试脚本的工具  
目录example是个例子  
使用示例：  
```
test.sh example namespace name kubeconfigfile
```
