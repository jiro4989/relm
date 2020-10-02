package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/docopt/docopt-go"
)

type CommandLineParam struct {
	Command string   `docopt:"<command>"`
	Args    []string `docopt:"<args>"`
}

type CommandLineInitParam struct {
	Init bool
}

type CommandLineEditParam struct {
	Edit   bool
	Editor string `docopt:"-e,--editor"`
}

type CommandLineInstallParam struct {
	Install          bool
	GitHubReleaseURL string `docopt:"<github_release_url>"`
}

type CommandLineUpdateParam struct {
	Update   bool
	Yes      bool     `docopt:"-y,--yes"`
	Releases []string `docopt:"<releases>"`
}

type CommandLineUpgradeParam struct {
	Upgrade  bool
	Yes      bool     `docopt:"-y,--yes"`
	Releases []string `docopt:"<releases>"`
}

type CommandLineUninstallParam struct {
	Uninstall bool
	OwnerRepo string `docopt:"<owner/repo>"`
}

type CommandLineListParam struct {
	List bool
}

const (
	appName = "relma"
	version = "v1.0.0"
	usage   = `relma manages GitHub Releases versioning.

usage:
  relma [options] <command> [<args>...]
  relma -h | --help
  relma --version

commands:
  init         initialize config file.
  edit         edit config file.
  install      install GitHub Releases.
  update       update installed version infomation.
  upgrade      upgrade installed GitHub Releases.
  uninstall    uninstall GitHub Releases.
  list         print installed GitHub Releases infomation.

options:
  -h, --help    print this help
  --version     print version
`

	usageInit = `usage: relma init [options]

options:
  -h, --help       print this help
`

	usageEdit = `usage: relma edit [options]

options:
  -h, --help               print this help
  -e, --editor=<editor>    using editor
`

	usageInstall = `usage: relma install [options] <github_release_url>

options:
  -h, --help       print this help
`

	usageUpdate = `usage: relma update [options] [<releases>...]

options:
  -h, --help       print this help
  -y, --yes        yes
`

	usageUpgrade = `usage: relma upgrade [options] [<releases>...]

options:
  -h, --help       print this help
  -y, --yes        yes
`

	usageUninstall = `usage: relma uninstall [options] <owner/repo>

options:
  -h, --help       print this help
`

	usageList = `usage: relma list [options]

options:
  -h, --help       print this help
`
)

func main() {
	os.Exit(Main(os.Args[1:]))
}

func Main(args []string) int {
	parser := &docopt.Parser{OptionsFirst: true}

	opts, err := parser.ParseArgs(usage, args, version)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	var clp CommandLineParam
	err = opts.Bind(&clp)
	if err != nil {
		panic(err)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	a := App{
		Config: Config{
			RelmaRoot: filepath.Join(home, "relma"),
		},
	}
	switch clp.Command {
	case "init":
		args := []string{clp.Command}
		opts, err := docopt.ParseArgs(usageInit, args, "")
		if err != nil {
			panic(err)
		}
		var clp CommandLineInitParam
		err = opts.Bind(&clp)
		if err != nil {
			panic(err)
		}

		err = a.CmdInit()
	case "edit":
		args := []string{clp.Command}
		args = append(args, opts["<args>"].([]string)...)
		opts, err := docopt.ParseArgs(usageEdit, args, "")
		if err != nil {
			panic(err)
		}
		var clp CommandLineEditParam
		err = opts.Bind(&clp)
		if err != nil {
			panic(err)
		}

		err = a.CmdEdit(&clp)
	case "install":
		args := []string{clp.Command}
		args = append(args, opts["<args>"].([]string)...)
		opts, err := docopt.ParseArgs(usageInstall, args, "")
		if err != nil {
			panic(err)
		}
		var clp CommandLineInstallParam
		err = opts.Bind(&clp)
		if err != nil {
			panic(err)
		}

		err = a.CmdInstall(clp.GitHubReleaseURL)
	case "update":
		args := []string{clp.Command}
		args = append(args, opts["<args>"].([]string)...)
		opts, err := docopt.ParseArgs(usageUpdate, args, "")
		if err != nil {
			panic(err)
		}
		var clp CommandLineUpdateParam
		err = opts.Bind(&clp)
		if err != nil {
			panic(err)
		}

		p := CmdUpdateParam{
			Yes:      clp.Yes,
			Releases: clp.Releases,
		}
		err = a.CmdUpdate(&p)
	case "upgrade":
		args := []string{clp.Command}
		args = append(args, opts["<args>"].([]string)...)
		opts, err := docopt.ParseArgs(usageUpgrade, args, "")
		if err != nil {
			panic(err)
		}
		var clp CommandLineUpgradeParam
		err = opts.Bind(&clp)
		if err != nil {
			panic(err)
		}

		p := CmdUpgradeParam{
			Yes: clp.Yes,
		}
		err = a.CmdUpgrade(&p)
	case "uninstall":
		args := []string{clp.Command}
		args = append(args, opts["<args>"].([]string)...)
		opts, err := docopt.ParseArgs(usageUninstall, args, "")
		if err != nil {
			panic(err)
		}
		var clp CommandLineUninstallParam
		err = opts.Bind(&clp)
		if err != nil {
			panic(err)
		}

		o := strings.Split(clp.OwnerRepo, "/")
		err = a.CmdUninstall(o[0], o[1])
	case "list":
		args := []string{clp.Command}
		args = append(args, opts["<args>"].([]string)...)
		opts, err := docopt.ParseArgs(usageList, args, "")
		if err != nil {
			panic(err)
		}
		var clp CommandLineListParam
		err = opts.Bind(&clp)
		if err != nil {
			panic(err)
		}

		err = a.CmdList(&clp)
	}
	if err != nil {
		panic(err)
	}

	return 0
}
