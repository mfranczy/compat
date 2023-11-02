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
	input, err := scanner.ConvertInputToMap(i)
	if err != nil {
		return err
	}

	matchCnt := 0
	for k, exVal := range input {
		if val, ok := c.Data[k]; ok && exVal == val {
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

	args := strings.Split(string(rawData), " ")
	for _, arg := range args {
		d := strings.Split(arg, "=")
		if len(d) > 1 {
			data[d[0]] = strings.TrimSpace(d[1])
		} else {
			data[d[0]] = ""
		}
	}

	return data, nil
}
