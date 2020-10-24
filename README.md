# Protobuffers

- Respositório: https://github.com/protocolbuffers/protobuf
- Instalação:   https://github.com/protocolbuffers/protobuf/blob/master/src/README.md

# Golang Protobuffers:

- Repositório: https://github.com/golang/protobuf
- Compilação: https://developers.google.com/protocol-buffers/docs/reference/go-generated

```cmd
    # Download
    go get -u -v github.com/golang/protobuf/...
```

```cmd
    # Windows install
    go install github.com\golang\protobuf\protoc-gen-go
```

```bash
    # Generate Code
    protoc -I=./messages/defs --go_out=. ./messages/defs/*.proto

    # -I=./messages/defs       -> indica o diretorio onde estão os arquivos .proto
    # --go_out=.               -> indica o diretorio onde serão gerados os arquivos .go, neste caso o diretório está especificado no arquivo .proto: option go_package = "messages/pb";
    # ./messages/defs/*.proto  -> indica quais os arquivos que vão ser gerados, neste caso será todos.
```