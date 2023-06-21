# Golang Backend Example

This Go [(Golang)](https://go.dev/) backend example project implements a [RESTful API](https://aws.amazon.com/what-is/restful-api/) for managing posts. Users can sign up, log in, create, retrieve, update, and delete posts using the provided endpoints. The application utilizes JSON payloads for data exchange and employs JWT tokens for secure authentication and authorization. With basic CRUD functionality and user authentication, it offers a flexible and secure solution for managing posts in a web environment.
## URL: https://old-sun-5714.fly.dev/

## Usage
I used [Fly.io](https://fly.io) to deploy and host this project. You can do the same if you want to host it on your own system.
First, clone this Git repository:
```bash
git clone https://github.com/safzanpirani/go_exampleBE
```
Move into the newly created directory:
```bash
cd go_exampleBE
```
Install flyctl in order to host this repository on Fly.io by running:
```bash
powershell -Command "iwr https://fly.io/install.ps1 -useb | iex"
```
After installing flyctl, run the following and authorize your Fly.io account:
```bash
flyctl auth login
```
Now, type the following to begin deployment:
```bash
flyctl launch
```
Follow the instructions on-screen to automatically be assigned a free deployment cluster to be able to host the project. This will take awhile.

The deployment will eventually fail. Follow the following commands and fill the the given data for the cluster you were assigned. We will use this data now in order to correctly configure our project to deploy properly.
```bash
flyctl secrets set DBNAME=~
flyctl secrets set HOST=HOSTNAME_URL~
flyctl secrets set PASSWORD=~
flyctl secrets set USER=~
```
Note that you must enter each command one by one and let the deployment start each time. As you enter the final command your deployment should be successful.

After deployment, go to your Fly.io dashboard and find your given hosted link where your deployed project can be accessed.
We will use Postman to interact with this URL.