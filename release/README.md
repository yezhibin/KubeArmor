### kubearmor镜像制作方法

```bash
#!/bin/sh

cd KubeArmor/KubeArmor && make
cp -rf KubeArmor/BPF ./ && cd BPF && make
cd release && sh prepack.sh
docker build -t kubearmor/kubearmor:debug_v1.0.0 . 
docker save kubearmor/kubearmor:debug_v1.0.0 -o kubearmor_debug_v1.0.0.tar
microk8s ctr image import kubearmor_debug_v1.0.1.tar
```
