# Hello OpenShift

A minimal Go program to demonstrate debugging capabilities of Go Toolset UBI.

## Deploy

Create a new project.
```
oc new-project hello-openshift
```

Deploy the application on top of ubi9/go-toolset.
```
oc new-app --image registry.access.redhat.com/ubi9/go-toolset~https://github.com/dbenoit17/hello-openshift
```

Create a service on port 8080.
```
oc expose deployment/hello-openshift --port=8080
```

Create a route to the service.
```
oc expose svc/hello-openshift
```
