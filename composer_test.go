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
