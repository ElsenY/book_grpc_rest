# book_grpc_rest
A library like app which use REST API as API Gateway, which use GRPC to calls microservices and communication between microservices

# How to run

1. clone this repo to your local machine
2. run `docker compose up --build -d`
3. API Gateway running at localhost:8080 and ready to serve
4. run `docker compose down --volumes` to stop the app from running and delete database volumes too

# ERD 
<img width="746" alt="image" src="https://github.com/user-attachments/assets/f517f9a2-9eef-49c1-a3f6-2b07b3af7260" />

# API Route
for the sake of simplicity, I have made the API route in postman [here](https://www.postman.com/maintenance-architect-99534403/elsen-public/collection/f5sz855/synapsis-book) and will explain the walkthrough based on this postman API collection

# API and APP Walkthrough
most of the routes are protected routes, so first thing we need to create user, and get the token by using the 'login' API in the user route

1. Use 'Register User' API at User Route to register user
2. Use 'Login' API at User Route so that you get the token
   
### Protected Routes : 
#### Author Route:
1. Use the token, go to 'register author' API at Author Route to register the current user as an author, so that you can insert a book

#### Book Route:
1. Go to 'Insert Book' API at Book route to insert book with the current user as the author
2. 'Borrow Book' API is used to borrow book and insert the record into 'borrow_records' table (will decrease book stock by 1 too)
3. 'Return Book' API is used to return borrowed book, and update the 'borrow_records' table (will increate book stock by 1 too)
4. 'Edit Stock' API is used to edit a book's stock

#### Category Route:
1. 'Insert Category' API is to insert a new category
2. 'link book with category' is to associate a book with a category, (will insert the assocation in the book_category table)

### Unprotected Routes:
#### User Route:
1. 'Register User' to register user
2. 'Login' to login and get token based on the user email and password (token containing user email data)

#### Book Route:
1. 'Book Recommendation' route will return the most borrowed book
2. 'Search Book' route will return book data based on the request's book title
   
# DockerHub Images
https://hub.docker.com/repository/docker/elseny/book/general



