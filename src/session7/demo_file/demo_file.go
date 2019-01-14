package demo_file

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"session7/entities"
	"strconv"
	"strings"
)

func Demo1() {
	file, err := os.Create("demo_file/a.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(file)
	}
}

func Demo2() {
	file, err := os.Stat("demo_file/a.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Name: ", file.Name())
		fmt.Println("Size(byte): ", file.Size())
		fmt.Println("Permission: ", file.Mode())
		fmt.Printf("Perm: %04o", file.Mode().Perm())
	}
}

func Demo3() {
	err := os.Remove("demo_file/a.txt")
	if err != nil {
		fmt.Println(err)
	}
}

func Demo4() {
	result, err := ioutil.ReadFile("demo_file/a.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(result))
	}
}

func Demo5() {
	file, err := os.Open("demo_file/a.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
		}
	}
	file.Close()
}

func Demo6() {
	file, err := os.Open("demo_file/products.csv")
	if err != nil {
		fmt.Println(err)
	} else {
		var products []entities.Product
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			result := strings.Split(line, ",")
			price, _ := strconv.ParseFloat(result[2], 64)
			quantity, _ := strconv.ParseInt(result[3], 10, 64)
			product := entities.Product{
				Id:       result[0],
				Name:     result[1],
				Price:    price,
				Quantity: quantity,
			}
			products = append(products, product)
		}
		fmt.Println("Product List:")
		for _, product := range products {
			fmt.Println(product.ToString())
			fmt.Println("--------------------------")
		}
	}
	file.Close()
}

func Demo7() {
	file, err := os.Create("demo_file/b.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		file.WriteString("Hello World!!!")
	}
	file.Close()
}

//append string into file
func Demo8() {
	file, err := os.OpenFile("demo_file/b.txt", os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Fprintf(file, "\nHi")
	}
	file.Close()
}
