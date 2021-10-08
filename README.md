# Tiny URL

Tiny Url is robust highly available, distributed system to generate tiny urls from long urls.
It uses zookeeper to co-ordinate between multiple instances of the service.

## Architecture


System have following components.
1. Tiny Service:
 
    This is the main application service. It exposes following endpoints
        
        1. POST TinyUrl(longUrl, slug) tinyUrl
        2. GET LongUrl(tinyUrl) longUrl
2. Zookeeper Cluster:
    
    This is used to co-ordinate between multiple instances of Tiny Service. It Keeps the track of keys used so far by all tiny services.
   
3. Redis:

    The Redis cache have been used to improve the read performance. It uses `Read Through` strategy.
   
4. Mongo:
    
    Mongo DB have been used for persistence storage. 

## Setup and Run

### Run the project
    1. run docker-compose -f docker-compose.yml up -d
    2. cd cmd/server && go run main.go 
    or
    2.1 run in docker container
        1. make docker
        2. docker run -it -p 8081:8081 tiny-url 
    
## Todo
1. setup simple project -> done
 -> 2 handler, serivce, dao -> done
 -> basic config -> done
-> docker, make file and zookeeper setup
   -> connect all this with dummy
   -> write the actual logic
