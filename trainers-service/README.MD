## List services
```sh
grpcurl --plaintext localhost:50051 list
```
returns
```
grpc.reflection.v1alpha.ServerReflection
protos.TrainersService
```
<br/>

## List methods

```sh
grpcurl --plaintext localhost:50051 list protos.TrainersService
```

returns
```
protos.TrainersService.Create
protos.TrainersService.Get
protos.TrainersService.List
protos.TrainersService.Update
```
<br/>

## Method details for GetTrainer

```sh
grpcurl --plaintext localhost:50051 describe protos.TrainersService.Get
```

returns
```
protos.TrainersService.Get is a method:
rpc Get ( .protos.GetTrainerRequest ) returns ( .protos.GetTrainerResponse );
```
<br/>

## Message details for CreateTrainerRequest

```sh
grpcurl --plaintext -msg-template localhost:50051 describe protos.CreateTrainerRequest
```

returns
```
protos.CreateTrainerRequest is a message:
message CreateTrainerRequest {
  string first_name = 1;
  string last_name = 2;
  int32 age = 3;
  int32 finished_games = 4;
}

Message template:
{
  "firstName": "",
  "lastName": "",
  "age": 0,
  "finishedGames": 0
}
```
<br/>

## Message details for ListTrainersRequest

```sh
grpcurl --plaintext -msg-template localhost:50051 describe pb.ListTrainersRequest
```

returns
```
protos.ListUsersRequest is a message:
message ListUsersRequest {
}

Message template:
{
  
}
```
<br/>

## Execute a request for CreateUser

```sh
grpcurl --plaintext -d '{"firstName": "John", "lastName": "Doe", "age": 21, "finishedGames": 6}' localhost:50051 protos.UserService.CreateUser
```
returns
```
{
  "id": "ObjectID(\"62d6b76b51868410b2c4c48c\")"
}
```
<br/>

## Execute a request for GetUser

```sh
grpcurl --plaintext -d '{"id": "62d6b76b51868410b2c4c48c"}' localhost:50051 protos.UserService.GetUser
```

returns
```
{
  "id": "62d6b76b51868410b2c4c48c",
  "firstName": "John",
  "lastName": "Doe",
  "age": 21,
  "finishedGames": 6
}
```
<br/>

## Execute a request for ListUsers

```sh
grpcurl --plaintext -d '{}' localhost:50051 protos.UserService.ListUsers
```

returns
```
{
  "users": [
    {
      "id": "62d6b76b51868410b2c4c48c",
      "firstName": "John",
      "lastName": "Doe",
      "age": 21,
      "finishedGames": 6
    },
    {
      "id": "62d6b8cd51868410b2c4c48d",
      "firstName": "Jack",
      "lastName": "Doe",
      "age": 50,
      "finishedGames": 5
    }
  ]
}
```

## Execute a request for UpdateUser

```sh
grpcurl --plaintext -d '{"id": "62d6b8cd51868410b2c4c48d", "firstName": "Jack", "lastName": "Sparrow", "age": 21, "finishedGames": 2}' localhost:50051 protos.UserService.UpdateUser
```

returns
```
{
  "id": "62d6b8cd51868410b2c4c48d", 
  "firstName": "Jack", 
  "lastName": "Sparrow", 
  "age": 21, 
  "finishedGames": 2
}
```

