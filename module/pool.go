package module

import (
	"log"
	"time"

	"github.com/panjf2000/ants/v2"
)

type goPool struct {
	pool *ants.Pool
	init bool
}

var GoPool = NewGoPool()

func NewGoPool() *goPool {
	g := goPool{}
	var err error
	g.pool, err = ants.NewPool(5000, ants.WithOptions(ants.Options{Nonblocking: true}), ants.WithExpiryDuration(30*time.Second))
	if err != nil {
		log.Println(err)
		g.init = false
	}
	g.init = true
	return &g
}
func (g *goPool) GetCount() int {
	return g.pool.Running()
}
func (g *goPool) Run(task func()) {
	if g.init {
		log.Println("[*] Thread Count: ", g.GetCount())
		if err := g.pool.Submit(task); err != nil {
			log.Println(err)
			g.pool.Release()
			g.pool.Reboot()
		}
	} else {
		go task()
	}
}
func (g *goPool) DisablePool() {
	g.pool = nil
	g.init = false
}
