docker build -t hello-world .
docker run -p 8080:8080 -it --rm --name hello-world-container hello-world