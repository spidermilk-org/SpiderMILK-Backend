import {createServer} from "http";
import dotenv from "dotenv";
dotenv.config();

const server = createServer((req, res) => {
    res.writeHead(200, {
        "Content-Type": "application/json",
        "Access-Control-Allow-Headers": "Content-Type",
        "Access-Control-Allow-Origin": "*"
    });
    res.end(JSON.stringify({
        responseText: "Hello World"
    }));
});

server.listen(process.env.PORT);
