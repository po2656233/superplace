# data-config组件
- 自定义数据源
- 读取数据
- 热更新数据

## Install

### Prerequisites
- GO >= 1.17

### Using go get
```
go get superplace/components/data-config@latest
```


## Quick Start
```
import dataConf "superplace/components/data-config"
```

```
package demo
import (
	"superman"
	dataConf "superplace/components/data-config"
)

// RegisterComponent 注册struct到data-config
func RegisterComponent() {
	dataConfig := dataConf.NewComponent()
	dataConfig.Register(
		&DropList,
		&DropOne,
	)

	//data-config组件注册到cherry引擎
	extend.RegisterComponent(dataConfig)
}

```

## example
- [示例代码跳转](../../examples/test_data_config)