package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/10is/pacmound"
	"github.com/10is/pacmound/agents"
	"github.com/10is/pacmound/agents/maximus"
)

func getPython() pacmound.Agent {
	return &agents.Random{}
}

func main() {
	var serve bool
	// flag.IntVar(&loops, "loops", 0, "")
	flag.BoolVar(&serve, "serve", true, "")
	flag.Parse()

	rand.Seed(time.Now().Unix())

	agent := &maximus.Agent{}
	// agent := &agents.Random{}
	fmt.Println(agent)

	getGopher := func() pacmound.Agent {
		return agent
	}

	// for i := 0; i < loops; i++ {
	// 	//fmt.Printf("loop %d\n", i)
	// 	pacmound.Level04(getGopher, getPython)
	// }
	fmt.Println(agent)

	if serve {
		mux := pacmound.NewGameMux(getGopher, getPython)
		log.Fatal(http.ListenAndServe(":8080", mux))
	}
}
