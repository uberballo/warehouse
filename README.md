## Warehouse  
## 1. Project purpose
Fetches product data from a bad api and displays it on a website.  
The bad api has some issues, so the backend fetches the data, handles it, saves it on the memory and serves it to the frontend. 
The bad api has 5min cache, so the backend polls data from there in 5 minute intervals.  
  
This was more focused on learning Go, as it's something that I'm really interested at. Feel free to make a pull request or raise a issue, if there are any parts where I could do differently.   

### 1.1 Implementation  
The application revolves around the Golang backend. The backend polls data from the bad api in 5 minute intervals, to get any new data.  
The frontend then polls data from the backend, and updates if there are any new products.  
  
As there wasn't any requirement for filtering, search, or pagination, I didn't implement any of those. All ~6000 thousand products are shown on one page.  
So I used React window for this. We only render the product cards that the user can see. The React window required some extra work, so the site could work on different window sizes.    
Otherwise, the backend serves build frontend.  


### 1.2 Technologies  
I used Golang on the backend, for the goroutines. As we had to fetch products from 3 categories and availability from ~7 manufacturers, goroutines are fast in this case.  
Fetching the previously mentioned takes from 15 to 25 seconds. Average being around 18 seconds.  
  
I used react window because rendering the site with 6000 element's takes around 2.2 seconds and having the backend build a new static website every 5 minutes seemed bit overkill, I opted for the window.  
For dynamic sizing, I had to add React-virtualized-auto-sizer and some window.eventListener's so the grid works on different sized screens. React window has only grids or lists. So with some index magic, I can pretend that the grid is like a flexbox. 


## 2. Architecture
The project consists of frontend and backend. Frontend makes requests for the data from the backend. Backend polls the data from the bad api. Super simple stuff  
  
Frontend <-- Backend <-- Bad api
## 3. Development environment

### 3.1. Prerequisites, and what to do first
Docker or Node.js, golang and yarn is required for running the application locally.   
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

Or without docker:  

Start the backend on the folder `server` with:  
`go run main.go`  

then install clients dependencies on the folder `client` with:  
`yarn`  
Start the client with:  
`yarn start`  

### 3.5. Access the application locally
Access the service from:  
`localhost:3000`


## 4. Production environment

### 4.1. Access  
The application is running on Heroku:  
[Application on Heroku](https://cool-smooth-warehouse.herokuapp.com/facemasks)  
If the service hasn't been in use for some time, Heroku has to spin up the application and this process may take some time. After the application starts, It will first fetch the product data without availability and updates them with the availability after some time.  

## 5. Continuous integration  
Nonexistent. Maybe the next step.  
