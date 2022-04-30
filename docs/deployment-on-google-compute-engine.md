# Deploying hugo-mx-gateway on Google App Engine

  - [Requirements](#requirements)
  - [Setup Procedure](#setup-procedure)

## Requirements
This procedure requires to have: 
 * an active GCP account.
 * [Google Cloud SDK](https://cloud.google.com/sdk) (gcloud) installed and configured on your work station.
 * The credentials configured for gcloud must have sufficient permissions to create an App Engine application.


## Setup Procedure
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
