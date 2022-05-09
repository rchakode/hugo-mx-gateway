# Configuration variables

Regardless of the deployment platform (Google App Engine, Kubernetes, Docker), the following configuration parameters must be provided when deploying hugo-mx-gateway.

* `SMTP_SERVER_ADDR`: Set the address of the SMTP server in the form of `host:port`. It's required that the SMTP server being supporting TLS.
* `SMTP_VERITY_CERT`: Tell whether the SMTP certificate should be validated against top level authorities. If you're using a self-signed certificate on the SMTP server, this value must be set to `false`.
* `SMTP_AUTHENTICATION_ENABLED`: Boolean (default: `true`) indicating whether SMTP authentication is required or not. If true, the variables `SMTP_CLIENT_USERNAME` and `SMTP_CLIENT_PASSWORD` are used the perform the authentication.
* `SMTP_CLIENT_USERNAME`: Set the username to connect to the SMTP server.
* `SMTP_CLIENT_PASSWORD`: Set the password to connect to the SMTP server.
* `CONTACT_REPLY_EMAIL`: Set an email address for the reply email. It's not necessary a valid email address; for example if don't want the user to reply the email, you can set something like `noreply@example.com`.
* `CONTACT_REPLY_BCC_EMAIL`: Sets an email address for bcc copy of the email sent to the user. This is useful for tracking and follow up.
* `DEMO_URL`: Specific for demo forms, it can be used to set the URL of the demo site that will be included to the user reply email (e.g. `https://demo.example.com/`). 
* `ALLOWED_ORIGINS`: Set a list of comma-separated list of domains that the `hugo-mx-gateway` App should trust. For security reason, only requests with an `Origin` header belonging to the defined list of origins will be accepted.
* `RECAPTCHA_PRIVATE_KEY` (optional): The [reCaptcha](https://www.google.com/recaptcha/intro/v3.html) private key.
* `TEMPLATE_DEMO_REQUEST_REPLY` (optional): Specify the path of the template to reply a demo request. The default templare is `templates/template_reply_demo_request.html`. The template is based on [Go Template](https://golang.org/pkg/text/template/). 
* `TEMPLATE_CONTACT_REQUEST_REPLY` (optional): Specify the path of the template to reply a contact request. The default templare is `templates/template_reply_contact_request.html`. The template is based on [Go Template](https://golang.org/pkg/text/template/). 
