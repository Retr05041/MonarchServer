package message

import (
	"strings"
)

func Parse(input string) string {
	if input == "" {
		return ""
	}
	splitInput := strings.Split(input, " ")

	if splitInput[0][0] == ':' {
		handlePrefix(splitInput[0])
        splitInput = splitInput[1:]
	}

    return handleCommand(string(splitInput[0]), splitInput[1:])
}

func handlePrefix(prefix string) {
	return
}

func handleCommand(cmd string, params []string) string {

    fullCmd := cmd
    for _, param := range params {
        fullCmd += " " + param
    }
    return fullCmd

}
