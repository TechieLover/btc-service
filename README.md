# Command to create a docker build
sudo docker build -t "btc" ~/work/src/vdart/btc-service/

# Command to run a container and mapping to external localhost 
sudo docker run -p 127.0.0.1:3000:3000 --env-file ~/work/src/vdart/btc-service/dev.env "btc"

# Access to service is http://localhost:3000/ping

# Command to get internal_ip of container by providing container_id , to get container_id use "sudo docker ps" 
sudo docker inspect --format '{{ .NetworkSettings.IPAddress }}' 'container_id'

# Access to service is http://internal_ip:3000/ping

# create a folder vdart "mkdir ~/work/src/vdart" then copy btc-service folder
# To run a service localally then Go to path "cd ~/work/src/vdart/btc-service" then run "./build.sh"  

# Open Endpoints wrt to localhost
1. http://localhost:3000/v1/currency/:symbol --- Accepted symbols are 'BTCUSD,ETHBTC'
2. http://localhost:3000/v1/allcurrency   -- To Get all currencies

