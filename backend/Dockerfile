#! this dockerfile must be built from the project root

FROM node:22-slim AS frontend-build

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable
WORKDIR /app

# Build the frontend
COPY "package.json" "pnpm-lock.yaml" ./
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile
COPY . .
RUN pnpm run build


# build the backend
FROM golang:1.22-alpine AS server-build
WORKDIR /app
COPY "./backend/go.sum" .
COPY "./backend/go.mod" .
RUN go mod download
COPY backend/ .
RUN go build -o server


# Final stage: Copy the frontend and Go server into a new image
FROM alpine:3.20
WORKDIR /app
COPY --from=frontend-build /app/dist ./dist
COPY --from=server-build /app/server .
COPY --from=server-build /app/.env ./.env
ENV ADDR="0.0.0.0"
ENV PORT=8080
ENV GIN_MODE=release
ENV DIST_DIR="/app/dist"
# Expose the port your Go server listens on
EXPOSE 8080
USER root
CMD ["./server"]