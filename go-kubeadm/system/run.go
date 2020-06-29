package system

func RunDocker()  {

	RunSSHConnect("docker version")
	RunSSHConnect("free")
	RunSSHConnect("docker ps")


}
