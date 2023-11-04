package linux

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sys/unix"

	"github.com/moby/moby/pkg/parsers/kernel"

	"compat/pkg/scanner"
)

const (
	KernelModulesScanner = "kernelModules"
	KernelDriversScanner = "kernelDrivers"

	modulesPath = "/sys/module"
)

var modules map[string]*Module

type Module struct {
	Version    string
	Parameters map[string]string
}

func initModules() error {
	if modules != nil {
		return nil
	}

	kernelVersion, err := kernel.GetKernelVersion()
	if err != nil {
		return err
	}

	dirs, err := os.ReadDir(modulesPath)
	if err != nil {
		return err
	}
	modules = make(map[string]*Module, len(dirs))

	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}
		modules[dir.Name()] = &Module{}
		modulePath := filepath.Join(modulesPath, dir.Name())

		v, err := os.ReadFile(filepath.Join(modulePath, "version"))
		ver := ""
		if err == nil {
			ver = string(v)
		} else {
			ver = kernelVersion.String()
		}
		modules[dir.Name()].Version = strings.TrimSpace(ver)

		paramsPath := filepath.Join(modulePath, "parameters")
		params, err := os.ReadDir(paramsPath)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return err
		}
		modules[dir.Name()].Parameters = make(map[string]string, len(params))
		for _, param := range params {
			c, err := os.ReadFile(filepath.Join(paramsPath, param.Name()))
			if err != nil {
				var pathErr *fs.PathError
				if errors.As(err, &pathErr) && pathErr.Err == unix.EPERM || pathErr.Err == unix.EACCES {
					// TODO: log about EPERM and EACCESS
					continue
				}
				return err
			}
			modules[dir.Name()].Parameters[param.Name()] = strings.TrimSpace(string(c))
		}
	}

	return nil
}

type KernelModules struct{}

func (m *KernelModules) Name() string {
	return KernelModulesScanner
}

func (m *KernelModules) Run(i interface{}) error {
	var input scanner.DynamicMap = i.(map[string]interface{})
	exModulesMatch := len(input)
	modulesMatch := 0

	for module := range input {
		params := input.Map(module)
		_, moduleMatch := modules[module]

		for param := range params {
			if input.Map(module).Val(param) != modules[module].Parameters[param] {
				moduleMatch = false
			}
		}

		if moduleMatch {
			modulesMatch++
		}
	}

	if exModulesMatch != modulesMatch {
		return fmt.Errorf("not all kernel modules are matched")
	}
	return nil
}

func NewKernelModules() (scanner.Scanner, error) {
	if err := initModules(); err != nil {
		return nil, err
	}
	return &KernelModules{}, nil
}

type KernelDrivers struct{}

func (m *KernelDrivers) Name() string {
	return KernelDriversScanner
}

func (m *KernelDrivers) Run(i interface{}) error {
	var input scanner.DynamicMap = i.(map[string]interface{})

	for module := range input {
		if modules[module]
		for version := range input.Map(module) {
			for param := range input.Map(module).Map(version) {
				fmt.Println(module, version, param)
			}
		}
	}
	return nil
}

func NewKernelDrivers() (scanner.Scanner, error) {
	if err := initModules(); err != nil {
		return nil, err
	}
	return &KernelDrivers{}, nil
}
