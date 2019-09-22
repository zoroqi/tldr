package main

import (
	"fmt"
	"os"

	"github.com/isacikgoz/gitin/prompt"
	"github.com/isacikgoz/tldr/cli"
	"github.com/isacikgoz/tldr/config"
	env "github.com/kelseyhightower/envconfig"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {

	clear := kingpin.Flag("clear-cache", "Clear local repository then clone github.com/tldr-pages/tldr").Short('c').Bool()
	update := kingpin.Flag("update", "Pulls the latest commits from github.com/tldr-pages/tldr").Short('u').Bool()
	static := kingpin.Flag("static", "Static mode, application behaves like a conventional tldr client.").Short('s').Default("false").Bool()
	random := kingpin.Flag("random", "Random page for testing purposes.").Short('r').Default("false").Bool()

	pages := kingpin.Arg("command", "Name of the command. (e.g. tldr grep)").Strings()

	kingpin.Version("tldr++ version 1.0-alpha")
	kingpin.Parse()

	config.StartUp(*clear, *update)

	var o prompt.Options
	err := env.Process("gitin", &o)
	exitIfError(err)

	exitIfError(run(*pages, &o, *static, *random))

}

func run(pages []string, opts *prompt.Options, static, random bool) error {
	p, err := cli.NewDefaultPrompt(pages, opts, static, random)
	if err != nil {
		return err
	}

	err = p.Run()
	if err != nil {
		return err
	}

	item := p.Selection()
	if item == nil {
		return nil
	}
	command := cli.SuggestCommand(item)

	return cli.ConfirmCommand(command)
}

func exitIfError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
