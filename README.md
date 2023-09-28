# Secure Development With Docker &mdash; DockerCon 2023 workshop

## Setup the workshop platform

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
