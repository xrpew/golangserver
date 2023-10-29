**Para inciar el proyecto**

Mandatorio tener correctamente: docker

<code> $ sh init.sh </code>

corre en localhost:80 por defecto

get /
curl localhost:80

crear usuario
curl -X POST -H "Content-Type: application/json" -d '{"Email": "valor@mail.io", "Password": "valorsecreto"}' localhost/signup

post /login
curl -X POST -H "Content-Type: application/json" -d '{"Email": "valor@mail.io", "Password": "valorsecreto"}' localhost/login

get /me
curl -X GET -H "Content-Type: application/json" -H "Authorization: TOKEN-LOGIN" -d '{"Email": "valor@mail.io", "Password": "valorsecreto"}' localhost/me