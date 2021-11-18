# To start the postgresql database using docker run the following command

`sudo docker run --name comments-api-database -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres`