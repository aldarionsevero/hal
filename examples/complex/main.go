package main

import (
	"github.com/danryan/hal"
	_ "github.com/danryan/hal/adapter/shell"
	"github.com/danryan/hal/handler"
	_ "github.com/danryan/hal/store/memory"
	"os"
	"fmt"
	"log"
	"time"
	//"reflect"
	//"fmt"
)

// HAL is just another Go package, which means you are free to organize things
// however you deem best.

// You can define your handlers in the same file...
var pingHandler = hal.Hear(`ping`, func(res *hal.Response) error {
	return res.Send("PONG")
})

var testouHandler = hal.Hear(`teste`, func(res *hal.Response) error {
	file, err := os.Create("result.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
	fmt.Fprintf(file, "opened")
	hal.Logger.Info("passou aqui")

	return res.Send("TESTOU")
})

var robot *hal.Robot

func run() int {
	robot, err := hal.NewRobot()
	//fmt.Println(reflect.TypeOf(robot))
	if err != nil {
		hal.Logger.Error(err)
		return 1
	}

	// Or define them inside another function...
	fooHandler := hal.Respond(`foo`, func(res *hal.Response) error {
		return res.Send("BAR")
	})

	tableFlipHandler := &hal.Handler{
		Method:  hal.HEAR,
		Pattern: `tableflip`,
		Run: func(res *hal.Response) error {
			return res.Send(`(╯°□°）╯︵ ┻━┻`)
		},
	}

	robot.Handle(
		pingHandler,
		testouHandler,
		fooHandler,
		tableFlipHandler,

		// Or stick them in an entirely different package, and reference them
		// exactly in the way you would expect.
		handler.Ping,

		// Or use a hal.Handler structure complete with usage...
		&hal.Handler{
			Method:  hal.RESPOND,
			Pattern: `SYN`,
			Usage:   `hal syn - replies with "ACK"`,
			Run: func(res *hal.Response) error {
				return res.Reply("ACK")
			},
		},

		// Or even inline!
		hal.Hear(`yo`, func(res *hal.Response) error {
			return res.Send("lo")
		}),
	)
	robot.Run()
//	for(true){
//
//	}
	return 0
}

func main() {
	os.Exit(run())
}
