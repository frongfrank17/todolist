# TODO LIST SERVICE
Folder Architecture 
```base
Hugeman
---configs
   |--config.go
---database
   |--database.go
---handler
   |--handler.go
---repository
   |--interface.go
   |--todo_repository.go
---service
   |--interface.go
   |--todo_service.go
---route
   |--route.go
---test
---main.go
```
Run App
```base
docker-compose up
```
**Created**

POST | http://localhost:3000/api/todo/created

```json
{
   title="Golang"
   description="GoLang Network"
   image="" //type file
}
```

**Get-List**

GET | http://localhost:3000/api/todo/lists?f=title&sorted=1

sorted
- 0="asc"
- 1="desc"
f= field ที่จะ sort เช่น title , description , status , createAt

**Get By type**

GET | localhost:3000/api/todo/list/:type/Golang
type = "title" || "description"

**Update By Field**

PUT | localhost:3000/api/todo/list/update/:field/:set/:id
field= "title" || "description"

**Update Status**

PUT | localhost:3000/api/todo/list/update_status/completed/:id

**Update Image**

PATCH | localhost:3000/api/todo/list/update/image

```json
{
   id="xxxx" 
   image="" //type file
}
```