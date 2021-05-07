package shutdown

import (
	"os"
	"os/signal"
)

func Hook(sig ...os.Signal) chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, sig...)
	return c
}
