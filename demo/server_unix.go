package demo

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"

	"virtual_fido"
)

func runServer(client virtual_fido.Client) {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		virtual_fido.Start(client)
		wg.Done()
	}()
	go func() {
		time.Sleep(500 * time.Millisecond)
		prog := exec.Command("usbip", "attach", "-r", "127.0.0.1", "-b", "2-2")
		prog.Stdin = os.Stdin
		prog.Stdout = os.Stdout
		prog.Stderr = os.Stderr
		err := prog.Run()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
		wg.Done()
	}()
	wg.Wait()
}
