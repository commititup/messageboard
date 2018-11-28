FROM golang:1.8

#make destination dir for app config file
RUN mkdir -p /etc/app

#make destination dir for backend binary to store
RUN mkdir -p /usr/bin

#make destination dir for frontend content to store
RUN mkdir -p /var/www/html

#copy the app config file to destination
COPY ./conf/app/config.json /etc/app/

#copy the backend binary to destination
COPY ./backend/bin/server /usr/bin

#copy the frontend to the destination
COPY ./frontend/dist /var/www/html
COPY ./frontend/index.html  /var/www/html
COPY ./frontend/src/assets /var/www/html/

#open port 
EXPOSE 4443
RUN ["chmod","+x","/usr/bin/server"]

#Run command
CMD ["server"]
