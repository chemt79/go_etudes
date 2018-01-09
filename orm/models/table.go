package models

import (
	"bytes"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type Stringer interface {
	String() string
}

func RenderTable(header *[]string, rows *[]Stringer) string {
	buffer := new(bytes.Buffer)
	table := tablewriter.NewWriter(buffer)
	table.SetHeader(*header)

	for _, r := range *rows {
		table.Append(strings.Split(r.String(), ","))
	}

	table.Render()
	return buffer.String()
}
