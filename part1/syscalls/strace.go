package main

import (
    "fmt"
    "os"
    "os/exec"
    "syscall"
"text/tabwriter"

    "github.com/seccomp/libseccomp-golang"

)

type syscallCounter []int

const maxSyscalls = 303

func (s syscallCounter) init() syscallCounter {
    s = make(syscallCounter, maxSyscalls)
    return s
}

func (s syscallCounter) inc(syscallID uint64) error {
    if syscallID > maxSyscalls {
        return fmt.Errorf("invalid syscall ID (%x)", syscallID)
    }

    s[syscallID]++
    return nil
}

func (s syscallCounter) print() {
    w := tabwriter.NewWriter(os.Stdout, 0, 0, 8, ' ', tabwriter.AlignRight|tabwriter.Debug)
    for k, v := range s {
        if v > 0 {
            name, _ := seccomp.ScmpSyscall(k).GetName()
            fmt.Fprintf(w, "%d\t%s\n", v, name)
        }
    }
    w.Flush()
}

func (s syscallCounter) getName(syscallID uint64) string {
    name, _ := seccomp.ScmpSyscall(syscallID).GetName()
    return name
}


func main() {
    var err error
    var regs syscall.PtraceRegs
    var ss syscallCounter
    ss = ss.init()
    pid := 8363
    exit := true

    for {
        if exit {
            err = syscall.PtraceGetRegs(pid, &regs)
            if err != nil {
                break
            }
            //fmt.Printf("%#v \n",regs)
            name := ss.getName(regs.Orig_rax)
            fmt.Printf("name: %s, id: %d \n", name, regs.Orig_rax)
            ss.inc(regs.Orig_rax)
        }

        /**
        http://www.linuxjournal.com/article/6100?page=0,1
        Here we are tracing the write system calls, and ls makes three write system calls. The call to ptrace, with a first argument of PTRACE_SYSCALL, makes the kernel stop the child process whenever a system call entry or exit is made. It's equivalent to doing a PTRACE_CONT and stopping at the next system call entry/exit.
        */
        err = syscall.PtraceSyscall(pid, 0)
        if err != nil {
            panic(err)
        }

        // http://www.linuxjournal.com/article/6100?page=0,1
        //The status variable in the wait call is used to check whether the child has exited. This is the typical way to check whether the child has been stopped by ptrace or was able to exit.
        _, err = syscall.Wait4(pid, nil, 0, nil)
        if err != nil {
            panic(err)
        }

        exit = !exit
    }

    ss.print()
}

