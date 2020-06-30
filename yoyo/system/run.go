package system


/**
	批量执行
 */
func RunDocker()  {

	RunSSHConnect("docker version")
	RunSSHConnect("free")
	RunSSHConnect("docker ps")

}
