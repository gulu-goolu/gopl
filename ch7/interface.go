package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

type Celsius float64
type wrapper struct {
	Celsius
}

// 转换
func (f *wrapper) Set(s string) error {
	var val float64
	fmt.Sscanf(s, "%f", &val)
	f.Celsius = Celsius(val)

	return nil
}

func (f *wrapper) String() string {
	return fmt.Sprintf("%f", f.Celsius)
}

// 注册函数
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := wrapper{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var cel = CelsiusFlag("temp", 0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*cel)
}
