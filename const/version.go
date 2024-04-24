package superConst

import (
	"fmt"
)

const (
	version = "1.0.0"
)

var logo = `game sever framework @v%s`

func GetLOGO() string {
	return fmt.Sprintf(logo, Version())
}

func Version() string {
	return version
}

const (
	DOT = "." //ActorPath的分隔符
)
