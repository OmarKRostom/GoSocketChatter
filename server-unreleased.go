package main
import(
	"fmt"
	"net"
	"os"
	"bufio"
	)

func main(){
	fmt.Println("WELCOME TO THE CHATTER 1.0")
	listener, err := net.Listen("tcp", "localhost:5000")
	if(err!=nil){
		fmt.Println("Error lunching the server",err.Error())
		return
	}
	
    for{
    	connection,err:=listener.Accept()
    	if(err!=nil){
    		fmt.Printf("Error accepting",err.Error())
    		return
    	}
    	go handleRequests(connection)
    }
}

func handleRequests(conn net.Conn){
	var HASRECIEVED bool = false
	for{
		

		buf:=make([]byte,512)
		_, err := conn.Read(buf)
		if (err != nil) {
			fmt.Println("Error reading", err.Error())
			return 
		} else {
			fmt.Printf("Received data: %v", string(buf))
			HASRECIEVED = true
		}

		if HASRECIEVED {
			connect,err2:=net.Dial("tcp","localhost:5050")
			if(err2!=nil) {
				fmt.Println("Error Connecting to server")
			}
			inputReader := bufio.NewReader(os.Stdin)
			message, _ := inputReader.ReadString('\n')
			_,err:=connect.Write([]byte(message))
			if(err!=nil){
				fmt.Println("Error sending",err.Error())
				return
			}
		}
	} 
}