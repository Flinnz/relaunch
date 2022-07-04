package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func printUsage() {
	fmt.Println("Usage: relaunch <-c <amount>> <-d <delay>> <cmd to relaunch>")
}

var count *uint
var delay *uint

func init() {
	count = flag.Uint("c", 0, "amount of relaunches. default is infinite(0)")
	delay = flag.Uint("d", 500, "delay in milliseconds. default is 500ms")
}

func skipArgs() []string {
	skip := 1
	flag.Visit(func(f *flag.Flag) { skip += 2 })
	return os.Args[skip:]
}

func decrement(n *uint) uint {
	valueBefore := *n
	*n--
	return valueBefore
}

func loop(count, delay uint, args []string) {
	for i := count; count == 0 || decrement(&i) > 0; {
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Millisecond * time.Duration(delay))
	}
}

func main() {
	flag.Parse()
	allArgs := skipArgs()
	if len(allArgs) <= 0 {
		printUsage()
		return
	}
	loop(*count, *delay, allArgs)
}
