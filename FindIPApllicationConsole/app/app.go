package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate will return the console application ready for execute
func Generate () *cli.App {
	app := cli.NewApp()

	app.Name = "Console Application"

	app.Usage = "Search Ips and Server Names"

	flags := []cli.Flag{
				cli.StringFlag{
					Name: "host",
					Value: "devbook.com.br",
				},
			}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "IPs search for addresses on the internet",
			Flags: flags ,
			Action: searchIp,
			
		},
		{
		   Name: "servers",
		   Usage: "Search name serves on the internet",
		   Flags: flags,
		   Action: searchServer,
		},
	}

	return app
}

func searchIp(c *cli.Context) {
	host := c.String("host")

	// net

	ips , err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _ , ip :=range ips {
		fmt.Println(ip)
	}
}

func searchServer(c * cli.Context) {
	host := c.String("host")

	servers , err := net.LookupNS(host)

		if err != nil {
		log.Fatal(err)
	}

		for _ , ip :=range servers {
		fmt.Println(ip)
	}
}