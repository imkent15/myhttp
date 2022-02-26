## myhttp
This is a tool which will make http request to the given web address and prints the checksum of the response body.
Capable of making parallel requests. User can provide a flag to indicate the maximum no of parallel requests

### Build Instructions
Clone the repository
 - https://github.com/imkent15/myhttp.git
Move inside the root directory
 - cd myhttp
Issue the below command to build the tool
 - go build .
 
### Executing the tool
User can execute this tool by providing the web addresses as arguments.
User can also set the -parallel flag to indicate the max no of parallel processors.
If the parallel flag is not set, the value will be defaulted to 10.
the below are few examples and their responses.
