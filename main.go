package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"

    "github.com/shirou/gopsutil/mem"
    "github.com/shirou/gopsutil/process"
)


func listProcesses() {
    procs, err := process.Processes()
    if err != nil {
        fmt.Println("Error retrieving processes:", err)
        return
    }

    fmt.Println("Running Processes:")
    for _, proc := range procs {
        name, _ := proc.Name() 
        pid := proc.Pid        
        fmt.Printf("PID: %d, Name: %s\n", pid, name)
    }
}

func killProcess(pid int32) {
    proc, err := process.NewProcess(pid)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    if err := proc.Kill(); err != nil {
        fmt.Printf("Failed to kill process %d: %s\n", pid, err)
    } else {
        fmt.Printf("Process %d killed successfully.\n", pid)
    }
}


func monitorMemory() {
    memStats, err := mem.VirtualMemory()
    if err != nil {
        fmt.Println("Error retrieving memory stats:", err)
        return
    }

    fmt.Printf("Total Memory: %v, Available Memory: %v, Used Memory: %v\n",
        memStats.Total, memStats.Available, memStats.Used)
}


func main() {
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Println("Welcome to the Process Manager CLI!")
    fmt.Println("Available commands: list, kill <PID>, memory, exit")

    for {
        fmt.Print("> ")
        scanner.Scan()                      
        input := scanner.Text()             
        command := strings.Fields(input)    

        if len(command) == 0 {
            continue 
        }

        switch command[0] {
        case "list":
            listProcesses() 

        case "kill":
            if len(command) < 2 {
                fmt.Println("Please provide a PID to kill.")
                continue
            }
            pidStr := command[1]
            pid, err := strconv.Atoi(pidStr)
            if err != nil {
                fmt.Println("Invalid PID:", pidStr)
                continue
            }
            killProcess(int32(pid)) 

        case "memory":
            monitorMemory() 

        case "exit":
            fmt.Println("Exiting the Process Manager CLI.")
            return

        default:
            fmt.Println("Unknown command. Available commands: list, kill <PID>, memory, exit")
        }
    }
}
