# ARB

## Archlinux r? build
It made sense at 7AM.

## The API
Basically:
```
/build/launch/{package} -- launches the build of a package
/build/complete (POST) -- marks the build as finished
/build/addURL (POST) -- adds URL to the DB (linked via the UUID)
/build/getURL/{UUID} -- get the urls of a specified build
/build/check/{UUID} -- get status info on a specific build
```
