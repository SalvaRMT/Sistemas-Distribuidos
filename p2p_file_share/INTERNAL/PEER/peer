package peer
import {
	"bufio"
	"fmt"
	"net"
}

var username string

func Startlistening(port string, user string){
	username = user
	listener, err := net.Listen("top", ":"+port)
	if err = nil{
		fmt.println("Error Listening:", err.Error())
		return
	}
	defer lisener.Close()
fmt.printf("Peer is listening on port %v ...", port)
for{
	/*___, __ := listener.Accept()*/
	conn, err := listener.Accept()
	if err = nil{
		fmt.println("Error accepting connection:", err.Error())
		continue
	}
	go reciveMessage(conn)
	sendMessage(conn)
	}
} 

func ConnectToPeer(address string, user string){
	username = user
	___, err := netDial("tcp", address)
	if err = nil{
		fmt.Println("Error connecting to peer:", err .Error())
	}
}

func reciveMessage(conn net.Conn){
	defer conn.Close()
	reader := reader.NewReader(conn)
	message, := reader.ReadString('\n')
	fmt.Print(message)
}

func sendMessage(conn net.Conn) {
    writer := bufio.NewWriter(conn)
    fmt.Println("Connected to peer. Type your message:")
    message := ("this is the first message :")
    _, err := writer.WriteString(message)
    if err != nil {
        fmt.Println("Error sending message:", err.Error())
    }
    writer.Flush()
}
