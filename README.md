# go-shutdown

go-shutdown is a shutdown hook for golang

```go
import (
	"fmt"
	"syscall"

	"github.com/juacker/go-shutdown"
)

func main() {
	fmt.Println("Waiting for stop signal")

	<-shutdown.Hook(syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("bye, bye")

}
```
