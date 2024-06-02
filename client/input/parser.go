package input

import (
	"errors"
	"flag"
	"fmt"
	"gogogo-tdl/commands"
)

type Parser struct {
}

type Command interface {
	Run() error
}

type CommandLineArgs struct {
	Command string
	Name    string
	Content string
}

func (i *Parser) Parse() (Command, error) {
	args, err := parseArguments()
	if err != nil {
		return nil, errors.New("error parsing input")
	}
	command, err := getCommandFromArgs(args)
	if err != nil {
		return nil, errors.New("error creating command")
	}
	return command, nil
}

func parseArguments() (*CommandLineArgs, error) {
	var args CommandLineArgs

	flag.StringVar(&args.Command, "c", "", "Command")
	flag.StringVar(&args.Name, "n", "", "Name")
	flag.StringVar(&args.Content, "c", "", "Content")

	flag.Parse()

	if len(args.Command) == 0 || len(args.Name) == 0 || len(args.Content) == 0 {
		fmt.Println("input must be -c command -n name -c content")
		return nil, fmt.Errorf("invalid input")
	}

	fmt.Printf("Received Command: %s\n", args.Command)
	fmt.Printf("Received Name: %s\n", args.Name)
	fmt.Printf("Received Content: %s\n", args.Content)

	return &args, nil
}

func getCommandFromArgs(args *CommandLineArgs) (Command, error) {
	switch args.Command {
	case "create":
		return &commands.CreateCommand{
			Name:    args.Name,
			Content: args.Content,
		}, nil
	case "translate":
		return &commands.TranslateCommand{
			Name: args.Name,
		}, nil
	}
	return nil, errors.New("invalid command")
}
