package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    
    "github.com/doroshka12/weather-app/internal/pkg/weather"
)

func main() {
    service := weather.NewOpenMeteoService()
    
    http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
        lat := r.URL.Query().Get("lat")
        lon := r.URL.Query().Get("lon")
        
        if lat == "" || lon == "" {
            http.Error(w, "missing lat or lon", http.StatusBadRequest)
            return
        }
        
        var latF, lonF float64
        fmt.Sscanf(lat, "%f", &latF)
        fmt.Sscanf(lon, "%f", &lonF)
        
        data, err := service.GetWeather(latF, lonF)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(data)
    })
    
    fmt.Println("HTTP server running on :8080")
    http.ListenAndServe(":8080", nil)
}