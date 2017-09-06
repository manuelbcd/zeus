# NODE DOCKERFILE - LATEST VERSION BY DEFAULT

ARG CODE_VERSION=latest
FROM node:${CODE_VERSION}

ARG PROJECT_NAME=project
RUN mkdir /opt/${PROJECT_NAME}

EXPOSE 80

WORKDIR /opt/${PROJECT_NAME}

RUN npm -g install nodemon

CMD npm start