/*
 *   Copyright (c) 2021 Anton Brekhov
 *   All rights reserved.
 */
package composer

import (
	"reflect"
	"testing"
)

func TestComposer_Compose(t *testing.T) {
	type fields struct {
		Name     string
		InitChan int
	}
	type args struct {
		initChan chan int
		fns      []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   reflect.Value
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Composer{
				Name:     tt.fields.Name,
				InitChan: tt.fields.InitChan,
			}
			if got := c.Compose(tt.args.initChan, tt.args.fns...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Composer.Compose() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckOneNum(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name  string
		args  args
		wantY int
	}{
		// TODO: Add test cases.
		{
			name:  "simple2x2",
			args:  args{x: 2},
			wantY: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotY := CheckOneNum(tt.args.x); gotY != tt.wantY {
				t.Errorf("CheckOneNum() = %v, want %v", gotY, tt.wantY)
			}
		})
	}
}
