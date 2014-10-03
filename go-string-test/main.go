package main

import (
	"fmt"
	"strings"
)

/*
   Arguments in the functions below may be an empty array or an array containing one array of strings.
   This is used only to state that the string array can be optionnal (some commands won't take any arguments)
   For commands who don't take arguments, just ignore them.
*/

var commandMapping = map[string]func(...[]string){
	"command": func(arguments ...[]string) {
		if len(arguments) > 0 {
			fmt.Println("I have arguments.")
			// Further argument management, command specific
			for _, element := range arguments[0] {
				fmt.Println(element)
			}
		} else {
			// No optionnal arguments provided, different behaviour
			// Example : AFK Command on b0tsec
			fmt.Println("I have no arguments.")
		}
	},
}

func main() {
	input := ""
	if _, err := fmt.Scan(&input); err != nil {
		panic(err)
	}
	if strings.HasPrefix(input, "!") {
		commandArray := strings.Fields(input[1:])
		command := commandArray[0]
		if commandCallback, ok := commandMapping[command]; ok {
			if len(commandArray) > 1 {
				fmt.Println("With arguments : ")
				commandCallback(commandArray[1:])
			}
			fmt.Println("Without arguments :")
			commandCallback()
		} else {
			fmt.Println("This command was not found")
		}
	}
}
