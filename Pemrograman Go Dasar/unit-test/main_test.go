package main

import "testing"
import "github.com/stretchr/testify/assert"

// cara menjalankan unit test dan assert:
// go test main.go main_test.go -v

// cara menjalankan benchmark:
// go test main.go main_test.go -bench=.

var (
	kubus  Kubus = Kubus{4}
	volumeSeharusnya float64 = 64
	luasSeharusnya   float64 = 96
	kelilingSeharusnya float64 = 48
)

func TestHitungVolume(t *testing.T) {
	t.Logf("Volume: %.2f", kubus.Volume())

	if kubus.Volume() != volumeSeharusnya {
		t.Errorf("SALAH! Seharusnya: %.2f", volumeSeharusnya)
	}
}

func TestHitungLuas(t *testing.T) {
	t.Logf("Luas: %.2f", kubus.Luas())

	if kubus.Luas() != luasSeharusnya {
		t.Errorf("SALAH! Seharusnya: %.2f", luasSeharusnya)
	}
}

func TestHitungKeliling(t *testing.T) {
	t.Logf("Keliling: %.2f", kubus.Keliling())

	if kubus.Keliling() != kelilingSeharusnya {
		t.Errorf("SALAH! Seharusnya: %.2f", kelilingSeharusnya)
	}
}

// Benchmark
func BenchmarkHitungLuas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kubus.Luas()
	}
}

// penggunaan package assert
func TestHitungVolume1(t *testing.T) {
    assert.Equal(t, kubus.Volume(), volumeSeharusnya, "perhitungan volume salah")
}

func TestHitungLuas1(t *testing.T) {
    assert.Equal(t, kubus.Luas(), luasSeharusnya, "perhitungan luas salah")
}

func TestHitungKeliling1(t *testing.T) {
    assert.Equal(t, kubus.Keliling(), kelilingSeharusnya, "perhitungan keliling salah")
}