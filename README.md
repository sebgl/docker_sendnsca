# docker_sendnsca
Send nsca checks results by running a docker image

## Usage

```
docker run \
    -e NSCA_SERVER=my.nsca.receiver \
    -e NSCA_PORT=5667 \
    -e NSCA_ENCRYPTION=1 \
    -e NSCA_KEY=my.encryption.key \
    sebgl/docker_sendnsca \
    --host hostname --service servicename --state 0 --msg ok
```