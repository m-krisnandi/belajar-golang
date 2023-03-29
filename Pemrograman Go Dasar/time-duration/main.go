package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	time.Sleep(5 * time.Second)

	duration := time.Since(start)

	fmt.Println("time elapsed in seconds: ", duration.Seconds())
	fmt.Println("time elapsed in minutes: ", duration.Minutes())
	fmt.Println("time elapsed in hours: ", duration.Hours())

	// kalkulasi durasi antara dua objek waktu
	t1 := time.Now()
	time.Sleep(5 * time.Second)
	t2 := time.Now()

	duration = t2.Sub(t1)

	fmt.Println("time elapsed in seconds: ", duration.Seconds())
	fmt.Println("time elapsed in minutes: ", duration.Minutes())
	fmt.Println("time elapsed in hours: ", duration.Hours())

	// konversi angka ke time.Duration
	// 12 * time.Minute             // 12 menit
	// 65 * time.Hour                 // 65 jam
	// 150000 * time.Milisecond     // 150k milidetik atau 150 detik
	// 45 * time.Microsecond         // 45 microdetik
	// 233 * time.Nanosecond         // 233 nano detik

	// // Jika ditampung ke variabel terlebih dahulu, maka pastikan tipe variabelnya adalah time.Duration.
	// var n time.Duration = 5
	// duration2 := n * time.Second

	// // atau casting konversi data numerik ke time.Duration
	// n2 := 5
	// duration3 := time.Duration(n2) * time.Second
}