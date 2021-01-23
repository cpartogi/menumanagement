## How To Setup Boilerplate In Local development

### Prerequisites : 
1. Golang version 1.14 or higher
2. Docker/PostgreSQL for database

### Steps : 
* Open file .env.example
* Save as .env.example to .env
* Change db user and password if you need it
* Open file main.example.yml
* Save as main.example.yml to main.yml
* Run this command to generate swagger api documentation : swag init -o api/docs 
* Run this command : make run
* From browser, open this address : http://localhost:8000/swagger/
