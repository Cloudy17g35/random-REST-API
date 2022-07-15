package converters

import (
	"log"
	"strconv"
	"strings"
)


func TrimString(s string, c string) string{
	return strings.Trim(s, c)
}

func SplitString(
	s string, 
	sep string) []string {
	
	splited := strings.Split(s , sep)
	return splited
}


func ConvertArrayOfStringsToFloats(s[]string) ([]float64, error) {
	t := []float64{}
	var empty_slice []float64
	for _, item := range s {
		j, err := strconv.ParseFloat(item, 64)
		if err != nil {
			log.Fatal(err)
			return empty_slice, err
		}
		t = append(t, j)
	}
	return t, nil
}

func StrToInt(s string) (int, error){
	n, e := strconv.Atoi(s)
	if e != nil {
		log.Fatal(e)
	}
	return n, nil
}