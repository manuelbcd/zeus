# NODE DOCKERFILE - LATEST VERSION BY DEFAULT

ARG CODE_VERSION=latest
FROM node:${CODE_VERSION}

ARG PROJECT_NAME=project
RUN mkdir /opt/${PROJECT_NAME}

EXPOSE 80

WORKDIR /opt/${PROJECT_NAME}

RUN npm -g install nodemon

CMD npm start

## CREATE COMMAND
##
## mkdir ~/projects/zeus
##
## docker build --build-arg CODE_VERSION=8.4 --build-arg PROJECT_NAME=zeus --tag node-zeus:8.4 -f dockerfile.node .
## 
## docker run -d -it -p 80:80 -v ~/projects/zeus:/opt/zeus --name node-zeus node-zeus:8.4