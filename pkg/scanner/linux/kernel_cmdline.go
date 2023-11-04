package linux

import (
	"fmt"
	"os"
	"strings"

	"compat/pkg/scanner"
)

const KernelCmdlineScanner = "kernelCmdline"

type KernelCmdline struct {
	Data map[string]string
}

func (c *KernelCmdline) Name() string {
	return KernelCmdlineScanner
}

func (c *KernelCmdline) Run(i interface{}) error {
	var input scanner.DynamicMap = i.(map[string]interface{})

	for param := range input {
		input.Val(param)
	}

	matchCnt := 0
	for param := range input {
		if val, ok := c.Data[param]; ok && input.Val(param) == val {
			matchCnt++
		}
	}
	if matchCnt != len(input) {
		return fmt.Errorf("not all kernel boot cmdline options are matched")
	}

	return nil
}

func NewKernelCmdline() (scanner.Scanner, error) {
	data, err := loadBootData()
	if err != nil {
		return nil, err
	}
	return &KernelCmdline{Data: data}, nil
}

func loadBootData() (map[string]string, error) {
	data := make(map[string]string)

	rawData, err := os.ReadFile("/proc/cmdline")
	if err != nil {
		return nil, err
	}

	args := strings.Split(strings.TrimSpace(string(rawData)), " ")
	for _, arg := range args {
		d := strings.Split(arg, "=")
		if len(d) > 1 {
			data[d[0]] = d[1]
		} else {
			data[d[0]] = ""
		}
	}

	return data, nil
}
