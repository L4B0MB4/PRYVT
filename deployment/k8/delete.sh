
kubectl delete -f ingress.yml 
kubectl delete -f identification_query.yml 
kubectl delete -f identification_command.yml 
kubectl delete -f evtsrc.yml 

echo "delete.sh: done"