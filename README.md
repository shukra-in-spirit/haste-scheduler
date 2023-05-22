# Haste
A simple alarm using Redis simple set and 1-d vector calculation.
An alarm service allows users or other services to persist `Alarm Item(s)` as time-based events in a mem-cached Priority Queue. When the `period` specified in the alarm expires, it sends the `message` to the `endpoint` specified in the alarm when it was initially set. The communication is supported for GRPC, REST and Message Queue. 

## Run in github codespaces
Open two terminal in Github Codespaces. The first terminal will be used to bring up the Haste service and the redis service. The second one is used to stress test the Haste service.


Step 1: Run `docker compose up`.
Step 2: Run `bash scripts\parallel_n.sh`.
