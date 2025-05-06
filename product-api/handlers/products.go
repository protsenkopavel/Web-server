package handlers

import (
	"awesomeProject/data"
	"log"
	"net/http"
)

type Product struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Product) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetAllProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}
