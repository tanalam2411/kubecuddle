# kubecuddle
kubecuddle k8s operator


### Installing the Operator SDK CLI

```bash
$ GO111MODULE=on go get -d github.com/operator-framework/operator-sdk
go: finding github.com/operator-framework/operator-sdk v0.12.0
go: downloading github.com/operator-framework/operator-sdk v0.12.0
go: extracting github.com/operator-framework/operator-sdk v0.12.0
go get: github.com/operator-framework/operator-sdk@v0.12.0 requires
	k8s.io/api@v0.0.0: reading k8s.io/api/go.mod at revision v0.0.0: unknown revision v0.0.0

$ cd $GOPATH/src/github.com/operator-framework/operator-sdk

~/go/src/github.com/operator-framework/operator-sdk$ git checkout master

~/go/src/github.com/operator-framework/operator-sdk$ make tidy

~/go/src/github.com/operator-framework/operator-sdk$ make install
```

Try `$ operator-sdk`:
```bash
$ operator-sdk 
An SDK for building operators with ease

Usage:
  operator-sdk [command]

Available Commands:
  add         Adds a controller or resource to the project
  alpha       Run an alpha subcommand
  build       Compiles code and builds artifacts
  completion  Generators for shell completions
  generate    Invokes specific generator
  help        Help about any command
  migrate     Adds source code to an operator
  new         Creates a new operator application
  olm-catalog Invokes a olm-catalog command
  print-deps  Print Golang packages and versions required to run the operator
  run         Runs a generic operator
  scorecard   Run scorecard tests
  test        Tests the operator
  up          Launches the operator
  version     Prints the version of operator-sdk

Flags:
  -h, --help      help for operator-sdk
      --verbose   Enable verbose logging

Use "operator-sdk [command] --help" for more information about a command.

```
---

1. Creating operator project by name `tgik-operator`:
```bash
$ cd ~/tanveer/k8s/kubecuddle/
~/tanveer/k8s/kubecuddle$ operator-sdk new tgik-operator --repo github.com/tanalam2411/kubecuddle
INFO[0000] Creating new Go operator 'tgik-operator'.    
INFO[0000] Created go.mod                               
INFO[0000] Created tools.go                             
INFO[0000] Created cmd/manager/main.go                  
INFO[0000] Created build/Dockerfile                     
INFO[0000] Created build/bin/entrypoint                 
INFO[0000] Created build/bin/user_setup                 
INFO[0000] Created deploy/service_account.yaml          
INFO[0000] Created deploy/role.yaml                     
INFO[0000] Created deploy/role_binding.yaml             
INFO[0000] Created deploy/operator.yaml                 
INFO[0000] Created pkg/apis/apis.go                     
INFO[0000] Created pkg/controller/controller.go         
INFO[0000] Created version/version.go                   
INFO[0000] Created .gitignore                           
INFO[0000] Validating project                           
go: finding github.com/operator-framework/operator-sdk master
go: downloading github.com/operator-framework/operator-sdk v0.12.1-0.20191127152849-96dd1d2d5de1
go: extracting github.com/operator-framework/operator-sdk v0.12.1-0.20191127152849-96dd1d2d5de1
INFO[0049] Project validation successful.               
INFO[0049] Project creation complete.   


~/tanveer/k8s/kubecuddle$ tree -I vendor
.
├── LICENSE
├── README.md
└── tgik-operator
    ├── build
    │   ├── bin
    │   │   ├── entrypoint
    │   │   └── user_setup
    │   └── Dockerfile
    ├── cmd
    │   └── manager
    │       └── main.go
    ├── deploy
    │   ├── operator.yaml
    │   ├── role_binding.yaml
    │   ├── role.yaml
    │   └── service_account.yaml
    ├── go.mod
    ├── go.sum
    ├── pkg
    │   ├── apis
    │   │   └── apis.go
    │   └── controller
    │       └── controller.go
    ├── tools.go
    └── version
        └── version.go

10 directories, 16 files

       
```

2. Add a new API for the custom resource `Tgik`
```bash
~/tanveer/k8s/kubecuddle$ cd tgik-operator/
~/tanveer/k8s/kubecuddle/tgik-operator$ operator-sdk add api --api-version=tan.github.com/v1alpha1 --kind=Tgik
INFO[0000] Generating api version tan.github.com/v1alpha1 for kind Tgik. 
INFO[0000] Created pkg/apis/tan/group.go                
INFO[0001] Created pkg/apis/tan/v1alpha1/tgik_types.go  
INFO[0001] Created pkg/apis/addtoscheme_tan_v1alpha1.go 
INFO[0001] Created pkg/apis/tan/v1alpha1/register.go    
INFO[0001] Created pkg/apis/tan/v1alpha1/doc.go         
INFO[0001] Created deploy/crds/tan.github.com_v1alpha1_tgik_cr.yaml 
INFO[0004] Created deploy/crds/tan.github.com_tgiks_crd.yaml 
INFO[0004] Running deepcopy code-generation for Custom Resource group versions: [tan:[v1alpha1], ] 
INFO[0012] Code-generation complete.                    
INFO[0012] Running OpenAPI code-generation for Custom Resource group versions: [tan:[v1alpha1], ] 
INFO[0020] Created deploy/crds/tan.github.com_tgiks_crd.yaml 
INFO[0020] Code-generation complete.                    
INFO[0020] API generation complete.   

~/tanveer/k8s/kubecuddle$ tree -I vendor
.
├── LICENSE
├── README.md
└── tgik-operator
    ├── build
    │   ├── bin
    │   │   ├── entrypoint
    │   │   └── user_setup
    │   └── Dockerfile
    ├── cmd
    │   └── manager
    │       └── main.go
    ├── deploy
    │   ├── crds
    │   │   ├── tan.github.com_tgiks_crd.yaml
    │   │   └── tan.github.com_v1alpha1_tgik_cr.yaml
    │   ├── operator.yaml
    │   ├── role_binding.yaml
    │   ├── role.yaml
    │   └── service_account.yaml
    ├── go.mod
    ├── go.sum
    ├── pkg
    │   ├── apis
    │   │   ├── addtoscheme_tan_v1alpha1.go
    │   │   ├── apis.go
    │   │   └── tan
    │   │       ├── group.go
    │   │       └── v1alpha1
    │   │           ├── doc.go
    │   │           ├── register.go
    │   │           ├── tgik_types.go
    │   │           ├── zz_generated.deepcopy.go
    │   │           └── zz_generated.openapi.go
    │   └── controller
    │       └── controller.go
    ├── tools.go
    └── version
        └── version.go

13 directories, 25 files
                 
```

Created following new files:
```bash
    ├── deploy
    │   ├── crds
    │   │   ├── tan.github.com_tgiks_crd.yaml
    │   │   └── tan.github.com_v1alpha1_tgik_cr.yaml
	
    ├── pkg
    │   ├── apis
    │   │   ├── addtoscheme_tan_v1alpha1.go

    │   │   └── tan
    │   │       ├── group.go
    │   │       └── v1alpha1
    │   │           ├── doc.go
    │   │           ├── register.go
    │   │           ├── tgik_types.go
    │   │           ├── zz_generated.deepcopy.go
    │   │           └── zz_generated.openapi.go
```

And updated :
```bash

    ├── deploy
    │   ├── role.yaml
```
Content added:
```yaml
- apiGroups:
  - tan.github.com
  resources:
  - '*'
  verbs:
  - '*'
```

3. Add a new controller that watches for `Tgik`:

```bash
~/tanveer/k8s/kubecuddle/tgik-operator$ operator-sdk add controller --api-version=tan.github.com/v1alpha1 --kind=Tgik
INFO[0000] Generating controller version tan.github.com/v1alpha1 for kind Tgik. 
INFO[0000] Created pkg/controller/tgik/tgik_controller.go 
INFO[0000] Created pkg/controller/add_tgik.go           
INFO[0000] Controller generation complete.  
```

4. Building and pushing the `tgik-operator` image to a public registry:
```bash
~/tanveer/k8s/kubecuddle/tgik-operator$ operator-sdk build on2411/tgik-operator:latest
INFO[0057] Building OCI image on2411/tgik-operator:latest 
Sending build context to Docker daemon     73MB
Step 1/7 : FROM registry.access.redhat.com/ubi8/ubi-minimal:latest
latest: Pulling from ubi8/ubi-minimal
645c2831c08a: Pull complete 
5e98065763a5: Pull complete 
Digest: sha256:32fb8bae553bfba2891f535fa9238f79aafefb7eff603789ba8920f505654607
Status: Downloaded newer image for registry.access.redhat.com/ubi8/ubi-minimal:latest
 ---> 469119976c56
Step 2/7 : ENV OPERATOR=/usr/local/bin/tgik-operator     USER_UID=1001     USER_NAME=tgik-operator
 ---> Running in e605ee299824
Removing intermediate container e605ee299824
 ---> 4c2cc47c3b4d
Step 3/7 : COPY build/_output/bin/tgik-operator ${OPERATOR}
 ---> f3a1cbafe378
Step 4/7 : COPY build/bin /usr/local/bin
 ---> c64cf4a61014
Step 5/7 : RUN  /usr/local/bin/user_setup
 ---> Running in 16cfd85c4293
+ mkdir -p /root
+ chown 1001:0 /root
+ chmod ug+rwx /root
+ chmod g+rw /etc/passwd
+ rm /usr/local/bin/user_setup
Removing intermediate container 16cfd85c4293
 ---> 59e279d56d8e
Step 6/7 : ENTRYPOINT ["/usr/local/bin/entrypoint"]
 ---> Running in ea4dcdc160de
Removing intermediate container ea4dcdc160de
 ---> 5965c19016db
Step 7/7 : USER ${USER_UID}
 ---> Running in 3c97c7726152
Removing intermediate container 3c97c7726152
 ---> 429f383c0636
Successfully built 429f383c0636
Successfully tagged on2411/tgik-operator:latest
INFO[0179] Operator build complete.                     


~/tanveer/k8s/kubecuddle/tgik-operator$ docker images | grep tgik
on2411/tgik-operator                          latest              429f383c0636        44 minutes ago      148MB

~/tanveer/k8s/kubecuddle/tgik-operator$ docker push on2411/tgik-operator:latest
The push refers to repository [docker.io/on2411/tgik-operator]
efa17ecf0121: Pushed 
f794914e5d9f: Pushed 
2902c328c4c0: Pushed 
a066f3d73913: Pushed 
26b543be03e2: Pushed 
latest: digest: sha256:6697eca97263c77c284bb9f91c501c5b9ae52ab71e85aaeb394d8086b49674f7 size: 1363
```

Image - https://hub.docker.com/repository/docker/on2411/tgik-operator

5. Update the operator manifest to use the built image name:
```bash
~/tanveer/k8s/kubecuddle/tgik-operator$ sed -i 's|REPLACE_IMAGE|on2411/tgik-operator:latest|g' deploy/operator.yaml
```

---

We're running k8's cluster using [kind](https://github.com/kubernetes-sigs/kind)
```bash
~/tanveer/k8s/kubecuddle/tgik-operator$ kubectl get nodes
NAME                 STATUS   ROLES    AGE    VERSION
kind-control-plane   Ready    master   2d8h   v1.15.3
```

6. Setup Service Account
```bash
~/tanveer/k8s/kubecuddle/tgik-operator$ kubectl create -f deploy/service_account.yaml 
serviceaccount/tgik-operator created

~/tanveer/k8s/kubecuddle/tgik-operator$ kubectl get serviceaccounts
NAME            SECRETS   AGE
default         1         2d8h
tgik-operator   1         5m55s

~/tanveer/k8s/kubecuddle/tgik-operator$ kubectl describe serviceaccounts tgik-operator
Name:                tgik-operator
Namespace:           default
Labels:              <none>
Annotations:         <none>
Image pull secrets:  <none>
Mountable secrets:   tgik-operator-token-77lgf
Tokens:              tgik-operator-token-77lgf
Events:              <none>

```

7. Setup RBAC
```bash
~/tanveer/k8s/kubecuddle/tgik-operator$ kubectl create -f deploy/role.yaml 
role.rbac.authorization.k8s.io/tgik-operator created

~/tanveer/k8s/kubecuddle/tgik-operator$ kubectl get role -A
NAMESPACE     NAME                                             AGE
default       tgik-operator                                    8s
kube-public   kubeadm:bootstrap-signer-clusterinfo             2d8h
kube-public   system:controller:bootstrap-signer               2d8h
kube-system   extension-apiserver-authentication-reader        2d8h
kube-system   kube-proxy                                       2d8h
kube-system   kubeadm:kubelet-config-1.15                      2d8h
kube-system   kubeadm:nodes-kubeadm-config                     2d8h
kube-system   system::leader-locking-kube-controller-manager   2d8h
kube-system   system::leader-locking-kube-scheduler            2d8h
kube-system   system:controller:bootstrap-signer               2d8h
kube-system   system:controller:cloud-provider                 2d8h
kube-system   system:controller:token-cleaner                  2d8h

~/tanveer/k8s/kubecuddle/tgik-operator$ kubectl describe role tgik-operator
Name:         tgik-operator
Labels:       <none>
Annotations:  <none>
PolicyRule:
  Resources                              Non-Resource URLs  Resource Names   Verbs
  ---------                              -----------------  --------------   -----
  pods                                   []                 []               [* get]
  deployments.apps                       []                 []               [* get]
  replicasets.apps                       []                 []               [* get]
  configmaps                             []                 []               [*]
  endpoints                              []                 []               [*]
  events                                 []                 []               [*]
  persistentvolumeclaims                 []                 []               [*]
  secrets                                []                 []               [*]
  services/finalizers                    []                 []               [*]
  services                               []                 []               [*]
  daemonsets.apps                        []                 []               [*]
  statefulsets.apps                      []                 []               [*]
  *.tan.github.com                       []                 []               [*]
  servicemonitors.monitoring.coreos.com  []                 []               [get create]
  deployments.apps/finalizers            []                 [tgik-operator]  [update]

```

Creating Role Bindings:
```bash
~/tanveer/k8s/kubecuddle/tgik-operator$ kubectl create -f deploy/role_binding.yaml 
rolebinding.rbac.authorization.k8s.io/tgik-operator created

~/tanveer/k8s/kubecuddle/tgik-operator$ kubectl get rolebindings -A
NAMESPACE     NAME                                                AGE
default       tgik-operator                                       13s
kube-public   kubeadm:bootstrap-signer-clusterinfo                2d8h
kube-public   system:controller:bootstrap-signer                  2d8h
kube-system   kube-proxy                                          2d8h
kube-system   kubeadm:kubelet-config-1.15                         2d8h
kube-system   kubeadm:nodes-kubeadm-config                        2d8h
kube-system   system::extension-apiserver-authentication-reader   2d8h
kube-system   system::leader-locking-kube-controller-manager      2d8h
kube-system   system::leader-locking-kube-scheduler               2d8h
kube-system   system:controller:bootstrap-signer                  2d8h
kube-system   system:controller:cloud-provider                    2d8h
kube-system   system:controller:token-cleaner                     2d8h

~/tanveer/k8s/kubecuddle/tgik-operator$ kubectl describe rolebindings tgik-operator
Name:         tgik-operator
Labels:       <none>
Annotations:  <none>
Role:
  Kind:  Role
  Name:  tgik-operator
Subjects:
  Kind            Name           Namespace
  ----            ----           ---------
  ServiceAccount  tgik-operator  

```

8. Setting up the `CRD`:
```bash
~/tanveer/k8s/kubecuddle/tgik-operator$ kubectl create -f deploy/crds/tan.github.com_tgiks_crd.yaml 
customresourcedefinition.apiextensions.k8s.io/tgiks.tan.github.com created

~/tanveer/k8s/kubecuddle/tgik-operator$ kubectl get crd
NAME                   CREATED AT
tgiks.tan.github.com   2019-11-28T10:36:23Z

~/tanveer/k8s/kubecuddle/tgik-operator$ kubectl describe crd tgiks.tan.github.com
Name:         tgiks.tan.github.com
Namespace:    
Labels:       <none>
Annotations:  <none>
API Version:  apiextensions.k8s.io/v1beta1
Kind:         CustomResourceDefinition
Metadata:
  Creation Timestamp:  2019-11-28T10:36:23Z
  Generation:          1
  Resource Version:    129968
  Self Link:           /apis/apiextensions.k8s.io/v1beta1/customresourcedefinitions/tgiks.tan.github.com
  UID:                 67a4a075-2f1a-4fbf-855e-97f102f820ce
Spec:
  Conversion:
    Strategy:  None
  Group:       tan.github.com
  Names:
    Kind:                   Tgik
    List Kind:              TgikList
    Plural:                 tgiks
    Singular:               tgik
  Preserve Unknown Fields:  true
  Scope:                    Namespaced
  Subresources:
    Status:
  Validation:
    Open APIV 3 Schema:
      Description:  Tgik is the Schema for the tgiks API
      Properties:
        API Version:
          Description:  APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
          Type:         string
        Kind:
          Description:  Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
          Type:         string
        Metadata:
          Type:  object
        Spec:
          Description:  TgikSpec defines the desired state of Tgik
          Type:         object
        Status:
          Description:  TgikStatus defines the observed state of Tgik
          Type:         object
      Type:             object
  Version:              v1alpha1
  Versions:
    Name:     v1alpha1
    Served:   true
    Storage:  true
Status:
  Accepted Names:
    Kind:       Tgik
    List Kind:  TgikList
    Plural:     tgiks
    Singular:   tgik
  Conditions:
    Last Transition Time:  2019-11-28T10:36:23Z
    Message:               no conflicts found
    Reason:                NoConflicts
    Status:                True
    Type:                  NamesAccepted
    Last Transition Time:  <nil>
    Message:               the initial names have been accepted
    Reason:                InitialNamesAccepted
    Status:                True
    Type:                  Established
  Stored Versions:
    v1alpha1
Events:  <none>

```

```bash
~/tanveer/k8s/kubecuddle/tgik-operator$ kubectl api-resources --api-group='tan.github.com' -o wide
NAME    SHORTNAMES   APIGROUP         NAMESPACED   KIND   VERBS
tgiks                tan.github.com   true         Tgik   [delete deletecollection get list patch create update watch]

~/tanveer/k8s/kubecuddle/tgik-operator$ kubectl get tgik
No resources found.

```

9. Deploy the `operator`:
```bash
~/tanveer/k8s/kubecuddle/tgik-operator$ kubectl create -f deploy/operator.yaml 
deployment.apps/tgik-operator created

~/tanveer/k8s/kubecuddle/tgik-operator$ kubectl get deployment -A
NAMESPACE     NAME            READY   UP-TO-DATE   AVAILABLE   AGE
default       tgik-operator   1/1     1            1           24m
kube-system   coredns         2/2     2            2           2d9h

~/tanveer/k8s/kubecuddle/tgik-operator$ kubectl get deployment tgik-operator -o yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  annotations:
    deployment.kubernetes.io/revision: "1"
  creationTimestamp: "2019-11-28T10:59:27Z"
  generation: 1
  name: tgik-operator
  namespace: default
  resourceVersion: "131708"
  selfLink: /apis/extensions/v1beta1/namespaces/default/deployments/tgik-operator
  uid: 35849b33-9856-402b-b247-888115239924
spec:
  progressDeadlineSeconds: 600
  replicas: 1
  revisionHistoryLimit: 10
  selector:
    matchLabels:
      name: tgik-operator
  strategy:
    rollingUpdate:
      maxSurge: 25%
      maxUnavailable: 25%
    type: RollingUpdate
  template:
    metadata:
      creationTimestamp: null
      labels:
        name: tgik-operator
    spec:
      containers:
      - command:
        - tgik-operator
        env:
        - name: WATCH_NAMESPACE
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.namespace
        - name: POD_NAME
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.name
        - name: OPERATOR_NAME
          value: tgik-operator
        image: on2411/tgik-operator:latest
        imagePullPolicy: Always
        name: tgik-operator
        resources: {}
        terminationMessagePath: /dev/termination-log
        terminationMessagePolicy: File
      dnsPolicy: ClusterFirst
      restartPolicy: Always
      schedulerName: default-scheduler
      securityContext: {}
      serviceAccount: tgik-operator
      serviceAccountName: tgik-operator
      terminationGracePeriodSeconds: 30
status:
  availableReplicas: 1
  conditions:
  - lastTransitionTime: "2019-11-28T10:59:58Z"
    lastUpdateTime: "2019-11-28T10:59:58Z"
    message: Deployment has minimum availability.
    reason: MinimumReplicasAvailable
    status: "True"
    type: Available
  - lastTransitionTime: "2019-11-28T10:59:27Z"
    lastUpdateTime: "2019-11-28T10:59:58Z"
    message: ReplicaSet "tgik-operator-6bcbff8b9f" has successfully progressed.
    reason: NewReplicaSetAvailable
    status: "True"
    type: Progressing
  observedGeneration: 1
  readyReplicas: 1
  replicas: 1
  updatedReplicas: 1

```

```bash
~/tanveer/k8s/kubecuddle/tgik-operator$ kubectl get pods
NAME                             READY   STATUS    RESTARTS   AGE
tgik-operator-6bcbff8b9f-cdvm7   1/1     Running   0          27m

~/tanveer/k8s/kubecuddle/tgik-operator$ kubectl logs --tail=100 tgik-operator-6bcbff8b9f-cdvm7
{"level":"info","ts":1574938797.9833722,"logger":"cmd","msg":"Operator Version: 0.0.1"}
{"level":"info","ts":1574938797.983402,"logger":"cmd","msg":"Go Version: go1.13.3"}
{"level":"info","ts":1574938797.9834096,"logger":"cmd","msg":"Go OS/Arch: linux/amd64"}
{"level":"info","ts":1574938797.9834156,"logger":"cmd","msg":"Version of operator-sdk: v0.12.0+git"}
{"level":"info","ts":1574938797.983658,"logger":"leader","msg":"Trying to become the leader."}
{"level":"info","ts":1574938798.266084,"logger":"leader","msg":"No pre-existing lock was found."}
{"level":"info","ts":1574938798.2706623,"logger":"leader","msg":"Became the leader."}
{"level":"info","ts":1574938798.5255582,"logger":"controller-runtime.metrics","msg":"metrics server is starting to listen","addr":"0.0.0.0:8383"}
{"level":"info","ts":1574938798.5257561,"logger":"cmd","msg":"Registering Components."}
{"level":"info","ts":1574938799.19932,"logger":"metrics","msg":"Metrics Service object created","Service.Name":"tgik-operator-metrics","Service.Namespace":"default"}
{"level":"info","ts":1574938799.4515955,"logger":"cmd","msg":"Could not create ServiceMonitor object","error":"no ServiceMonitor registered with the API"}
{"level":"info","ts":1574938799.4516606,"logger":"cmd","msg":"Install prometheus-operator in your cluster to create ServiceMonitor objects","error":"no ServiceMonitor registered with the API"}
{"level":"info","ts":1574938799.45168,"logger":"cmd","msg":"Starting the Cmd."}
{"level":"info","ts":1574938799.4526892,"logger":"controller-runtime.controller","msg":"Starting EventSource","controller":"tgik-controller","source":"kind source: /, Kind="}
{"level":"info","ts":1574938799.4532652,"logger":"controller-runtime.controller","msg":"Starting EventSource","controller":"tgik-controller","source":"kind source: /, Kind="}
{"level":"info","ts":1574938799.4534595,"logger":"controller-runtime.manager","msg":"starting metrics server","path":"/metrics"}
{"level":"info","ts":1574938799.4535701,"logger":"controller-runtime.controller","msg":"Starting Controller","controller":"tgik-controller"}
{"level":"info","ts":1574938799.5538788,"logger":"controller-runtime.controller","msg":"Starting workers","controller":"tgik-controller","worker count":1}

```