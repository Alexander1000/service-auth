package trap

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Trap struct {
	trap chan os.Signal
}

func NewTrap() *Trap {
	trap := make(chan os.Signal, 1)
	signal.Notify(trap, syscall.SIGINT, os.Interrupt, syscall.SIGTERM)
	return &Trap{
		trap: trap,
	}
}

func (t *Trap) WaitShutdown(ctx context.Context) error {
	select {
	case <-t.trap:
		log.Println("termination signal caught")
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
