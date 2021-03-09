package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/deepmap/oapi-codegen/pkg/util"
	"github.com/godpeny/go-codegen/pkg/codegen"
)

func errExit(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}

var (
	flagPackageName string
	flagGenerate    string
	flagOutputFile  string
)

type configuration struct {
	PackageName     string            `yaml:"package"`
	GenerateTargets []string          `yaml:"generate"`
	OutputFile      string            `yaml:"output"`
	IncludeTags     []string          `yaml:"include-tags"`
	ExcludeTags     []string          `yaml:"exclude-tags"`
	TemplatesDir    string            `yaml:"templates"`
	ImportMapping   map[string]string `yaml:"import-mapping"`
	ExcludeSchemas  []string          `yaml:"exclude-schemas"`
}

func main() {

	flag.StringVar(&flagPackageName, "package", "", "The package name for generated code")
	flag.StringVar(&flagGenerate, "generate", "types,client,server,spec",
		`Comma-separated list of code to generate; valid options: "types", "client", "chi-server", "server", "spec", "skip-fmt", "skip-prune"`)
	flag.StringVar(&flagOutputFile, "o", "", "Where to output generated code, stdout is default")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Please specify a path to a OpenAPI 3.0 spec file")
		os.Exit(1)
	}

	cfg := configFromFlags()

	// If the package name has not been specified, we will use the name of the
	// swagger file.
	if cfg.PackageName == "" {
		path := flag.Arg(0)
		baseName := filepath.Base(path)
		// Split the base name on '.' to get the first part of the file.
		nameParts := strings.Split(baseName, ".")
		cfg.PackageName = codegen.ToCamelCase(nameParts[0])
	}

	for _, g := range cfg.GenerateTargets {
		switch g {
		case "client":
			opts.GenerateClient = true
		case "chi-server":
			opts.GenerateChiServer = true
		case "types":
			opts.GenerateTypes = true
		default:
			fmt.Printf("unknown generate option %s\n", g)
			flag.PrintDefaults()
			os.Exit(1)
		}
	}

	swagger, err := util.LoadSwagger(flag.Arg(0))
	if err != nil {
		errExit("error loading swagger spec\n: %s", err)
	}

	code, err := codegen.Generate(swagger, cfg.PackageName, opts)
	if err != nil {
		errExit("error generating code: %s\n", err)
	}

	if cfg.OutputFile != "" {
		err = ioutil.WriteFile(cfg.OutputFile, []byte(code), 0644)
		if err != nil {
			errExit("error writing generated code to file: %s", err)
		}
	} else {
		fmt.Println(code)
	}
}

// configFromFlags parses the flags.
func configFromFlags() *configuration {
	var cfg configuration

	if cfg.PackageName == "" {
		cfg.PackageName = flagPackageName
	}
	if cfg.GenerateTargets == nil {
		cfg.GenerateTargets = util.ParseCommandLineList(flagGenerate)
	}
	if cfg.OutputFile == "" {
		cfg.OutputFile = flagOutputFile
	}
	return &cfg
}
