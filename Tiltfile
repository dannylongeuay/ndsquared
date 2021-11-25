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
