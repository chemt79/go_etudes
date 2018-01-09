package models

import "testing"

func TestRenderTable(t *testing.T) {
	type args struct {
		header *[]string
		rows   *[]Stringer
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RenderTable(tt.args.header, tt.args.rows)
		})
	}
}
