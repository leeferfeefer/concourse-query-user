#!/bin/sh

 cd /usr/local/concourse

./bin/concourse worker \
     --work-dir ./work_dir \
     --tsa-host localhost:2222 \
     --tsa-public-key ~/Desktop/concourse-docker/keys/web/tsa_host_key.pub \
     --tsa-worker-private-key ~/Desktop/concourse-docker/keys/worker/worker_key
