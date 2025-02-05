# Gatekeeper Plugin

## English

Gatekeeper is a Traefik middleware plugin that verifies the `Authorization` header by sending a GET request to an external URL. If the external URL returns a `200 OK` response, the request is allowed to proceed; otherwise, the request is blocked and an error response is returned.

### Features
- Checks for the presence of an `Authorization` header.
- Sends a GET request with the header to an external URL.
- If the external service returns `200 OK`, the request proceeds.
- If the external service returns any other response, the request is blocked.

### Installation

1. Clone this repository to your local machine.
2. Add the plugin to the Traefik configuration:

   ```yaml
   experimental:
     plugins:
       gatekeeper:
         moduleName: github.com/fatiharge/gatekeeper
         version: v0.1.0
   ```

3. Configure the plugin in the dynamic configuration:

   ```yaml
   http:
     middlewares:
       gatekeeper:
         plugin:
           gatekeeper:
             externalURL: "https://auth.example.com/validate"
             authHeader: "Authorization"
   ```

4. Reload Traefik to apply the changes.

### Configuration

- `externalURL`: The URL where the GET request will be sent for validation.
- `authHeader`: The name of the header to check (default is `Authorization`).