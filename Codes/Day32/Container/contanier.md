
### Step 1: Create Dockerfile -- Layers as per project
```bash
cat > Dockerfile << 'EOF'
FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
EOF
```

### Step 2: Verify Dockerfile was created
```bash
ls -la | grep Dockerfile
cat Dockerfile
```

### Step 3: Create docker-compose.yml

## Why We Created `docker-compose.yml` in Our Project
`docker-compose.yml` is created to **manage multiple containers together** as a single application. In our project, we have **two containers** that need to work together:

1. **PostgreSQL container** 
2. **Go application container**

## 🎯 Without Docker Compose (Complex & Manual)

Without `docker-compose.yml`, we would need to run **multiple commands** every time:

```bash
# 1. Create network
docker network create app_network

# 2. Run PostgreSQL container
docker run -d \
  --name gorm_postgres \
  --network app_network \
  -e POSTGRES_USER=arvinder.pal \
. . .

# 3. Build Go image
docker build -t gorm-gin-postgres-app .

# 4. Run Go container
docker run -d \
  --name gorm_app \
  --network app_network \
  -e DB_HOST=postgres \
  -e DB_USER=arvinder.pal \
. . .

# 5. To stop: need to stop both separately
docker stop gorm_app gorm_postgres
docker rm gorm_app gorm_postgres
```

**Problems with manual approach:**
- Need to remember 10+ commands
- Easy to forget network setup
- Containers might start in wrong order
- No automatic restart if they crash
- Hard to share setup with team

## ✅ With Docker Compose

With `docker-compose.yml`, everything is defined in **one file**:

```yaml
version: '3.8'

services:
  postgres:
    image: postgres:15-alpine
    container_name: gorm_postgres
    environment:
      POSTGRES_USER: arvinder.pal
      POSTGRES_PASSWORD: dockerpass123
      POSTGRES_DB: gorm_gin_db
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
    networks:
      - app_network

  app:
    build: .
    container_name: gorm_app
    ports:
      - "8083:8080"
    depends_on:
      - postgres
    environment:
      DB_HOST: postgres
      DB_USER: arvinder.pal
      DB_PASSWORD: dockerpass123
      DB_NAME: gorm_gin_db
    networks:
      - app_network

volumes:
  postgres_data:

networks:
  app_network:
```

Then just **one command** does everything:

```bash
# Start everything
docker-compose up -d

# Stop everything
docker-compose down

# View logs from both containers
docker-compose logs -f

# Rebuild and restart
docker-compose up -d --build
```

## 📋 What Docker Compose Does For Us

| Feature | Without Compose | With Compose |
|---------|----------------|--------------|
| **Start containers** | Multiple `docker run` commands | One `docker-compose up` |
| **Stop containers** | Multiple `docker stop` commands | One `docker-compose down` |
| **Network setup** | Manual `docker network create` | Automatic network creation |
| **Container order** | Must start DB first manually | `depends_on` ensures DB starts first |
| **Environment variables** | Set in each `docker run` | Defined once in YAML |
| **Volume management** | Manual volume creation | Defined in YAML |
| **Restart policy** | Set per container | Set once for all services |
| **Configuration sharing** | Email commands to team | Share one YAML file |


### 1. **Simplified Development**
```bash
# New developer joining the team? Just run:
git clone <repo>
cd project
docker-compose up --build
# Everything runs automatically!
```

### 2. **Proper Container Ordering**
```yaml
depends_on:
  - postgres  # Go app waits for database to start
```
Without this, Go app might start before PostgreSQL and crash.

### 3. **Automatic Network**
```yaml
networks:
  - app_network
```
Both containers can talk to each other using service names (`postgres` instead of IP address).

### 4. **Data Persistence**
```yaml
volumes:
  - postgres_data:/var/lib/postgresql/data
```
Database data survives container restarts and crashes.

### 5. **Environment Variables**
```yaml
environment:
  DB_HOST: postgres
  DB_USER: arvinder.pal
  DB_PASSWORD: dockerpass123
```
---
### Step 4: Verify docker-compose.yml was created
```bash
cat docker-compose.yml
```

### Step 5: Create .env file for Docker
```bash
cat > .env << 'EOF'
DB_HOST=postgres
DB_PORT=5432
DB_USER=arvinder.pal
DB_PASSWORD=dockerpass123
DB_NAME=gorm_gin_db
DB_SSLMODE=disable
SERVER_PORT=8080
EOF
```

### Step 6: Ensure go.mod and go.sum exist
```bash
go mod tidy
ls -la go.mod go.sum
```

## Part 2: Build and Start Containers

### Step 2.1: Build and start containers
```bash
docker-compose up --build
```
(Press Ctrl+C to stop, or run in background next time)

### Step 2.2: Start containers in background (detached mode)
```bash
docker-compose up -d --build
```

### Step 2.3: List running containers
```bash
docker ps
```

### Step 2.4: List all containers
```bash
docker ps -a
```

### Step 2.5: List Docker images
```bash
docker images
```

### Step 2.6: List volumes
```bash
docker volume ls
```

### Step 2.7: List networks
```bash
docker network ls
```

## 📂 Part 3: Inside Container Commands

### Step 3.1: Enter Go app container
```bash
docker exec -it gorm_app sh
```

**Once inside the app container, run:**
```bash
# List files in current directory
ls -la

# Print working directory
pwd

# Change directory
cd /app
pwd

# List files in /app
ls -la

# View source code
cat main.go

# View environment variables
env

# Check Go version
go version

# Check if app is running
ps aux

# Exit container
exit
```

### Step 3.2: Enter PostgreSQL container
```bash
docker exec -it gorm_postgres sh
```

**Once inside PostgreSQL container, run:**
```bash
# List files
ls -la

# Print working directory
pwd

# Change to data directory
cd /var/lib/postgresql/data
ls -la

# Exit container
exit
```

### Step 3.3: Access PostgreSQL directly (without entering shell)
```bash
docker exec -it gorm_postgres psql -U arvinder.pal -d gorm_gin_db
```

**Once inside psql, run:**
```sql
-- List all tables
\dt

-- View items
SELECT * FROM items;

-- Insert a test item
INSERT INTO items (name, price) VALUES ('Test Item', 99.99);

-- View inserted item
SELECT * FROM items;

-- Exit psql
\q
```

## Part 4: Copy Files Between Host and Container

### Step 4.1: Copy file from host to container
```bash
# Create a test file on host
echo "Hello from host" > test.txt

# Copy to container
docker cp test.txt gorm_app:/app/test.txt

# Verify it's there
docker exec gorm_app cat /app/test.txt
```

### Step 4.2: Copy file from container to host
```bash
# Copy from container to host
docker cp gorm_app:/app/main.go ./main.go.backup

# Verify on host
ls -la main.go.backup
cat main.go.backup
```

### Step 4.3: Copy directory from container to host
```bash
docker cp gorm_app:/app ./app-backup
ls -la app-backup/
```

## Part 5: Container Lifecycle Management

### Step 5.1: Stop containers
```bash
docker-compose stop
```

### Step 5.2: Start stopped containers
```bash
docker-compose start
```

### Step 5.3: Restart containers
```bash
docker-compose restart
```

### Step 5.4: Stop and remove containers (keep volumes)
```bash
docker-compose down
```

### Step 5.5: Stop and remove everything (including volumes)
```bash
docker-compose down -v
```

### Step 5.6: Rebuild and start fresh
```bash
docker-compose up -d --build
```

## Part 6: Monitoring and Debugging

### Step 6.1: View logs
```bash
# View all logs
docker-compose logs

# View specific service logs
docker-compose logs app
docker-compose logs postgres

# Follow logs in real-time
docker-compose logs -f

# View last 50 lines
docker-compose logs --tail=50 app
```

### Step 6.2: View container resource usage
```bash
docker stats
docker stats gorm_app gorm_postgres
```

### Step 6.3: Inspect container details
```bash
docker inspect gorm_app
docker inspect gorm_postgres | grep IPAddress
```

### Step 6.4: Check container health
```bash
docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"
```

## 🌐 Part 7: Testing API (Outside Container)

### Step 7.1: Test API endpoints
```bash
# GET all items
curl http://localhost:8080/api/items

# POST create item
curl -X POST http://localhost:8080/api/items \
  -H "Content-Type: application/json" \
  -d '{"name":"Laptop","price":999.99}'

# GET single item
curl http://localhost:8080/api/items/1

# UPDATE item
curl -X PUT http://localhost:8080/api/items/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Gaming Laptop","price":1299.99}'

# DELETE item
curl -X DELETE http://localhost:8080/api/items/1
```

## Part 8: Database Operations

### Step 8.1: Run SQL queries directly
```bash
# View all items
docker exec gorm_postgres psql -U arvinder.pal -d gorm_gin_db -c "SELECT * FROM items;"

# Count items
docker exec gorm_postgres psql -U arvinder.pal -d gorm_gin_db -c "SELECT COUNT(*) FROM items;"

# Insert via command line
docker exec gorm_postgres psql -U arvinder.pal -d gorm_gin_db -c "INSERT INTO items (name, price) VALUES ('Direct Insert', 49.99);"

# Delete all items
docker exec gorm_postgres psql -U arvinder.pal -d gorm_gin_db -c "DELETE FROM items;"
```

## Part 9: Cleanup Commands

### Step 9.1: Remove specific container
```bash
docker stop gorm_app
docker rm gorm_app
```

### Step 9.2: Remove specific image
```bash
docker rmi gorm-gin-postgres-app
```

### Step 9.3: Remove unused volumes
```bash
docker volume prune
```

### Step 9.4: Remove everything unused
```bash
docker system prune -a
```

### Step 9.5: Remove specific volume
```bash
docker volume rm gorm-gin-postgres_postgres_data
```