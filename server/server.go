package main

import (
	"log"
	"net"
)

func main() {

	//Incia el servidor
	server, err := net.Listen("tcp", ":7878")

	if err != nil {
		panic(err)
	} else {
		log.Println("Servidor conectado localhost:7878")
	}

	contadorClientes := 0

	for {
		log.Println("Esperando un cliente...")
		contadorClientes++

		//Se conecta un cliente

		client, err := server.Accept()
		if err != nil {
			panic(err)
		}

		if contadorClientes < 2 {
			log.Println("Hay", contadorClientes, "usario conectado")
		} 
		if contadorClientes >= 2 {
			log.Println("Hay", contadorClientes, "usarios conectados")
		}

		//Administrador de coneccion
		go managerconnection(client)
	}
}

func managerconnection(cliente net.Conn)  {
	userConnection = append(userConnection ,cliente)

	for {
		//Leemos el mensaje
		var buff = make([]byte, 1000)

		_, err := cliente.Read(buff)
		if err != nil {
			panic(err)
		}

		//Le mandamos el mansaje al usuario
		writeMessageAllUsers(buff)
	}
}

func writeMessageAllUsers(mensaje []byte)  {
	for _, c := range userConnection {
		c.Write(mensaje)
	}
}