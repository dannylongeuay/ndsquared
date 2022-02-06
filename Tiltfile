load('ext://helm_remote', 'helm_remote')

helm_remote('redis',
            repo_name='redis',
            repo_url='https://charts.bitnami.com/bitnami',
            namespace='local',
            set=[
              'global.redis.password=supersecurepassword',
            ],
)

helm_remote('external-secrets',
            release_name='core',
            repo_name='external-secrets',
            repo_url='https://charts.external-secrets.io',
            namespace='core',
            create_namespace=True,
            version='0.4.1',
)

docker_build(
  'portfolio-app',
  './apps/portfolio/',
  entrypoint="VITE_SERVER_HMR_PORT=8000 npm run tilt",
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

k8s_yaml(
  [
    'gitlab-secret.yaml',
  ]
)

k8s_yaml(
  helm(
    "helm/portfolio/",
    name="local",
    namespace="local",
    values=[
      "helm/values.common.local.yaml",
      "helm/values.portfolio.local.yaml",
    ],
  )
)

k8s_yaml(
  helm(
    "helm/api/",
    name="local",
    namespace="local",
    values=[
      "helm/values.common.local.yaml",
      "helm/values.api.local.yaml",
    ],
  )
)

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
  "core-external-secrets",
  labels=["secrets"],
  objects=[
    "core:namespace",
    "clustersecretstores.external-secrets.io:customresourcedefinition",
    "externalsecrets.external-secrets.io:customresourcedefinition",
    "secretstores.external-secrets.io:customresourcedefinition",
    "core-external-secrets:serviceaccount",
    "core-external-secrets-leaderelection:role",
    "core-external-secrets-controller:clusterrole",
    "core-external-secrets-view:clusterrole",
    "core-external-secrets-edit:clusterrole",
    "core-external-secrets-leaderelection:rolebinding",
    "core-external-secrets-controller:clusterrolebinding",
    "gitlab-secret:secret",
  ]
)

k8s_resource(
  "local-portfolio",
  links=["http://localhost:8000/"],
  labels=["portfolio"],
  objects=[
    "local-portfolio:ingress",
    "local-portfolio:externalsecret",
    "local-portfolio:secretstore",
  ],
  resource_deps=[
    'core-external-secrets',
  ],
)

k8s_resource(
  "local-api",
  links=["http://api.localhost:8000/"],
  labels=["api"],
  objects=[
    "local-api:ingress",
    "local-api:externalsecret",
    "local-api:secretstore",
    "local-api-fake:externalsecret",
    "local-api-fake:secretstore",
  ],
  resource_deps=[
    'core-external-secrets',
  ],
)

k8s_resource(
  "local-api-worker",
  labels=["api"],
  resource_deps=[
    'core-external-secrets',
    'redis-replicas',
  ],
)

# docker_build(
#   'wedding-app',
#   './apps/wedding/',
#   entrypoint="VITE_SERVER_HMR_PORT=8000 npm run tilt",
#   live_update=[
#     sync('./apps/wedding/', '/app')
#   ],
# )

# k8s_yaml(
#   helm(
#     "helm/wedding/",
#     name="local",
#     namespace="ndsquared",
#     values=[
#       "helm/values.common.local.yaml",
#       "helm/values.wedding.local.yaml",
#     ],
#   )
# )

# k8s_resource(
#   "local-wedding",
#   links=["http://wedding.localhost:8000/"],
#   labels=["wedding"],
#   objects=[
#     "local-wedding:ingress",
#   ],
# )
