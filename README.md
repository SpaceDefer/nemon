
![output-onlinepngtools](https://user-images.githubusercontent.com/63122405/160476270-d64b5f68-2706-4991-bafa-48fe3ef16b56.png)


# Project Submission for SIH 2022 
Demo code for SIH 2020 submission. An app(SaaS) to remotely manage all the worker computers on a workstation from the Admin computers.

**Problem Statement :** To track list of software installed in the PC’s / Workstation’s attached to the network. \
**PS Code :** GR917 \
**Team Name :** Space_Defer. 

## Setup and run instructions 
To run this code locally, clone this repository and you can start *coordinator* or *worker* using following commands:

Please make sure that the coordinator and worker are connected on the same network(LAN).

To run coordinator :
```
go run main.go --mode coordinator
```

To run worker : 
```
go run main.go --mode worker
```

Once the coordinator and worker processes are running, the coordinator will send heartbeats to the workers every 10 seconds and gets a list of all the 
installed apps on the PC using RPCs.

## Video Demo of the Project : 

## UI Demo of the Project : 
The following UI prototype can be found [here](https://www.figma.com/proto/GYTEPNoBelBPkCCZvGyrnQ/Dashboard?node-id=214%3A481&scaling=contain&page-id=0%3A1&starting-point-node-id=1%3A6).

![Screenshot 2022-03-29 at 1 17 35 AM](https://user-images.githubusercontent.com/63122405/160475272-032cf8a3-5195-45fc-9829-4974a0aa45e9.png)
![Screenshot 2022-03-29 at 1 17 16 AM](https://user-images.githubusercontent.com/63122405/160475211-0443ef44-fcd1-46b0-9d9b-b95162ff36d7.png)


## Team Details
* Ujjwal Kadam ( 2019UCS2022 )
* Tarush Bajaj ( 2019UCS2015 )
* Tanmay Gairola ( 2019UCS2063 )
* Neha Yadav ( 2019UCS2008 )
* Atishek Kumar (2019UCS2027)
* Anish Chaudhary (2019UCS2050)

> Mentors : Dr. Mohinder Pal Singh Bhatia and Mr. Manu Sheel Gupta

