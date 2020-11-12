package command

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/urfave/cli"
)

var NewSIDCommand = cli.Command{
	Name:  "new",
	Usage: "New snow id",
	Flags: []cli.Flag{
		cli.Uint64Flag{
			Name:  "n",
			Usage: "node number",
		},
	},
	Action: func(context *cli.Context) error {
		nodeNumber := context.Int64("n")
		node, err := snowflake.NewNode(nodeNumber)
		if err != nil {
			return err
		}
		fmt.Printf("SID is : %+v\r\n", node.Generate().Int64())
		return nil
	},
}
