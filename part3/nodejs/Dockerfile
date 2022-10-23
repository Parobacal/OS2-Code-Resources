FROM node:lts-alpine

WORKDIR /usr/src/app

COPY package.json package-lock.json ./
RUN npm install

COPY . .
RUN npm install

EXPOSE 4000

CMD ["node", "index.js"]