package main
import(
	"fmt"
	"net"
	"os"
	"bufio"
)

func main(){
    fmt.Println("This the client send stuff to server")
	connect,err:=net.Dial("tcp","localhost:5000")
	if(err!=nil){
		fmt.Println("Error dailing",err.Error())
		return
	}
	fmt.Println("Start sending to the server")

	//CREATE CLIENT SERVER PART
	listener, err2 := net.Listen("tcp", "localhost:5050")
	if(err2!=nil){
		fmt.Println("Error lunching the server in client",err.Error())
		return
	}


	for{
		inputReader := bufio.NewReader(os.Stdin)
		message, _ := inputReader.ReadString('\n')
		_,err:=connect.Write([]byte(message))
		if(err!=nil){
			fmt.Println("Error sending",err.Error())
			return
		}
		connection,err:=listener.Accept()
    	if(err!=nil){
    		fmt.Printf("Error accepting",err.Error())
    		return
    	}
		go HandleRequests(connection)
	}
}

func HandleRequests(conn net.Conn) {
	for {
		buf:=make([]byte,512)
		_, err := conn.Read(buf)
		if (err != nil) {
			fmt.Println("Error reading", err.Error())
			return 
		} else {
			fmt.Printf("Received data: %v", string(buf))
		}
	}
}