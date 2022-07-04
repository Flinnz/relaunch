# Simple Command Relaunch in Go
I really should just make shell script for this, but Windows exists. Coverage is 0% because tests is for cowards. 

Written in 5 minutes.
## Build:
    go build
## Usage: 
    relaunch [OPTION] [CMD TO RELAUNCH]
## Options: 
### -c 
----

    Amount of relaunches. 0 is infinite. Default is 0.
    
### -d 
----
    Delay between relaunches in milliseconds. Default 500