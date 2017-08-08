package fid

import (
	"testing"

	"github.com/inazo1115/bitarray"
)

const (
	_0 = false
	_1 = true
)

func TestRank(t *testing.T) {

	bits := bitarray.NewBitArrayWithInit(
		[]bool{_0, _1, _1, _0, _1, _1, _0, _0, _1, _0, _0})
	fid := NewFID(bits)

	tests := []struct {
		val      bool
		idx      int
		expected int
	}{
		{_0, 0, 0},
		{_0, 1, 1},
		{_0, 2, 1},
		{_0, 3, 1},
		{_0, 4, 2},
		{_0, 5, 2},
		{_0, 6, 2},
		{_0, 7, 3},
		{_0, 8, 4},
		{_0, 9, 4},
		{_0, 10, 5},
		{_1, 0, 0},
		{_1, 1, 0},
		{_1, 2, 1},
		{_1, 3, 2},
		{_1, 4, 2},
		{_1, 5, 3},
		{_1, 6, 4},
		{_1, 7, 4},
		{_1, 8, 4},
		{_1, 9, 5},
		{_1, 10, 5},
	}

	for _, test := range tests {

		bits_actual, err := bits.Rank(test.val, test.idx)
		if err != nil {
			t.Errorf(err.Error())
		}

		fid_actual, err := fid.Rank(test.val, test.idx)
		if err != nil {
			t.Errorf(err.Error())
		}

		if fid_actual != bits_actual {
			t.Errorf("Rank(%v, %v) => '%v', want '%v'",
				test.val, test.idx, fid_actual, bits_actual)
		}
		if fid_actual != test.expected {
			t.Errorf("Rank(%v, %v) => '%v', want '%v'",
				test.val, test.idx, fid_actual, test.expected)
		}
	}
}

func TestSelect(t *testing.T) {

	bits := bitarray.NewBitArrayWithInit(
		[]bool{_0, _1, _1, _0, _1, _1, _0, _0, _1, _0, _0})
	fid := NewFID(bits)

	tests := []struct {
		val      bool
		ith      int
		expected int
	}{
		{_0, 0, 0},
		{_0, 1, 3},
		{_0, 2, 6},
		{_0, 3, 7},
		{_0, 4, 9},
		{_0, 5, 10},
		{_1, 0, 1},
		{_1, 1, 2},
		{_1, 2, 4},
		{_1, 3, 5},
		{_1, 4, 8},
	}

	for _, test := range tests {

		bits_actual, err := bits.Select(test.val, test.ith)
		if err != nil {
			t.Errorf(err.Error())
		}

		fid_actual, err := fid.Select(test.val, test.ith)
		if err != nil {
			t.Errorf(err.Error())
		}

		if fid_actual != bits_actual {
			t.Errorf("Select(%v, %v) => '%v', want '%v'",
				test.val, test.ith, fid_actual, bits_actual)
		}
		if fid_actual != test.expected {
			t.Errorf("Select(%v, %v) => '%v', want '%v'",
				test.val, test.ith, fid_actual, test.expected)
		}
	}
}
