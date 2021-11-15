load("ext://helm_remote", "helm_remote")

helm_remote(
  "ingress-nginx",
  repo_name="ingress-nginx",
  repo_url="https://kubernetes.github.io/ingress-nginx",
  release_name="dev",
  namespace="ingress-nginx",
  version="4.0.6",
)

k8s_yaml(
  helm(
    "helm/portfolio/",
    name="dev",
    namespace="ndsquared",
  )
)

k8s_resource(
  "dev-ingress-nginx-controller",
  new_name="dev-ingress-nginx",
  port_forwards=["8000"], 
  objects=[
    "nginx",
    "dev-ingress-nginx:clusterrole",
    "dev-ingress-nginx:clusterrolebinding",
    "dev-ingress-nginx:role",
    "dev-ingress-nginx:rolebinding",
    "dev-ingress-nginx:serviceaccount",
    "dev-ingress-nginx-admission:clusterrole",
    "dev-ingress-nginx-admission:clusterrolebinding",
    "dev-ingress-nginx-admission:role",
    "dev-ingress-nginx-admission:rolebinding",
    "dev-ingress-nginx-admission:serviceaccount",
    "dev-ingress-nginx-controller:configmap",
    "dev-ingress-nginx-admission:validatingwebhookconfiguration",
  ],
  labels=["ingress"],
)

k8s_resource(
  "dev-ingress-nginx-admission-create",
  labels=["jobs"],
)

k8s_resource(
  "dev-ingress-nginx-admission-patch",
  labels=["jobs"],
)

k8s_resource(
  "dev-portfolio",
  links=["http://localhost:8000/"],
  labels=["portfolio"],
  objects=[
    "dev-portfolio:ingress",
  ],
  resource_deps=[
    "dev-ingress-nginx",
  ]
)

docker_build(
  'portfolio-app',
  './apps/portfolio/',
  dockerfile='./apps/portfolio/Dockerfile',
  entrypoint="npm run tilt",
  live_update=[
    sync('./apps/portfolio/', '/app')
  ],
)
