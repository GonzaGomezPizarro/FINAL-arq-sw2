# Use the official Node.js image as base
FROM node:latest

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json to the workspace
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application source code to the workspace
COPY . .

# Expose port if necessary
EXPOSE 8080

# Command to start the frontend server
CMD ["npm", "start"]
