
.DEFAULT_GOAL := run

.PHONY: setup
setup:
	mysql -uroot < mysql/database.sql && mysql -unoonde_w noonde_api -pnoonde_dev < mysql/tables.sql && zsh elastic/index.sh

.PHONY: database
database:
	mysql -uroot < mysql/database.sql

.PHONY: run
run:
	go run cmd/api/*.go

.PHONY: elastic
elastic:
	zsh elastic/index.sh

.PHONY: table
table:
	mysql -unoonde_w noonde_api -pnoonde_dev < mysql/tables.sql

.PHONY: spacemarket
spacemarket:
	go run cmd/spacemarket/*.go

.PHONY: instabase
instabase:
	go run cmd/instabase/*.go

.PHONY: ssh
ssh:
	ssh -i ~/.ssh/yunling-test.pem ec2-user@13.114.179.185
