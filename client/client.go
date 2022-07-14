package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() { 
	log.Println("Iniciando chat...")

	log.Println("Escribe un mensaje")

	//Se conecta al servidor
	serverCnn, err := net.Dial("tcp", ":7878")
	if err != nil {
		panic(err)
	}

	//Escucha el servidor

	go readMessage(serverCnn)

	//Enviar mensajes

	writeMessage(serverCnn)
}

func readMessage(conn net.Conn) {
	var scanner = bufio.NewScanner(os. Stdin)

	for scanner.Scan() {
		//Una vez cada vez que haga enter
		var texto = scanner.Text()

		_, err := conn.Write([]byte(texto))
		if err != nil {
			panic(err)
		}
	}
}

func writeMessage(cnn net.Conn) {
	//Leer mensajes

	for {
		var mensaje = make([]byte, 1000)

		_, err := cnn.Read(mensaje)
		if err != nil {
			panic(err)
		}

		//Imprimir mensaje

		fmt.Println("El mensaje del servidor es:", string(mensaje))
	}
}