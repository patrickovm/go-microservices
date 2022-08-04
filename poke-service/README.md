## List services
```sh
grpcurl --plaintext localhost:50052 list
```
returns
```
grpc.reflection.v1alpha.ServerReflection
protos.PokeService
```
<br/>

## List methods

```sh
grpcurl --plaintext localhost:50052 list protos.PokeService
```

returns
```
protos.PokeService.Create
protos.PokeService.GetAllFromOneTrainer
protos.PokeService.Update
```
<br/>

## Message details for CreatePokemonRequest

```sh
grpcurl --plaintext -msg-template localhost:50052 describe protos.CreatePokemonRequest
```

returns
```
protos.CreatePokemonRequest is a message:
message CreatePokemonRequest {
  string trainer_id = 1;
  string name = 2;
  bool has_badge = 3;
}

Message template:
{
  "trainerId": "",
  "name": "",
  "hasBadge": false
}
```
<br/>

## Message details for GetAllFromOneTrainerRequest

```sh
grpcurl --plaintext -msg-template localhost:50052 describe protos.GetAllFromOneTrainerRequest
```

returns
```
protos.GetAllFromOneTrainerRequest is a message:
message GetAllFromOneTrainerRequest {
  string trainer_id = 1;
}

Message template:
{
  "trainerId": ""
}
```
<br/>

## Message details for UpdatePokemonRequest

```sh
grpcurl --plaintext -msg-template localhost:50052 describe protos.UpdatePokemonRequest
```

returns
```
protos.UpdatePokemonRequest is a message:
message UpdatePokemonRequest {
  int64 id = 1;
  string trainer_id = 2;
  string name = 3;
  bool has_badge = 4;
}

Message template:
{
  "id": "0",
  "trainerId": "",
  "name": "",
  "hasBadge": false
}
```
<br/>

## Execute a request for Create

```sh
grpcurl --plaintext -d '{"trainerId": "62d7178cb910060d33849b39", "name": "Machamp", "hasBadge": true}' localhost:50052 protos.PokeService.Create
```

returns
```
{
  "id": "1"
}
```
<br/>

## Execute a request for GetAllFromOneTrainer
```sh
grpcurl --plaintext -d '{"trainerId": "62d7178cb910060d33849b39"}' localhost:50052 protos.PokeService.GetAllFromOneTrainer
```

returns
```
{
  "pokelist": [
    {
      "id": "1",
      "trainerId": "62d7178cb910060d33849b39",
      "name": "Machamp",
      "hasBadge": true
    },
    {
      "id": "2",
      "trainerId": "62d7178cb910060d33849b39",
      "name": "Pichu"
    }
  ]
}
```
<br/>

## Execute a request for Update

```sh
grpcurl --plaintext -d '{"id": "1", "trainerId": "62d7178cb910060d33849b39", "name": "Pichu", "hasBadge": true}' localhost:50052 protos.PokeService.Update
```

returns
```
{
  "trainerId": "62d7178cb910060d33849b39",
  "name": "Pichu",
  "hasBadge": true
}
```
