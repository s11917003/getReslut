package mt19937rand

import (
	"math/rand"
	"sync"
	"time"

	"github.com/seehuhn/mt19937"
)

var rng *rand.Rand
var mu sync.Mutex

func Init() {
	rng = rand.New(mt19937.New()) //MT19937 梅森旋轉法亂數
	rng.Seed(time.Now().UnixNano())
}

func Intn(n int) int {
	var res int
	mu.Lock() //避免多執行序打亂mt19937內部指標
	res = rng.Intn(n)
	defer mu.Unlock()
	return res
}

func Float32() float32 {
	var res float32
	mu.Lock() //避免多執行序打亂mt19937內部指標
	res = rng.Float32()
	defer mu.Unlock()
	return res
}
