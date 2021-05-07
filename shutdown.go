package shutdown

import (
	"os"
	"os/signal"
)

func Hook(sig ...os.Signal) {
	c := make(chan os.Signal, 1)
	defer close(c)

	signal.Notify(c, sig...)

	<-c
}
