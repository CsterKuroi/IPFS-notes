for host in $(docker ps -aq)
do
echo The state is $host
docker exec -it $host sh -c "ipfs $*"
done
