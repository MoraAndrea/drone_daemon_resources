echo "Build Application..."
go build -o app .

echo "Create image docker..."
docker build -t drone-daemon-resources:first .

echo "Load image on kind..."
kind load docker-image drone-daemon-resources:first --name cluster1

# kubectl run --rm -i drone-daemon-resources --image=drone-daemon-resources:first -n demo
# kubectl apply -f drone-daemon-resources.yml
# kind load docker-image drone-daemon-resources:first --name cluster1