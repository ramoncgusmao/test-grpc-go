para iniciar o projeto precisa ter o o go instalado

precisa instalar o protobuf-compiler tbm
```
sudo apt install protobuf-compiler 
```
apos usar o comando

```
go mod init github.com/<nome_usuario>/<nome_projeto>

```

apos adicionar e instalar os pacotes gRPC

```
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

para crias os proto stubs rodar o comando

```
protoc --proto_path=proto/ proto/*.proto --plugin=$(go env GOPATH)/bin/protoc-gen-go-grpc --go-grpc_out=. --go_out=.
```

caso der erro executar esse comando no terminal

```
export PATH="$PATH:$(go env GOPATH)/bin"
```