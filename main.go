package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	interval int
)
func init () {
	flag.IntVar(&interval, "i", 1, "interval seconds")
}
func usage() {
	fmt.Fprintf(os.Stderr, "usage: gwatch -i [interval] seconds")
	os.Exit(-2)
}


func main() {
	flag.Usage = usage
	flag.Parse()

args := flag.Args()

	if len(args) < 1 {
		usage()
	}

	stoper := make(chan os.Signal, 2)
	signal.Notify(stoper, os.Interrupt, syscall.SIGTERM)
	f(args)

	timer := time.NewTicker(time.Duration(interval)*time.Second)
	for {
		select {
		case <- timer.C:
			f(args)
		case <- stoper:
			timer.Stop()
			os.Exit(0)
		}
	}
}

func f(args []string) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	cmd = exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	fmt.Println("interval", interval, " second,command: ", strings.Join(args, " "))
	output, _ := cmd.Output()
	fmt.Print(string(output))
}
