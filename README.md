# Secure Development With Docker &mdash; DockerCon 2023 workshop

<details><summary><h2>Setup the workshop platform</h2></summary>

### Requirements

- git
- Docker Desktop 4.24.0 or greater
- Docker Hub account `$ORG`

Docker Desktop must be configured to use containerd.  In Docker Desktop, go to
the Settings (⚙️ icon) &gt; Features in development and make sure the box next
to "Use containerd for pulling and storing images". If you changed the setting,
click the "Apply & Restart" button.

![Docker Desktop settings](./ss/desktop-containerd.png)

### Prerequisites

> *Conventions:*
> `$ORG` is the name of the Docker Hub account you will use.
> It can be a personal or a team one. Better if you have full ownership on it.

1. Clone the repository on your local machine
2. (optional) export `ORG` environment variable so you can more easily copy/paste commands

   ```console
   export ORG=<your organization namespace>
   ```
3. Configure organization for Docker Scout

   ```console
   docker scout config organization $ORG
   ```
4. Copy `env.dist` file to `.env`
5. Edit `.env` and set `NAMESPACE` to `$ORG`
6. Enroll your organization to Docker Scout

   ```console
   docker scout enroll $ORG
   ```
7. Checkout Hands-On #1

   ```console
   git checkout hands-on-1
   ```
8. Build demo images

   ```console
   docker compose --profile images build
   ```

   > This command will build two images we will explore.
   > To know more about how they are built look at
   > [`./docker-compose.yml`](./docker-compose.yml) and
   > [`./backend/Dockerfile`](./backend/Dockerfile) that
   > contains the build definitions.

   > In case of network issues, you can also build the following
   > image that is prebuilt and don't need extra dependencies.

   ```console
   docker compose --profile low_network build
   ```
</details>

<details><summary><h2>Hands-on #1: Remediating vulnerabilities</h2></summary>

### Reset git repository

```console
git reset --hard hands-on-1
```

### Base Image Fix (Docker Desktop)

1. Open Docker Desktop and select the image `$ORG/scout-demo-service:v1`

   ![](./ss/layer_view_scout-demo-service_v1.png)
2. Take the time to explore the different information displayed on this page

   - Image hierarchy, layers and the Images tab
   - Vulnerabilities
   - Packages

3. Select your base image and explore vulnerabilities specific to the base image.
   Then select "Recommendations for base image…" in the "Recommended fixes"
   dropdown in the upper right portion of the window.

   ![](./ss/select-base-image.png)

4. Select _Change base image_ and set the current image to 3.14

   ![](./ss/change-base-image.png)

5. Open [`frontend/Dockerfile`](./frontend/Dockerfile) in your favorite file editor
   and apply the "Tag is preferred tag" recommendation, i.e., change the `FROM` line to

   ```dockerfile
   FROM alpine:3.18
   ```

6. (optional) Update the tag to `v2` in [`docker-compose.yml`](./docker-compose.yml)
7. Rebuild the image

   ```console
   docker compose --profile scout-demo-service build
   ```
8. Open the image inside Desktop and see the impact of your change

   ![](./ss/layer_view_scout-demo-service_v2.png)

### Base Image Fix (Docker Scout CLI)

1. Run `docker scout cves` command against the image you just built:

   ```console
   docker scout cves $ORG/scout-demo-service:v1
   ```

   ![](./ss/scout-demo-service-v1-cves.png)
2. (optional) Explore the filters options from `docker scout cves` command
3. (optional) See vulnerable packages only

   ```console
   docker scout cves --format only-packages --only-vuln-packages $ORG/scout-demo-service:v1
   ```

   ![](./ss/scout-demo-service-v1-only-vuln-packages.png)

4. Explore base image recommendations using `docker scout recommendations`

   -> specify the tag upfront, like we selected it in Docker Desktop

   ```console
   docker scout recommendations --tag 3.14 $ORG/scout-demo-service:v1
   ```

   ![](./ss/scout-demo-service-v1-recommendations.png)

5. Apply the "Tag is preferred tag" recommendation to the [`frontend/Dockerfile`](./frontend/Dockerfile)
6. (optional) Update the tag to `v2` in [`docker-compose.yml`](./docker-compose.yml)
7. Rebuild the image

   ```console
   docker compose --profile scout-demo-service build
   ```
8. Run `docker scout cves` command against the image you just built and see the changes:

   ```console
   docker scout cves $ORG/scout-demo-service:v2
   ```

   ![](./ss/scout-demo-service-v2-cves.png)
9. Compare the two images to see the differences:

   ```console
   docker scout compare $ORG/scout-demo-service:v2 --to $ORG/scout-demo-service:v1
   ```

   ![](./ss/scout-demo-service-compare.png)

### Application Dependency Fix

1. Explore the still existing vulnerabilities: (Desktop or CLI)
   - find vulnerable package
   - find fix version
2. Update [`package.json`](./frontend/package.json) to upgrade the dependency
3. (optional) Update the tag to `v3` in [`docker-compose.yml`](./docker-compose.yml)
4. Rebuild the image
5. Explore the built image to see what changed
   - using Desktop
   - using CLI

### Extra Command

1. Get a quick overview of your image, including vulnerabilities and recommendations

   ```console
   docker scout quickview $ORG/scout-demo-service:v1
   ```

### Extra Image

Repeat the above steps for the `$ORG/scout-demo-service-back:v1` image (or any other image you have).

</details>

## Hand-on #2: Using Docker Scout to connect your data model

### Reset git repository

```console
git reset --hard hands-on-2
```

### Registry - https://scout.docker.com Integration

#### Create repositories

Go to https://hub.docker.com and create repositories for the images you will push

- `$ORG/scout-demo-service`
- `$ORG/scout-demo-service-back`

#### Enable repositories for Docker Scout

_Option 1: Use Docker Scout Dashboard_

1. Go to https://scout.docker.com
2. Select your organization in the dropdown next to your user
3. Open the settings menu (⚙️ icon) and select _Repository settings_
4. Select the repository to enable and enable it

_Option 2: Docker Hub Integration using the CLI_

```console
docker scout repo enable $ORG/scout-demo-service
```

```console
docker scout repo enable $ORG/scout-demo-service-back
```

### Push images

1. Push an image to Hub

   ```console
   docker push $ORG/scout-demo-service:v1
   ```

2. Browse https://scout.docker.com and see your images (this might take up to a few minutes)

   ![](./ss/pushes_images.png)

### Analyze images

Reproduce the exploratory steps from _Hands-on #1_ on https://scout.docker.com.
Find vulnerabilities, package information and compare your images.

You can build and push the different versions of the images you previously built (with vulnerabilities or with fixes).
Or simply push the fixes you built in the first hands-on exercise.

```console
docker push $ORG/scout-demo-service:v3
```

### Record images to an environment

1. Record the image to an environment (adapt to your registry)

   ```console
   docker scout environment staging registry://$ORG/scout-demo-service:v1
   ```

   This will explicitly record the image that has been pushed to a registry, to the environment `staging`.

   ![](./ss/staging.png)

2. Investigate environment information from the CLI

   ```console
   docker scout environment
   ```

   ```console
   docker scout environment staging
   ```

3. Compare your fixed local image to the one recorded as part of the `staging` environment

   ```console
   docker scout compare local://$ORG/scout-demo-service:v3 --to-env staging
   ```

   ![](./ss/scout-demo-service-compare-staging.png)
4. Browse https://scout.docker.com

   - find images recorded to an environment
   - find vulnerabilities and packages in an environment
   - compare images between versions and/or environments

   ![](./ss/scout-demo-service-ui-compare.png)

5. Record `v3` to the staging environment

   ```console
   docker scout environment staging registry://$ORG/scout-demo-service:v3
   ```

6. Verify you have no vulnerabilities in `staging` at https://scout.docker.com &gt;
   Vulnerabilities and select the staging environment

7. Browse available integrations on https://scout.docker.com &gt;
   Settings (⚙️ icon) &gt; Integrations
