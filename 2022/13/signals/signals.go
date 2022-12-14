package signals

import (
	"encoding/json"
	"log"
	"reflect"
)

type comparable struct {
	kind        reflect.Kind
	actualValue interface{}
	intValue    int
	listValue   []*comparable
}

func convertSingleIntToList(sign interface{}) (output []interface{}) {
	signFloat, isInt := sign.(float64)
	if isInt {
		return []interface{}{signFloat}
	}
	// assuming that it's already list
	for _, signElement := range sign.([]interface{}) {
		output = append(output, signElement)
	}
	return output
}

func compare(element1List, element2List []interface{}, indent string) string {
	log.Printf("%s- Compare '%+v' vs '%+v'", indent, element1List, element2List)
	if len(element1List) == 0 && len(element2List) > 0 {
		log.Printf("%s- (1)Left side ran out of items, so inputs are in the right order", indent)
		return "true"
	} else if len(element2List) == 0 && len(element1List) > 0 {
		log.Printf("%s- (2)Right side ran out of items, so inputs are NOT in the right order", indent)
		return "false"
	}
	var i int
	for i = range element1List {
		element1 := element1List[i]
		if i >= len(element2List) {
			log.Printf("%s- Right side ran out of items, so inputs are NOT in the right order", indent)
			return "false"
		}
		element2 := element2List[i]
		if reflect.DeepEqual(element1, element2) {
			log.Println("elements are deeply equal")
			continue
		}

		element1Type := reflect.TypeOf(element1)
		element2Type := reflect.TypeOf(element2)
		if element1Type.Kind() != element2Type.Kind() {
			if element1Type.Kind() == reflect.Float64 {
				element1 = convertSingleIntToList(element1)
			} else {
				element2 = convertSingleIntToList(element2)
			}

			compareResult := compare(convertSingleIntToList(element1), convertSingleIntToList(element2), indent+"    ")
			if compareResult == "next" {
				continue
			} else {
				return compareResult
			}
		} else if element1Type.Kind() == reflect.Slice && element2Type.Kind() == reflect.Slice {
			compareResult := compare(convertSingleIntToList(element1), convertSingleIntToList(element2), indent+"    ")
			if compareResult == "next" {
				continue
			} else {
				return compareResult
			}
		} else if element1Type.Kind() == reflect.Float64 && element2Type.Kind() == reflect.Float64 {
			log.Printf("%s- Compare '%+v' vs '%+v'", indent+"    ", element1, element2)
			if element1.(float64) < element2.(float64) {
				log.Printf("%s- Left side is smaller, so inputs are in the right order", indent+"        ")
				return "true"
			}
			if element1.(float64) > element2.(float64) {
				log.Printf("%s- Right side is smaller, so inputs are not in the right order", indent+"        ")
				return "false"
			}
		}

	}
	if i+1 <= len(element2List) {
		log.Printf("%s- Left side ran out of items, so inputs are in the right order", indent)
		return "true"
	}
	return "next"
}

//	func compare(element1, element2 interface{}) string {
//	   log.Printf("- Compare '%+v' vs '%+v'", element1, element2)
//		element1Type := reflect.TypeOf(element1)
//		element2Type := reflect.TypeOf(element2)
//
//
//		return "next"
//	}
func IsRightOrder(signalString1, signalString2 string) bool {
	var signal1, signal2 interface{}
	err := json.Unmarshal([]byte(signalString1), &signal1)
	if err != nil {
		log.Fatalf("unmarsalling '%s' - %v", signalString1, err)
	}

	err = json.Unmarshal([]byte(signalString2), &signal2)
	if err != nil {
		log.Fatalf("unmarsalling '%s' - %v", signalString2, err)
	}
	signalList1 := convertSingleIntToList(signal1)
	signalList2 := convertSingleIntToList(signal2)
	if compare(signalList1, signalList2, "") == "true" {
		return true
	}
	return false
}
