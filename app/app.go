package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

var defaultHost string

func init() {
	defaultHost = "google.com"
}

// `Generate` returns the command line application ready to be runned.
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "server-info"
	app.Usage = "Search IP Address and Names on internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: defaultHost,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ips",
			Usage:  "Search IP Address in internet",
			Flags:  flags,
			Action: ips,
		},
		{
			Name:   "servers",
			Usage:  "Search Server Names in internet",
			Flags:  flags,
			Action: servers,
		},
	}

	return app
}

func ips(context *cli.Context) {
	host := context.String("host")

	ips, erro := net.LookupIP(host)
	verifyError(erro)

	fmt.Printf("The following ips were found for the host '%s' \n", host)
	for _, ip := range ips {
		fmt.Printf("- %s \n", ip)
	}
}

func servers(context *cli.Context) {
	host := context.String("host")

	servers, erro := net.LookupNS(host)
	verifyError(erro)

	fmt.Printf("The following servers were found for the host '%s' \n", host)
	for _, server := range servers {
		fmt.Printf("- %s \n", server.Host)
	}
}

func verifyError(erro error) {
	if erro != nil {
		log.Fatal(erro)
	}
}
