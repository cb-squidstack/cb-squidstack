
PROJECT_NAME = cb-squidstack

up:
	docker-compose up --build -d
	docker-compose ps

down:
	docker-compose down

logs:
	docker-compose logs -f

healthcheck:
	@echo "🔍 Checking squid-ui on localhost..."
	@curl -s -o /dev/null -w "%{http_code} http://localhost:3000/ -> / \n" http://localhost:3000/
	@curl -s -o /dev/null -w "%{http_code} http://localhost:3000/squidui -> /squidui \n" http://localhost:3000/squidui
	@echo "🔍 Checking nautilus-inventory on localhost..."
	@curl -s -o /dev/null -w "%{http_code} http://localhost:8081/health -> /health \n" http://localhost:8081/health
	@curl -s -o /dev/null -w "%{http_code} http://localhost:8081/inventory -> /inventory \n" http://localhost:8081/inventory
	@echo "✅ Health check (external endpoints) complete."

status:
	docker-compose ps
