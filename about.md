# rest-api

# user-service

# REST API

GET /users -- list of users -- 200(все ок), 404(сущность не найдена), 500(не смогли обработать запрос)
GET /users/:id -- users by id -- 200, 404, 500
POST /users/:id -- create user -- 204(no content), 4xx, Header Location: url
PUT /users/:id -- fully update user -- 200/204, 404, 400, 500
PATCH /users/:id -- partially update user -- 200/204, 404, 400, 500
DELETE /users/:id -- delete user bi id -- 200/204, 404, 400
