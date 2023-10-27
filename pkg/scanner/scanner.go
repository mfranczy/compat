package scanner

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
	Run(interface{}) error
}

type Initialized map[string]Scanner
type InitFunc func() (Scanner, error)
