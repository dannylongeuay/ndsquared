load("ext://helm_remote", "helm_remote")

helm_remote("traefik",
  repo_name="traefik",
            repo_url="https://helm.traefik.io/traefik",
  version="10.6.0",
  set=[
    "service.type=ClusterIP"
  ]
)

k8s_resource("traefik", port_forwards=["8000", "9000"])
