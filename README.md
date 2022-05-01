# Calculating-Software-Application-s-Copy
## Abstract
As a deliverable to a Coding Test, which is to calculate the total copy of application software. The original requirements can be seen in the file named requirements.pdf under sub-folder requirements for more elaboration. The software is implemented in ubiquitous language Golang of version 1.17.1 (go1.17.1 windows/amd64).

## Design
### Set Theory
Following set diagram illustrates philosophy behind the design. In scanning input csv file, only records associated with specified ApplicationID will be retrieved for calculation.
![Set](image/set.png "Set") 
Among the retrieved records, there will be further categorisation by UserID. The remainder of criteria of calculation are ComputerID and ComputerType.
