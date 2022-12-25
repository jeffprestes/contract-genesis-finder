# Buscador do numero do bloco onde um Contrato Inteligente (Smart Contract) foi criado

## Como usar

Após clonar esse repo a sua máquina e tendo o compilador Go instalado, crie um arquivo chamado *rpcserver.txt* e dentro dele coloque a URL do servidor RPC que você irá utilizar para se conectar a blockchain.

**Importante** : Este servidor deve estar conectado a mesma rede Blockchain onde você acredita estar o Contrato Inteligente.

Depois edite o arquivo main.go com o endereço do contrato ou dos endereços dos contratos que você quer obter a informação.

Por fim, compile e execute ou apenas execute `go run *.go`

