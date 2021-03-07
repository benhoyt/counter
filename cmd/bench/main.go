// Benchmark Counter on real-world data.

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"runtime/pprof"
	"sort"

	"github.com/benhoyt/counter"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: bench [-cpuprofile] map|counter\n")
		flag.CommandLine.PrintDefaults()
	}
	cpuprofile := flag.Bool("cpuprofile", false, "create a CPU profile")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		flag.CommandLine.Usage()
		os.Exit(1)
	}
	mode := args[0]

	if *cpuprofile {
		f, err := os.Create("cpuprofile")
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not create CPU profile: %v\n", err)
			os.Exit(1)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			fmt.Fprintf(os.Stderr, "could not start CPU profile: %v\n", err)
			os.Exit(1)
		}
		defer pprof.StopCPUProfile()
	}

	switch mode {
	case "map":
		doMap()

	case "counter":
		doCounter()

	default:
		fmt.Fprintf(os.Stderr, "invalid mode %q\n", mode)
		flag.CommandLine.Usage()
		os.Exit(1)
	}
}

func doMap() {
	offset := 0
	buf := make([]byte, 64*1024)
	counts := make(map[string]int)
	for {
		n, err := os.Stdin.Read(buf[offset:])
		if err != nil && err != io.EOF {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if n == 0 {
			break
		}
		chunk := buf[:offset+n]
		lastLF := bytes.LastIndexByte(chunk, '\n')
		process := chunk
		if lastLF != -1 {
			process = chunk[:lastLF]
		}

		start := -1
		for i, c := range process {
			if c >= 'A' && c <= 'Z' {
				c = c + ('a' - 'A')
				process[i] = c
			}
			if start >= 0 {
				if c == ' ' || c == '\n' {
					counts[string(process[start:i])]++
					start = -1
				}
			} else {
				if c != ' ' && c != '\n' {
					start = i
				}
			}
		}
		if start >= 0 && start < len(process) {
			counts[string(process[start:])]++
		}

		if lastLF == -1 {
			offset = 0
		} else {
			remaining := chunk[lastLF+1:]
			copy(buf, remaining)
			offset = len(remaining)
		}
	}

	var ordered []Count
	for word, count := range counts {
		ordered = append(ordered, Count{word, count})
	}
	sort.Slice(ordered, func(i, j int) bool {
		return ordered[i].Count > ordered[j].Count
	})
	for _, count := range ordered {
		fmt.Println(string(count.Word), count.Count)
	}
}

type Count struct {
	Word  string
	Count int
}

func doCounter() {
	offset := 0
	buf := make([]byte, 64*1024)
	var counts counter.Counter
	for {
		n, err := os.Stdin.Read(buf[offset:])
		if err != nil && err != io.EOF {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if n == 0 {
			break
		}
		chunk := buf[:offset+n]
		lastLF := bytes.LastIndexByte(chunk, '\n')
		process := chunk
		if lastLF != -1 {
			process = chunk[:lastLF]
		}

		start := -1
		for i, c := range process {
			if c >= 'A' && c <= 'Z' {
				c = c + ('a' - 'A')
				process[i] = c
			}
			if start >= 0 {
				if c == ' ' || c == '\n' {
					counts.Inc(process[start:i], 1)
					start = -1
				}
			} else {
				if c != ' ' && c != '\n' {
					start = i
				}
			}
		}
		if start >= 0 && start < len(process) {
			counts.Inc(process[start:], 1)
		}

		if lastLF == -1 {
			offset = 0
		} else {
			remaining := chunk[lastLF+1:]
			copy(buf, remaining)
			offset = len(remaining)
		}
	}

	ordered := counts.Items()
	sort.Slice(ordered, func(i, j int) bool {
		return ordered[i].Count > ordered[j].Count
	})
	for _, count := range ordered {
		fmt.Println(string(count.Key), count.Count)
	}
}
