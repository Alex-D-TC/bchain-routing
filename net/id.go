package net

import (
	"math/big"

	"secondbit.org/wendy"
)

type NodeID wendy.NodeID

func NodeIDFromBytes(source []byte) (NodeID, error) {
	wendyID, err := wendy.NodeIDFromBytes(source)
	return NodeID(wendyID), err
}

// String returns the hexadecimal string encoding of the NodeID.
func (id NodeID) String() string {
	return wendy.NodeID(id).String()
}

// Equals tests two NodeIDs for equality and returns true if they are considered equal, false if they are considered inequal. NodeIDs are considered equal if each digit of the NodeID is equal.
func (id NodeID) Equals(other NodeID) bool {
	return wendy.NodeID(id).Equals(wendy.NodeID(other))
}

// Less tests two NodeIDs to determine if the ID the method is called on is less than the ID passed as an argument. An ID is considered to be less if the first inequal digit between the two IDs is considered to be less.
func (id NodeID) Less(other NodeID) bool {
	return wendy.NodeID(id).Less(wendy.NodeID(other))
}

// CommonPrefixLen returns the number of leading digits that are equal in the two NodeIDs.
func (id NodeID) CommonPrefixLen(other NodeID) int {
	return wendy.NodeID(id).CommonPrefixLen(wendy.NodeID(other))
}

// Diff returns the difference between two NodeIDs as an absolute value. It performs the modular arithmetic necessary to find the shortest distance between the IDs in the (2^128)-1 item nodespace.
func (id NodeID) Diff(other NodeID) *big.Int {
	return wendy.NodeID(id).Diff(wendy.NodeID(other))
}

// RelPos uses modular arithmetic to determine whether the NodeID passed as an argument is to the left of the NodeID it is called on (-1), the same as the NodeID it is called on (0), or to the right of the NodeID it is called on (1) in the circular node space.
func (id NodeID) RelPos(other NodeID) int {
	return wendy.NodeID(id).RelPos(wendy.NodeID(other))
}

// Base10 returns the NodeID as a base 10 number, translating each base 16 digit.
func (id NodeID) Base10() *big.Int {
	return wendy.NodeID(id).Base10()
}

// MarshalJSON fulfills the Marshaler interface, allowing NodeIDs to be serialised to JSON safely.
func (id NodeID) MarshalJSON() ([]byte, error) {
	return wendy.NodeID(id).MarshalJSON()
}

// UnmarshalJSON fulfills the Unmarshaler interface, allowing NodeIDs to be unserialised from JSON safely.
func (id *NodeID) UnmarshalJSON(source []byte) error {
	wendyID := wendy.NodeID(*id)
	return (&wendyID).UnmarshalJSON(source)
}

// Digit returns the ith 4-bit digit in the NodeID. If i >= 32, Digit panics.
func (id NodeID) Digit(i int) byte {
	return wendy.NodeID(id).Digit(i)
}
