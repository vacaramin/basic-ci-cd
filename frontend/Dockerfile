
# Using lighter image for running the application build
FROM nginx:alpine

# Copying the build files to the nginx server
COPY . /usr/share/nginx/html

# Exposing the port
EXPOSE 80



# Running the nginx server
CMD ["nginx", "-g", "daemon off;"]