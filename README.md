# RegisterService(test)
This service is designed to register with the Stalcraft API

The service takes on the task of user authorization separately from your main application.
With gRPC, you establish a connection to it and transfer data.

! The service must be located at a separate address, this address is sent by the link that you send to users in the redirect_uri parameter. You also have to generate the uuid for the parameter state by reference yourself! Link example: https://exbo.net/oauth/authorize?client_id=[your_id_client]&redirect_uri=[your_link_on_service_register]&scope=&response_type=code&state=[uuid] 

The service accepts a nickname, state and region as input.
Returns a token and refresh  token.

the nickname you get from the user. The service compares the nickname and nickname associated with the token, if they match, then the user is the one whose nickname introduced himself.

region you also get from the user.

To use this service, you will need minimal experience with gRPS, docker and ideally golang, but you can do without it.
For a quick start, use a ready-made dockerfile.

It is enough to change the line in it ```CMD ["./main", "", "", ""]```

Add the following parameters to the quotes:
1)client_id
2)client_secret
3)redirect_uri

You get them from the telegram bot Starcraft API.

to launch docker image:
1) ```sudo docker build -t regservice .```
2) ```docker run -p 3001:3001 -p 3002:3002 regservice```
  
proto file - ```protoc  --go_out=api/grpc --go-grpc_out=api/grpc api/proto/reg.proto```

If there is a gRPC error, try using ```<export PATH="$PATH:$(go env GOPATH)/bin">```

This service uses two handlers on ports 3001(tcp) and 3002(http). 
If desired, they can be changed.
3001 - Handles requests from your application.
3002 - Processes user requests.


The service can return errors inside the token and update the token.
404 is a default error.
303 - the nickname sent did not match the nickname of the token owner.
101 - the user's waiting time has been exceeded (the service is waiting for the user to register, this is given 1 minute).


Code customization:
If you know golang or you are not afraid of gRPC, you can modify the service:
1) Change the received and sent data
2) User waiting time 
3) ports (it's easy here)
4) Any other things :)