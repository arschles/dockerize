# dockerize

**Note**: This repository is mostly empty. It's a big TODO and is mostly for holding my ideas.

The point of this project is to build a server that can download a repository from GitHub and then do the following:

1. Look for a `.dockerize.yml` file, which will have:
  1. Locations of files to download from Amazon S3, and where to store each one locally
  2. A glob of files to send to the Docker daemon
  3. The name(s) of the Docker image to build from the glob. These will also all be pushed.
2. After it's downloaded everything, it sends all of the files in the glob (from 1.2) to the Docker daemon, then builds and pushes all of the images listed in 1.3
