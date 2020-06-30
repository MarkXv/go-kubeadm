package system

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)
import (
	ssh "golang.org/x/crypto/ssh"

)

func SSHConnect(user, password, host string, port int) (*ssh.Session, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)

	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	hostKeyCallback := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}

	clientConfig = &ssh.ClientConfig{
		Config:            ssh.Config{},
		User:              user,
		Auth:              auth,
		HostKeyCallback:   hostKeyCallback,

	}

	addr = fmt.Sprintf("%s:%d", host, port)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	return session, nil
}
func RunSSHConnect(cmd string) {
	var stdOut, stdErr bytes.Buffer
	session, err :=	SSHConnect("root", "xv13294600841.." , "49.234.198.239", 22)

	if err != nil {
		log.Fatal(err)
	}
	defer func() {


		session.Close()
		//recover ()

	}()

	session.Stdout = &stdOut
	session.Stderr = &stdErr

	//session.Run("docker run -dit nginx")
	session.Run(cmd)

	fmt.Println(stdOut.String()+"\n")


	ret, err := strconv.Atoi(strings.Replace(stdOut.String(),"\n", "" ,  -1) )
	if err != nil {
		//抛出异常
		fmt.Println(err)
		//panic(err)
	}

	fmt.Printf("%d,%s\n" , ret , stdErr.String())

}

