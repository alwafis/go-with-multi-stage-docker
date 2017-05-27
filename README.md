# Prerequisite
---------------
## Docker
### Windows
https://docs.docker.com/docker-for-windows/install/#start-docker-for-windows
### Ubuntu
https://docs.docker.com/engine/installation/linux/ubuntu/
### Mac
https://docs.docker.com/docker-for-mac/install/


## Docker-Compose
https://docs.docker.com/compose/install/

For docker compose alternatively you can use ```pip install docker-compose``` to install it.


# Usage
You can simply just run following command from your terminal from this directory:
- Build your docker-compose file with ```docker-compose build```
- Up your docker-compose image ```docker-compose up```

# Simple Go Lang Function With Multi Stage Docker
--------------
This project created to show how efficient is docker multi-staging

## Go File Processing
golang file for this project is placed inside `{project_dir}/files/code`.
Inside the go file there is a simple function to create and read a file.

## Testing
run `docker-compose build` then `docker-compose up` from your terminal inside `{project_dir}`

## Size Comparison
run `docker images` and data bellow will appear

| REPOSITORY        | TAG           | IMAGE  | ID           | CREATED  | SIZE |
| ------------- |:-------------:|:-------------:| -----:|:-------------:| -----:|-----:|
| deployment_image      | latest | 1e065ab4f082 | - | 6 minutes ago | 5.83MB |
| builder_image      | <none>      |  b199d5e3483e | centered      |   6 minutes ago |   261MB |
