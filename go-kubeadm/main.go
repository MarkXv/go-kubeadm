package main

import (
	"fmt"
	"go-kubeadm/system"
)


func main() {
	fmt.Println("go-kubeadm")

	//system.RunSSHConnect()
	system.RunSSHConnect("systemctl start docker ")
	//for i := 0; i < 10; i ++ {
	//	system.RunSSHConnect("docker ps" );
	//}

	system.RunDocker()




}

