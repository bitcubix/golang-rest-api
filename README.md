# RESTful API written in GO
[![PkgGoDev](https://pkg.go.dev/badge/github.com/gabrielix29/go-rest-api)](https://pkg.go.dev/github.com/gabrielix29/go-rest-api)

## Installation
We provide you with two ways to install the api.

### Installation with docker
docker-compose up

### Installation with bash script
- create a debian10 system with ssh and standard systemtools
- get a domain name and create dns A record to your WAN IP
- on your router: create port forwarding for http and https to your debian10 system 
- copy the bash script to your debian10 system and execute it as normal user:
[bash script](https://github.com/MystixCode/install_go_api)

## API endpoints
- `GET /books/` list of all books
- `POST /books/` create new book
- `GET /books/:id` get single book
- `PUT /books/:id` update book
- `DELETE /books/:id` delete single book


- `GET /author/` list of all author
- `POST /author/` create new author
- `GET /author/:id` get single author
- `PUT /author/:id` update author
- `DELETE /author/:id` delete single author

