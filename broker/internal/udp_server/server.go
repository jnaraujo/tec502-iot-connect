package udp_server

import (
	"broker/internal/cmd"
	"broker/internal/storage/responses"
	"broker/internal/storage/sensors"
	"fmt"
	"log"
	"strconv"
)

// NewServer cria um novo servidor UDP.
// Também é responsável por lidar com as requisições recebidas.
func NewServer(addr string, port int) {
	fmt.Printf("Starting UDP server on %s:%d", addr, port)

	// Cria um novo servidor UDP - ainda sem ouvir
	udpServer := NewUDPServer(fmt.Sprintf("%s:%d", addr, port))
	defer udpServer.Close()

	// Define a função que lida com as mensagens recebidas pelo servidor UDP
	udpServer.HandleRequest(func(addr, content string) {
		cmd, err := cmd.Encode(content) // Decodifica o comando recebido
		if err != nil {
			fmt.Println("Erro ao encodar o comando", err)
			return
		}

		// Verifica se o sensor existe
		if sensors.FindSensorAddrById(cmd.IdFrom) == "" {
			fmt.Printf("O sensor (%s - %s) não foi encontrado\n", addr, cmd.IdFrom)
			return
		}

		// Se uma resposta para o sensor ainda não foi criada, cria uma
		if responses.FindBySensorId(cmd.IdFrom).SensorID == "" {
			responses.Create(cmd.IdFrom, cmd.Command)
		}

		value, err := strconv.ParseFloat(cmd.Content, 64) // Converte o conteúdo do comando para float64
		if err != nil {
			fmt.Println(err)
			return
		}

		responses.AddContent(cmd.IdFrom, value) // Adiciona o conteúdo do comando à resposta do sensor
	})

	err := udpServer.Listen() // Inicia o servidor UDP
	if err != nil {
		log.Fatal("Error starting UDP server:", err)
	}
}
