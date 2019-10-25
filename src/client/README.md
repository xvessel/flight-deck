# deck client 

deck 命令行工具

## component

### create component

```shell script
➜ ./dkctl component init etcd
➜ tree etcd 
etcd
├── CREATE
├── DELETE
├── PRE
├── READY
├── SPEC.yml
├── UPDATE
└── UPDATE_CHECK

0 directories, 7 files
```

### list all components

```shell script
➜ ./dkctl component list     
example hello
```

### get one component detail

```shell script
➜ ./dkctl component get example
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
                "Description": "访问port"
            }
        ]
    }

```