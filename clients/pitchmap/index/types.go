// This file was generated by nomgen.
// To regenerate, run `go generate` in this package.

package main

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

// ListOfPitch

type ListOfPitch struct {
	l types.List
}

type ListOfPitchIterCallback (func (p Pitch) (stop bool))

func NewListOfPitch() ListOfPitch {
	return ListOfPitch{types.NewList()}
}

func ListOfPitchFromVal(p types.Value) ListOfPitch {
	return ListOfPitch{p.(types.List)}
}

func (l ListOfPitch) NomsValue() types.List {
	return l.l
}

func (l ListOfPitch) Equals(p ListOfPitch) bool {
	return l.l.Equals(p.l)
}

func (l ListOfPitch) Ref() ref.Ref {
	return l.l.Ref()
}

func (l ListOfPitch) Len() uint64 {
	return l.l.Len()
}

func (l ListOfPitch) Empty() bool {
	return l.Len() == uint64(0)
}

func (l ListOfPitch) Get(idx uint64) Pitch {
	return PitchFromVal(l.l.Get(idx))
}

func (l ListOfPitch) Slice(idx uint64, end uint64) ListOfPitch {
	return ListOfPitch{l.l.Slice(idx, end)}
}

func (l ListOfPitch) Set(idx uint64, v Pitch) ListOfPitch {
	return ListOfPitch{l.l.Set(idx, v.NomsValue())}
}

func (l ListOfPitch) Append(v ...Pitch) ListOfPitch {
	return ListOfPitch{l.l.Append(l.fromElemSlice(v)...)}
}

func (l ListOfPitch) Insert(idx uint64, v ...Pitch) ListOfPitch {
	return ListOfPitch{l.l.Insert(idx, l.fromElemSlice(v)...)}
}

func (l ListOfPitch) Remove(idx uint64, end uint64) ListOfPitch {
	return ListOfPitch{l.l.Remove(idx, end)}
}

func (l ListOfPitch) RemoveAt(idx uint64) ListOfPitch {
	return ListOfPitch{(l.l.RemoveAt(idx))}
}

func (l ListOfPitch) fromElemSlice(p []Pitch) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v.NomsValue()
	}
	return r
}

// MapOfStringToListOfPitch

type MapOfStringToListOfPitch struct {
	m types.Map
}

type MapOfStringToListOfPitchIterCallback (func(k types.String, v ListOfPitch) (stop bool))

func NewMapOfStringToListOfPitch() MapOfStringToListOfPitch {
	return MapOfStringToListOfPitch{types.NewMap()}
}

func MapOfStringToListOfPitchFromVal(p types.Value) MapOfStringToListOfPitch {
	return MapOfStringToListOfPitch{p.(types.Map)}
}

func (m MapOfStringToListOfPitch) NomsValue() types.Map {
	return m.m
}

func (m MapOfStringToListOfPitch) Equals(p MapOfStringToListOfPitch) bool {
	return m.m.Equals(p.m)
}

func (m MapOfStringToListOfPitch) Ref() ref.Ref {
	return m.m.Ref()
}

func (m MapOfStringToListOfPitch) Empty() bool {
	return m.m.Empty()
}

func (m MapOfStringToListOfPitch) Len() uint64 {
	return m.m.Len()
}

func (m MapOfStringToListOfPitch) Has(p types.String) bool {
	return m.m.Has(p)
}

func (m MapOfStringToListOfPitch) Get(p types.String) ListOfPitch {
	return ListOfPitchFromVal(m.m.Get(p))
}

func (m MapOfStringToListOfPitch) Set(k types.String, v ListOfPitch) MapOfStringToListOfPitch {
	return MapOfStringToListOfPitchFromVal(m.m.Set(k, v.NomsValue()))
}

// TODO: Implement SetM?

func (m MapOfStringToListOfPitch) Remove(p types.String) MapOfStringToListOfPitch {
	return MapOfStringToListOfPitchFromVal(m.m.Remove(p))
}

func (m MapOfStringToListOfPitch) Iter(cb MapOfStringToListOfPitchIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(types.StringFromVal(k), ListOfPitchFromVal(v))
	})
}

// MapOfStringToString

type MapOfStringToString struct {
	m types.Map
}

type MapOfStringToStringIterCallback (func(k types.String, v types.String) (stop bool))

func NewMapOfStringToString() MapOfStringToString {
	return MapOfStringToString{types.NewMap()}
}

func MapOfStringToStringFromVal(p types.Value) MapOfStringToString {
	return MapOfStringToString{p.(types.Map)}
}

func (m MapOfStringToString) NomsValue() types.Map {
	return m.m
}

func (m MapOfStringToString) Equals(p MapOfStringToString) bool {
	return m.m.Equals(p.m)
}

func (m MapOfStringToString) Ref() ref.Ref {
	return m.m.Ref()
}

func (m MapOfStringToString) Empty() bool {
	return m.m.Empty()
}

func (m MapOfStringToString) Len() uint64 {
	return m.m.Len()
}

func (m MapOfStringToString) Has(p types.String) bool {
	return m.m.Has(p)
}

func (m MapOfStringToString) Get(p types.String) types.String {
	return types.StringFromVal(m.m.Get(p))
}

func (m MapOfStringToString) Set(k types.String, v types.String) MapOfStringToString {
	return MapOfStringToStringFromVal(m.m.Set(k, v))
}

// TODO: Implement SetM?

func (m MapOfStringToString) Remove(p types.String) MapOfStringToString {
	return MapOfStringToStringFromVal(m.m.Remove(p))
}

func (m MapOfStringToString) Iter(cb MapOfStringToStringIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(types.StringFromVal(k), types.StringFromVal(v))
	})
}

// Pitch

type Pitch struct {
	m types.Map
}

func NewPitch() Pitch {
	return Pitch{types.NewMap()}
}

func PitchFromVal(v types.Value) Pitch {
	return Pitch{v.(types.Map)}
}

// TODO: This was going to be called Value() but it collides with root.value. We need some other place to put the built-in fields like Value() and Equals().
func (s Pitch) NomsValue() types.Map {
	return s.m
}

func (s Pitch) Equals(p Pitch) bool {
	return s.m.Equals(p.m)
}

func (s Pitch) Ref() ref.Ref {
	return s.m.Ref()
}

func (s Pitch) Z() types.Float64 {
	return types.Float64FromVal(s.m.Get(types.NewString("Z")))
}

func (s Pitch) SetZ(p types.Float64) Pitch {
	return PitchFromVal(s.m.Set(types.NewString("Z"), p))
}

func (s Pitch) X() types.Float64 {
	return types.Float64FromVal(s.m.Get(types.NewString("X")))
}

func (s Pitch) SetX(p types.Float64) Pitch {
	return PitchFromVal(s.m.Set(types.NewString("X"), p))
}

