# Calculating-Software-Application-s-Copy
## Abstract
As a deliverable to a Coding Test, which is to calculate the total copy of application software. Please refer to original ![requirements](requirements/requirements.pdf "requirements")  for more elaborate requirements. The software is implemented in ubiquitous language Golang of version 1.17.1 (go1.17.1 windows/amd64).

## Design
### Set Theory
Following set diagram illustrates philosophy behind the design. In scanning input csv file, only records associated with specified ApplicationID will be retrieved for calculation.
![Set](image/set.png "Set") 
Among the retrieved records, there will be further categorisation by UserID. The remainder of criteria of calculation are ComputerID and ComputerType.

### Approach of Reading Records - All into memory or one at a time
Initially I have tried reading records were read all at once to memory before being parsed for calculation. In the case of processing a relatively large csv file (sizes 0.99Gb) file I received, according to Task Manager of my Win7 laptop machine, the app eats up nearly 1.2GB of memory, reaching upper limit of the laptop’s memory. In light of even bigger CSV files, the memory resources could become an issue. <br/>
As a matter of fact, it’s a trade-off between resources (memory) and performance. Provided memory is unlimited, then loading all records once into memory before parsing and calculation could be ideal. But memory is not unlimited. <br/>
Hence the adjustment is that up to pre-configured number, which is currently is currently 50, of records are loaded into buffered channel at a time. The buffered channel serves the communication between the process of loading CSV file and main program. With such change, tests indicate that the memory consumption stops peaking to its maximum.

