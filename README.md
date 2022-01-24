# NDSquared

Personal portfolio site

www.ndsquared.net

## Secrets

Secrets are managed with [SealedSecrets](https://github.com/bitnami-labs/sealed-secrets).

**New secret creation:**

First create a generic kubernetes secret. Make sure to update the `name` and `namespace` accordingly.

```
make generic-secret
```

Next `base64` encode secrets and copy them into the generated generic secret.

```
echo 'secretvaluehere' | base64 | xclip -selection clipboard
```

Lastly, convert the generic secret into a sealed secret, which is safe to be version controlled.

```
make seal-secret
```

**Update existing secrets:**

This process is the same as new secret creation with the additional step of copying secrets from the generated `sealedsecret.yaml` into an existing sealed secret.

> Ensure the generic secret's `name` and `namespace` are the same as existing sealed secret.
