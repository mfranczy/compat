package validate

import (
	"flag"
	"fmt"
	"github.com/mfranczy/compat/pkg/validate"
	"runtime"

	"github.com/mfranczy/compat/pkg/scanner"
	"github.com/mfranczy/compat/schema"

	// Register scanners
	_ "github.com/mfranczy/compat/pkg/scanner/linux"
)

const (
	CmdName = "validate"

	artifactPathArg = "artifact-path"
)

func Run(args []string) error {
	var artifactPath string

	f := flag.NewFlagSet(CmdName, flag.ExitOnError)
	f.StringVar(&artifactPath, artifactPathArg, "", "os path to artifact")
	if err := f.Parse(args); err != nil {
		return err
	}

	ic, err := schema.LoadSchema(artifactPath)
	if err != nil {
		return err
	}

	var sc scanner.OsScanner
	switch runtime.GOOS {
	case scanner.Linux.String():
		sc = scanner.Scanners.Linux
		break
	case scanner.Illumos.String():
		sc = scanner.Scanners.Linux
		break
	case scanner.Windows.String():
		sc = scanner.Scanners.Windows
		break
	default:
		return fmt.Errorf("not supported OS: %s", runtime.GOOS)
	}

	res, err := validate.Run(ic, sc)
	if err != nil {
		return err
	}

	// TEST
	fmt.Println(res)

	return nil
}
