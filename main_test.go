package main

import "testing"

func BenchmarkIftrue(b *testing.B)  { Iftrue(b.N) }
func BenchmarkIffalse(b *testing.B) { Iffalse(b.N) }
