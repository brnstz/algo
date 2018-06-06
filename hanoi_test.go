package algo

import (
	"log"
	"testing"
)

func TestHanoi(t *testing.T) {
	moves := TowersOfHanoi(10)
	log.Println(moves)
	t.Fatal()

}
