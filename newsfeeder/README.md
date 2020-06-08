Small example of how to put together a REST API in Go using the Gin web framework.  

1. Run the application (open the terminal inside of the folder "newsfeeder" and run the command "make")
2. Open a browser and call the following URL http://localhost:8080/ping
3. Test it with the browser console,
    e.g. await fetch ("/newsfeed", {
            method: "POST",
            headers: {'content-type' : 'application/json'},
            body: JSON.stringify({
                title: 'Here comes the sun',
                post: 'Sun, sun, sun, here it comes'
            })
         })
4. Have a look at the logs

Based on a [YouTube Tutorial](https://www.youtube.com/watch?v=LOn1GUsjOF4&list=WL&index=6)