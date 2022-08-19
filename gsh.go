package main

// program imports
import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

// commands lenght structure
type LockedArray [7]string

// alias structure
type Alias struct {
	Name		string
	Substitute	string
}

// environment structure
type Env struct {
	HOSTNAME	string
	HOME		string
	USER		string
	SHELL		string
	PATH		string	
	PAGER		string
	VISUAL		string
	PS1		string
	LANG		string
	PWD		string
}

// main script, it loops into an infinite command prompt.
func main() {

	// environment variables
	envvars := Env{
		HOSTNAME:	"gomputer",
		HOME:		"/home/gopher",
		USER:		"gopher",
		SHELL:		"/bin/ash",
		PATH:		"/usr/sbin:/usr/bin:/sbin:/bin",
		PAGER:		"less",
		VISUAL:		"vi",
		PS1:		"\\h:\\w\\$",
		LANG:		"en_US.UTF-8",
		PWD:		"/home/gopher",
	}

	// command list
	commands := LockedArray {
		"clear",
		"exit",
		"help",
		"hostname",
		"whoami",
		"pwd",
		"alias",
	}

	// program loop
	for {
		var isReal int
		prompted := readingInput(&envvars)
		// splitting the full command to isolate chains, arguments and so on
		splitPrompt := strings.Fields(prompted)
		// incase you just typed 'Enter', it sets 'isReal' to '2'
		if len(splitPrompt) == 0 {
			isReal = 2
		} else if len(splitPrompt) > 0 {
			isReal = isCommandReal(&commands, splitPrompt[0])
		}
		// 2: command is empty just redo the loop
		// 1: command is real launch 'commandEnforcer()'
		// 0: command isn't real print an error
		if isReal == 2 {
			continue;
		} else if isReal == 1 {
			commandEnforcer(&envvars, &commands, splitPrompt)
		} else if isReal == 0 {
			fmt.Printf("%s: %s: not found\n", envvars.SHELL, splitPrompt[0])
		}
	}
}

// input reader, it removes the '\n' at every ends
func readingInput(env *Env) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s:~%s$ ", env.HOSTNAME, env.PWD)
	prompt, _ := reader.ReadString('\n')
	// trim only if the string is longer than one character
	if len(prompt) > 1 {
		// removing '\n' character
		prompt = strings.TrimSuffix(prompt, "\n")
	}

	return prompt
}

// check if the command exists or not
func isCommandReal(arr *LockedArray, command string) int {
	val := 0
	for cmd := range arr {
		if arr[cmd] == command {
			val ++
		}
	}
	return val
}

/*
all commands are added after this comment
if you want to add your own, don't forget to add an 'else if' statement into the 'commandEnforcer()' function
*/

// clear command
func clear(parameters []string) {
	helpMessage := `BusyBox v1.35.0 (2022-08-01 15:14:44 UTC) multi-call binary.

Usage: clear

Clear screen`
	if len(parameters) > 1 {
		fmt.Printf("%s\n", helpMessage)
	} else {
		fmt.Printf("\x1bc")
	}
}

// exit command
func exit() {
	os.Exit(0)
}

// pwd command
func pwd(current string, parameters []string) {
	fmt.Printf("%s\n", current)
}

// help command
func help() {
	fmt.Print(`Built-in commands:
------------------
        . : [ [[ alias bg break cd chdir command continue echo eval exec
        exit export false fg getopts hash help history jobs kill let
        local printf pwd read readonly return set shift source test times
        trap true type ulimit umask unalias unset wait`,"\n")
}

// hostname command
func hostname(name string, parameters []string) {
	fmt.Printf("%s\n", name)
}

// whoami command
func whoami(name string, parameters []string) {
	helpMessage := `BusyBox v1.35.0 (2022-08-01 15:14:44 UTC) multi-call binary.

Usage: whoami

Print the user name associated with the current effective user id`
	if len(parameters) > 1 {
		fmt.Printf("%s\n", helpMessage)
	} else {
		fmt.Printf("%s\n", name)
	}
}

// verify which command was prompted, launch the command
func commandEnforcer(env *Env, arr *LockedArray, chain []string) {
	// declaring the command without the parameters
	cmd := chain[0]
	// conditional segment to call the right func/command
	if cmd == "help" {
		help()
	} else if cmd == "hostname" {
		hostname(env.HOSTNAME, chain)
	} else if cmd == "exit" {
		exit()
	} else if cmd == "whoami" {
		whoami(env.USER, chain)
	} else if cmd == "clear" {
		clear(chain)
	} else if cmd == "pwd" {
		pwd(env.PWD, chain)
	}
}
