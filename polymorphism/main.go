package main

import "fmt"


func  main() {
	project1 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
	project2 := FixedBilling{projectName: "project 2", biddedAmount: 10000}
	project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}

	incomeStreams := []Income{project1, project2, project3}

	calculateNetIncome(incomeStreams)

	fmt.Println()
	//constant example
	const mine = 4.5
	fmt.Printf("this is the type: %T\n", mine)

	type theStruct struct {
		Name string
		Age int
	}
	newer := new(theStruct)
	newer.Name = "Ade"
	fmt.Println("the struct: ", newer)

}

type Income interface {
		calculate() int
		source() string
	}

	type FixedBilling struct {
		projectName string
		biddedAmount int
	}
	type TimeAndMaterial struct  {
		projectName string
		noOfHours int
		hourlyRate int
	}

func (fb FixedBilling) calculate() int  {
	return fb.biddedAmount
}
func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

//using polymorphism for calculation based on the array of variables of interface type
func calculateNetIncome(ic []Income) {
	var netincome int = 0
	for _, income := range ic {
		fmt.Printf("Income from %s = $%d\n", income.source(), income.calculate())
		netincome += income.calculate()
	}
	fmt.Printf("Net income of organization = $%d", netincome)
}
