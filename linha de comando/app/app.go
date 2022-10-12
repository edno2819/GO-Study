package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Aplicação de linha de comando
func Gerar() *cli.App {
	app := cli.NewApp()

	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca de IPs Nomes de servidores na internet"
	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca IPS de endereços na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				},
				cli.Int64Flag{
					Name: "qtd",
				},
			},
			Action: buscarIps,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")
	qtd := c.Int("qtd")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	var max_len int

	if len(ips) < qtd {
		max_len = len(ips)
	} else {
		max_len = qtd
	}

	for _, ip := range ips[:max_len] {
		fmt.Println(ip)
	}

}
