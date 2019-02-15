# practice-01
practice-01

## Add a command
`http POST localhost:3030/API/create command=ls flag=notrequired path=.`



## basic

Create a git repo, and commit your code into it - just like you would do on a normal work day `check`

The server can be started and listens on localhost:3030 `check`

###Create the following endpoints:

GET /fizzbuzz/{number} - Depending on the value of number, the server should return “fizz” it’s divisible by 3, “buzz” if divisible by 5 and “fizzbuzz” if divisible by both. `check`

POST /count - The request body should contain a JSON object with one field called “data”. The value of data should be a string. The server takes the string from the “data” field and returns an associative array (hash, map, dictionary, however the implementation is called in the language) where the keys are the letters of the input string and the values are the number of occurrences of that letter. (e. g. if the input string is “aaa” then the response is ‘a: 3’) `check`

POST /run - The request body should contain a JSON object with one field called “command”. The value of command should be a string. The server takes the command and runs it using a system call, waits for it to finish and returns the output. (e. g. if the value of command is “pwd” then the return value will be the current working directory) `check`

Create a Dockerfile and run the server in a container `check`

Create a shell script that builds the docker image and runs the container `check`

## advanced 

Create a git repo, and commit your code into it - just like you would do on a normal work day `check`

The server can be started and listens on localhost:3030 `check`
 
Endpoint: new Task (a simple shell command) can be added through a POST request `check`

when a new Task is added execute the specified command (example command: ls -alh) and store the log/output of the command for the Task `check`

Endpoint: list the Tasks `check`

Endpoint: get the log/output of a specified Task `check`

Store the tasks and their outputs in the memory `nope`

//Write a reasonable amount of unit tests

###Extras (not in order, feel free to choose any of these)

//Simple CLI tool to manage the tasks & query logs - commands for the endpoints of the server (create, list and get logs)

Use a database for storing the task data `check`

//Separate Worker, which fetches the task/command from the Server, executes it and sends the logs etc. back to the server. In short: remove the command execution from the Server.

Run the Task/command in a docker container `check`

Run the server in a docker container `check`

docker-compose config so that the server can be started with a simple docker-compose up `check`

//Instead of a REST-JSON API, use gRPC