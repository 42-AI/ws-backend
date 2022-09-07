# User guide
- You need a username and password on the database of the GraphL server, if you want one contact someone with an admin account.  
- Create an account on https://wandb.ai/site and get an API key (in settings), it will be used to get the results of your jobs.
- You need a docker image for your projects with the setps you want to execute on the workstation as entrypoint or cmd in the dockerfile.  
Create your docker image and upload it to https://hub.docker.com/, it must NOT contain any secret.
- In order to get back the results of your tasks, you need to add code executed by your container to save the results on your wandb account.  
A minimal python script that saves the content of a directory called `models` to your wandb account:  
```
import wandb  
wandb.init(project="project_name")  
wandb.save("models/*", policy="now")
```  
You can also use more advanced saving options from the SDK of wandb : https://docs.wandb.ai/.  
- Go to http://54.77.14.151:8080/playground.  
- Follow the steps in [readme](https://github.com/42-AI/ws-backend#login) for login :  
copy the following request to log in your user :
```
query login {
  login (id: "your_user_id", pwd: "your-password") {
    ... on Token {
      token
      userId
      username
      isAdmin
    }
    ... on Error {
      code
      message
    }
  }
}
```
You should get a response similar to this :
```
{
  "data": {
    "login": {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MTczODMzNDksInVzZXJfaWQiOiJkZjljNDYzZC00ZmIwLTRmYzAtYTU5OC00YmQ3NzEzMzg2ZDAifQ.Xj_rUGIB7l90kiXD_U12ni2kf9U-afARaCZKbEao-oU",
      "userId": "df9c463d-4fb0-4fc0-a598-4bd7713386d0",
      "username": "your_user_id"
    }
  }
}
```
Copy the `token` value.
- Use the login token as header variable (cf readme [create a user](https://github.com/42-AI/ws-backend/tree/gs/doc/user-tutorial#create-a-user)) in `HTTP HEADERS` :
```
{
  "auth": "your_token"
}
```
and [create a task](https://github.com/42-AI/ws-backend/tree/gs/doc/user-tutorial#create-a-machine-learning-task) with your wandb API key as env variable :
```
mutation createTask {
  create_task(input:{env:"WANDB_API_KEY=putYourKeyHere", docker_image:"jjauzion/wandb-test", dataset:"s3//"}) {
    id
    user_id
        created_at
        started_at
        ended_at
        status
    job { dataset, docker_image, env }
  }
}
```
- You should get the results of your job on your wandb account.  
