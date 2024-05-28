// Package byteseq implements a randomised byte sequence.
//
// Returns values in the byte range of 0-255, but returns each
// value only once. Will return an error if an attempt is made
// to get a value from an 'exhausted' sequence.
package byteseq

import (
	"fmt"
	"math/rand"
)

// A RandomByteSeq returns byte values randomly and only once each.
// Byte values can be marked as already having been consumed when
// creating a new struct.
type RandomByteSeq struct {
	consumedBytes   [32]byte // Bitmap for each of 256 possible values
	remainingValues int
}

// NewRandomSeq returns a new RandomByteSeq. consumedBytes is a slice of
// optional values to exclude from the sequence. Internally, they are considered
// as having already been "consumed".
func NewRandomSeq(consumedBytes []byte) *RandomByteSeq {
	seq := &RandomByteSeq{}

	seq.remainingValues = 256

	// Were any already-consumed bytes specified?
	if len(consumedBytes) > 0 {
		for _, b := range consumedBytes {
			seq.consumeByte(b)
		}
	}
	return seq
}

// determine if a given byte value is marked as consumed in the internal bitmap structures
func (r *RandomByteSeq) valueHasBeenConsumed(b byte) bool {
	// determine if a given byte value is marked as consumed in the internal bitmap structures
	// Get the correct byte of our bitmap from top 5 bits
	byteIndex := b & 0xF8 >> 3
	testByte := r.consumedBytes[byteIndex]

	// Now, from the bottom 3 bits, check the appropriate bit
	bitMask := byte(1 << (b & 0x07))

	return testByte&bitMask != 0
}

func (r *RandomByteSeq) consumeByte(b byte) {
	// marks a given byte value as having been consumed in the internal bitmap structures.
	byteIndex := b & 0xF8 >> 3
	bitMask := byte(1 << (b & 0x07))

	r.consumedBytes[byteIndex] |= bitMask
	r.remainingValues--
}

// HasMore is intended to check if the RandomByteSeq has any more values
// that it can return.
func (r *RandomByteSeq) HasMore() bool {
	return r.remainingValues > 0
}

// NextValue returns either the next random byte value that hasn't been
// previously returned, or will return an error if an attempt is made
// to get a value from an exhausted sequence. The HasMore function
// can be used to avoid this.
func (r *RandomByteSeq) NextValue() (byte, error) {
	// Are there any more values available?
	if r.remainingValues == 0 {
		return 0, fmt.Errorf("sequence has been exhausted")
	}

	for {
		valueByte := byte(rand.Int() & 0xFF)

		if !r.valueHasBeenConsumed(valueByte) {
			r.consumeByte(valueByte)
			return valueByte, nil
		}
	}
}
