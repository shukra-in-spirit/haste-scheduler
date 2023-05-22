from datetime import datetime
import time
from queue import PriorityQueue

r = datetime(year=2023, month=5, day=19, hour=3, minute=00, second=00)
license_input = [
    {"period": 10, "id": "A"},
    {"period": 8, "id": "B"},
    {"period": 5, "id": "C"},
    {"period": 9, "id": "D"},
    {"period": 4, "id": "E"},
    {"period": 3, "id": "F"},
    {"period": 5, "id": "G"},
    {"period": 7, "id": "H"},
    {"period": 4, "id": "I"},
    {"period": 7, "id": "J"},
    {"period": 10, "id": "AA"},
    {"period": 8, "id": "BB"},
    {"period": 5, "id": "CC"},
    {"period": 9, "id": "DD"},
    {"period": 4, "id": "EE"},
    {"period": 3, "id": "FF"},
    {"period": 5, "id": "GG"},
    {"period": 7, "id": "HH"},
    {"period": 4, "id": "II"},
    {"period": 7, "id": "JJ"}
]
pq = PriorityQueue()


def activateLicense(license):
    print("Adding license: " + license["id"])
    t_now = datetime.now()
    score = int((t_now - r).total_seconds()) + license["period"]
    pq.put((score, license["id"]))


# for l in license_input:
#     addLicense(l)


for n in range(0, 20):
    time.sleep(1)
    print("day " + str(n+1))
    activateLicense(license_input[n])
    min = pq.queue[0]
    while (int((datetime.now() - r).total_seconds())) >= min[0]:
        print("license expired: " + str(pq.get()))
        min = pq.queue[0]
    print("")
    # print(min[0])
    # print(pq.queue)
