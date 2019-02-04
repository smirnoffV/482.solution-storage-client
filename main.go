package main

import (
	"482.solutions-node-storage-client/client"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	config := client.NewConfiguration()

	fmt.Println("Launching client...")

	conn, err := net.Dial("tcp", net.JoinHostPort(config.StorageHost, config.StoragePort))
	if err != nil {
		panic("error connecting to server")
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Input your command: ")

		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading from input")
			continue
		}

		command = BuildRequest(command)

		fmt.Fprintf(conn, command+"\n")

		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("error reading response")
			continue
		}
		fmt.Println(FormatResponse(message))
	}
}

func BuildRequest(input string) string {
	input = strings.TrimSuffix(input, "\n")
	sep := strings.Split(input, " ")

	inpStr := &Input{
		Method: sep[0],
	}

	if isset(sep, 1) {
		inpStr.Key = sep[1]
	}

	if isset(sep, 2) {
		inpStr.Value = sep[2]
	}

	return inpStr.BuildMessage()
}

func FormatResponse(message string) string {
	message = strings.TrimSuffix(message, "\n")
	var response Output
	if err := json.Unmarshal([]byte(message), &response); err != nil {
		fmt.Println("error reading response")
	}
	return response.Value
}

type Input struct {
	Method string `json:"-"`
	Key    string `json:"key"`
	Value  string `json:"value"`
}

func (i *Input) BuildMessage() string {

	body, err := json.Marshal(i)

	if err != nil {
		log.Println("error processing request")
	}

	return fmt.Sprintf("%s||%s", i.Method, body)
}

type Output struct {
	Value string `json:"value"`
}

func isset(arr []string, index int) bool {
	return (len(arr) > index)
}
