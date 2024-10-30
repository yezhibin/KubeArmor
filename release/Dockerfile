FROM alpine:3.20 AS kubearmor

ENV DEBUG=true

RUN echo "@community http://dl-cdn.alpinelinux.org/alpine/edge/community" | tee -a /etc/apk/repositories

RUN apk --no-cache update
RUN apk add apparmor@community apparmor-utils@community bash

COPY KubeArmor/kubearmor /KubeArmor/kubearmor
COPY opt/kubearmor/BPF/*.o /opt/kubearmor/BPF/
COPY KubeArmor/templates/* /KubeArmor/templates/

ENTRYPOINT ["/KubeArmor/kubearmor"]
