package command

import (
	"fmt"
	"github.com/urfave/cli"
	"time"
)

var ParseTimeCommand = cli.Command{
	Name:  "parse",
	Usage: "Parse snow id to time",
	Flags: []cli.Flag{
		cli.Uint64Flag{
			Name:  "i",
			Usage: "snow id",
		},
	},
	Action: func(context *cli.Context) error {
		sid := context.Int64("i")
		ec := int64(1288834974657)
		step2 := sid - 4096
		step1 := step2 >> 22
		unixMill := step1 + ec
		utcTime := unixMill2Time(unixMill).UTC()
		locTime := utcTime.Local()
		fmt.Printf("SID is : %+v\r\nUnixMill is : %+v\r\nUTC time is : %+v\r\nLocal time is : %+v\r\n", sid, unixMill, utcTime, locTime)
		return nil
	},
}

func unixMill2Time(unixMill int64) time.Time {
	//second
	sec := unixMill / (1e3)
	//nano second
	nsec := (unixMill - sec*1e3) * 1e6
	return time.Unix(sec, nsec)
}
