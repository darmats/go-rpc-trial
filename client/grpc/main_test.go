package main

import "testing"

func BenchmarkRun1(b *testing.B) {
	if err := Run1(b.N, 0); err != nil {
		b.Fatal(err)
	}
}

func BenchmarkRun2(b *testing.B) {
	if err := Run2(b.N, 0); err != nil {
		b.Fatal(err)
	}
}

func BenchmarkRun3(b *testing.B) {
	if err := Run3(b.N, 0); err != nil {
		b.Fatal(err)
	}
}

//func BenchmarkRun4(b *testing.B) {
//	if err := Run4(b.N, 0); err != nil {
//		b.Fatal(err)
//	}
//}

func TestRun4(t *testing.T) {
	fatalIfError(t, Run4(100, 0))
	//fatalIfError(t, Run4(500, 0))
}

func fatalIfError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
