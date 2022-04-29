
# Sample of contact form for Hugo

From the source tree, the file `samples/hugo-partial-contact-form.html` contains a sample HTML form for Hugo. It can be used for both contact and demo requests.

![Screenshot of a successful submission](../screenshots/form-submitted-and-processed-with-success.png)

Open the file in your favorite editor and review it.

Notice that the form is configured to be rendered specifically according to a Hugo parameter named **tags**, which is actually a **list of tags**. If the parameter holds a tag named `contact` then, the form will be rendered as a contact form. Otherwise, it'll be rendered as a demo form.

The integration works as follows:
 * Copy the HTML form content in your target **Hugo HTML template**. 
 * Modify the `<form>` tag to make the **action** point to the URL of the sendmail backend deployed previously.
   * On Google App Engine, the endpoint shall be: https://PROJECT-ID.REGION.r.appspot.com/sendmail. Replace `PROJECT-ID` and `REGION`, repectively, with the GCP's project ID and the deployment region.
   * On Kubernetes, the in-cluster endpoint shall be: http://hugo-mx-gateway.hugo-mx-gateway.svc.cluster.local/sendmail
   * On Docker, the endpoint shall be: http://DOCKER-HOST:8080/sendmail. Replace `DOCKER-HOST` with the IP address or the hostname of the Docker machine.
 * Edit the **Hugo Markdown content** of the target contact/demo page to ensure that the **tags** parameter holds a appropriate value (i.e. `contact` for a contact form, or `demo` for a demo request form).

Here is an example of header for a Hugo Markdown page. The `tags` parameter holds a tag named `contact`) meaning that the page will be rendered as a contact request form.

  ```
  ---
  title: "Contact Us"
  description: "Contact request page"
  date: 2020-04-25T00:20:27+02:00
  tags: [contact]
  ---
  ```