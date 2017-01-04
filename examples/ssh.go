package main

import (
	"bufio"
	gexpect "github.com/ThomasRooney/gexpect"
	"golang.org/x/crypto/ssh"
	"log"
)

func main() {
	log.Printf("Testing SSH... ")

	client, err := ssh.Dial("tcp", "alt.org:22", &ssh.ClientConfig{
		User: "nethack",
	})
	if err != nil {
		log.Fatalf("Failed to connect to SSH server: %s", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
	}
	defer session.Close()

	sshOut, err := session.StdoutPipe()
	if err != nil {
		log.Fatalf("Failed to get stdout pipe: %s", err)
	}

	sshIn, err := session.StdinPipe()
	if err != nil {
		log.Fatalf("Failed to get stdin pipe: %s", err)
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 38400,
		ssh.TTY_OP_OSPEED: 38400,
	}

	if err := session.RequestPty("xterm", 0, 200, modes); err != nil {
		log.Fatalf("Request for pty failed: %s", err)
	}

	if err := session.Shell(); err != nil {
		log.Fatalf("Request for shell failed: %s", err)
	}

	exp := gexpect.NewExpectIO(bufio.NewReader(sshOut), bufio.NewWriter(sshIn))

	exp.Expect("=> ")
	exp.Send("s")
	exp.Expect("This is the nethack.alt.org public NetHack server.")
	exp.Expect("=> ")
	exp.Send("q")
	exp.Expect("=> ")
	exp.Send("q")
	log.Printf("Success\n")
}
