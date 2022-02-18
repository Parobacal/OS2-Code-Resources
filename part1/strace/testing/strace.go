package main

import (
	"fmt"
	"syscall"
)

func main() {
	var err error
	var wstat syscall.WaitStatus
	var regs syscall.PtraceRegs
	var ss syscallCounter
	ss = ss.init()

	var pid = 57076
	exit := true

	erx := syscall.PtraceAttach(pid)
	if err != nil {
		fmt.Print("Attach")
		fmt.Print(erx)
	}

	_, err = syscall.Wait4(pid, &wstat, 0, nil)
	if err != nil {
		fmt.Printf("wait %d err %s\n", pid, err)
		fmt.Println(err)
	}

	err = syscall.PtraceSetOptions(pid, syscall.PTRACE_O_TRACESYSGOOD)
	if err != nil {
		fmt.Println("Ptrace set options")
		panic(err)
	}

	for {
		if exit {
			err = syscall.PtraceGetRegs(pid, &regs)
			if err != nil {
				break
			}
			name := ss.getName(regs.Orig_rax)
			fmt.Printf("name: %s, id: %d \n", name, regs.Orig_rax)
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
