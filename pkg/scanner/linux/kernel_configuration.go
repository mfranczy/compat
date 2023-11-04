package linux

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/moby/moby/pkg/parsers/kernel"

	"compat/pkg/scanner"
)

const KernelConfigurationScanner = "kernelConfiguration"

type KernelConfiguration struct {
	Data map[string]string
}

func (c *KernelConfiguration) Name() string {
	return KernelConfigurationScanner
}

func (c *KernelConfiguration) Run(i interface{}) error {
	var input scanner.DynamicMap = i.(map[string]interface{})

	matchCnt := 0
	for config := range input {
		if val, ok := c.Data[config]; ok && input.Val(config) == val {
			matchCnt++
		}
	}
	if matchCnt != len(input) {
		return fmt.Errorf("not all kernel configuration options are matched")
	}
	return nil
}

func NewKernelConfiguration() (scanner.Scanner, error) {
	data, err := loadConfiguration()
	if err != nil {
		return nil, err
	}
	return &KernelConfiguration{Data: data}, nil
}

func loadConfiguration() (map[string]string, error) {
	data := make(map[string]string)
	kv, err := kernel.GetKernelVersion()
	if err != nil {
		return nil, err
	}

	// TODO: check more files, not only /boot/config-$(uname -r)
	path := fmt.Sprintf("/boot/config-%s", kv)
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	s := bufio.NewScanner(fd)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		t := s.Text()
		if strings.HasPrefix(t, "#") {
			continue
		}
		c := strings.Split(t, "=")
		if len(c) >= 2 {
			data[c[0]] = c[1]
		}
	}

	if err = fd.Close(); err != nil {
		return nil, err
	}
	return data, nil
}
