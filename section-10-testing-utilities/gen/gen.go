package gen

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var m sync.Mutex
var count = 0

var firstNames = []string{
	"Jon",
	"Bob",
	"Mark",
	"Jimmy",
	"Holly",
	"Alice",
	"Eve",
}

func Email() string {
	m.Lock()
	defer m.Unlock()
	rand.Seed(time.Now().UnixNano())
	name := firstNames[rand.Intn(len(firstNames)-1)]
	ret := fmt.Sprintf("%s-%d@example.com", name, count)
	count++
	return ret
}
