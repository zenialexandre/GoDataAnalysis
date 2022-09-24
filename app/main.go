package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func main() {
	var asterisk string = "*"

	gotaSeries()
	addSpace(asterisk)
	gotaDataFrame()
	addSpace(asterisk)
	gotaDataFrameStruct()
	addSpace(asterisk)
	readingJson()
	addSpace(asterisk)
	readingCSV()
	addSpace(asterisk)
	readingCSVFile()
	addSpace(asterisk)
	readingJsonFile()
}

func addSpace(asterisk string) {
	fmt.Println(strings.Repeat(asterisk, 45))
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

func readingJson() {
	newJson := `[
		{
			"Name": "John",
			"Age": 44,
			"Favorite Color": "Red",
			"Height(ft)": 6.7
		},
		{
			"Name": "Mary",
			"Age": 40,
			"Favorite Color": "Blue",
			"Height(ft)": 5.7
		}
	]`

	newJsonDf := dataframe.ReadJSON((strings.NewReader(newJson)))
	fmt.Println(newJsonDf)
}

func readingCSV() {
	newCSV := `
		Name, Age, Favorite Color, Height(ft)
		John, 44, Red, 6.7
		Mary, 40, Blue, 5.7`

	newCSVDf := dataframe.ReadCSV(strings.NewReader(newCSV))
	fmt.Println(newCSVDf)
}

func readingCSVFile() {
	file, err := os.Open("frequencyTableA.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	df := dataframe.ReadCSV(file)
	fmt.Println(df)

	// reading only the first two rows
	rows := df.Subset([]int{0, 2})
	fmt.Println(rows)

	// selecting columns by index and by name
	firstTwoColumns := df.Select([]int{0, 2})
	namedColumns := df.Select([]string{"Classes"})
	fmt.Println(firstTwoColumns)
	fmt.Println(namedColumns)
}

func readingJsonFile() {
	file, err := os.Open("Expense [MConverter.eu].json")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	df := dataframe.ReadJSON(file)
	fmt.Println(df)
}
