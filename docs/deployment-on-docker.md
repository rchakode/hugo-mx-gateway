# Deploying hugo-mx-gateway on Docker

As described in the below procedure, an instance of `hugo-mx-gateway` can be quickly started on any machine running Docker.

First review the [configuration variables](./configuration-variables.md). 

Then apply the following command while setting the configuration variables appropriately:

  ```
  docker run -d \
    --publish 8080:8080 \
    --name 'hugo-mx-gateway' \
    -e SMTP_SERVER_ADDR="smtp.example.com:465" \
    -e SMTP_SKIP_VERIFY_CERT=false \
    -e SMTP_CLIENT_USERNAME="postmaster@example.com" \
    -e SMTP_CLIENT_PASSWORD="postmasterSecretPassWord" \
    -e CONTACT_REPLY_EMAIL="noreply@example.com" \
    -e CONTACT_REPLY_BCC_EMAIL="contact@example.com" \
    -e DEMO_URL="https://demo.example.com/" \
    -e ALLOWED_ORIGINS="127.0.0.1,example.com" \
    docker.io/rchakoder/hugo-mx-gateway
  ```

Check that the container is up and functionning.

  ```
  curl http://127.0.0.1:8080/
  ```
The output in case of success shall be `{"status": "ok"}`.
