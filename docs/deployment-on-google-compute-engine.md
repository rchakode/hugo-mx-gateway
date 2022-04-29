# Deploying hugo-mx-gateway on Google App Engine


This requires to have an active GCP account and [Google Cloud SDK](https://cloud.google.com/sdk) (gcloud) installed on your work station.

* Create/select a GCP project to deploy `hugo-mx-gateway`. 
  
  Note that each GCP project can hold only a single App Engine instance. Several applications can be co-hosted as services for the root App Engine instance. In this case, a new application has to be declared as `service` in the `app.yaml` file.
* Create the Google App Engine configuration file 
  ```
  cp app.yaml.sample app.yaml
  ```
* Open the `app.yaml` file with your favorite editor.
* Edit the configuration variables as described [here](#configuration-variables).
* Start the deployment
  ```
  make deploy-gcp
  ```
* Check that `hugo-mx-gateway` is up and working
  ```
  curl https://PROJECT-ID.REGION.r.appspot.com/ 
  ```
  Replace `PROJECT-ID` with the GCP project ID, and `REGION` with the deployment region.*
  
  The output in case of success shall be `{"status": "ok"}`.
