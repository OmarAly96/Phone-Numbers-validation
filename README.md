# Phone-Numbers

To run the backend

  cd backend
  
  go run main.go
  
To run the frontend

  cd frontend/phone-numbers-client
  
  npm install
  
  ng serve
  
Main API: 

  GET: http://localhost:8080/api/phone-numbers
  
  queryparams: limit,offset,state,country
  
  example: 
  
    http://localhost:8080/api/phone-numbers?limit=5&offset=0&state=0&country=Morocco,Cameroon
