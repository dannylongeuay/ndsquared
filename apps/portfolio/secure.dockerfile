FROM node:16-alpine

USER node

WORKDIR /app

RUN chown -R node:node /app

COPY --chown=node:node package.json package-lock.json ./

RUN npm ci

COPY --chown=node:node . .

RUN npm run build
