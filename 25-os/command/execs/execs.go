package main

import "syscall"
import "os"
import "os/exec"

func main() {
	// dir()
	cmd() //calls command.go
}
func cmd() {
	binary, lookErr := exec.LookPath("../command")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"../command", "-b=true", "-f=59.95"}
	env := os.Environ()
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}

}
func dir() {
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"ls", "-a", "-l", "-h"}

	// `Exec` also needs a set of [environment variables](environment-variables)
	// to use. Here we just provide our current
	// environment.
	env := os.Environ()

	// If this call is successful, the execution of our process will end
	// here and be replaced by the /bin/ls -a -l -h
	// process. If there is an error we'll get a return value.
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}

}
