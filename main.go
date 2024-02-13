package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func generateNumbers(debut, fin int) []string {
	var numbers []string
	for i := debut; i <= fin; i++ {
		numbers = append(numbers, fmt.Sprintf("0%010d", i))
	}
	return numbers
}

func main() {
	telma34 := generateNumbers(3400000000, 3499999999)

	telma38 := generateNumbers(3800000000, 3899999999)

	data := map[string]interface{}{
		"telma34": telma34,
		"telma38": telma38,
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}

	file, err := os.Create("Telma.json")
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Erreur ", err)
		return
	}

	fmt.Println("success.")
}