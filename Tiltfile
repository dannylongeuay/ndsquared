docker_build(
  'portfolio-app',
  './apps/portfolio/',
  entrypoint="npm run tilt",
  live_update=[
    sync('./apps/portfolio/', '/app')
  ],
)

docker_build(
  'wedding-app',
  './apps/wedding/',
  entrypoint="npm run tilt",
  live_update=[
    sync('./apps/wedding/', '/app')
  ],
)

docker_build(
  'api-app',
  './apps/api/',
  entrypoint="uvicorn main:app --host=0.0.0.0 --reload",
  live_update=[
    sync('./apps/api/', '/usr/src')
  ],
)

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

k8s_yaml(
  helm(
    "helm/wedding/",
    name="dev",
    namespace="ndsquared",
    values=[
      "helm/values.common.local.yaml",
      "helm/values.wedding.local.yaml",
    ],
  )
)

k8s_yaml(
  helm(
    "helm/api/",
    name="dev",
    namespace="ndsquared",
    values=[
      "helm/values.common.local.yaml",
      "helm/values.api.local.yaml",
    ],
  )
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
  "dev-wedding",
  links=["http://wedding.localhost:8000/"],
  labels=["wedding"],
  objects=[
    "dev-wedding:ingress",
  ],
)

k8s_resource(
  "dev-api",
  links=["http://api.localhost:8000/"],
  labels=["api"],
  objects=[
    "dev-api:ingress",
  ],
)
