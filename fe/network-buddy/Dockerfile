#FROM node:10 as builder
#
## Copying angular folder from local directory to Educative directory
#COPY . /usr/local/app
#
## Installing Angular cli and node modules in angular directory
#RUN     npm install -g @angular/cli &&\
#        cd /usr/local/app/ &&\
#        ng build --prod
#
#WORKDIR cd /usr/local/app/dist
#RUN ls -la
#WORKDIR cd /usr/local/app/
#RUN ls -la
##RUN ng build --prod


FROM nginxinc/nginx-unprivileged

COPY dist/ /usr/share/nginx/html

RUN ls -la /usr/share/nginx/html
