package main

import (
	"bufio"
	"bytes"
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
	fmt.Println("filename", cfg.FileName)
	fmt.Println("model", cfg.Model)

	f, err := os.Open(cfg.FileName)
	if err != nil {
		fmt.Printf("open file err: %s\n", err)
		return
	}
	defer f.Close()
	t := engine.CreateInstance(cfg.Model)
	defer func() {
		err := t.Destroy()
		if err != nil {
			fmt.Printf("Destroy err:%v\n", err)
			return
		}
	}()
	countTotal, _ := StatsCount(cfg.FileName)
	count := 0
	output := ""
	br := bufio.NewReader(f)
	for {
		data, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		count++
		if count == countTotal {
			t.Process(string(data), 1)
		} else {
			t.Process(string(data), 0)
		}
		fmt.Println("output data is ", output)
	}
}

func StatsCount(fileName string) (c int, err error) {
	fi, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file err: %s\n", err)
		return
	}
	// 32K cache
	buf := make([]byte, 32*1024)
	count := 1
	lineSep := []byte{'\n'}
	for {
		c, err := fi.Read(buf)
		count += bytes.Count(buf[:c], lineSep)
		switch {
		case err == io.EOF:
			return count, nil
		case err != nil:
			return count, err
		}
	}
}