package std

import (
	"fmt"
	"os"
	"strings"
)

// this file content is all the standard commands and functions available natively in the gsh shell
// it's by convention not modified because all of the standard functions can be commented out in the config file to be removed from the program, it's also the same for operators and so on

// clear
// the clear screen command
func clear(parameters []string) Descriptor {
	// checking if we need to prompt the helpMessage or just clear the screen
	if len(parameters) > 0 {
		// help message of the command
		//fmt.Printf("\nUsage: clear\n\nClear screen\n")
		// return values
		ret := Descriptor{
			Id:     1,
			OutMsg: "\nUsage: clear\n\nClear screen\n",
		}
		return ret
	} else {
		// clear the screen
		ret := Descriptor{
			Id:     1,
			OutMsg: "\033[H\033[2J",
		}
		return ret
	}
}

// exit
// the exit command wich exit gsh
func exit(parameters []string) Descriptor {
	if len(parameters) > 0 {
		ret := Descriptor{
			Id:     2,
			ErrMsg: "exit: Illegal number: %s\n" + parameters[0],
		}
		return ret
	} else {
		fmt.Printf("\033[H\033[2J")
		os.Exit(0)
		ret := Descriptor{
			Id:     1,
			OutMsg: "\n",
		}
		return ret
	}
}

// help
// the help command
func help(parameters []string) Descriptor {
	// declare a variable for our loop, after this we concatenate the whole command list
	var list string
	fmt.Println(parameters)
	if len(parameters) > 0 && parameters[0] != "help" {
		// appends just the first index of parameters to avoid errors
		parameters = []string{parameters[0], "--help"}
		cmd := Descriptor{
			Id:     0,
			InMsg:  parameters,
    	}
		// calls IOE to call the desired helper
		return cmd
	} else {
		for word, _ := range cmdMap {
			list = list + word + " "
		}
		// put the whole command list into the output command string
		ret := Descriptor{
			Id:     1,
			OutMsg: "Built-in commands:\n------------------\n\t%s\n" + list,
		}
		return ret
	}
}

func commandChopper(parameters []string) map[int][]string {
	listMap := make(map[int][]string);
	for index, command := range parameters {
		var splt []string = strings.SplitN(command, " ", 2)
		listMap[index] = splt
	}
	return listMap
}
/*
func pipe(parameters []string) Descriptor {
	var helpMessage string = "\nUsage:\t[command1] | [command2]\n";
	if len(parameters) >= 3 {
		ret := Descriptor{
			Id:     2,
			OutMsg: "\nPipe Error: This command takes only two parameters" + helpMessage,
		}
	} else if len(parameters) == 0 {
		// help message of the command
		// return values
		ret := Descriptor{
			Id:     1,
			OutMsg: helpMessage,
		}
	} else {
		commandList := commandChopper(parameters)
		// launch the first command
		var firstCommandResult string = cmdMap[commandList[0][0]](commandList[0][1:]);
		// send result into the second command
		if len(commandList) == 0 {
			var secondCommandResult Descriptor = cmdMap[commandList[1][0]](firstCommandResult.OutMsg);
		} else {
			var secondCommandResult Descriptor = cmdMap[commandList[1][0]](string(commandList[1][1:] + firstCommandResult.OutMsg));
		}
		// return values
		ret := Descriptor{
			Id: 1,
			OutMsg: secondCommandResult,
		}
	}
	return ret
}*/
