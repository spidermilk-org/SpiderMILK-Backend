export function commandHandler(command) {
    switch(command) {
        case "help":
            return "\"clear\" to clear the screen<br>\"exit\" to exit the game<br>\"help\" to print help message";
        case "clear":
        case "exit":
            return "Don't do that";
        default:
            return "Unknown command";
    }
}