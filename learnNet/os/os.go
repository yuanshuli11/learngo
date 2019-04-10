package main

import (
	"fmt"
	"os"
)

func main(){
	Hostname,_ := os.Hostname()
	fmt.Printf("Hostname: %q\n", Hostname)

	os.Setenv("TEST_MODE","test")
	env := os.Getenv("TEST_MODE")
	fmt.Printf("env: %q\n", env)

	//获取进程id
	pid :=os.Getpid()
	fmt.Printf("pid: %d\n", pid)
	//父进程id
	ppid :=os.Getppid()
	fmt.Printf("ppid: %d\n", ppid)
}