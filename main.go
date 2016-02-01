// utility for determining single origin
// placeholder for http request

package main

import (
    "fmt"
    "os"
    "time"
    "github.com/codegangsta/cli"
)

func main(){
	app				:= cli.NewApp()

	app.Commands	= []cli.Command{
		{
			Name:	'request',
			Usage:	'Starts curl like request',
			Action: func(c *cli.Context) {
				f := highlander.NewReq()
			}
	    }
	}

	app.Run(os.Args)
}