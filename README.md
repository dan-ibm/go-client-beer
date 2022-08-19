Go-client implementation to get beer's styles
=========================
Import package
```
go get -u github.com/dan-ibm/go-client-beer
```
## Example
```go
package main

import (
	"fmt"
	"github.com/dan-ibm/go-client-beer/client"
	"log"
)

func main() {
	c, err := client.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	// All styles
	data, err := c.GetStylesByName("Roggenbier")
	if err != nil {
		log.Fatal(err)
	}

	for _, d := range data {
		fmt.Println(d.Info())
	}
	
	// Style by name
	data, err = c.GetStylesByName("Roggenbier")
	if err != nil {
		log.Fatal(err)
	}
	

	for _, d := range data {
		fmt.Println(d.Info())
	}
}

```
