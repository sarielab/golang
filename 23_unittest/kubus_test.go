package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify"
)

var (
	kubus              Kubus   = Kubus{4}
	volumeSeharusnya   float64 = 64
	luasSeharusnya     float64 = 96
	kelilingSeharusnya float64 = 48
)

//HOW TO RUN
//go test
//or with benchmarking....
//go test kubus.go kubus_test.go -bench=.

/*
	testify
	- assert Berisikan	tools	standar	untuk	testing
	- http Berisikan	tools	untuk	keperluan	testing	http
	- mock Berisikan	tools	untuk	mocking	object
	- require Sama	seperti	assert,	hanya	saja	jika	terjadi	fail	pada	saat	test	akan
	- menghentikan	eksekusi	program
	- suite Berisikan	tools	testing	yang	berhubungan	dengan	struct	dan	method
*/

func TestHitungVolume(t *testing.T) {
	// t.Logf("Volume: %.2f", kubus.Volume())
	//
	// if kubus.Volume() != volumeSeharusnya {
	// 	t.Errorf("SALAH! harusnya %.2f", volumeSeharusnya)
	// }
	assert.Equal(t, kubus.Volume(), volumeSeharusnya, "Perhitungan volume salah")
}

func TestHitungLuas(t *testing.T) {
	t.Logf("Luas	:	%.2f", kubus.Luas())
	if kubus.Luas() != luasSeharusnya {
		t.Errorf("SALAH!	harusnya	%.2f", luasSeharusnya)
	}
}

func TestHitungKeliling(t *testing.T) {
	t.Logf("Keliling	:	%.2f", kubus.Keliling())
	if kubus.Keliling() != kelilingSeharusnya {
		t.Errorf("SALAH!	harusnya	%.2f", kelilingSeharusnya)
	}
}

func BenchmarkHitungLuas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kubus.Luas()
	}
}
