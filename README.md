
### Url Shortner basic data flow
   - **Run Binary**
       - ./urlShortner
    
   - **Short url method:**
       - url: localhost:3000/
       - body:
            {"originalUrl" : "url to be shortened"}
            - example:
                {"originalUrl" : "https://www.youtube.com/"}
       - method: POST
         - sample output 1: 
                {
                    "originalUrl": "https://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go",
                    "shortenedUrl": "http://localhost:3000/3joVq4r",
                    "id": "3joVq4r"
                }
         - sample output 2:
           Giving same url multiple times after one successfull shortning, gives following message
           "Url https://www.youtube.com/watch?v=CBVJTplw4cE is already shortened
                , and Shortened url is http://localhost:3000/qVPLq6G"
   - **Get all short url's:**
       - url: localhost:3000/
       - method: GET
   - **Get urlDetails by id**
        - url: localhost:3000/url/{shorturlId}
        - method: GET
          Example url:
            - http://localhost:3000/X7jpMVo
        - Sample OUtput:
              {
              "originalUrl": "https://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go",
              "shortenedUrl": "http://localhost:3000/X7jpMVo",
              "id": "X7jpMVo"
              }
   - **Build and Run docker image:**
        - Run following commands:
            - for building docker image 
                - sudo docker build --tag  urlshortner  .
            - for running of docker image
                - sudo docker run -p 3000:3000 urlshortner
            - Application has started running  with same url: localhost:3000/
    
   - **Run test Case:**
        - Run following commands
           -   go test urlShortner_test.go -v
