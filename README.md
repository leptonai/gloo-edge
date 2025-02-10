Lepton:

Release:
```
IMAGE_REPO="public.ecr.aws/n0k2r8k8/gloo" VERSION="0.0.1-alpha.1" GOARCH=amd64 make gloo-docker
aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws/leptonai
docker push public.ecr.aws/n0k2r8k8/gloo/gloo:0.0.1-alpha.1
```
