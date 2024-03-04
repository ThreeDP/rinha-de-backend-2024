
SOURCES_PATH		=		./requirements/
COMPOSE				=		$(addprefix $(SOURCES_PATH), docker-compose.yml)
# REDIS_PATH			=		$(addprefix $(SOURCES_PATH), redis/data)
# DB_PATH				=		$(addprefix $(SOURCES_PATH), db/data)
CONTAINER_IMAGES	=		db postgres:16.1-alpine node
CONTAINER_PROCESS	=		api-one api-two nginx db
CONTAINER_VOLUMES	=		nginx-volume db-volume

all: build up

build:
# mkdir -p $(DB_PATH) $(REDIS_PATH)
	docker-compose -f $(COMPOSE) build

up:
	docker-compose -f $(COMPOSE) up -d

down:
	docker-compose -f $(COMPOSE) down

clean: down
	docker volume rm $(CONTAINER_VOLUMES)
#	docker rm $(CONTAINER_PROCESS)

fclean: down clean
	docker rmi $(CONTAINER_IMAGES)

.PHONY: all build up down clean fclean