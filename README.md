# ARB

## Archlinux r? build
It made sense at 7AM.

## The API
Basically:
```
/build/launch (POST) -- launches the build of a package
/build/complete (POST) -- marks the build as finished
/build/addURL (POST) -- adds URL to the DB (linked via the UUID)
/build/getURL/{UUID} -- get the urls of a specified build
/build/check/{UUID} -- get status info on a specific build
```

## Configuration
Configuration is done via [env viar](https://en.wikipedia.org/wiki/Environment_variable):
* MAIN_HOST: IP/address of your main host where this API runs
* S3_HOST: IP/address of you S3/minio cluster
* S3_REGION: region of the minio/s3 cluster (default for minio: us-east-1)
* S3_KEY: access key for S3/minio
* S3_SECRET: secret key fo S3/minio

