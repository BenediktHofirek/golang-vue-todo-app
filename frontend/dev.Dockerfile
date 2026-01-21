FROM node:24-slim AS base
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
ENV CI=true
RUN corepack enable

# For healthcheck
RUN apt-get update && \
    apt-get install -y --no-install-recommends curl && \
    rm -rf /var/lib/apt/lists/*

COPY pnpm-lock.yaml pnpm-workspace.yaml /app/

WORKDIR /app

RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm fetch

RUN pnpm add -g nodemon

CMD ["/bin/sh", "-c", "nodemon --watch pnpm-lock.yaml --watch vite.config.ts --exec \"pnpm install --frozen-lockfile && pnpm run dev\""]
