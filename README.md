# Application Development with Go  
Go is one of the most beautiful programming languages existing. It's clear, simple and fast.

Here you'll find some small programs I created or copied while learning Go.

## Newsfeeder  
Small example of how to put together a REST API in Go using the Gin web framework.
### ⚙️ Setting Up 
1. To run the application open the terminal inside of the folder `/newsfeeder` and run the command `make`
2. Open a browser and call the following URL http://localhost:8080/ping
3. Test it with the browser console or something like postman,  
    e.g. w/ a `POST` request
```
await fetch ("/newsfeed", {
            method: "POST",  
            headers: {'content-type' : 'application/json'},  
            body: JSON.stringify({  
                title: 'Here comes the sun',  
                post: 'Sun, sun, sun, here it comes'  
            })  
           })
``` 
4. Have a look at the logs!

Based on a [YouTube Tutorial](https://www.youtube.com/watch?v=LOn1GUsjOF4&list=WL&index=6)

