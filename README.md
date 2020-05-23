![](https://img.shields.io/github/license/rchakode/hugo-mx-gateway.svg?label=License)
[![Actions Status](https://github.com/rchakode/hugo-mx-gateway/workflows/Build/badge.svg)](https://github.com/rchakode/hugo-mx-gateway/actions)
![](https://img.shields.io/docker/pulls/rchakode/hugo-mx-gateway.svg?label=Docker%20Pulls)


- [Overview](#overview)
  - [How it Works](#how-it-works)
  - [Screenshots](#screenshots)
- [Configuration variables](#configuration-variables)
- [Deployment options](#deployment-options)
  - [Deployment on Google Appp Engine](#deployment-on-google-appp-engine)
    - [Before you begin](#before-you-begin)
    - [Configure and deploy the App](#configure-and-deploy-the-app)
  - [Deployment on a Kubernetes cluster](#deployment-on-a-kubernetes-cluster)
    - [Installation using Helm 3 (i.e. without tiller)](#installation-using-helm-3-ie-without-tiller)
    - [Installation using Kubectl](#installation-using-kubectl)
  - [Deployment on Docker](#deployment-on-docker)
  - [Test the App](#test-the-app)
- [Samples of Hugo Contact Forms](#samples-of-hugo-contact-forms)
- [Developers' corner](#developers-corner)
- [License & Copyrights](#license--copyrights)
- [Support & Contributions](#support--contributions)


# Overview
Did you ever experience building a static website (e.g. using [Hugo](https://gohugo.io/)) or whatever alternative, and stuck when coming the time to add a contact or a demo request form?

You're at the place.

`hugo-mx-gateway` provides a HTTP POST endpoint that can be bound to your HTML template to handle user contact and demo requests in a simple yet powerful way.

## How it Works
`hugo-mx-gateway` is built upon a simple request form handling workflow:

* You create a HTML form with a POST action pointing towards the `hugo-mx-gateway` service.
* The `hugo-mx-gateway` service is a RESTful HTTP POST endpoint backed by an application easily deployable on [Google App Engine](https://cloud.google.com/appengine), on Kubernetes, or on Docker. Hereafter we also refer to it as **the App**.
* For each form request, the App retrieves information submitted by the user (email, subject, message details...), **automatically generates and sends a _templated email_ to the user**, while **bcc**ing a copy of that email to an address that you do define for tracking and follow up.
* Once a request is processed (upon success or failure), the App copes with the reply towards the origin web site (static page) by redirecting the browser to the request page with additionnal URL parameters describing the status of the processing (e.g. `/contact.html?status=success&message=request%20submitted`). With this, you can then add a few lines of Javascript to retrieve and display the reply message on the page.
* The App is shipped with a sample HTML form including some common fields for contact and demo requests, as well as a sample Javascript code to handle the processing response. That said, this is a open source software, so you're free to adapt it for your specific use cases.

## Screenshots
This screenshot show an example of a form successfully submitted and handled by the backend, which replied with the message in green.

![Screenshot of a successful submission](./screenshots/form-submitted-and-processed-with-success.png)

# Configuration variables
According to your deployment approach (Google App Engine, Kubernetes or Docker), you must provide the following configuration parameters as environment variables:

* `SMTP_SERVER_ADDR`: Set the address of the SMTP server in the form of `host:port`. It's required that the SMTP server being supporting TLS.
* `SMTP_VERITY_CERT`: Tell if the `hugo-mx-gateway` App should validate the SMTP certificate against valid authorities. If you're using a self-signed certificate on the SMTP server, this value must be set to `false`.
* `SMTP_CLIENT_USERNAME`: Set the username to connect to the SMTP server.
* `SMTP_CLIENT_PASSWORD`: Set the password to connect to the SMTP server.
* `CONTACT_REPLY_EMAIL`: Set an email address for the reply email. It's not necessary a valid email address, for example if don't want the user to reply you can use something like `noreply@example.com`.
* `CONTACT_REPLY_BCC_EMAIL`: Sets an email address for bcc copy of the email sent to the user. This is useful for tracking and follow up.
* `DEMO_URL`: Specific for demo forms, it can be used to set the URL of the demo site that will be included to the user reply email (e.g. `https://demo.example.com/`). 
* `ALLOWED_ORIGINS`: Set a list of comma-separated domains that the `hugo-mx-gateway` App shoudl trust. This is for security reason to filter requests. Only requests with an `Origin` header belonging to the defined origins will be accepted, through it's only required that the request has a valid `Referer` header. It's expected in the future to these request filtering and admission rules.
* `TEMPLATE_DEMO_REQUEST_REPLY`: Specify the path of the email template to reply to demo requests. The default templare used in described in the file `templates/template_reply_demo_request.html`
* `TEMPLATE_CONTACT_REQUEST_REPLY`: Specify the path of the email template to reply to contact requests. The default templare used in described in the file `templates/template_reply_contact_request.html`.

# Deployment options

## Deployment on Google Appp Engine

### Before you begin
To deploy the `hugo-mx-gateway` App on Google App Engine, make sure that you have a active GCP account and:
* Install the [Google Cloud SDK](https://cloud.google.com/sdk) (gcloud) installed on your work station.
* Create/select a GCP project to deploy the App. Note that each GCP project can hold only a single App Engine instance.

### Configure and deploy the App
* Create the Google App Engine configuration file 
  ```
  cp app.yaml.sample app.yaml
  ```
* Open the `app.yaml` file with your favorite editor.
* Set the configuration variables as described [here](#configuration-variables).
* Apply the deployment as follows:
  ```
  make deploy-gcp
  ```

On success the sendmail POST endpoint shall be reachable at the address: `https://<project_id>.<region>.r.appspot.com/sendmail`

Replace:
* `<project_id>` with the GCP project ID.
* `<region>` with the region of the App Engine instance. 


## Deployment on a Kubernetes cluster
<a name="deployment-on-a-kubernetes-cluster"></a>

There is a [Helm chart](./helm/) to ease the deployment on Kubernetes using Helm or `kubectl`. 

The Helm based deployment has been validated with Helm 3, i.e. without `Tiller`.

Either way, check the [values.yaml](./helm/values.yaml) file to set the [configuration options](#configuration-variables) according to your specific settings.

> **Security Context:**
> `hugo-mx-gateway`'s pod is deployed with a unprivileged security context by default. However, if needed, it's possible to launch the pod in privileged mode by setting the Helm configuration value `securityContext.enabled` to `false`.

In the next deployment commands, it's assumed that the target namespace `hugo-mx-gateway` does exist. Otherwise create it first, or, alternatively, adapt the commands to use any other namespace of your choice.

### Installation using Helm 3 (i.e. without tiller)
<a name="installation-using-helm-3-ie-without-tiller"></a>

Helm 3 does not longer require to have [`tiller`](https://v2.helm.sh/docs/install/).

As a consequence the below command shall work with a fresh installation of `hugo-mx-gateway` or a former version installed with Helm 3. There is a [known issue](https://github.com/helm/helm/issues/6850) when there is already a version **not** installed with Helm 3.

```
helm upgrade --namespace hugo-mx-gateway --install hugo-mx-gateway helm/
```

### Installation using Kubectl
<a name="installation-using-kubectl"></a>
This approach requires to have the Helm client (version 2 or 3) installed to generate a raw template for kubectl.

```
$ helm template hugo-mx-gateway --namespace hugo-mx-gateway helm/ | kubectl apply -f -
```

## Deployment on Docker
`hugo-mx-gateway` is released as a Docker image. So you can quickly start an instance of the service by running the following command:

```
$ docker run -d \
   --publish 8080:8080 \
   --name 'hugo-mx-gateway' \
   -e SMTP_SERVER_ADDR="smtp.example.com:465" \
   -e SMTP_VERITY_CERT=true \
   -e SMTP_CLIENT_USERNAME="postmaster@example.com" \
   -e SMTP_CLIENT_PASSWORD="postmasterSecretPassWord" \
   -e CONTACT_REPLY_EMAIL="noreply@example.com" \
   -e CONTACT_REPLY_BCC_EMAIL="contact@example.com" \
   -e DEMO_URL="https://demo.example.com/" \
   -e ALLOWED_ORIGINS="127.0.0.1,example.com" \
   rchakode/hugo-mx-gateway
```

In this command, you SHOULD adapt the values of configuration variables as described [here](#configuration-variables).

## Test the App
```
curl -H'Origin: http://example.com'  \
    -H'Referer: example.com' \
    -H'Content-Type: application/x-www-form-urlencoded' \
    -d 'target=contact' \
    -XPOST https://<project_id>.<region>.r.appspot.com/sendmail 
```

# Samples of Hugo Contact Forms
See in `./samples/`.


# Developers' corner
To build the stack
```
make build
```


# License & Copyrights
This tool (code and documentation) is licensed under the terms of Apache 2.0 License. Read the [LICENSE](LICENSE) file for more details on the license terms.

The tool may inlcude third-party libraries provided with their owns licenses and copyrights, but always compatible with the Apache 2.0 License terms of use.

# Support & Contributions
We encourage feedback and do make our best to handle any troubles you may encounter when using this tool.

Here is the link to submit issues: https://github.com/rchakode/hugo-mx-gateway/issues.

New ideas are welcomed, please open an issue to submit your idea if you have any one.

Contributions are accepted subject that the code and documentation be released under the terms of Apache 2.0 License.

To contribute bug patches or new features, please use the Github Pull Request model.