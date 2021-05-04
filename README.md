# Logram

This package allows you to log errors in telegram channel. First you need to contact the BotFather to create your logger bot. Then create a channel where your bot will log errors. Save bot token and channel id.

## Installing

```shell
go get github.com/iluxaorlov/logram
```

## Usage
```go
package main

import (
	"flag"
	"github.com/iluxaorlov/logram/pkg/logram"
	"io"
	"log"
	"os"
)

var token string
var chatId int64

func init() {
	flag.StringVar(&token, "token", "", "bot token")
	flag.Int64Var(&chatId, "chatId", 0, "chat id")
}

func main() {
	flag.Parse()

	lg := logram.NewWriter(token, chatId)

	log.SetOutput(io.MultiWriter(os.Stdout, lg))

	log.Println("Some error occurred")
}
```
