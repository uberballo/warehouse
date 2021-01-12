## Warehouse  
## 1. Project purpose
Fetches product data from a bad api and displays it on a website.  
The api has some issues, so the backend fetches the data, handles it and serves it to the frontend. 
The bad api has 5min cache, so the backend polls data from there in 5 minute intervals.  
This was more focused on learning Go, as it's something that I'm really interested at.  

### 1.1 Implementation  
The application revolves around the Golang backend. The backend polls data from the bad api in 5 minute intervals, to get any new data.  
The frontend then polls data from the backend, and updates if there are any new products.  
As there wasn't any requirement for filtering, search, or pagination, I didn't implement any of those. All ~6000 thousand products are shown on one page.  

### 1.2 Technologies  
I used Golang on the backend, for the goroutines. As we had to fetch products from 3 categories and availability from ~7 manufacturers, goroutines are fast in this case.  
Fetching the previously mentioned takes from 15 to 25 seconds. Average being around 18 seconds.  
On the frontend I used React and React Window to 'render' all the products. Rendering 6000 products on their own took around 2.2 seconds and it was quite tedious when changing categories. As searching wasn't required and I wanted to avoid pagination, this was the best option for now. 

## 2. Architecture
The project consists of frontend and backend. Frontend makes requests for the data from the backend. Backend polls the data from the bad api.  
  
Frontend <-- Backend <-- Bad api
## 3. Development environment

### 3.1. Prerequisites, and what to do first
Docker or Node.js and golang is required for running the application locally.   
For Testing, Golang is required.  
Tested versions:  
Node v15, Golang 1.15.6.  
### 3.2. Run tests  
Run Server tests by running the following in the `server` folder:  
`go test -v ./...`

### 3.4. Start the application locally  
Build the image:  
`docker build -t warehouse .`  
 
Run the image:  
`docker run -p 3000:8080 warehouse`  

### 3.5. Access the application locally
Access the service from:  
`localhost:3000`


## 4. Production environment

### 4.1. Access  
The application is running on Heroku:  
[Application on Heroku](http://localhost:3000/beanies)  
The service takes some time to start, as it fetches the data from the bad api.

## 5. Continuous integration  
Nonexistent. Maybe the next step.  
