# Build stage for GUI application
FROM golang:1.23-bullseye AS builder

# Set working directory
WORKDIR /app

# Install build dependencies for Fyne GUI (requires CGO)
RUN apt-get update && apt-get install -y \
    gcc \
    pkg-config \
    libgl1-mesa-dev \
    libxrandr-dev \
    libxinerama-dev \
    libxcursor-dev \
    libxi-dev \
    libxext-dev \
    libxfixes-dev \
    libx11-dev \
    libxxf86vm-dev \
    && rm -rf /var/lib/apt/lists/*

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the GUI application with CGO enabled
RUN CGO_ENABLED=1 GOOS=linux go build -v -o ai-launcher ./cmd/gui

# Final stage
FROM alpine:latest

# Install runtime dependencies for GUI application
RUN apk --no-cache add \
    ca-certificates \
    tzdata \
    mesa-gl \
    libx11 \
    libxrandr \
    libxinerama \
    libxcursor \
    libxi \
    dbus \
    fontconfig

# Create non-root user
RUN addgroup -g 1001 appuser && \
    adduser -D -s /bin/sh -u 1001 -G appuser appuser

# Set working directory
WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/ai-launcher .

# Change ownership to appuser
RUN chown appuser:appuser /app/ai-launcher

# Switch to non-root user
USER appuser

# Expose port (if needed)
EXPOSE 8080

# Environment variables for GUI
ENV DISPLAY=:0
ENV QT_QPA_PLATFORM=minimal

# Health check (simplified for GUI app)
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD pgrep ai-launcher || exit 1

# Run the application
ENTRYPOINT ["./ai-launcher"]