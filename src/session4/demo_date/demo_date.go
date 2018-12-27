package demo_date

import (
	"fmt"
	"time"
)

func Demo() {
	today := time.Now()
	fmt.Println("today:", today)
	fmt.Println("year:", today.Year())

	month := today.Month()
	fmt.Println("month:", month)
	fmt.Println("month:", int(month))

	fmt.Println("day: ", today.Day())
}

func Demo2() {
	today := time.Now()
	fmt.Println(today.Format("02/01/2006 15:04:05"))
}

func Demo3() {
	s := "27/12/2018"
	d, err := time.Parse("02/01/2006", s)
	if err != nil {
		fmt.Println("Error")
	} else {
		fmt.Println(d.Format("02/01/2006"))
	}
}

func Demo4() {
	today := time.Now()
	time1 := today.AddDate(0, 0, 2)
	fmt.Println("time 1:", time1.Format("02/01/2006"))
	time2 := today.Add(2 * 24 * time.Hour)
	fmt.Println("time 2:", time2.Format("02/01/2006"))
}

/*
Cho 1 ngay sinh kieu chuoi:
1. Khai bao kiem tra hom nay co phai la sinh nhat ko?
2. Khai bao ham tinh tuoi dua vao ngay sinh dc cung cap
*/

func CheckBirthday(bt string) {
	today := time.Now()
	todayDay := today.Day()
	todayMonth := int(today.Month())

	d, err := time.Parse("02/01/2006", bt)
	if err != nil {
		fmt.Println("Error")
	} else {
		if int(d.Month()) == todayMonth && d.Day() == todayDay {
			fmt.Println("Happy Birth Day!!!")
		}
	}
}

func AgeCal(_birth string) {
	today := time.Now()
	birth, err := time.Parse("02/01/2006", _birth)
	if err != nil {
		fmt.Println("Error")
	} else {
		age := today.Year() - birth.Year()
		fmt.Println("Ban", age, "tuoi")
	}
}
