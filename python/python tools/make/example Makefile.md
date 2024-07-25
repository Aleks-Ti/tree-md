# Example Makefile

```makefile
ifeq (revision,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif

.PHONY:

start:
	poetry run uvicorn \
		--reload \
		--host $$HOST \
		--port $$PORT \
		"$$APP_MODULE"

poetry_install_main:
	poetry env use python3.11 && poetry install --only main --no-root

poetry_install_dev:
	poetry env use python3.11 && poetry install --only dev

pdb:
	docker run --name=pdb \
	 			-e SSL_MODE='disable'\
				-e POSTGRES_USER=$$PG_USER\
				-e POSTGRES_PASSWORD=$$PG_PASSWORD\
				-e POSTGRES_DB=$$PG_DB\
				-e TZ=GMT-3\
				-p $$PG_PORT:5432 -d --rm postgres:alpine

redis_db:
	docker run --name=redis_db \
				-p 6379:6379 \
				-d --rm redis

revision:
	poetry run alembic revision --autogenerate -m $(RUN_ARGS)

migrate:
	poetry run alembic upgrade head

stop_pdb:
	docker stop project_db

load_dump:
	cat dumps/dump.sql | docker exec -i project_db psql -U podvig -d podvig
```
