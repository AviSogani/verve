- Mutex lock has been used. This ensures that multiple requests at the same time will not create
  consistency issues.

- For extension #2, I would have used Redis as a common data store for the requestIdMap. This would ensure that multiple
  instances of the app behind a load balancer would be able to access the same requestIdMap. Redis can be seen as single
  threaded which will ensure the mutex aspect. Additionally, Redis' master slave model and sentinel can be used to
  achieve eventual consistency and full availability. [Architecture](architecture.png) showcases the client-server
  models with Redis that can be used for the workflow.

- When deploying, container orchestration can be done using Kubernetes for scaling the service behind a load balancer.
  The load balancer can be set up using Google Cloud Console and GKE can be used.

- Future Scope: A CDN or caching layer can be added for requests with repeated Ids. The caching can be made specific to
  query parameters. If ‘id’ and ‘endpoint’ are repeated, then the same response can be used.

- For extensions, classes can be implemented for alternative workflow. Or features can be used which can be mapped to
  pick the workflow.

- Docker image is pushed to https://github.com/users/AviSogani/packages?repo_name=verve
