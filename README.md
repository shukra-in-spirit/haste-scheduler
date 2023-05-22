# Haste
A simple alarm using Redis simple set and 1-d vector calculation.
An alarm service allows users or other services to persist `Alarm Item(s)` as time-based events in a mem-cached Priority Queue. When the `period` specified in the alarm expires, it sends the `message` to the `endpoint` specified in the alarm when it was initially set. The communication is supported for GRPC, REST and Message Queue. 

### Run in github codespaces
Open two terminal in Github Codespaces. The first terminal will be used to bring up the Haste service and the redis service. The second one is used to stress test the Haste service.


Step 1: Run `docker compose up` in the first terminal. Wait for it to complete. The gin server should come up.

Step 2: Run `bash scripts\parallel_n.sh` in the second terminal. The response can be watched in the first terminal.

### Stress Test Output (22/5/23)
![image](https://github.com/shukra-in-spirit/haste-scheduler/assets/104008671/2258285d-613e-4c87-a910-d2b948ee99dc)

Current Accuracy is at below 1 second error on 128m memory and 100m cpu for service pod.
Redis configuration: 256m memory, 300m cpu
