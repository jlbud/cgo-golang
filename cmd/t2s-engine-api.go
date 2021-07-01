package main

import (
	"bufio"
	"cgo-golang/engine"
	"fmt"
	"github.com/guonaihong/clop"
	"io"
	"os"
)

// EngineConfig
type EngineConfig struct {
	FileName string `clop:"--file-name" usage:"vpr audio file" valid:"required"`
	Model    string `clop:"--model" usage:"vpr model path" valid:"required"`
}

func main() {
	cfg := &EngineConfig{}
	err := clop.Bind(cfg)
	if err != nil {
		fmt.Printf("clop err:%v\n", err)
		return
	}
	fmt.Println("cfg filename is ", cfg.FileName)
	fmt.Println("cfg model is ", cfg.Model)

	f, err := os.Open(cfg.FileName)
	if err != nil {
		fmt.Printf("open file err: %s\n", err)
		return
	}
	defer f.Close()
	b, err := engine.CreatePostBaseHandle(cfg.Model)
	if err != nil {
		fmt.Printf("CreatePostBaseHandle err:%v\n", err)
		return
	}
	defer func() {
		err := b.Destroy()
		if err != nil {
			fmt.Printf("Destroy err:%v\n", err)
			return
		}
	}()
	ps, err := engine.CreatePostSession(b)
	if err != nil {
		fmt.Printf("CreatePostSession err:%v\n", err)
		return
	}
	countTotal := StatsCount(cfg.FileName)
	count := 0
	br := bufio.NewReader(f)
	for {
		output := ""
		data, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if len(data) == 0 {
			continue
		}
		fmt.Printf("input json:\n%s\n", string(data))
		count++
		if count == countTotal {
			output, err = ps.Process(string(data), 1)
		} else {
			output, err = ps.Process(string(data), 0)
		}
		if err != nil {
			fmt.Printf("Process err:%v\n", err)
			return
		}
		fmt.Printf("output data is:\n%s\n", output)
		fmt.Printf("\n")
	}
}

func StatsCount(fileName string) int {
	fi, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("StatsCount open file err: %s\n", err)
		return 0
	}
	defer fi.Close()
	reader := bufio.NewReader(fi)
	count := 0
	for {
		str, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if len(str) == 0 {
			continue
		}
		count++
	}
	return count
}
