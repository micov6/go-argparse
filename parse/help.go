package parse

import (
	"fmt"
	"strings"
)

var HelpCommand = Command{
	code:        "help",
	description: "Show help",
	aliases:     []string{"--help", "-h"},
}

var HelpChildCommand = Command{
	code:        "help",
	description: "Show help",
	aliases:     []string{"--help", "-h"},
}

func helpToString(help helpInfo) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("%v\n\n", help.description))
	// TODO format
	if help.usageSuffix != "" {
		sb.WriteString(fmt.Sprintf("  usage: %v %v\n\n", help.code, help.usageSuffix))
	}

	if help.childrenName != "" {
		sb.WriteString(fmt.Sprintf("  %v:\n", help.childrenName))
	}
	for _, child := range help.children {
		sb.WriteString(fmt.Sprintf("    %v - %v\n", child.code, child.description))
	}

	return sb.String()
}

func (hp *App) Help() helpInfo {
	children := make([]helpInfo, len(hp.commands))

	for k, v := range hp.commands {
		children[k] = v.Help()
	}

	return helpInfo{
		code:         hp.code,
		description:  hp.description,
		usageSuffix:  "[command] [subcommand] [...parameters]",
		childrenName: "commands",
		children:     children,
	}
}

func (hp *Command) Help() helpInfo {
	var description string
	if len(hp.aliases) <= 0 {
		description = hp.description
	} else {
		description = fmt.Sprintf("%v. Alternatives: %v", hp.description, strings.Join(hp.aliases, ", "))
	}

	children := make([]helpInfo, len(hp.children))
	for k, v := range hp.children {
		children[k] = v.Help()
	}

	return helpInfo{
		code:        hp.code,
		description: description,
		usageSuffix:  "[subcommand] [...parameters]",
		childrenName: "subcommands",
		children:    children,
	}
}
