package application

import (
	"os"
	"fmt"
)

const MEMO_INPUT_LEN = 3
const SAVE_COMMAND_INPUT_POS = 1
const MEMO_INPUT_POS = 2
const SAVE_COMMAND = "save"

const CONFIG_INPUT_LEN = 4
const SET_INPUT_POS = 1
const SET_COMMAND_INPUT_POS = 2
const COMMAND_CONTENT_INPUT_POS = 3
const SET_COMMAND = "set"
const SET_API_COMMAND = "api"

// Handle is the entry of the cli.
func Handle() {
	input := os.Args
	length := len(input)
	switch length {
	case MEMO_INPUT_LEN:
		if input[SAVE_COMMAND_INPUT_POS] != SAVE_COMMAND {
			fmt.Println("Invalid command, use save to send a memo to flomo.")
			return
		}
		memo := input[MEMO_INPUT_POS]
		if len(memo) == 0 {
			fmt.Println("The memo you typed is blank, try to say some thing.")
			return
		}
		SendMemo(memo)
	case CONFIG_INPUT_LEN:
		if input[SET_INPUT_POS] != SET_COMMAND {
			fmt.Println("Invalid command, use set to configure the flomo-cli.")
			return
		}
		if input[SET_COMMAND_INPUT_POS] == SET_API_COMMAND {
			commandContent := input[COMMAND_CONTENT_INPUT_POS]
			if commandContent == "" {
				fmt.Println("The command content is blank, try [flomo set api].")
			}
			SaveConfig(FlomoConfig{Api: commandContent})
			return
		}
		fmt.Println("Nothing to set, try [flomo set api].")
	default:
		fmt.Println("Invalid input. Use set to configure the cli or use save to type a new memo.")
	}
}
