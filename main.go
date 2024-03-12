package main

import (
	"fmt"
	"strconv"
)

func main() {
	var operand1 string
	var operator string
	var operand2 string

	fmt.Scan(&operand1, &operator, &operand2)

	result, class, err := calculate(operand1, operator, operand2)
	if err != nil {
		fmt.Println("Error:", err)
	}
	if class == "Rome" {
		res := convertToArabic(result)
		fmt.Println("Result:", res)
	}
	if class == "Arabic" {
		fmt.Println("Result:", result)
	}
}

func calculate(op1, operator, op2 string) (int, string, error) {
	var class string
	operand1, operand2, err := convertOperandArabic(op1, op2)
	class = "Arabic"
	if err != nil {
		class = "Rome"
		fmt.Println(class)
		operand1, operand2, err = convertOperandRome(op1, op2)
		if err != nil {
			panic("not valid value")
		}
	}

	switch operator {
	case "+":
		return operand1 + operand2, class, nil
	case "-":
		return operand1 - operand2, class, nil
	case "*":
		return operand1 * operand2, class, nil
	case "/":
		if operand1 == 0 || operand2 == 0 {
			return 0, class, fmt.Errorf("division by zero")
		}
		return operand1 / operand2, class, nil
	default:
		return 0, class, fmt.Errorf("unsupported operator: %s", operator)
	}
}

func convertOperandArabic(op1, op2 string) (operand1, operand2 int, err error) {
	operand1, err = strconv.Atoi(op1)
	if err != nil {
		return 0, 0, err
	}

	operand2, err = strconv.Atoi(op2)
	if err != nil {
		return 0, 0, err
	}

	if operand1 > 10 || operand2 > 10 {
		panic("number greater than 10")
	}

	return operand1, operand2, err
}

func convertOperandRome(op1, op2 string) (operand1, operand2 int, err error) {
	romanMap := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
		"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}

	if val, ok := romanMap[op1]; ok {
		operand1 = val
	}

	if val, ok := romanMap[op2]; ok {
		operand2 = val
	}

	if operand1 == 0 || operand2 == 0 {
		err = fmt.Errorf("not found")
	}

	return operand1, operand2, err
}

func convertToArabic(res int) string {
	arabicToRoman := map[int]string{
		1: "I", 2: "II", 3: "III", 4: "IV", 5: "V",
		6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
		11: "XI", 12: "XII", 13: "XIII", 14: "XIV", 15: "XV",
		16: "XVI", 17: "XVII", 18: "XVIII", 19: "XIX", 20: "XX",
		21: "XXI", 22: "XXII", 23: "XXIII", 24: "XXIV", 25: "XXV",
		26: "XXVI", 27: "XXVII", 28: "XXVIII", 29: "XXIX", 30: "XXX",
		31: "XXXI", 32: "XXXII", 33: "XXXIII", 34: "XXXIV", 35: "XXXV",
		36: "XXXVI", 37: "XXXVII", 38: "XXXVIII", 39: "XXXIX", 40: "XL",
		41: "XLI", 42: "XLII", 43: "XLIII", 44: "XLIV", 45: "XLV",
		46: "XLVI", 47: "XLVII", 48: "XLVIII", 49: "XLIX", 50: "L",
		51: "LI", 52: "LII", 53: "LIII", 54: "LIV", 55: "LV",
		56: "LVI", 57: "LVII", 58: "LVIII", 59: "LIX", 60: "LX",
		61: "LXI", 62: "LXII", 63: "LXIII", 64: "LXIV", 65: "LXV",
		66: "LXVI", 67: "LXVII", 68: "LXVIII", 69: "LXIX", 70: "LXX",
		71: "LXXI", 72: "LXXII", 73: "LXXIII", 74: "LXXIV", 75: "LXXV",
		76: "LXXVI", 77: "LXXVII", 78: "LXXVIII", 79: "LXXIX", 80: "LXXX",
		81: "LXXXI", 82: "LXXXII", 83: "LXXXIII", 84: "LXXXIV", 85: "LXXXV",
		86: "LXXXVI", 87: "LXXXVII", 88: "LXXXVIII", 89: "LXXXIX", 90: "XC",
		91: "XCI", 92: "XCII", 93: "XCIII", 94: "XCIV", 95: "XCV",
		96: "XCVI", 97: "XCVII", 98: "XCVIII", 99: "XCIX", 100: "C",
	}

	if val, ok := arabicToRoman[res]; ok {
		return val
	}
	return ""
}
