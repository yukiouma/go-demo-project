package appmanage

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
)

func ReceiveSignal(ctx context.Context, signals []os.Signal) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, signals...)

	// waiting signal or context done
	select {
	case sig := <-c:
		fmt.Println()
		log.Printf("receive signal: \n\t%v", sig)
		if sig == os.Interrupt {
			return errors.New(sig.String())
		}
		return nil
	case <-ctx.Done():
		err := ctx.Err()
		log.Printf("stop receiving signal: \n\t%v", err)
		close(c)
		return errors.New(err.Error())
	}
}
