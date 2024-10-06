.PHONY: all build run fclean re start stop list generate_certs copy_certs

COMPOSE_FILE = docker-compose.yml
# DOCKER = docker

DOCKER = sudo docker

CERT_DIR = $(HOME)/.local/share/mkcert
BACKEND_CERT_DIR = ./back-end
FRONTEND_CERT_DIR = ./front-end
CERTS = localhost

all: run

generate_certs: 
	@echo "Generating certificates..."
	mkcert $(CERTS)

copy_certs: generate_certs
	@echo "Copying certificates..."
	@cp $(CERT_DIR)/* $(BACKEND_CERT_DIR)/
	@cp $(CERT_DIR)/* $(FRONTEND_CERT_DIR)/

build: copy_certs
	@echo "Building all images..."
	$(DOCKER)-compose build

run: build
	@echo "Running all containers..."
	$(DOCKER)-compose up -d
	@echo "Application is running in detached mode, use make stop to stop the running containers"

start:
	@echo "Starting containers..."
	$(DOCKER)-compose start

stop:
	@echo "Stopping containers..."
	$(DOCKER)-compose stop

list:
	@echo "Listing all containers..."
	$(DOCKER) ps -a

fclean: stop
	@echo "Removing all stopped containers..."
	$(DOCKER)-compose down --rmi all --volumes --remove-orphans

re: fclean run
