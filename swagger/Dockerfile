FROM swaggerapi/swagger-ui:latest

COPY swagger_mixin /usr/share/nginx/html/swagger

# SWAGGER_JSON 会影响nginx root，默认取SWAGGER_JSON文件的目录为root,此文件须存在才会改变root
# https://github.com/swagger-api/swagger-ui/blob/master/docker/run.sh
ENV SWAGGER_JSON /usr/share/nginx/html/swagger/swagger.json
ENV URLS "[ { url: \"flamefatex/example-api/swagger.json\", name: \"flamefatex/example-api\" } ]"
