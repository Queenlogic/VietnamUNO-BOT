package main

import (
	"fmt"
)

func main() {
	REDCARDS := [10]int {0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	BLUECARDS := [10]int {0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	GREENCARDS := [10]int {0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	YELLOWCARDS := [10]int {0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	SPECIALCOLORCARDS := [5]string {"Đổi Chiều", "Cộng 2", "Cấm", "Bay màu", "Đừng dừng lại"}
	BLACKCARDS := [8]string {"Tạo nghiệp", "Vận đổi sao dời", "Lầy", "Nín", "Cười bò", "Nghiệp quật", "Wild Card", "Cộng 2x2"}
}

func AddCard(cardtype string, color string, number int) {
	if (cardtype == "Color Card") {
		switch color {
			case "Red" : REDCARDS = append(REDCARDS, number)
			case "Blue" : BLUECARDS = append(BLUECARDS, number)
			case "Green" : GREENCARDS = append(GREENCARDS, number)
			case "Yellow" : YELLOWCARDS = append(YELLOWCARDS, number)
			default : fmt.Println("Kiểm tra lại thông số bận điền")
		}
	}
	
	if (cardtype == "Special Color Card") {
		switch color {
			case "Đổi chiều" : SPECIALCOLORCARDS = append(SPECIALCOLORCARDS, "Đổi chiều")
			case "Cộng 2" : SPECIALCOLORCARDS = append(SPECIALCOLORCARDS, "Cộng 2")
			case "Cấm" : SPECIALCOLORCARDS = append(SPECIALCOLORCARDS, "Cấm")
			case "Bay màu" : SPECIALCOLORCARDS = append(SPECIALCOLORCARDS, "Bay màu")
			case "Đừng dừng lại" : SPECIALCOLORCARDS = append(SPECIALCOLORCARDS, "Đừng dừng lại")
			default : fmt.Println("Kiểm tra lại thông số bận điền")
		}
	}

	if (cardtype == "Black Card") {
		switch color {
		   case "Tạo nghiệp" : BLACKCARDS = append(BLACKCARDS, "Tạo nghiệp")
		   case "Vận đổi sao dời" : BLACKCARDS = append(BLACKCARDS, "Vận đổi sao dời")
		   case "Lầy" : BLACKCARDS = append(BLACKCARDS, "Lầy")
		   case "Nín" : BLACKCARDS = append(BLACKCARDS, "Nín")
		   case "Cười bò" : BLACKCARDS = append(BLACKCARDS, "Cười bò")
		   case "Nghiệp quật" : BLACKCARDS = append(BLACKCARDS, "Nghiệp quật")
		   case "Wild Card" : BLACKCARDS = append(BLACKCARDS, "Wild Card")
		   case "Cộng 2x2" : BLACKCARDS = append(BLACKCARDS, "Cộng 2x2")
		   default : fmt.Println("Kiểm tra lại thông số bận điền")
		}
	}
}
