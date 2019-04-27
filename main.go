package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/urfave/cli"

	. "github.com/logrusorgru/aurora"
)

func main() {
	app := cli.NewApp()
	app.Name = "JWT CLI"
	app.Usage = `Encode/Decode jwt in terminal`
	app.Version = "1.0.0"
	//  Global Flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "alg",
			Value: "HS256",
			Usage: "Signed method",
		},
		cli.StringFlag{
			Name:  "secret, s",
			Value: "",
			Usage: "Secret key",
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.NArg() <= 0 {
			return fmt.Errorf("%v", Red("invalid argument"))
		}
		parseToken(c.Args().Get(0), c.String("secret"), c.String("alg"))
		return nil
	}
	app.Commands = []cli.Command{
		{
			Name:    "sign",
			Aliases: []string{"s"},
			Usage:   "sign new jwt token",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "payload, p",
					Value: "{}",
					Usage: "JSON payload of jwt token",
				},
			},
			Action: func(c *cli.Context) error {
				data := c.String("payload")
				// alg := flag.String("alg", "HS256", "Get Signing method from this str, default = HS256")
				// secret := flag.String("secret", "", "secret key")
				// flag.Parse()
				jwtData := jwt.MapClaims{}
				err := json.Unmarshal([]byte(data), &jwtData)
				if err != nil {
					return fmt.Errorf("%v", Red(err))
				}
				method := jwt.GetSigningMethod(c.GlobalString("alg"))
				if method == nil {
					return fmt.Errorf("%v", Red("invalid alg"))
				}
				secret := c.GlobalString("secret")
				token, err := signToken(jwtData, method, secret)
				if err != nil {
					return fmt.Errorf("%v", Red(err))
				}
				fmt.Println(token)
				parseToken(token, secret, method.Alg())
				return nil
			},
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
