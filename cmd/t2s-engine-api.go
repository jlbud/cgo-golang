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
	fmt.Println("cfg filename is ", cfg.FileName)
	fmt.Println("cfg model is ", cfg.Model)

	f, err := os.Open(cfg.FileName)
	if err != nil {
		fmt.Printf("open file err: %s\n", err)
		return
	}
	defer f.Close()
	t, err := engine.CreateInstance(cfg.Model)
	if err != nil {
		fmt.Printf("CreateInstance err:%v\n", err)
	}
	defer func() {
		err := t.Destroy()
		if err != nil {
			fmt.Printf("Destroy err:%v\n", err)
			return
		}
	}()
	countTotal, _ := StatsCount(cfg.FileName)
	count := 0
	br := bufio.NewReader(f)
	for {
		output := ""
		data, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if len(string(data)) == 0 {
			continue
		}
		fmt.Println("json: ", string(data))
		count++
		if count == countTotal {
			output = t.Process(string(data), 1)
		} else {
			output = t.Process(string(data), 0)
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
