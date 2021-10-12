package main

import (
	"fmt"
)

func main() {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)

}

// Паттерн Строитель также используется, когда нужный продукт сложный и требует нескольких шагов для построения.
// В таких случаях несколько конструкторных методов подойдут лучше, чем один громадный конструктор.
// При использовании пошагового построения объектов потенциальной проблемой является выдача клиенту частично построенного нестабильного продукта.
// Паттерн "Строитель" скрывает объект до тех пор, пока он не построен до конца.
