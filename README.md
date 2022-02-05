# servermedia
ServerMedia

### Create file main.go

```
package main

import (
	"github.com/lordbasex/servermedia"
)

func main() {

	servermedia.Start()

}
```

### Build or Run

```
go mod init servermedia
go get -v
go build -o servermedia
chmod 777 servermedia
./servermedia
#go run main.go
```
