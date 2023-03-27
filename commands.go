package main

func commandHandler(command string) string {
	switch(command) {
		case "help":
		case "clear":
		case "exit":
			return "Don't do that"
	}
	return "Command not found"
}
