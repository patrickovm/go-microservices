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
grpcurl --plaintext -msg-template localhost:50051 describe protos.ListTrainersRequest
```

returns
```
protos.ListTrainersRequest is a message:
message ListTrainersRequest {
}

Message template:
{
  
}
```
<br/>

## Execute a request for Create

```sh
grpcurl --plaintext -d '{"firstName": "John", "lastName": "Doe", "age": 21, "finishedGames": 6}' localhost:50051 protos.TrainersService.Create
```
returns
```
{
  "id": "ObjectID(\"62d7178cb910060d33849b39\")"
}
```
<br/>

## Execute a request for Get

```sh
grpcurl --plaintext -d '{"id": "62d7178cb910060d33849b39"}' localhost:50051 protos.TrainersService.Get
```

returns
```
{
  "id": "62d7178cb910060d33849b39",
  "firstName": "John",
  "lastName": "Doe",
  "age": 21,
  "finishedGames": 6
}
```
<br/>

## Execute a request for List

```sh
grpcurl --plaintext -d '{}' localhost:50051 protos.TrainersService.List
```

returns
```
{
  "trainers": [
    {
      "id": "62d7178cb910060d33849b39",
      "firstName": "John",
      "lastName": "Doe",
      "age": 21,
      "finishedGames": 6
    },
    {
      "id": "62d718a7b910060d33849b3a",
      "firstName": "Jack",
      "lastName": "Doe",
      "age": 21,
      "finishedGames": 6
    }
  ]
}
```

## Execute a request for Update

```sh
grpcurl --plaintext -d '{"id": "62d718a7b910060d33849b3a", "firstName": "Jack", "lastName": "Sparrow", "age": 21, "finishedGames": 2}' localhost:50051 protos.TrainersService.Update
```

returns
```
{
  "id": "62d718a7b910060d33849b3a",
  "firstName": "Jack",
  "lastName": "Sparrow",
  "age": 21,
  "finishedGames": 2
}
```

