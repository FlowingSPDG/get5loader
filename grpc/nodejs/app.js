// node app.js --addr 127.0.0.1:50055 --matchid 100 --steamid 76561198072054549

const grpc = require('grpc')
const protoLoader = require('@grpc/proto-loader')
const path = require('path')
const PROTO_PATH = path.resolve(__dirname,"../../server/proto/get5-web-go.proto")
const argv = require('argv');

argv.option({
	  name: 'addr',
	  short: 'a',
	  type : 'string',
	  description :'gRPC target address and port.',
	  example: "127.0.0.1:50055"
})

argv.option({
  name: 'matchid',
  short: 'm',
  type : 'number',
  description :'Streaming API MatchID.',
  example: 0
})

argv.option({
  name: 'steamid',
  short: 's',
  type : 'string',
  description :'User SteamID64',
  example: "76561198072054549"
})

const args = argv.run().targets
 
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
    const client = new Get5Proto.Get5(args[0], grpc.credentials.createInsecure())
    let user_req = Get5Proto.GetUserRequest
    user_req.steamid = args[2]
    await client.GetUser(user_req, function(err, response) {
      if(err){
        console.error(err)
        return
      }
      console.log(response);
    });

    let stream_req = Get5Proto.MatchEventRequest
    stream_req.matchid = args[1]

    let stream = await client.MatchEvent(stream_req);
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