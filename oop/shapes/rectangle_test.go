package shapes

import (
	"reflect"
	"testing"
)

func TestNewRectangle(t *testing.T) {
	type args struct {
		x  int
		y  int
		x1 int
		y1 int
	}
	tests := []struct {
		name string
		args args
		want *Rectangle
	}{
		{
			name: "It should create Rectangle object",
			args: args{1, 2, 3, 4},
			want: &Rectangle{Shape{1, 2}, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRectangle(tt.args.x, tt.args.y, tt.args.x1, tt.args.y1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_String(t *testing.T) {
	tests := []struct {
		name string
		r    *Rectangle
		want string
	}{
		{
			name: "It should returb object string",
			r:    NewRectangle(1, 2, 3, 4),
			want: "Rectangle x:1, y:2, x1:3, y1:4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.String(); got != tt.want {
				t.Errorf("Rectangle.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_Parent(t *testing.T) {
	tests := []struct {
		name string
		r    *Rectangle
		want *Shape
	}{
		{
			name: "It should return parent of object",
			r:    NewRectangle(1, 2, 3, 4),
			want: &Shape{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Parent(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rectangle.Parent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_SetX1(t *testing.T) {
	type fields struct {
		Shape Shape
		x1    int
		y1    int
	}
	type args struct {
		x1 int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "It should set Rectangle.X1",
			fields: fields{Shape{1, 2}, 3, 4},
			args:   args{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rectangle{
				Shape: tt.fields.Shape,
				x1:    tt.fields.x1,
				y1:    tt.fields.y1,
			}
			r.SetX1(tt.args.x1)

			if r.X1() != tt.args.x1 {
				t.Errorf("Doesn't set Rectangle.x1")
			}
		})
	}
}

func TestRectangle_SetY1(t *testing.T) {
	type fields struct {
		Shape Shape
		x1    int
		y1    int
	}
	type args struct {
		y1 int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "It should set Rectangle.Y1",
			fields: fields{Shape{1, 2}, 3, 4},
			args:   args{4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rectangle{
				Shape: tt.fields.Shape,
				x1:    tt.fields.x1,
				y1:    tt.fields.y1,
			}
			r.SetY1(tt.args.y1)

			if r.Y1() != tt.args.y1 {
				t.Errorf("Doesn't set Rectangle.y1")
			}
		})
	}
}

func TestRectangle_X1(t *testing.T) {
	type fields struct {
		Shape Shape
		x1    int
		y1    int
	}
	tests := []struct {
		name   string
		fields fields
		wantX1 int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rectangle{
				Shape: tt.fields.Shape,
				x1:    tt.fields.x1,
				y1:    tt.fields.y1,
			}
			if gotX1 := r.X1(); gotX1 != tt.wantX1 {
				t.Errorf("Rectangle.X1() = %v, want %v", gotX1, tt.wantX1)
			}
		})
	}
}

func TestRectangle_Y1(t *testing.T) {
	type fields struct {
		Shape Shape
		x1    int
		y1    int
	}
	tests := []struct {
		name   string
		fields fields
		wantY1 int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rectangle{
				Shape: tt.fields.Shape,
				x1:    tt.fields.x1,
				y1:    tt.fields.y1,
			}
			if gotY1 := r.Y1(); gotY1 != tt.wantY1 {
				t.Errorf("Rectangle.Y1() = %v, want %v", gotY1, tt.wantY1)
			}
		})
	}
}
