package middleware

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"gopkg.in/flosch/pongo2.v3"
)

func init() {
	pongo2.RegisterFilter("humanize_bytes", func(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
		fmt.Printf("ddd%d\n", in.Integer())
		return pongo2.AsValue(humanize.IBytes(uint64(in.Integer()))), nil
	})

}
