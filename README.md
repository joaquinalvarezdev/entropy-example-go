# Shannon entropy implementation on go 

Sample implementation of shannon entropy on go as a simple api endpoint.

Upload a file and specify the block size that's going to be used to calculate the entropy. (If not specified it will default to 1024).

Validate results using https://planetcalc.com/2476/

Todo-list:
* [x] Basic funcional example
* [ ] Unit testing
* [ ] Bit of refactor for better code structure


## How To Use ?
* Download latest release according to your OS from https://github.com/joaquinalvarezdev/entropy-example-go/releases/latest 
* Open postman or any api testing tool, create a POST request to point http://localhost:8000/API/entropy and in the body tab add two attributes:
  - file: any file you want to test. **MANDATORY**
  - blocksize: blocksize in bytes to chunk the file. **OPTIONAL**
