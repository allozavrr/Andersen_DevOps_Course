## Hello, my friend.

* **What is it?**

  This is a small app for output "Hello, world!" via http. Application listen port 8080 and saves logs.

* **How it works?**

  For use this application install next packages:

  *go get github.com/gin-gonic/gin*

Then compile application from same folder.

And build app: 

  *go build Hello_World.go*

And run app:

  *./helloworld*
  

* **OK, and what about Docker?**

You can install app with Docker, friend!

Download this repository and build image

   *docker build. -t {image_name}*

And run the container

   *docker run -d -p 8080: 8080 {image_name}*


* **OK, and what about Kubernetes?**

[You can download minikube](https://github.com/kubernetes/minikube) and [kubectl](https://kubernetes.io/docs/tasks/tools/)
  
Then do it:

   *minikube start
   minikube addons enable ingress
   kubectl config use-context minikube*

Testing:
   
   *kubectl get deployment
   kubectl get service
   kubectl get ingress*

And add domen **advent.test** to local file **/etc/hosts**

Have a good day and fun!
