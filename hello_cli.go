package main

import(
	"os"
	"fmt"
	"github.com/codegangsta/cli"
)

func main(){

	author := cli.Author{"Marcus Willock", "crazcalm@gmail.com"}

	app := cli.NewApp()
	app.Name = "boom"
	app.Usage = "make an explosive entrance"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "lang",
			Value: "english",
			Usage: "Language for the greeting",

		},
	}
	app.Commands = []cli.Command{
		{
			Name: "describeit",
			Aliases: []string{"d"},
			Usage: "use it to see a description",
			Description: "This is how we describe describeit the function",
			Action: func(c *cli.Context){
				fmt.Printf("I like to describe things")
			},
		},
	}	

	app.Authors = []cli.Author{author}
	app.Action = func(c *cli.Context){
		name := "someone"
		if len(c.Args()) > 0{
			name = c.Args()[0]
		}
		if c.String("lang") == "spanish" {
			println("Hola", name)
		} else {
			println("Hello", name)
		}
	}

	app.Run(os.Args)
}