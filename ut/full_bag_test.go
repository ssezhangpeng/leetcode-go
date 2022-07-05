package ut

import (
	"fmt"
	"testing"
)

func TestFullBag(t *testing.T) {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	bag := 4

	price := fullBag(weight, value, bag)

	fmt.Println(price)
}
