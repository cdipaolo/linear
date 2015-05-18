package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func abs(n float64) float64 {
	if n < 0 {
		return -1 * n
	}

	return n
}

func h(theta, x []float64) float64 {
	if len(theta) != len(x)+1 {
		panic(fmt.Sprintf("Theta <len: %v> should be 1 larger than x <len: %v>", len(theta), len(x)))
	}

	// add constant term
	hypothesis := theta[0]

	// take dot product of x and theta for rest
	for i := range x {
		hypothesis += theta[i+1] * x[i]
	}

	return hypothesis
}

// J corresponds to the least squares cost function of the data
func J(theta []float64, xData [][]float64, y []float64) float64 {
	var cost float64
	for i := range xData {
		cost += (h(theta, xData[i]) - y[i]) * (h(theta, xData[i]) - y[i])
	}

	return cost / float64(2*len(xData))
}

// dj is the partial derivative of J(θ) with respect to θ_j
func dj(j int64, theta []float64, xData [][]float64, y []float64) float64 {
	var cost float64
	if j == 0 {
		for i := range xData {
			cost += (h(theta, xData[i]) - y[i])
		}
	} else {
		for i := range xData {
			cost += (h(theta, xData[i]) - y[i]) * xData[i][j-1]
		}
	}

	return cost / float64(len(xData))
}

// regression takes in the matrix of feature data, a vector of result data,
//     and the learning rate alpha and returns the parameter vector θ as well
//     as the final cost vector J(θ)
func regression(xData [][]float64, y []float64, alpha float64) ([]float64, float64, error) {
	theta := make([]float64, len(xData[0])+1)

	j := J(theta, xData, y)
	jNew := float64(0) // update in the loop later

	iterations := 0
	for abs(jNew-j) > 1e-6 && iterations < 100000 {
		// check if diverging
		if iterations > 100 && jNew > 1e6 {
			fmt.Printf("\nDIFFERENCE: %v\n", jNew)
			return nil, 0, fmt.Errorf("Error: regression with learning rate <α = %v> diverging", alpha)
		}

		thetaNew := make([]float64, len(theta))
		for i := range theta {
			thetaNew[i] = theta[i] - alpha*dj(int64(i), theta, xData, y)
		}

		j = jNew
		jNew = J(thetaNew, xData, y)
		theta = thetaNew

		iterations++
	}

	return theta, jNew, nil
}

func main() {
	// get the CSV filepath from the 'f' flag
	// eg. `linear -f='~/data.csv'
	path := flag.String("f", "/Users/cdipaolo/Data/LSD_Math/lsd.csv", "CSV filepath for data to be linearly regressed")
	fmt.Printf("CSV File Path: %v\n", *path)

	// open the csv file
	dataFile, err := os.Open(*path)
	if err != nil {
		panic(fmt.Sprintf("Error opening file: %v", err))
	}
	defer dataFile.Close()

	reader := csv.NewReader(dataFile)
	var data [][]float64
	var y []float64

	// read data from csv
	for {
		record, err := reader.Read()

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error reading CSV file: %v", err)
			return
		}
		var row []float64

		// convert all row values to flaot64's
		for i, val := range record {
			float, err := strconv.ParseFloat(val, 64)
			if err != nil {
				fmt.Printf("Error converting data value <%v> to float64: %v", val, err)
				break
			}

			if i < len(record)-1 {
				row = append(row, float)
			} else {
				y = append(y, float)
			}
		}

		data = append(data, row)
	}

	// print converted feature data and parameter data
	fmt.Printf("Data: %v\ny: %v", data, y)

	// regress data at multiple alpha values
	var theta []float64
	j := float64(1e6)

	alpha := float64(10)
	for i := 0; i < 10; i++ {
		thetaNew, jNew, err := regression(data, y, alpha)
		alpha /= 3

		if err != nil {
			fmt.Printf("\nError regressing data:\n\t%v\n", err)
			continue
		}

		fmt.Printf("\n\nj(θ) = %v\nTheta with <α = %v>: %v", jNew, alpha, thetaNew)

		// if new j is less, set current best hypothesis
		if jNew < j {
			j = jNew
			theta = thetaNew
		}
	}

	// print final hypothesis
	fmt.Printf("\n\nFinal theta with j(θ) = %v\n%v\n", j, theta)
}
