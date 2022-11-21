package main

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name     string
		args     args
		wantHead *ListNode
	}{
		{
			name: "test1",
			args: args{
				l1: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
				l2: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
			wantHead: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 8,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHead := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(gotHead, tt.wantHead) {
				t.Errorf("addTwoNumbers() = %v, want %v", gotHead, tt.wantHead)
			}
		})
	}
}
