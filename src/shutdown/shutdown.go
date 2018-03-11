package shutdown

import (
	"os"
    "os/signal"
	"syscall"
)

var waitstop chan struct{}
var sdwnFunctions []func()

func init() {
    waitstop = make(chan struct{}, 1)
}

// Wait - wait system halting events
func Wait() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    select {
        case <-c:
        case <-waitstop:
        break;
    }
    close(waitstop)
    close(c)
    // call shutdown functions
    for i := range sdwnFunctions {
        f := sdwnFunctions[len(sdwnFunctions) - i - 1]
        f()
   }
}

// RegFunction - to registry function will called at app shutdown event
func RegFunction(f func()) {
    sdwnFunctions = append(sdwnFunctions, f)
}

// Stop - stops the application
func Stop() {
    waitstop <- struct{}{}
}
