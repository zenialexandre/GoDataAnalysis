package main

import (
	"fmt"
	"strings"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func main() {
	var asterisk string = "*"

	gotaSeries()
	fmt.Println(strings.Repeat(asterisk, 45))
	gotaDataFrame()
	fmt.Println(strings.Repeat(asterisk, 45))
	gotaDataFrameStruct()
}

func gotaSeries() {
	// with a slice
	fmt.Println(series.New([]string{"z", "y", "d", "e"}, series.String, "column"))

	// with a map
	a := map[string]series.Type{
		"A": series.String,
		"D": series.Bool,
	}
	fmt.Println(a)
}

func gotaDataFrame() {
	df := dataframe.New(
		series.New([]string{"a", "b", "c", "d", "e"}, series.String, "alphas"),
		series.New([]int{5, 4, 3, 2, 1}, series.Int, "numbers"),
		series.New([]string{"a1", "b2", "c3", "d4", "e5"}, series.String, "alnums"),
		series.New([]bool{true, false, true, true, false}, series.Bool, "state"),
	)
	fmt.Println(df)
}

func gotaDataFrameStruct() {
	type Dog struct {
		Name       string
		Color      string
		Height     int
		Vaccinated bool
	}

	dogs := []Dog{
		{"Buster", "Black", 56, false},
		{"Jake", "White", 61, false},
		{"Bingo", "Brown", 50, true},
		{"Gray", "Cream", 68, false},
	}

	dogsDf := dataframe.LoadStructs(dogs)
	fmt.Println(dogsDf)
	queryDogsDf(dogsDf)        // querying a DataFrame
	queryDogsDfColumns(dogsDf) // querying columns
}

func queryDogsDf(dogsDf dataframe.DataFrame) {
	fmt.Printf("Dimensions: ")
	fmt.Println(dogsDf.Dims())
	fmt.Print("Types: ")
	fmt.Println(dogsDf.Types())
	fmt.Print("Names: ")
	fmt.Println(dogsDf.Names())
	fmt.Print("Number of rows: ")
	fmt.Println(dogsDf.Nrow())
	fmt.Print("Number of columns: ")
	fmt.Println(dogsDf.Ncol())
}

func queryDogsDfColumns(dogsDf dataframe.DataFrame) {
	col := dogsDf.Col("Name") // selecting a column

	fmt.Printf("Verifying column values: %v\n", col.IsNaN())
	fmt.Printf("Getting column mean: %v\n", col.Mean())
	fmt.Printf("Copy of the column values: %v\n", col.Copy())
	fmt.Printf("There is a null value? %v\n", col.HasNaN())
	fmt.Printf("Records: %v\n", col.Records())
}
