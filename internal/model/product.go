package model

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name" binding:"required, min=2"`
	Price float64 `json:"price" binding:"required, gt=0"`
}

/*

                 Client
                    │
                    ▼
             cmd/server/main.go
                    │
                    ▼
               /ping route

internal/model
       │
       └── Product

*/
