# Secure Development With Docker - Dockercon 23 workshop

## Setup the Workshop Platform

### Requirements

- docker
- docker compose
- docker scout cli
- git
- Docker Hub account `$ORG`

### Prerequisites

Conventions:

- `$ORG` is the name of the Docker Hub account you will use.
It can be a personal or a team one. Better if you have full ownership on it.
- `$WS_PATH` is the path on your local machine where you will clone
the workshop Git repository.

1. Clone the repository on your local machine, to `$WS_PATH`
2. Copy `.env.config` file to `.env` (no need to change now).

## Hands-on #1: Remediating Vulnerabilities

### Workshop Images

1. Build demo images

   ```
   $ docker compose --profile images build
   ```

   > This command will build two images we will explore.
   > To know more about how they are built look at
   > [`./docker-compose.yml`](./docker-compose.yml) and
   > [`./backend/Dockerfile`](./backend/Dockerfile) that
   > contains the build definitions.

   > In case of network issues, you can also build the following
   > image that is prebuilt and don't need extra dependencies.

   ```
   $ docker compose --profile low_network build
   ```

2. Open Docker Desktop and select the image `dc23/scout-demo-service:v1`

    ![](./ss/layer_view_scout-demo-service_v1.png)
3. Take the time to explore the different information displayed on this page

   - Image hierarchy, layers and the Images tab
   - Vulnerabilities
   - Packages

### Base Image Fix (Docker Desktop)

1. Select your base image and explore vulnerabilities specific to the base image

   ![](./ss/select-base-image.png)

2. Select _Change base image_ and pick the current image as it's defined in [`./frontend/Dockerfile`](./frontend/Dockerfile)

   ![](./ss/change-base-image.png)

3. Apply the recommendation to the [`Dockerfile`](./frontend/Dockerfile)
4. (optional) Update the tag to `v2` in [`docker-compose.yml`](./docker-compose.yml)
5. Rebuild the image

   ```
   $ docker compose --profile scout-demo-service build
   ```
6. Open the image inside Desktop and see the impact of your change

   ![](./ss/layer_view_scout-demo-service_v2.png)

### Base Image Fix (Docker Scout CLI)

1. Run `docker scout cves` command against the image you just built:

   ```
   $ docker scout cves dc23/scout-demo-service:v1
   ```

   ![](./ss/scout-demo-service-v1-cves.png)
2. (optional) Explore the filters options from `docker scout cves` command
3. (optional) See vulnerable packages only

   ```
   $ docker scout cves --format only-packages --only-vuln-packages dc23/scout-demo-service:v1
   ```

   ![](./ss/scout-demo-service-v1-only-vuln-packages.png)

4. Explore base image recommendations using `docker scout recommendations`

   -> specify the tag upfront, like we selected it in Docker Desktop

   ```
   $ docker scout recommendations --tag 3.14 dc23/scout-demo-service:v1
   ```

   ![](./ss/scout-demo-service-v1-recommendations.png)

5. Apply the recommendation to the [`Dockerfile`](./frontend/Dockerfile)
6. (optional) Update the tag to `v2` in [`docker-compose.yml`](./docker-compose.yml)
7. Rebuild the image

   ```
   $ docker compose --profile scout-demo-service build
   ```
8. Run `docker scout cves` command against the image you just built and see the changes:

   ```
   $ docker scout cves dc23/scout-demo-service:v2
   ```

   ![](./ss/scout-demo-service-v2-cves.png)
9. Compare the two images to see the differences:

   ```
   $ docker scout compare dc23/scout-demo-service:v2 --to dc23/scout-demo-service:v1
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

   ```
   $ docker scout quickview dc23/scout-demo-service:v1
   ```

### Extra Image

Repeat the above steps for the `dc23/scout-demo-service-back:v1` image (or any other image you have).
