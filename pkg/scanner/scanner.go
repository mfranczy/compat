package scanner

import (
	"fmt"
)

type OS string

func (o OS) String() string {
	return string(o)
}

const (
	Linux   OS = "linux"
	Illumos OS = "illumos"
	Windows OS = "windows"
)

type Scanner interface {
	Name() string
	ID() string
	Group() string
	Input() interface{}
	Run() error
}

type Creator func(id string, group string, input interface{}) (Scanner, error)
type OsScanner map[string]Creator

// TODO: get rid of the struct and provide scanners based on build
type OsScanners struct {
	Linux   OsScanner
	Illumos OsScanner
	Windows OsScanner
}

var Scanners OsScanners

func init() {
	Scanners.Linux = make(OsScanner)
	Scanners.Illumos = make(OsScanner)
	Scanners.Windows = make(OsScanner)
}

func Register(os OS, name string, creator Creator) {
	switch os {
	case Linux:
		Scanners.Linux[name] = creator
		break
	case Illumos:
		Scanners.Illumos[name] = creator
		break
	case Windows:
		Scanners.Windows[name] = creator
		break
	default:
		panic(fmt.Sprintf("not supported OS: %s", os))
	}
}
