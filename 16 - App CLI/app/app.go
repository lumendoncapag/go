package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Função Gerar, terar retorno do Tipo App
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Buscar IPs de Nomes de Servidor na Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			//Exemplo ip --host amazon.com.br
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	//Comandos
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca de IPs de Endereços na Internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o Nome dos Servidores na Internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app

}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	// pacote net
	if ips, erro := net.LookupIP(host); erro != nil {
		log.Fatal(erro)
	} else {
		fmt.Printf("Ips pertencentes ao host %s :\n", host)
		for _, ip := range ips {
			fmt.Println(ip)
		}
	}

}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	if servidores, erro := net.LookupNS(host); erro != nil {
		log.Fatal(erro)
	} else {
		fmt.Printf("Servidores pertencentes ao host %s :\n", host)
		for _, servidor := range servidores {
			fmt.Println(servidor.Host)
		}
	}
}
