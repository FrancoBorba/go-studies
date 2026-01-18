package main

import (
	"console/app"
	"fmt"
	"log"
	"os"
)

func main()  {


	fmt.Println("Start")

	application := app.Generate()

	err := application.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
