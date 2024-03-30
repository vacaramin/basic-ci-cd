
# Using lighter image for running the application build
FROM nginx:alpine

# Copying the build files to the nginx server
COPY . /usr/share/nginx/html

# Exposing the port
EXPOSE 3000

# Modify NGINX configuration to listen on port 3000
RUN sed -i -e 's/80/3000/g' /etc/nginx/conf.d/default.conf

# Running the nginx server
CMD ["nginx", "-g", "daemon off;"]