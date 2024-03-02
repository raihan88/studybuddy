# Description: 
This is a very simple note taking CLI app. You can add a new word and it's corresponding note or definition. And you can see the list also. 

# How to Run:
Download the repository on your local machine. Make sure you have go and cobra-cli installed in your machine. Then run the following commands: 

# Commands: 
```
go build
go install
```
These two commands will build the project and install it as studybuddy CLI app in your local machine. 

To use the app, you will need these commands:
```
studybuddy init
```
This will initialize a database to store your notes. 
```
studybuddy note new
```
will help you to add a new note. And, 
```
studybuddy note list
```
will help you to see all the notes you have taken so far. 
