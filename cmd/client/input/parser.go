package input

import (
	"errors"
	"file-editor/api"
	"file-editor/cmd/client/commands"
	"flag"
)

const (
	SaveCommand = "save"
	ReadCommand = "read"
	ReadAll     = "read-all"
	FindCommand = "find"
	HelpCommand = "help"
)

var (
	ErrNoArgs       = errors.New("no arguments")
	ErrInvalidInput = errors.New("invalid input")
)

type Parser struct {
}

type Command interface {
	Run(api.TextEditorClient) (string, error)
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
	inputSplit := splitFields(input)
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
	if len(input) == 0 {
		return nil, ErrInvalidInput
	}
	command := input[0]
	switch command {
	case SaveCommand:
		if len(input) != 3 {
			return nil, ErrInvalidInput
		}
		return &CommandLineArgs{
			Command: command,
			Name:    input[1],
			Body:    input[2],
		}, nil
	case ReadCommand:
		if len(input) != 2 {
			return nil, ErrInvalidInput
		}
		return &CommandLineArgs{
			Command: command,
			Name:    input[1],
		}, nil
	case ReadAll:
		if len(input) != 1 {
			return nil, ErrInvalidInput
		}
		return &CommandLineArgs{
			Command: command,
		}, nil
	case FindCommand:
		if len(input) != 3 {
			return nil, ErrInvalidInput
		}
		return &CommandLineArgs{
			Command: command,
			Name:    input[1],
			Body:    input[2],
		}, nil
	case HelpCommand:
		if len(input) != 1 {
			return nil, ErrInvalidInput
		}
		return &CommandLineArgs{
			Command: command,
		}, nil
	}
	return nil, ErrInvalidInput
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

	return &args, nil
}

func getCommandFromArgs(args *CommandLineArgs) (Command, error) {
	switch args.Command {
	case SaveCommand:
		return &commands.SaveFileCommand{
			Name:    args.Name,
			Content: args.Body,
		}, nil
	case ReadCommand:
		return &commands.ReadCommand{
			Name: args.Name,
		}, nil
	case ReadAll:
		return &commands.ReadAllCommand{}, nil
	case FindCommand:
		return &commands.FindCommand{
			Name: args.Name,
			Text: args.Body,
		}, nil
	case HelpCommand:
		return &commands.HelpCommand{}, nil
	}
	return nil, errors.New("invalid command")
}
