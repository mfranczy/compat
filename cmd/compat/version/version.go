package version

import (
	"fmt"

	v1 "compat/pkg/types/v1"
)

const (
	CmdName = "version"
)

func Run() {
	fmt.Println(v1.Version)
}
