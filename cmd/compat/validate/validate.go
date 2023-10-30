package validate

import (
	"flag"
	"fmt"

	"github.com/mfranczy/compat/pkg/scanner"
	pkgvalidate "github.com/mfranczy/compat/pkg/validate"
	"github.com/mfranczy/compat/schema"

	// Register scanners (TODO: register based on build)
	_ "github.com/mfranczy/compat/pkg/scanner/linux"
)

const (
	CmdName = "validate"

	artifactPathArg = "artifact-path"
	workersNumArg   = "workers-num"
)

func Run(args []string) error {
	var (
		artifactPath string
		workersNum   int
	)

	f := flag.NewFlagSet(CmdName, flag.ExitOnError)
	f.StringVar(&artifactPath, artifactPathArg, "", "os path to artifact")
	f.IntVar(&workersNum, workersNumArg, 1, "number of workers processing the schema")
	if err := f.Parse(args); err != nil {
		return err
	}

	if artifactPath == "" {
		return fmt.Errorf("--artifact-path arg is required")
	}

	ic, err := schema.LoadSchema(artifactPath)
	if err != nil {
		return err
	}

	res, err := pkgvalidate.Run(ic, scanner.Scanners, workersNum)
	if err != nil {
		return err
	}

	// TEST
	fmt.Println("Validation results:")
	fmt.Println()
	for k, v := range res {
		for k1, v1 := range v {
			if k1 == k {
				fmt.Printf("%s: %t\n", k, v1)
			} else {
				fmt.Printf("%s.%s: %t\n", k, k1, v1)
			}
		}
		fmt.Println()
	}

	// my hardware is intel
	// cpu X
	//  - intel X - break early
	//  - amd

	// cpu - arm64
	// - intel
	// - amd

	// cpu - Failed

	// cpu
	// 	- intel - Passed
	//
	// vfio - Passed

	// compat validate --artifact-path <local>
	// compat validate <image>

	// compat
	// compat validate docker.io/busybox:v1.20

	// compat - is a main tool to validate and create image compatibility artifacts,
	//  - validate-schema - runs schema validation
	//  - validate-host - runs validation on the host

	return nil
}
