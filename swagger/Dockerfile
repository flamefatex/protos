FROM swaggerapi/swagger-ui:latest

COPY swagger_mixin /usr/share/nginx/html/swagger

ENV SWAGGER_JSON /usr/share/nginx/html/swagger/traffic-control-api.swagger.json
ENV URLS "[ { url: \"swagger/traffic-control-api.swagger.json\", name: \"traffic-control-api\" }, {url: \"swagger/ns-agent.swagger.json\", name: \"ns-agent\"}, {url: \"swagger/config.swagger.json\", name: \"config\"}]"
