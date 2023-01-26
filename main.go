package main

import(
	"fmt"
	
)

func main()  {
	
	ns := GetNutritionalScore(NutritionalData{
		Energy: 				EnergyFromKcal(200), 
		Sugars: 				SugarGram(10),
		SaturatedFattyAcids: 	SaturatedFattyAcids(2) ,
		Sodium: 				SodiumMilligram(500),
		Fruits: 				FruitsPercent(60),
		Fiber: 					FiberGram(4),
		Protien: 				ProtienGram(2),
	}, Food)

	fmt.Printf("Nutrional Score: %d\n", ns.Value)
	fmt.Printf("NutriScore: %s\n", ns.GetNutriScore())
}