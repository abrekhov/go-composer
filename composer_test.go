/*
 *   Copyright (c) 2021 Anton Brekhov
 *   All rights reserved.
 */
package composer

import (
	"testing"
)

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
		{
			name:  "simple6",
			args:  args{x: 6},
			wantY: 24,
		},
		{
			name:  "simple20",
			args:  args{x: 20},
			wantY: 80,
		},
		{
			name:  "simple3",
			args:  args{x: 3},
			wantY: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotY := CheckOneNum(tt.args.x); gotY != tt.wantY {
				t.Logf("Mock:%v, Real:%v", gotY, tt.wantY)
				t.Errorf("CheckOneNum() = %v, want %v", gotY, tt.wantY)
			}
		})
	}
}

func TestCheckChildOneNum(t *testing.T) {
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
			name:  "simple20",
			args:  args{x: 20},
			wantY: 80,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotY := CheckChildOneNum(tt.args.x); gotY != tt.wantY {
				t.Errorf("CheckChildOneNum() = %v, want %v", gotY, tt.wantY)
			}
		})
	}
}
