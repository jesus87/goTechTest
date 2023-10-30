package main

type Resumen struct {
	Total         float64            `json:"total"`
	ComprasPorTDC map[string]float64 `json:"comprasPorTDC"`
	NoCompraron   int                `json:"nocompraron"`
	CompraMasAlta float64            `json:"compraMasAlta"`
}
type Compra struct {
	ClientID int     `json:"clientId"`
	Phone    string  `json:"phone"`
	Nombre   string  `json:"nombre"`
	Compro   bool    `json:"compro"`
	TDC      string  `json:"tdc"`
	Monto    float64 `json:"monto"`
	Date     string  `json:"date"`
}
