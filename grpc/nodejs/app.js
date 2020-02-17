const grpc = require('grpc')
const protoLoader = require('@grpc/proto-loader')
const path = require('path')
const PROTO_PATH = path.resolve(__dirname,"../../server/proto/get5-web-go.proto")
 
const packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    }
)

async function main() {
    const Get5Proto = grpc.loadPackageDefinition(packageDefinition)
    const client = new Get5Proto.Get5('127.0.0.1:50055', grpc.credentials.createInsecure())
    let user_req = Get5Proto.GetUserRequest
    user_req.steamid = "76561198072054549"
    client.GetUser(user_req, function(err, response) {
      if(err){
        console.error(err)
        return
      }
      console.log(response);
    });

    let stream_req = Get5Proto.MatchEventRequest
    stream_req.matchid = 100

    let stream = client.MatchEvent(stream_req);
    stream.on('data', function(data) {
      console.log(data)
    });
    stream.on('end', function() {
      // The server has finished sending
    });
    stream.on('error', function(e) {
      // An error has occurred and the stream has been closed.
    });
    stream.on('status', function(status) {
      console.log(status)
    });
  }
  
  main();