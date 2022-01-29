module myapp

go 1.18

replace github.com/zLeki/Celeritas => ../celeritas

require github.com/zLeki/Celeritas v0.0.0-20220129141345-424bae94c0a3

require (
	github.com/go-chi/chi/v5 v5.0.7 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
)
