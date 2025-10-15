package main

import (
	"context"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/ncraft-io/tilestream/service-go/pkg/tilestream/mbtiles"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

// merge --src  --dest
// info
func main() {
	app := &cli.App{
		Name:  "mbtiles",
		Usage: "mbtiles utility",
		Commands: []*cli.Command{
			{
				Name:    "merge",
				Aliases: []string{"m"},
				Usage:   "merge mbtiles",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "src", Aliases: []string{"s"}},
					&cli.StringFlag{Name: "dest", Aliases: []string{"d"}},
				},
				Action: func(ctx *cli.Context) error {
					src := ctx.String("src")
					dest := ctx.String("dest")
					return mbtiles.Merge(src, dest)
				},
			},
			{
				Name:    "info",
				Aliases: []string{"i"},
				Usage:   "show the info of the mbtiles",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "file", Aliases: []string{"f"}},
				},
				Action: func(ctx *cli.Context) error {
					file := ctx.String("file")
					info, err := mbtiles.NewFromFile(file).Info(context.Background())
					if err != nil {
						return err
					}
					json, _ := jsoniter.MarshalIndent(info, "", "    ")
					fmt.Println("\n" + string(json))
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
