#!/bin/sh

version=$1

if [ "$version" == "" ];then
    version=debug_v1.0.0
fi

docker build -t docker.io/kubearmor/kubearmor:${version} -f Dockerfile_selinux . 
docker save kubearmor/kubearmor:${version} -o kubearmor_${version}.tar

exist=`microk8s ctr image ls | grep "${version}"`
if [ "$exist" != "" ];then
    microk8s ctr image delete docker.io/kubearmor/kubearmor:${version}
    echo "[WARN] delete older docker.io/kubearmor/kubearmor:${version}"
fi

microk8s ctr image import kubearmor_${version}.tar
echo "kubearmor/kubearmor:${version} import to microk8s ok"
