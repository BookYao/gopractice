/**
 * @Author: BookYao
 * @Description: 摄氏度和华氏转换的命令行方法
 * @File:  celsiusFlag
 * @Version: 1.0.0
 * @Date: 2020/8/23 23:41
 */

package main

import (
	"flag"
	"fmt"
)

type celsius float64

type celsiusFlag struct  {
	celsius
}

type Value interface {
	 String() string
	 Set(string) error
}

func Celsius(c celsius) celsius {
	return (c*9/5+32)
}

func FToC(f celsius) celsius {
	return (f-32)*5/9
}

func (f *celsiusFlag) String() string {
	return fmt.Sprint(f.celsius)
}

func (f *celsiusFlag) Set(s string) error {
	var value celsius
	var unit string

	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C":
	f.celsius = Celsius(value)
	return nil
	case "F":
	f.celsius = FToC(celsius(value))
	return nil
	}
	return fmt.Errorf("invalid temperature: %q", s)
}

func CelsiusFlag(name string, value celsius, usage string) *celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.celsius
}

/* Usage: .celsiutFlag -temp  80C */
func main() {
	var temp celsiusFlag
	var s string

	err := temp.Set("20C")
	if err != nil {
		fmt.Printf("temperature translate error! %v\n", err)
		return
	}
	fmt.Println(temp.celsius)
	s = temp.String()
	fmt.Printf("s: %s\n", s)

	err = temp.Set("100F")
	if err != nil {
		fmt.Printf("temperature translate error! %v\n", err)
		return
	}
	fmt.Println(temp.celsius)
	s = temp.String()
	fmt.Printf("s: %s\n", s)

	var Temp = CelsiusFlag("temp", 20.0, "The temperatue")
	flag.Parse()
	fmt.Println(*Temp)
}

  