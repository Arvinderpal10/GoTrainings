# 📚 Dockerfile KEYWORDS - Complete Guide with Examples

## 🎯 What is a Dockerfile?
A Dockerfile is a recipe that tells Docker how to build your application image. Each keyword is an instruction.

---

## 📋 Complete List of Dockerfile Keywords

### 1️⃣ FROM - Sets the base image
Purpose: Specifies the starting image to build upon

```dockerfile
# Examples
FROM ubuntu:22.04
FROM alpine:latest
FROM node:18-alpine
FROM python:3.11-slim
FROM golang:1.21-alpine
FROM nginx:alpine
```

Why: Every container starts from an existing image

---

### 2️⃣ WORKDIR - Sets working directory
Purpose: Changes the current directory inside container (like `cd`)

```dockerfile
# Examples
WORKDIR /app
WORKDIR /usr/src/app
WORKDIR /home/node/app
WORKDIR /var/www/html
```

Why: All following commands will run from this directory

---

### 3️⃣ COPY - Copies files from host to container
Purpose: Moves files/folders from your computer into the container

```dockerfile
# Syntax: COPY <source> <destination>

# Examples
COPY . /app                    # Copy everything to /app
COPY package.json .            # Copy single file
COPY go.mod go.sum ./          # Copy multiple files
COPY --chown=node:node . /app  # Copy with ownership
COPY *.txt /app/files/         # Copy all txt files
```

Why: Your application code needs to be inside the container

---

### 4️⃣ ADD - Advanced version of COPY
Purpose: COPY + automatic archive extraction + remote URLs

```dockerfile
# Examples
ADD . /app                     # Same as COPY
ADD archive.tar.gz /app/       # Automatically extracts
ADD https://example.com/file /tmp/  # Downloads from URL
ADD --chown=node:node . /app   # With ownership
```

Difference from COPY:
- COPY = simple file copy
- ADD = can extract tar files and download URLs

Best practice: Use COPY unless you need ADD's extra features

---

### 5️⃣ RUN - Executes commands during build
Purpose: Runs commands to install packages, create directories, etc.

```dockerfile
# Examples
RUN apt-get update && apt-get install -y python3
RUN npm install
RUN go mod download
RUN pip install -r requirements.txt
RUN mkdir -p /app/logs
RUN chmod +x /app/script.sh
```

Why: Prepares your environment (install dependencies, create folders)

---

### 6️⃣ CMD - Default command when container starts
Purpose: Defines what runs when container starts (can be overridden)

```dockerfile
# Examples
CMD ["node", "app.js"]
CMD ["python", "app.py"]
CMD ["./main"]
CMD ["npm", "start"]
CMD ["java", "-jar", "app.jar"]

# Shell form (not recommended)
CMD node app.js
```

Important: Only one CMD per Dockerfile. Last one wins.

Override example: `docker run myimage python test.py`

---

### 7️⃣ ENTRYPOINT - Main command (cannot be overridden easily)
Purpose: Defines the executable that always runs

```dockerfile
# Examples
ENTRYPOINT ["node"]
CMD ["app.js"]              # CMD becomes default arguments

ENTRYPOINT ["python"]
CMD ["app.py"]

ENTRYPOINT ["./main"]

# With CMD as arguments
ENTRYPOINT ["docker-entrypoint.sh"]
CMD ["--help"]
```

Difference from CMD:
- CMD = can be overridden
- ENTRYPOINT = always runs

---

### 8️⃣ EXPOSE - Documents which port to publish
Purpose: Informs which port the container listens on (documentation only)

```dockerfile
# Examples
EXPOSE 8080
EXPOSE 3000
EXPOSE 80
EXPOSE 5432
EXPOSE 8080/tcp
EXPOSE 8080/udp
```

Note: This doesn't actually publish the port. Use `-p` flag with `docker run`.

---

### 9️⃣ ENV - Sets environment variables
Purpose: Defines environment variables inside container

```dockerfile
# Examples
ENV NODE_ENV=production
ENV PORT=8080
ENV DB_HOST=localhost
ENV DB_USER=admin DB_PASSWORD=secret
ENV PATH="/myapp/bin:${PATH}"
```

Why: Configure your application without changing code

---

### 🔟 ARG - Build-time variables
Purpose: Variables only available during image building

```dockerfile
# Examples
ARG VERSION=latest
ARG DEBIAN_FRONTEND=noninteractive
ARG NODE_VERSION=18

# Use with --build-arg
RUN echo "Building version ${VERSION}"

# Build command:
# docker build --build-arg VERSION=1.2.3 -t myimage .
```

Difference from ENV: ARG disappears after build, ENV stays in container

---

### 1️⃣1️⃣ VOLUME - Creates a mount point for persistent data
Purpose: Marks a directory for persistent storage

```dockerfile
# Examples
VOLUME /data
VOLUME /var/lib/postgresql/data
VOLUME ["/app/logs", "/app/uploads"]
VOLUME /app/node_modules
```

Why: Data in volumes survives container deletion

---

### 1️⃣2️⃣ USER - Changes user for subsequent commands
Purpose: Switches to a non-root user (security best practice)

```dockerfile
# Examples
USER node
USER 1000
USER appuser:appgroup

# Create user first
RUN useradd -m appuser
USER appuser
```

Why: Running as root is a security risk

---

### 1️⃣3️⃣ LABEL - Adds metadata to image
Purpose: Attaches key-value pairs for organization

```dockerfile
# Examples
LABEL version="1.0.0"
LABEL maintainer="john@example.com"
LABEL description="My awesome app"
LABEL com.example.version="1.0"
LABEL org.opencontainers.image.created="2024-01-01"
```

Why: Helps organize and filter images

---

### 1️⃣4️⃣ HEALTHCHECK - Tests if container is working
Purpose: Docker checks if application is healthy

```dockerfile
# Examples
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD curl -f http://localhost:8080/health || exit 1

HEALTHCHECK CMD node health.js

HEALTHCHECK NONE  # Disable healthcheck
```

Why: Docker can restart unhealthy containers

---

### 1️⃣5️⃣ SHELL - Changes default shell
Purpose: Changes the shell used for shell form commands

```dockerfile
# Examples
SHELL ["/bin/bash", "-c"]
SHELL ["/bin/sh", "-c"]
SHELL ["powershell", "-Command"]

RUN echo "Hello"  # Uses specified shell
```

---

### 1️⃣6️⃣ STOPSIGNAL - Sets system call to stop container
Purpose: Defines which signal to use for stopping

```dockerfile
# Examples
STOPSIGNAL SIGTERM
STOPSIGNAL SIGINT
STOPSIGNAL SIGKILL
```

---

### 17 ONBUILD - Triggers when image is used as base
Purpose: Delayed execution when image is used in another Dockerfile

```dockerfile
# Examples
ONBUILD COPY . /app
ONBUILD RUN npm install
ONBUILD ADD . /usr/src

# In child Dockerfile:
FROM parent-image  # ONBUILD commands run here
```

---