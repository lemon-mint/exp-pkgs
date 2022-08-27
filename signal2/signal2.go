package signal2

import (
	"os"
	"os/signal"
	"syscall"
)

// WaitForInterrupt waits for an os.Interrupt (Ctrl+C) or SIGTERM (kill) and then return
func WaitForInterrupt() {
	var ch chan os.Signal = make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
	signal.Stop(ch)
}

// WFI waits for an os.Interrupt (Ctrl+C) or SIGTERM (kill) and then return
func WFI() { WaitForInterrupt() }
