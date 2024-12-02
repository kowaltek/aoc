package main

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// generate random data for bar chart
func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < 8; i++ {
		items = append(items, opts.LineData{Value: i * (7 - i)})
	}
	return items
}

func main() {
	// create a new bar instance
	bar := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Distances x time spent preparing",
		Subtitle: "For total time limit = 10ms",
	}))

	// Put data into instance
	bar.SetXAxis([]int{0, 1, 2, 3, 4, 5, 6, 7}).
		AddSeries("Distance", generateLineItems())
	// Where the magic happens
	f, _ := os.Create("bar.html")
	bar.Render(f)
}
