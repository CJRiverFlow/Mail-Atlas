FROM node:16-alpine

WORKDIR /app

COPY package.json package-lock.json ./

RUN npm install

COPY . /app

RUN npm run build

EXPOSE 8080 5173

# CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0", "--port", "5173"]
CMD ["npm", "run", "server-docker"]