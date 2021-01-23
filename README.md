## How To Setup In Local development

### Prerequisites : 
1. Golang version 1.14 or higher
2. PostgreSQL for database

### Steps : 
* Open file .env.example
* Save as .env.example to .envmo
* Open file main.example.yml
* Save as main.example.yml to main.yml
* Run this command to generate swagger api documentation : swag init -o api/docs 
* Run this command : make local
* From browser, open this address : http://localhost:8000/swagger/
