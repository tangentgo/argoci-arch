``` bash
argocd app create argo-demo \
    --repo https://github.com/tangentgo/argoci-arch.git \
    --path deploy \
    --dest-server https://kubernetes.default.svc \
    --dest-namespace default \
    --sync-policy automated
```