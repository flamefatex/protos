FROM swaggerapi/swagger-ui:latest

COPY swagger_mixin /usr/share/nginx/html/swagger

ENV SWAGGER_JSON /usr/share/nginx/html/swagger/traffic-control-api.swagger.json
ENV URLS "[ { url: \"flamefate/example-api/swagger.json\", name: \"example-api\" } ]"
