FROM node:16.14-alpine as builder

WORKDIR /app

COPY package.json ./

RUN npm install --production --silent

COPY . .

RUN npm run build

FROM nginx:1.20

COPY --from=builder /app/build /usr/share/nginx/html
COPY frontend-nginx.conf /etc/nginx/conf.d/default.conf