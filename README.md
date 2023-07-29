# Practice Backend

I create a simple server to create article and get articles.

Prerequisite:
- go 1.17
- docker
- postman

What you need to do:
1. clone this repo
2. open terminal/bash -> execute `docker compose up`
3. open terminal/bash -> execute `go run main.go`
4. the server will be running

## Database Structure
Table articles
- id (int, primary key)
- author (string)
- title (string)
- body (string)
- created (datetime)

## API Contract:
1. Create Article
   ```POST http://localhost:3000/article```
   ```
   Request Body
   {
    "author": "john doe",
    "title": "the best song in the world",
    "body": "lorem ipsum"
   }
   ```
2. Get Articles
   ```GET http://localhost:3000/articles```
   ```
   Query Param
   - query="best song" // filter from title or body
   - author="john doe" // filter from author
   ```

## Flow
1. Create Article
   ```receive request -> parsing -> validate -> save into database -> store into redis -> return response```
2. Get Articles
   ```
   receive request -> parsing -> get article from redis
     -> if article found -> return response
     -> if article not found -> get from database -> store into redis -> return response
   ```
