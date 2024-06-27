package input

import (
	"errors"
	"file-editor/api"
	"file-editor/cmd/client/commands"
	"strconv"
)

const (
	NewFileCommand        = "new"
	SaveCommand           = "save"
	EditCommand           = "edit"
	ReadCommand           = "read"
	ReadAll               = "read-all"
	FindCommand           = "find"
	HelpCommand           = "help"
	DeleteTextCommand     = "delete"
	FindAndReplaceCommand = "find-replace"
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
	Args    []string
}

// func (i *Parser) ParseFromArgs() (Command, error) {
// 	args, err := parseArguments()
// 	if err != nil {
// 		return nil, err
// 	}
// 	command, err := getCommandFromArgs(args)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return command, nil
// }

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
	case NewFileCommand:
		if len(input) > 3 || len(input) < 2 {
			return nil, ErrInvalidInput
		}
		args := input[1:]
		return &CommandLineArgs{
			Command: command,
			Args:    args,
		}, nil
	case DeleteTextCommand:
		if len(input) != 4 {
			return nil, ErrInvalidInput
		}
		args := input[1:]
		return &CommandLineArgs{
			Command: command,
			Args:    args,
		}, nil
	case SaveCommand:
		if len(input) != 2 {
			return nil, ErrInvalidInput
		}
		args := input[1:]
		return &CommandLineArgs{
			Command: command,
			Args:    args,
		}, nil
	case EditCommand:
		if len(input) != 2 {
			return nil, ErrInvalidInput
		}
		args := input[1:]
		return &CommandLineArgs{
			Command: command,
			Args:    args,
		}, nil
	case ReadCommand:
		if len(input) != 2 {
			return nil, ErrInvalidInput
		}
		args := input[1:]
		return &CommandLineArgs{
			Command: command,
			Args:    args,
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
		args := input[1:]
		return &CommandLineArgs{
			Command: command,
			Args:    args,
		}, nil
	case FindAndReplaceCommand:
		if len(input) != 4 {
			return nil, ErrInvalidInput
		}
		args := input[1:]
		return &CommandLineArgs{
			Command: command,
			Args:    args,
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

// func parseArguments() (*CommandLineArgs, error) {
// 	// var args CommandLineArgs

// 	// flag.StringVar(&args.Command, "c", "", "Command")
// 	// flag.StringVar(&args.Name, "1", "", "Name")
// 	// flag.StringVar(&args.Body, "2", "", "Body")

// 	// flag.Parse()

// 	// if len(args.Command) == 0 {
// 	// 	return nil, ErrNoArgs
// 	// }

// 	// return &args, nil
// 	return nil, nil
// }

func getCommandFromArgs(args *CommandLineArgs) (Command, error) {
	switch args.Command {
	case DeleteTextCommand:
		start, err := strconv.Atoi(args.Args[1])
		if err != nil {
			return nil, err
		}
		length, err := strconv.Atoi(args.Args[2])
		if err != nil {
			return nil, err
		}
		return &commands.DeleteTextCommand{
			Name:   args.Args[0],
			Start:  start,
			Length: length,
		}, nil
	case SaveCommand:
		return &commands.SaveFileCommand{
			Name: args.Args[0],
		}, nil
	case EditCommand:
		return &commands.EditCommand{
			Name: args.Args[0],
		}, nil
	case ReadCommand:
		return &commands.ReadCommand{
			Name: args.Args[0],
		}, nil
	case ReadAll:
		return &commands.ReadAllCommand{}, nil
	case FindCommand:
		return &commands.FindCommand{
			Name: args.Args[0],
			Text: args.Args[1],
		}, nil
	case FindAndReplaceCommand:
		return &commands.FindAndReplaceCommand{
			Name:    args.Args[0],
			Find:    args.Args[1],
			Replace: args.Args[2],
		}, nil
	case NewFileCommand:
		if len(args.Args) == 1 {
			return &commands.NewFileCommand{
				Name:    args.Args[0],
				Content: "",
			}, nil
		} else if len(args.Args) == 2 {
			return &commands.NewFileCommand{
				Name:    args.Args[0],
				Content: args.Args[1],
			}, nil
		}
	case HelpCommand:
		return &commands.HelpCommand{}, nil
	}
	return nil, errors.New("invalid command")
}
