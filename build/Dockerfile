FROM centos:centos7

RUN mkdir -p /home/data1/webhook/logs
ADD ./admission-webhook /home/data1/webhook
ADD ./control.sh /home/data1/webhook
WORKDIR /home/data1/webhook
ENTRYPOINT ["/bin/bash","control.sh","start"]
