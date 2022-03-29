# Pokemon REST API

The app can return basic information about a given Pokémon e.g. name, description, habitat etc. If the Pokémon is legendary or its habitat is a cave, the description can get Yoda translated. 
All others can get Shakespeare translated. If the Pokémon doesn't have English description, or we have reached the maximum API calls to `funtranslations.com` then the standard description will be returned


## Running app locally with Docker

Note: make sure Docker is installed and running and from the root directory run the following commands:

- `docker build -t pokemon .` This will build the image and run all tests
- `docker run -p 80:80 -d <docker_image_id>` Starts a container in detached mode

Example request basic information:

```
curl -X GET "localhost/api/v2/pokemon/mewtwo"
```

Response:
```
{
  "name": "mewtwo",
  "desc": "Psychic power has augmented its muscles. It has a grip strength of one ton and can sprint a hundred meters in two seconds flat!",
  "habitat": "rare",
  "isLegendary": true,
  "apiLimit": false
}
```

Example request Pokemon translated description:

```
curl -X GET "localhost/api/v2/pokemon/translated/abra"
```

Response:

```
{
  "name": "abra",
  "desc": "Abra can teleport in its catch but a wink. Apparently the moo deeply abra sleeps,  the farther its teleportations wend.",
  "habitat": "urban",
  "isLegendary": false,
  "apiLimit": false
}
```

## Running app locally with Minikube

Note: make sure Minikube is installed and running and from the root directory run the following commands:

- `kubectl create -f deploy-app.yaml` The YAML will create a deployment with a single replica and a service with NodePort. The image will get pulled from the Docker Hub
- `minikube service pokemon-api-service --url` Starts a tunnel and returns localhost ip with an ephemeral port
- `curl -X GET http://127.0.0.1:<port>/api/v2/pokemon/[translated]/<pokemon_name>` Returns information as documented above

## Exceeding the API limit

As part of the response body the field `apiLimit` is set to true if the API calls to `funtranslations.com` are exceeded and will be reset in 56 min

## Running in production

- Because of the API limitations, a paid account will need to be created with `funtranslations.com` and the code extended to support the access token
- Instead of running locally the app will have to be deployed to a cloud provider e.g. AWS or GCP for high availability and scalability
- The service NodePort would need to be replaced with an Ingress controller, so we don't expose the service to the outside world but instead an endpoint which routes traffic to backend services

test3

