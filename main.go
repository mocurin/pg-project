package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"sort"
	"sync"
	"syscall"
	"time"

	"github.com/mocurin/pg-project/internal"
)

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	go func() {
		http.ListenAndServe(":8080", r)
	}()

	BruteCnt := 0.0
	Brute := 0.0
	SeqCnt := 0.0
	Seq := 0.0
	ParCnt := 0.0
	Par := 0.0

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	defer func() {
		fmt.Printf("\nbrute: %fns;\nseq: %fns;\npar: %fns;\n", Brute/BruteCnt, Seq/SeqCnt, Par/ParCnt)
	}()

	go func() {
		<-c
		fmt.Printf("\nbrute: %fns;\nseq: %fns;\npar: %fns;\n", Brute/BruteCnt, Seq/SeqCnt, Par/ParCnt)
		os.Exit(1)
	}()

	// loop:
	for {
		fp := internal.Field(17).RandomPolynomial(129)
		wg := sync.WaitGroup{}

		fmt.Println(fp)

		var r1 []int
		wg.Add(1)
		go func() {
			defer wg.Done()
			start := time.Now()
			r1 = internal.FactorizeSequential(fp)
			Seq += float64(time.Since(start).Nanoseconds())
			SeqCnt += 1
			sort.Slice(r1, func(i, j int) bool {
				return r1[i] > r1[j]
			})
		}()

		var r2 []int
		wg.Add(1)
		go func() {
			defer wg.Done()
			start := time.Now()
			r2 = internal.FactorizeBrute(fp)
			Brute += float64(time.Since(start).Nanoseconds())
			BruteCnt += 1
			sort.Slice(r2, func(i, j int) bool {
				return r2[i] > r2[j]
			})
		}()

		var r3 []int
		wg.Add(1)
		go func() {
			defer wg.Done()
			start := time.Now()
			r3 = internal.FactorizeParralel(fp)
			Par += float64(time.Since(start).Nanoseconds())
			ParCnt += 1
			sort.Slice(r3, func(i, j int) bool {
				return r3[i] > r3[j]
			})
		}()

		wg.Wait()

		// if len(r1) != len(r2) || len(r2) != len(r3) {
		// 	break
		// }

		// fmt.Println(r1, r3)

		// for i := 0; i < len(r1); i++ {
		// 	if r1[i] != r2[i] || r2[i] != r3[i] {
		// 		break loop
		// 	}
		// }
	}
}
