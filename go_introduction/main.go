package main

import (
	"fmt"

	"github.com/go_introduction/countries"
	"github.com/go_introduction/users"
	"github.com/google/uuid"
)

func GoIntroductions() {
	//var gretting string = "Hello there"
	gretting := "Hello there"

	fmt.Println(gretting)

	//Asi se puede inicializar un array
	//var numbers [5]int
	//numbers[2] = 8

	//Usando un Slice -> es como un array pero este es dinamico

	numbers := []int{1, 5, 8, 10}
	numbers = append(numbers, 2, 7) //agregando 2 elementos al array numbers
	fmt.Println(numbers)

	for i := 0; i < 2; i++ {
		fmt.Println("Indice:", i)
	}

	for i, number := range numbers {
		fmt.Println("Indice:", i, "Value:", number)
	}

	for _, number := range numbers {
		fmt.Println("Value:", number)
	}

	//Tomando porciones del slice [inicio: final], y son excluyentes
	//[1 5 8 10 2 7]
	fmt.Println(numbers[1:3]) //[5 8]
	fmt.Println(numbers[1:])  //[5 8 10 2 7]
	fmt.Println(numbers[:4])  //[1 5 8 10]
	//
	fruits := []string{"Manzana", "Banano", "Sandia", "Melon"}
	//Mostrar unicamente las que no terminan en "a"
	for _, fruit := range fruits {
		index := len(fruit) - 1
		letter := fruit[index:]
		if letter == "a" {
			continue
		}

		fmt.Println("Fruit:", fruit)
	}

	//Usando un paquete externo

	uID := uuid.New()
	fmt.Println(uID)

	countries.Add("CHINA")
	countries.Add("MEXICO")
	countries.Add("USA")
	countries.Add("JAPON")
	countries.SetMyCountry("MEXICO")
	countries.List()

	//Estructuras

	martha := users.User{Id: 1, Name: "martha"}
	pedro := users.User{Id: 2, Name: "pedro"}
	juan := users.User{Id: 3, Name: "juan"}
	martha.SayHello()

	martha.AddFriend(pedro)
	martha.AddFriend(juan)

	martha.ListFriends()

	//Apuntadores
	println("========================")
	num1 := 1
	refNum := &num1
	fmt.Println("Valor de variable num1:", num1, "Ref:", &num1)
	fmt.Println("Referencia de num1:", refNum)
	fmt.Println("Referencia de num1:", *refNum)

	//Interfaces
	var frank users.Cashier = users.NewEmployee("frank")
	var robert users.Admin = users.NewEmployee("robert")

	total := frank.CalcularTotal(10, 10, 1)
	fmt.Println(total)

	robert.DeactivateEmployee(frank)
	total2 := frank.CalcularTotal(10, 10, 1)
	fmt.Println(total2)

}
