package http

import "testing"

func BenchmarkRun1(b *testing.B) {
	if err := Run1(b.N, 0); err != nil {
		b.Fatal(err)
	}
}

//func BenchmarkRun2(b *testing.B) {
//	if err := Run2(b.N, 0); err != nil {
//		b.Fatal(err)
//	}
//}

func TestRun2(t *testing.T) {
	fatalIfError(t, Run2(100, 0))
	//fatalIfError(t, Run2(500, 0))
}

func fatalIfError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
