# Makefile
.EXPORT_ALL_VARIABLES:	

GO111MODULE=on
GOPROXY=direct
GOSUMDB=off

run:
	@echo "########## Run subscriber-storage-manager"
	@go run cmd/main.go

build:
	@echo "########## Build subscriber-storage-manager "
	@go build -trimpath -ldflags="-s -w" -o cmd/cadastro_funcionario cmd/main.go
	@echo "buid completo..."

#test:
#	@echo "########## Executando Tests"
#	@sleep 1
#	@go test -gcflags=-l gitlab.engdb.com.br/pmid/newfast/subscriber-storage-manager/internal/pkg/repository/mongo
#
#test-v:
#	@echo "########## Executando Tests"
#	@sleep 1
#	@go test -gcflags=-l gitlab.engdb.com.br/pmid/newfast/subscriber-storage-manager/internal/pkg/repository/mongo -v
