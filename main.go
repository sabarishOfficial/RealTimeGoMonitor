package main

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

// Monitor Variables
type Monitor struct {
	CpuMetrics    float64
	MemoryMetrics float64
	Message       string
}

// Global function for error handling
func checkErrors(err error) {
	if err != nil {
		panic(err)
	}
}

// index page in monitor metrics
func homePage(w http.ResponseWriter) {
	tmp := template.Must(template.ParseFiles("templates/index.html"))
	err := tmp.Execute(w, tmp)
	if err != nil {
		return
	}
}

// collection of metrics
func getData(w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")

	cpumetric, err := cpu.Percent(0, false)
	checkErrors(err)
	memmetric, err := mem.VirtualMemory()
	checkErrors(err)
	message := ""

	if cpumetric[0] > 80 || memmetric.UsedPercent > 80 {
		message = "Cpu High Usage !!"
	}

	data := Monitor{
		CpuMetrics:    cpumetric[0],
		MemoryMetrics: memmetric.UsedPercent,
		Message:       message,
	}

	err = json.NewEncoder(w).Encode(data)
	checkErrors(err)
}

// routePath handler functions
func indexHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.RequestURI() {
	case "/":
		homePage(w)
	case "/metrics":
		getData(w)
	}
}

// main functions
func main() {
	http.HandleFunc("/", indexHandler)

	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))

}
