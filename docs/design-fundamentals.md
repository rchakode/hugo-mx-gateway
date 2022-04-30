# Design Fundamentals of Hugo-mx-gateway 

`hugo-mx-gateway` is built upon a simple request handling workflow:

* Create an HTML form with a POST action pointing towards the `hugo-mx-gateway` service. This service enables a RESTful POST endpoint backed by an application built in Go. The application is released as container image along with manifests to ease deployment on [Google App Engine](https://cloud.google.com/appengine), Kubernetes, and Docker environments.
* For each user request, `hugo-mx-gateway` automatically retrieves information submitted by the user (email, subject, message details...), then **generates and sends** a **templated email** (based on [Go Template](https://golang.org/pkg/text/template/)) to the user-provided email address, while **bcc**ing a copy of that email to an address that you can define for internal tracking and follow up.
* Once a request is processed (upon success or failure), `hugo-mx-gateway` handles the reply back towards the calling static page by redirecting the browser to the origin page with additional URL parameters describing the completion status of the processing (e.g. `/contact.html?status=success&message=request%20submitted`). The parameters can then be easily retrieved and shown to the user, e.g. with a few lines of Javascript within the static page.
* The project is shipped with a sample HTML form intending to cover a basic use case of contact and demo requests. That said, this is a open source software, so you're free to adapt it for your specific use cases.
