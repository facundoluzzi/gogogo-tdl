package input

import (
	"errors"
	"file-editor/commands"
	"flag"
	"fmt"
	"strings"
)

var (
	ErrNoArgs       = errors.New("no arguments")
	ErrInvalidInput = errors.New("invalid input")
)

type Parser struct {
}

type Command interface {
	Run() error
	Print()
}

type CommandLineArgs struct {
	Command string
	Name    string
	Body    string
}

func (i *Parser) ParseFromArgs() (Command, error) {
	args, err := parseArguments()
	if err != nil {
		return nil, err
	}
	command, err := getCommandFromArgs(args)
	if err != nil {
		return nil, err
	}
	return command, nil
}

func (i *Parser) Parse(input string) (Command, error) {
	inputSplit := strings.Fields(input)
	args, err := parseSlice(inputSplit)
	if err != nil {
		return nil, err
	}
	command, err := getCommandFromArgs(args)
	if err != nil {
		return nil, err
	}
	return command, nil
}

func parseSlice(input []string) (*CommandLineArgs, error) {
	if len(input) != 3 {
		fmt.Println("input must be command name body")
		return nil, ErrInvalidInput
	}

	return &CommandLineArgs{
		Command: input[0],
		Name:    input[1],
		Body:    input[2],
	}, nil
}

func parseArguments() (*CommandLineArgs, error) {
	var args CommandLineArgs

	flag.StringVar(&args.Command, "c", "", "Command")
	flag.StringVar(&args.Name, "n", "", "Name")
	flag.StringVar(&args.Body, "b", "", "Body")

	flag.Parse()

	if len(args.Command) == 0 {
		return nil, ErrNoArgs
	}

	if len(args.Command) == 0 || len(args.Name) == 0 || len(args.Body) == 0 {
		fmt.Println("input must be -c command -n name -b body")
		return nil, fmt.Errorf("invalid input")
	}
	return &args, nil
}

func getCommandFromArgs(args *CommandLineArgs) (Command, error) {
	switch args.Command {
	case "create":
		return &commands.CreateCommand{
			Name:    args.Name,
			Content: args.Body,
		}, nil
	case "translate":
		return &commands.TranslateCommand{
			Name: args.Name,
		}, nil
	}
	return nil, errors.New("invalid command")
}