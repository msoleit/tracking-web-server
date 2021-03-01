# Tracking Web Server 

### /ping

- returns 200 Ok if `/tmp/ok exists`
- returns 503 otherwise

### /img
- returns 1*1 gif image.

------------------
## Ideas for more concurrency and scalability
- containerize the application (e.g using docker).
- Add a heartbeat mechanism.
- Add a threshold of active connections and spin new containers (we can use an exponential formula)
- use kubernetes to orchestrate the service containers.
