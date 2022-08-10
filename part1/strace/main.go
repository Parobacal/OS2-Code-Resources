package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"bufio"
	"strings"
)

func main() {
	fmt.Println("")
	fmt.Println("################### STRACE SYSTEM ###################")
	fmt.Println("|----------------------------------------------------|")
	fmt.Println("|    Enter command or exit to finish the execution:  |")
	fmt.Println("|----------------------------------------------------|")
	fmt.Println("")
	for {
		fmt.Print(">> ")
		com := bufio.NewScanner(os.Stdin)
		if com.Scan() {
			if com.Text() == "exit" {
				break
			} else {
				strace(strings.Fields(com.Text()))
			}
		}
	}
}

func strace(command []string){
	var regs syscall.PtraceRegs
	var ss syscallCounter

	ss = ss.init()

	cmd := exec.Command(command[0], command[1])
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Ptrace: true,
	}

	cmd.Start()
	err := cmd.Wait()
	if err != nil {
		fmt.Printf("Wait returned: %v\n", err)
	}

	pid := cmd.Process.Pid
	exit := true

	for {
		if exit {
			err = syscall.PtraceGetRegs(pid, &regs)
			if err != nil {
				break
			}
			ss.inc(regs.Orig_rax)
		}

		err = syscall.PtraceSyscall(pid, 0)
		if err != nil {
			panic(err)
		}

		_, err = syscall.Wait4(pid, nil, 0, nil)
		if err != nil {
			panic(err)
		}

		exit = !exit
	}

	ss.print()
}
