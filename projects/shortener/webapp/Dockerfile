# Start from a base image with Node.js pre-installed, using Alpine Linux for minimal size
FROM node:21-alpine

# Set the working directory inside the container
WORKDIR /usr/src/app

# Copy package.json and package-lock.json files to the working directory
COPY package*.json ./

# Copy the rest of the application code
COPY . .

# Install npm dependencies
RUN npm install 

# Expose the port your app runs on
EXPOSE 3000

# Build the next
RUN npm run build

# Command to run the application
CMD npm start
