#!/bin/sh

rm -rf KubeArmor opt
mkdir -p ./KubeArmor/templates
mkdir -p ./opt/kubearmor/BPF

\cp -rf /usr/src/KubeArmor/KubeArmor/kubearmor ./KubeArmor/kubearmor
\cp -rf /usr/src/KubeArmor/BPF/*.o ./opt/kubearmor/BPF/
\cp -rf /usr/src/KubeArmor/KubeArmor/templates/* ./KubeArmor/templates/

echo "pack ok"
