package main

import (
	"context"
	"fmt"
	"log"
	"os"

	jsoniter "github.com/json-iterator/go"
	"github.com/ncraft-io/tilestream/service-go/pkg/tilestream/mbtiles"
	"github.com/urfave/cli/v2"
)

func newApp() *cli.App {
	return &cli.App{
		Name:  "mbtiles",
		Usage: "mbtiles utility",
		Commands: []*cli.Command{
			{
				Name:    "merge",
				Aliases: []string{"m"},
				Usage:   "merge an MBTiles file or directory recursively (later paths overwrite duplicate tiles)",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "src", Aliases: []string{"s"}, Required: true, Usage: "source MBTiles file or directory"},
					&cli.StringFlag{Name: "dest", Aliases: []string{"d"}, Required: true, Usage: "destination .mbtiles file"},
				},
				Action: func(ctx *cli.Context) error {
					src := ctx.String("src")
					dest := ctx.String("dest")
					return mbtiles.MergeContext(ctx.Context, src, dest)
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

}

func main() {
	if err := newApp().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
