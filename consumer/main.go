package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:2222")
	if err != nil {
		fmt.Println("Failed to connect to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Receive message from the server
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Failed to read message:", err)
		os.Exit(1)
	}

	// Print the received message
	fmt.Println("Message from server:", string(buffer[:n]))
}

// func main() {
// 	host := "127.0.0.1:2222"
// 	user := "<username>"
// 	pwd := "<password>"
// 	pKey := []byte("<privateKey>")
//
// 	var err error
// 	var signer ssh.Signer
//
// 	signer, err = ssh.ParsePrivateKey(pKey)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
//
// 	conf := &ssh.ClientConfig{
// 		User:            user,
// 		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
// 		Auth: []ssh.AuthMethod{
// 			ssh.Password(pwd),
// 			ssh.PublicKeys(signer),
// 		},
// 	}
//
// 	var conn *ssh.Client
//
// 	conn, err = ssh.Dial("tcp", host, conf)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	defer conn.Close()
//
// 	var session *ssh.Session
// 	var stdin io.WriteCloser
// 	var stdout, stderr io.Reader
//
// 	session, err = conn.NewSession()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	defer session.Close()
//
// 	stdin, err = session.StdinPipe()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
//
// 	stdout, err = session.StdoutPipe()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
//
// 	stderr, err = session.StderrPipe()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
//
// 	wr := make(chan []byte, 10)
//
// 	go func() {
// 		for {
// 			select {
// 			case d := <-wr:
// 				_, err := stdin.Write(d)
// 				if err != nil {
// 					fmt.Println(err.Error())
// 				}
// 			}
// 		}
// 	}()
//
// 	go func() {
// 		scanner := bufio.NewScanner(stdout)
// 		for {
// 			if tkn := scanner.Scan(); tkn {
// 				rcv := scanner.Bytes()
//
// 				raw := make([]byte, len(rcv))
// 				copy(raw, rcv)
//
// 				fmt.Println(string(raw))
// 			} else {
// 				if scanner.Err() != nil {
// 					fmt.Println(scanner.Err())
// 				} else {
// 					fmt.Println("io.EOF")
// 				}
// 				return
// 			}
// 		}
// 	}()
//
// 	go func() {
// 		scanner := bufio.NewScanner(stderr)
//
// 		for scanner.Scan() {
// 			fmt.Println(scanner.Text())
// 		}
// 	}()
//
// 	session.Shell()
//
// 	for {
// 		fmt.Println("$")
//
// 		scanner := bufio.NewScanner(os.Stdin)
// 		scanner.Scan()
// 		text := scanner.Text()
//
// 		wr <- []byte(text + "\n")
// 	}
// }
