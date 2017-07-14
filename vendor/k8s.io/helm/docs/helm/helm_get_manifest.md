## helm get manifest

download the manifest for a named release

### Synopsis



This command fetches the generated manifest for a given release.

A manifest is a YAML-encoded representation of the Kubernetes resources that
were generated from this release's chart(s). If a chart is dependent on other
charts, those resources will also be included in the manifest.


```
helm get manifest [flags] RELEASE_NAME
```

### Options

```
      --revision int32   get the named release with revision
```

### Options inherited from parent commands

```
      --debug                     enable verbose output
      --home string               location of your Helm config. Overrides $HELM_HOME (default "~/.helm")
      --host string               address of tiller. Overrides $HELM_HOST
      --kube-context string       name of the kubeconfig context to use
      --tiller-namespace string   namespace of tiller (default "kube-system")
```

### SEE ALSO
* [helm get](helm_get.md)	 - download a named release

###### Auto generated by spf13/cobra on 17-May-2017