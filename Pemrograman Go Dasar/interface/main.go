package main

import (
	"fmt"
	"math"
)

// embed interface
type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type hitung interface {
	// luas() float64
	// keliling() float64

	// embed
	hitung2d
	// hitung3d
}

type hitung2 interface {
	// luas() float64
	// keliling() float64

	// embed
	hitung2d
	hitung3d
}

// kubus memiliki method yang sama dengan persegi
type kubus struct {
	sisi float64
}

func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

func (k *kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}

func (k *kubus) keliling() float64 {
	return k.sisi * 12
}

// lingkaran dan persegi memiliki method yang sama
type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

// fungsi main
func main() {
	var bangunDatar hitung
	
	bangunDatar = persegi{10.0}

	fmt.Println("===== Persegi =====")
	fmt.Println("Luas\t\t:", bangunDatar.luas())
	fmt.Println("Keliling\t:", bangunDatar.keliling())

	bangunDatar = lingkaran{14.0}
	// mengakses jariJari() dari lingkaran
	var bangunLingkaran lingkaran = lingkaran{bangunDatar.(lingkaran).diameter}

	fmt.Println("===== Lingkaran =====")
	fmt.Println("Luas\t\t:", bangunDatar.luas())
	fmt.Println("Keliling\t:", bangunDatar.keliling())

	// fmt.Println("Jari-jari\t:", bangunDatar.(lingkaran).jariJari())
	// atau
	fmt.Println("Jari-jari\t:", bangunLingkaran.jariJari())

	// kubus
	var bangunRuang hitung2 = &kubus{4}

	fmt.Println("===== Kubus =====")
	fmt.Println("Luas\t\t:", bangunRuang.luas())
	fmt.Println("Keliling\t:", bangunRuang.keliling())
	fmt.Println("Volume\t\t:", bangunRuang.volume())
}