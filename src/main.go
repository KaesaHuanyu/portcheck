package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "portcheck"
	app.Usage = "portcheck"
	app.Version = "1.0"
	app.Action = Run
	app.Author = "Kaesa Li"
	app.Email = "kaesalai@gmail.com"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "address",
			Usage:  "Input the check host and port, just like host:port.",
			EnvVar: "ADDRESS",
		},
	}
	app.Run(os.Args)
}

// Run initializes the driver
func Run(c *cli.Context) {
	address := mustGetStringVar(c, "address")
	addresses := getSlice(address)
	fmt.Println(addresses)
	for _, ad := range addresses {
		for {
			fmt.Printf("making TCP connection to %s ...", ad)
			err := portCheck(ad)
			if err != nil {
				log.Println(err)
			} else {
				fmt.Println("Success!")
				break
			}
			time.Sleep(2 * time.Second)
		}
	}
}

func errExit(code int, format string, val ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", val...)
	os.Exit(code)
}

func mustGetStringVar(c *cli.Context, key string) string {
	v := strings.TrimSpace(c.String(key))
	if v == "" {
		errExit(1, "%s must be provided", key)
	}
	return v
}

func portCheck(address string) error {
	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err == nil {
		conn.Close()
		return nil
	}
	if err, ok := err.(net.Error); ok && err.Timeout() {
		log.Fatalf("timeout when making TCP connection: %+v", err)
		return errors.New(fmt.Sprintf("timeout when making TCP connection: %+v", err))
	}
	return errors.New(fmt.Sprintf("failure to make TCP connection: %+v", err))
}

func getSlice(addresses string) []string {
	address := []byte(addresses)
	new := []string{}
	var in int
	for index, ad := range address {
		if ad == ',' {
			new = append(new, string(address[in:index]))
			in = index + 1
		} else if index == len(addresses)-1 {
			new = append(new, string(address[in:]))
		}
	}
	return new
}
