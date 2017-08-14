package fid

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/inazo1115/bitarray"
)

type FID struct {
	bits     *bitarray.BitArray
	lBitSize int
	sBitSize int
	lMap     []int
	sMap     []int
	pMap     []int
}

func NewFID(bits *bitarray.BitArray) *FID {
	lBitSize, sBitSize, lMap, sMap, pMap := build(bits)
	return &FID{bits, lBitSize, sBitSize, lMap, sMap, pMap}
}

func build(bits *bitarray.BitArray) (int, int, []int, []int, []int) {

	lg := int(math.Log2(float64(bits.Size())))
	lBitSize := lg * lg
	sBitSize := lg / 2

	lMap := make([]int, divAndCeil(bits.Size(), lBitSize))
	sMap := make([]int, divAndCeil(bits.Size(), sBitSize))
	pMap := make([]int, sBitSize*sBitSize)

	for i := 0; i < len(lMap); i++ {
		lMap[i], _ = bits.Rank(true, lBitSize*i)
	}
	for i := 0; i < len(sMap); i++ {
		tmp, _ := bits.Rank(true, sBitSize*i)
		sMap[i] = tmp - lMap[(i*sBitSize)/lBitSize]
	}
	for i := 0; i < len(pMap); i++ {
		pMap[i] = strings.Count(strconv.FormatInt(int64(i), 2), "1")
	}

	return lBitSize, sBitSize, lMap, sMap, pMap
}

func divAndCeil(x, y int) int {
	if x%y == 0 {
		return x / y
	}
	return (x / y) + 1
}

func (fid *FID) Access(idx int) (bool, error) {
	return fid.bits.Get(idx)
}

func (fid *FID) Rank(val bool, idx int) (int, error) {
	if idx >= fid.bits.Size() {
		return 0, fmt.Errorf("out of index: %d >= %d", idx, fid.bits.Size())
	}
	t := fid.lMap[idx/fid.lBitSize] + fid.sMap[idx/fid.sBitSize] + fid.pMap[idx%fid.sBitSize]
	if val {
		return t, nil
	} else {
		return idx - t, nil
	}
}

func (fid *FID) Select(val bool, ith int) (int, error) {

	r, _ := fid.Rank(val, fid.bits.Size()-1)
	a, _ := fid.Access(fid.bits.Size() - 1)

	if r < ith || (r < ith+1 && a != val) {
		return 0, fmt.Errorf("bits doesn't have %d + 1 %t", ith, val)
	}
	return fid.select_(val, ith, 0, fid.bits.Size())
}

func (fid *FID) select_(val bool, ith, from, to int) (int, error) {

	mid := (from + to) / 2
	r, _ := fid.Rank(val, mid)
	a, _ := fid.Access(mid)

	if r == ith && a == val {
		return mid, nil
	} else if r > ith {
		return fid.select_(val, ith, from, mid)
	} else {
		return fid.select_(val, ith, mid, to)
	}

	panic("unreach here")
}

func (fid *FID) Bits() *bitarray.BitArray {
	return fid.bits
}

func (fid *FID) String() string {
	ret := "FID {\n"
	ret += fmt.Sprintf("  bits: %v\n", fid.bits.String())
	ret += fmt.Sprintf("  lBitSize: %v\n", fid.lBitSize)
	ret += fmt.Sprintf("  sBitSize: %v\n", fid.sBitSize)
	ret += fmt.Sprintf("  lMap: %v\n", fid.lMap)
	ret += fmt.Sprintf("  sMap: %v\n", fid.sMap)
	ret += fmt.Sprintf("  pMap: %v\n", fid.pMap)
	ret += "}"
	return ret
}
