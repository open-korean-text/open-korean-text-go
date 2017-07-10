# open-korean-text-go

### To get packages
run these operations on terminal or console

```
go get github.com/open-korean-text/open-korean-text-go
```

You might get some errors, but don't worry!
That's not problem.

### Run
```go
package main

import (
	"fmt"

	processor "github.com/open-korean-text/open-korean-text-go/processor"
)

func main() {
	result := processor.Normalize("만듀 먹것니? 먹겄서? 먹즤?")
	fmt.Println(result)
}
```


### And...
'Tokenizer' and 'Phrase Extractor' will be built!
