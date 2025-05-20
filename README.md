# 101 TidyBeaverAPI

## Api Request
### http://localhost:9090/api/random-response?count=5

## Running the API on background
### nohup go run main.go > output.log 2>&1 &

## Kill the API's background process 
### ps aux | grep main.go   # Find PID
### kill <PID>
