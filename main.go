package main

import (
	"fmt"

	tugas2 "sandboxhsi3.com/golang/tugas"
)

func main() {
	nikArn241, err := tugas2.GeneratorNIK("male", 2024, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(nikArn241)
	}

	nikArn192, err := tugas2.NikBerikut("ARN192-00051", 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(nikArn192)
	}

	nikArn151, err := tugas2.NikBerikut("ARN151-02024", 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(nikArn151)
	}

	nikArt211, err := tugas2.GeneratorNIK("female", 2021, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(nikArt211)
	}

	nikArt161, err := tugas2.NikBerikut("ART161-01076", 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(nikArt161)
	}

	nikArt232, err := tugas2.NikBerikut("ART232-00376", 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(nikArt232)
	}

	niks := append(nikArn241, nikArn192...)
	niks = append(niks, nikArn151...)
	niks = append(niks, nikArt232...)
	niks = append(niks, nikArt161...)
	niks = append(niks, nikArt211...)
	fmt.Println(tugas2.KelompokHalaqah(niks))
}
