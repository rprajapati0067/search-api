# Search-api
1. Restaurant API implemented with 3 nearest restaurant.
2. Parking Spots API implemented with 3 nearest spots
3. Charging Stations API implemented with 3 nearest spots

# TODO
  1. Have to call API concurrently.
  2. Enable Caching of APIS.
  3. Have to write test cases.
  4. Package and deliver the final outcome as a Docker container.

# Steps to Run API
Open cmd in the project directory
run command go run main.go
Application starts on 
http://localhost:8080
Hit below endpoint by providing at and cat parameters
http://localhost:8080/search?at=52.4708,13.3281&cat=restaurant
