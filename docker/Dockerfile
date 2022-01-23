FROM lacework/datacollector:latest AS agent-build-image

FROM nginx:1.21.0
EXPOSE 80
WORKDIR /app

COPY --from=agent-build-image  /var/lib/backup/*/datacollector /var/lib/lacework/datacollector
ARG LW_AGENT_ACCESS_TOKEN 
RUN mkdir -p /var/log/lacework/                 && \
    mkdir -p /var/lib/lacework/config/          && \
    echo '{"tokens": {"accesstoken": "'${LW_AGENT_ACCESS_TOKEN}'"}}' > /var/lib/lacework/config/config.json