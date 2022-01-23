
- ### Url Shortner basic data flow

    - **Short url method:**
       - url: localhost:3000/shortUrl/
       - body:
            {"originalUrl" : "url to be shortened"}
            - example:
                {"originalUrl" : "https://www.youtube.com/"}
       - method: POST
         - sample output 1: 
            Successfully Shortened  the Url: {
                "originalUrl": "https://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go",
                "shortenedUrl": "http://localhost:3000/3joVq4r",
                "id": "3joVq4r"
            }
         - sample output 2:
           Giving same url multiple times after one successfull shortning, gives following message
           "Url https://www.youtube.com/watch?v=CBVJTplw4cE is already shortened
                , and Shortened url is http://localhost:3000/qVPLq6G"
    - **Get all short url's:**
       - url: localhost:3000/getUrl/
       - method: GET
    - **Build and Run docker image:**
        - Run following commands:
            - for building docker image 
                - sudo docker build --tag  urlshortner  .
            - for running of docker image
                - sudo docker run -p 3000:3000 urlshortner
            - Application has started running  with same url: localhost:3000/
