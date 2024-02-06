package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: solo <PORT> <COMMAND>")
		os.Exit(1)
	}

	port := os.Args[1]
	if !isValidPort(port) {
		fmt.Println("Port invalide:", port)
		os.Exit(1)
	}

	ln, err := net.Listen("tcp", "127.0.0.1:"+port)
	if err != nil {
		fmt.Println("Impossible de se lier au port, une autre instance semble être en cours d'exécution")
		os.Exit(2)
	}
	defer ln.Close()

	command := os.Args[2:]
	cmd := exec.Command("/bin/bash", "-c", command[0]) //, command[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println("Erreur lors de l'exécution de la commande:", err)
		os.Exit(1)
	}
}

func isValidPort(port string) bool {
	p, err := strconv.Atoi(port)
	return err == nil && p > 0 && p < 65536
}


