#!/bin/zsh

export MACHINEONE=$(kubectl get pod -l app=machine-1 -o jsonpath={.items..metadata.name} -n=machines)
export MACHINETWO=$(kubectl get pod -l app=machine-2 -o jsonpath={.items..metadata.name} -n=machines)
export MACHINETHREE=$(kubectl get pod -l app=machine-3 -o jsonpath={.items..metadata.name} -n=machines)

for (( ; ; ))
do
    kubectl exec "$MACHINEONE" -c machine-1 -- curl -sV -v machine-2.machines.svc.cluster:local:8001 -n=machines
    kubectl exec "$MACHINETWO" -c machine-2 -- curl -sV -v machine-3.machines.svc.cluster:local:8001 -n=machines
    kubectl exec "$MACHINETHREE" -c machine-3 -- curl -sV -v machine-1.machines.svc.cluster:local:8001 -n=machines
    sleep 2s
done
