---
title: "kf doctor"
slug: kf-doctor
url: /docs/general-info/kf-cli/commands/kf-doctor/
---
## kf doctor

Doctor runs validation tests against one or more components

### Synopsis

Doctor runs tests one or more components to validate them.

If no arguments are supplied, then all tests are run.
If one or more arguments are suplied then only those components are run.

Possible components are: buildpacks, cluster

```
kf doctor [COMPONENT...] [flags]
```

### Examples

```
  kf doctor cluster
```

### Options

```
  -h, --help   help for doctor
```

### Options inherited from parent commands

```
      --config string       config file (default is $HOME/.kf)
      --kubeconfig string   kubectl config file (default is $HOME/.kube/config)
      --namespace string    kubernetes namespace
```

### SEE ALSO

* [kf](/docs/general-info/kf-cli/commands/kf/)	 - kf is like cf for Knative
