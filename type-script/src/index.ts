import { WebSocketServer } from "ws";

function main() {
  const wss = new WebSocketServer({ port: 8080 });
  console.log("Hello from ts server");
}
main();
