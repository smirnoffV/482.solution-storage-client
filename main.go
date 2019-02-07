package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const EXIT_COMMAND = "exit"

func main() {

	host := flag.String("h", "127.0.0.1", "storage server host")
	port := flag.String("p", "8080", "storage server port")

	flag.Parse()

	fmt.Println("Launching client...")

	conn, err := net.Dial("tcp", net.JoinHostPort(*host, *port))
	if err != nil {
		log.Println("error connecting to server")
		os.Exit(0)
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Input your command: ")

		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading from input")
			continue
		}

		command = strings.Replace(command, "\r", "", -1)
		command = strings.Replace(command, "\n", "", -1)

		if command == EXIT_COMMAND {
			conn.Close()
			os.Exit(1)
		}

		command = BuildRequest(command)
		fmt.Fprintf(conn, command+"\n")

		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("error reading response")
			break
		}
		fmt.Println(FormatResponse(message))
	}
}

func BuildRequest(input string) string {
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
