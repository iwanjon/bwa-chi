# bwa-chi

## Campaign Donation
_to run this project midtrans account is needded and save the apikey in .env locate next to main.go file_
_this project based on BWA Course (API ONLY), rewrite using golang chi-router_
_another version like gin, echo can be found in other repo with prefix name bwa (for this momemt)_

1. run go mod tidy
2. create postgres database
3. create .env file next to main.go --> google about godotenv
4. add ServerKey, host, port, userdb, password, dbname in .env-->(serverKey is midtrans key, google about how to use it, the other variable host, port etc is db postgres variable)
5. install golang-migrate 
6. run migration --> only run upgrade 
7. run main.go file
8. postman collection is attached next to main.go file

migration example:
- for installation 
go install -tags 'postgres,mysql,mongodb' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
- for create migration file
migrate  create -ext sql -dir db/migrations create_table_users
- for upgrade
migrate -database "postgresql://postgres:a1@localhost:5432/bwastartup" -dir db/migrations up
- for downgrade
migrate -database "postgresql://postgres:a1@localhost:5432/bwastartup" -dir db/migrations down
