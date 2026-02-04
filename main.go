package main

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"img_url"`
}

var productList []Product

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	encoder := json.NewEncoder(w)
	err := encoder.Encode(productList)
	if err != nil {
		http.Error(w, "Failed to encode products", http.StatusInternalServerError)
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/products", getProducts)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}

func init() {
	// Initialization code here
	prd1 := Product{ID: 1, Title: "Product 1", Description: "Description 1", Price: 10.0, ImgURL: "http://example.com/img1.jpg"}
	prd2 := Product{ID: 2, Title: "Product 2", Description: "Description 2", Price: 20.0, ImgURL: "http://example.com/img2.jpg"}
	prd3 := Product{ID: 3, Title: "Product 3", Description: "Description 3", Price: 30.0, ImgURL: "http://example.com/img3.jpg"}
	prd4 := Product{ID: 4, Title: "Product 4", Description: "Description 4", Price: 40.0, ImgURL: "http://example.com/img4.jpg"}
	prd5 := Product{ID: 5, Title: "Product 5", Description: "Description 5", Price: 50.0, ImgURL: "http://example.com/img5.jpg"}
	prd6 := Product{ID: 6, Title: "Product 6", Description: "Description 6", Price: 60.0, ImgURL: "http://example.com/img6.jpg"}
	productList = append(productList, prd1, prd2, prd3, prd4, prd5, prd6)
}
