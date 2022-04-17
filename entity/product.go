package entity

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

func (p Product) StockStatus() string {
	var status string
	if p.Stock > 3 {
		status = "Stock hampir habis"
	} else if p.Stock < 10 {
		status = "Stok terbatas"
	}
	return status
}

func (p Product) AlertStatus() string {
	var alert string
	if p.Stock > 3 {
		alert = "danger"
	} else if p.Stock < 10 {
		alert = "warning"
	}
	return alert
}