load('ext://helm_remote', 'helm_remote')

helm_remote('redis',
            repo_name='redis',
            repo_url='https://charts.bitnami.com/bitnami',
            namespace='ndsquared',
            set=[
              'global.redis.password=supersecurepassword',
            ],
)

helm_remote('sealed-secrets',
            release_name='core',
            repo_name='sealed-secrets',
            repo_url='https://bitnami-labs.github.io/sealed-secrets',
            namespace='core',
            create_namespace=True,
            set=[
              'secretName=sealed-secrets-key-local',
            ],
            version='2.0.0',
)

docker_build(
  'portfolio-app',
  './apps/portfolio/',
  entrypoint="npm run tilt",
  live_update=[
    sync('./apps/portfolio/', '/app')
  ],
)


docker_build(
  'api-app',
  './apps/api/',
  entrypoint="uvicorn app.main:app --host=0.0.0.0 --reload",
  live_update=[
    sync('./apps/api/', '/usr/src')
  ],
)

docker_build(
  'api-app-worker',
  './apps/api/',
  entrypoint="dramatiq -v --watch . app.main:broker",
  live_update=[
    sync('./apps/api/', '/usr/src')
  ],
)

k8s_yaml('helm/sealed-secrets-key-local.yaml')
k8s_yaml('helm/api-secret-local.yaml')

k8s_yaml(
  helm(
    "helm/portfolio/",
    name="dev",
    namespace="ndsquared",
    values=[
      "helm/values.common.local.yaml",
      "helm/values.portfolio.local.yaml",
    ],
  )
)

# Filter out production sealed secrets
sealedsecret_yaml, rest = filter_yaml(
  helm(
    "helm/api/",
    name="dev",
    namespace="ndsquared",
    values=[
      "helm/values.common.local.yaml",
      "helm/values.api.local.yaml",
    ],
  ),
  kind='SealedSecret'
)

k8s_yaml(rest)

k8s_resource(
  "redis-master",
  labels=["redis"],
  objects=[
    'redis:serviceaccount',
    'redis-configuration:configmap',
    'redis-health:configmap',
    'redis-scripts:configmap',
    'redis:secret',
  ],
)

k8s_resource(
  "redis-replicas",
  labels=["redis"],
)

k8s_resource(
  "core-sealed-secrets",
  labels=["sealed-secrets"],
  objects=[
    "core:namespace",
    "sealedsecrets.bitnami.com:customresourcedefinition:default",
    "sealedsecrets.bitnami.com:customresourcedefinition:core",
    "core-sealed-secrets:serviceaccount",
    "core-sealed-secrets-key-admin:role",
    "core-sealed-secrets-service-proxier:role",
    "secrets-unsealer:clusterrole",
    "core-sealed-secrets-key-admin:rolebinding",
    "core-sealed-secrets-service-proxier:rolebinding",
    "core-sealed-secrets:clusterrolebinding",
    "sealed-secrets-key-local:secret",
  ],
)

k8s_resource(
  "dev-portfolio",
  links=["http://localhost:8000/"],
  labels=["portfolio"],
  objects=[
    "dev-portfolio:ingress",
  ],
)

k8s_resource(
  "dev-api",
  links=["http://api.localhost:8000/"],
  labels=["api"],
  objects=[
    "dev-api:ingress",
    "dev-api-secret-local:sealedsecret",
  ],
)

k8s_resource(
  "dev-api-worker",
  labels=["api"],
)

# docker_build(
#   'wedding-app',
#   './apps/wedding/',
#   entrypoint="npm run tilt",
#   live_update=[
#     sync('./apps/wedding/', '/app')
#   ],
# )

# k8s_yaml(
#   helm(
#     "helm/wedding/",
#     name="dev",
#     namespace="ndsquared",
#     values=[
#       "helm/values.common.local.yaml",
#       "helm/values.wedding.local.yaml",
#     ],
#   )
# )

# k8s_resource(
#   "dev-wedding",
#   links=["http://wedding.localhost:8000/"],
#   labels=["wedding"],
#   objects=[
#     "dev-wedding:ingress",
#   ],
# )
