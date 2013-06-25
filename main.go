package main

import (
	"flag"
	"github.com/microcosm-cc/gosnippet/server"
	"log"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
)

var (
	// http://blog.golang.org/2011/06/profiling-go-programs.html
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
)

func main() {
	// Use as many procs as there are
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Enable CPU profiling if the flag is set
	// -cpuprofile=logfile.log
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()

		// Catch CTRL+C and stop the profiling
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			for sig := range c {
				log.Printf("captured %v, stopping profiler and exiting..", sig)
				pprof.StopCPUProfile()
				os.Exit(1)
			}
		}()
	}

	server.StartServer()
}
