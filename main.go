package main

import (
	"log"
	"os"
)

func main() {
	log.Println("Obtendo definição de servidor RPC...")
	tmp, err := os.ReadFile("rpcserver.txt")
	if err != nil {
		log.Fatalf("[main] Could not read rpcserver due to this %s error \n", err)
	}
	rpcServer := cleanContentBytes(tmp)
	log.Println("Iniciando as buscas através de:", rpcServer)
	cFinder, err := NewContractFinder(rpcServer)
	if err != nil {
		log.Fatalln("Erro ao criar o buscador:", err.Error())
	}

	contratosABuscar := []string{"0x31Ac27d901912A21747979526e96CC8F0Be67130"}

	for _, contrato := range contratosABuscar {
		log.Println("Iniciando as buscas do contrato", contrato)
		blocoDeCriacao, err := cFinder.GetContractCreationBlock(contrato)
		if err != nil {
			log.Println("Erro", err.Error(), " ao buscar o bloco de criacao do contrato", contrato)
			continue
		}
		log.Println("O contrato", contrato, "foi criado no bloco", blocoDeCriacao)
	}
}
