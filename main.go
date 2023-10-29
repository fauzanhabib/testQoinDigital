package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var N, M int
	fmt.Print("Masukkan total pemain: ")
	fmt.Scan(&N)
	fmt.Print("Masukkan total dadu : ")
	fmt.Scan(&M)
	dadu(N,M)
}
func dadu(N,M int){
	// Inisialisasi generator angka acak dengan nilai seed yang berubah ( berdasarkan waktu saat ini).
	rand.Seed(time.Now().UnixNano())

	players := make([]int, N)
	scores := make([]int, N)

	// Inisialisasi pemain dengan dadu
	for i := 0; i < N; i++ {
		players[i] = M
	}

	//loop tidak terbatas
	for {
		for i := 0; i < N; i++ {
			if players[i] > 0 {
				for j := 0; j < players[i]; j++ {
					dice := rand.Intn(6) + 1
					fmt.Printf("Pemain %d melempar dadu %d\n", i+1, dice)
					switch dice {
					case 6:
						//Dadu angka 6 akan dikeluarkan dari permainan dan ditambahkan sebagai poin bagi pemain tersebut.
						scores[i]++
					case 1:
						//Dadu angka 1 akan diberikan kepada pemain yang duduk disampingnya.
						nextPlayer := (i + 1) % N
						players[nextPlayer]++
					}
				}
				players[i] = 0
			}
		}

		remainingPlayers := 0
		for i := 0; i < N; i++ {
			if players[i] > 0 {
				remainingPlayers++
			}
		}

		if remainingPlayers == 1 {
			var winner int
			for i := 0; i < N; i++ {
				if players[i] > 0 {
					winner = i
					break
				}
			}
			fmt.Printf("Pemain %d menang dengan skor %d poin!\n", winner+1, scores[winner])
			break
		}else{
			fmt.Printf("Pemain tidak ada yang menang,")
			break
		}
	}
}