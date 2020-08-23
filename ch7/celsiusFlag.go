/**
 * @Author: BookYao
 * @Description:
 * @File:  celsiusFlag
 * @Version: 1.0.0
 * @Date: 2020/8/23 23:41
 */

package main

import (
	"fmt"
)

type celsius float64

type celsiusFlag struct  {
	celsius
}

/*type Value interface {
	 String() string
	 Set(string) error
}*/

func Celsius(c celsius) celsius {
	return (c*9/5+32)
}

func FToC(f celsius) celsius {
	return (f-32)*5/9
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

func main() {
	var temp celsiusFlag
	err := temp.Set("20C")
	if err != nil {
		fmt.Printf("temperature translate error! %v\n", err)
		return
	}
	fmt.Println(temp.celsius)

	err = temp.Set("100F")
	if err != nil {
		fmt.Printf("temperature translate error! %v\n", err)
		return
	}
	fmt.Println(temp.celsius)

}

  