# TidyBeaver's API code

## This code has been implemented in the repository "TidyBeaver".
## For simplicity, the code for this repository is a standalone version of the API that the Log Aggregator TidyBeaver uses.

# Simple instructions:

### Api Request
#### http://localhost:9090/api/random-response?count=$
##### Change the $ with an integer number for the quantity of Logs you'd like to receive back.

### How to run the API on the background
### Use this command on the terminal once it is within the scope of main.go.
#### nohup go run main.go > output.log 2>&1 &

### How to kill the API's background process 
### Use this command on the terminal once it is within the scope of main.go.
#### ps aux | grep main.go   # Find PID
#### kill <PID>
