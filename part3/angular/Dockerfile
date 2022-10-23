FROM node:lts-alpine AS builder

WORKDIR /usr/src/app

COPY /new-app/package.json new-app/package-lock.json ./
RUN npm install

COPY new-app .

RUN npm run build


FROM nginx:alpine
COPY --from=builder /usr/src/app/dist/new-app/ /usr/share/nginx/html
COPY ./nginx-custom.conf /etc/nginx/conf.d/default.conf

EXPOSE 80