You need an username and password on the database of the GraphL server, if you need one contact someone with an admin account.
Create an account on https://wandb.ai/site and get an API key (in settings), it will be used to get the results of your jobs.
You need a docker image for your projects with the setps you want to execute on the workstation as entrypoint or cmd in the dockerfile.
In order to get back the results of your tasks, you need to add code to save the results on your wandb account.
A minimal python script that save the content of a directory to your wandb account:
`import wandb                                                                    
wandb.init(project="sbb2")                                                                                
wandb.save("models/*", policy="now")`.
You can also use more advanced saving options from the SDK of wandb : https://docs.wandb.ai/.
Create your docker image and upload it to https://hub.docker.com/, it must not contain any secret.
Go to http://54.77.14.151:8080/playground.
Follow the steps in readme : https://github.com/42-AI/ws-backend#readme for login.
Use the login token as env variable as well as your wandb API key, and create a task (readme).
You should get the results of your job on your wandb accont.
