package shapes

import (
	"reflect"
	"testing"
)

func TestNewShape(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want *Shape
	}{
		{
			name: "It should create Shape object",
			args: args{1, 2},
			want: &Shape{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewShape(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewShape() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShape_X(t *testing.T) {
	tests := []struct {
		name string
		s    *Shape
		want int
	}{
		{
			name: "It should return 1",
			s:    NewShape(1, 2),
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.X(); got != tt.want {
				t.Errorf("Shape.X() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShape_Y(t *testing.T) {
	tests := []struct {
		name string
		s    *Shape
		want int
	}{
		{
			name: "It should return 2",
			s:    NewShape(1, 2),
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Y(); got != tt.want {
				t.Errorf("Shape.Y() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShape_SetX(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		s    *Shape
		args args
	}{
		{
			name: "It should set Shape.x",
			s:    NewShape(1, 2),
			args: args{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.SetX(tt.args.x)
			if tt.s.X() != tt.args.x {
				t.Errorf("Doesn't set Shape.x")
			}
		})
	}
}

func TestShape_SetY(t *testing.T) {
	type args struct {
		y int
	}
	tests := []struct {
		name string
		s    *Shape
		args args
	}{
		{
			name: "It should set Shape.y",
			s:    NewShape(1, 2),
			args: args{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.SetY(tt.args.y)
			if tt.s.Y() != tt.args.y {
				t.Errorf("Doesn't set Shape.y")
			}
		})
	}
}

func TestShape_String(t *testing.T) {
	tests := []struct {
		name string
		s    *Shape
		want string
	}{
		{
			name: "It should return object string",
			s:    NewShape(1, 2),
			want: "Shape x:1, y:2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("Shape.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
