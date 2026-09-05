package model

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
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
