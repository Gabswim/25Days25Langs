FROM node:22-slim
WORKDIR /app
RUN npm install -g tsx
COPY solution.ts .
CMD ["tsx", "solution.ts"]