import { createServer } from "http";
import dotenv from "dotenv";
dotenv.config();
import { commandHandler } from "./commands.mjs";

const server = createServer((req, res) => {
    let body = "";
    req.on("data", (chunk) => {
        body += chunk.toString();
    });
    req.on("end", () => {
        handleRequest(res, body);
    });
});

function handleRequest(res, body) {
    const { action, jwt } = JSON.parse(body);
    res.writeHead(200, {
        "Content-Type": "application/json",
        "Access-Control-Allow-Headers": "Content-Type",
        "Access-Control-Allow-Origin": "*"
    });
    res.end(JSON.stringify({
        responseText: commandHandler(action)
    }));
}

server.listen(process.env.PORT);
