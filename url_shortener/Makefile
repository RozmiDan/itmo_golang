.PHONY: build run

build:
	docker build -t myapp .

run: build
	@echo "Запуск контейнера..."
	# Если контейнер с таким именем существует, удаляем его:
	-docker rm -f myapp_container 2>/dev/null || true
	# Запускаем контейнер в фореграунд режиме (вывод в терминал):
	docker run --rm --name myapp_container -p 8080:8080 myapp
