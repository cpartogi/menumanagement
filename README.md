## How To Setup Boilerplate In Local development

### Prerequisites : 
1. Golang version 1.14 or higher
2. Docker/PostgreSQL for database

### Steps : 
* Open file .env.example
* Save as .env.example to .env
* Open file docker-compose.yml (if using docker)
* Change db user and password if you need it
* Open file main.example.yml
* Modify database username and password, use the same value from docker-compose.yml
* Save as main.example.yml to main.yml
* Start Docker
* Run this command : docker-compose up
* After docker with postgres already running, open new terminal
* Run this command to generate swagger api documentation : swag init -o api/docs 
* Run this command : make run
* From browser, open this address : http://localhost:8000/swagger/
* To create tables, you need to open sql client (for example DBeaver) and run sql scripts in folder /scripts/migrations/localdev from boilerplate repository 