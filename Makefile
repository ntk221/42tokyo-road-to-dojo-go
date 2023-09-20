.PHONY: up down clean logs mysql seed deleteData

up:
	docker-compose up -d

down:
	docker-compose down

clean:
	docker-compose down -v

logs:
	docker-compose logs

mysql:
	docker exec -it mysql mysql -h localhost -u myuser -p game

migrate:
	mysqldef -u myuser -p mypassword -h 127.0.0.1 -P 3306 game < ./_tools/mysql/schema.sql

seed:
	mysql -h 127.0.0.1 -u myuser game -p < seed.sql

deleteData:
	mysql -h 127.0.0.1 -u myuser game -p < dropTables.sql