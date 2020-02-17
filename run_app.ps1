docker build -t hello-world .
docker run -p 8090:8090 -it --rm --name hello-world-container hello-world