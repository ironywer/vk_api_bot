#!/bin/bash

echo "Обновляем строку подключения к БД..."
jq '.SqlSettings.DataSource = "postgres://mmuser:mmuser_password@db:5432/mattermost?sslmode=disable&connect_timeout=10&binary_parameters=yes"' \
    /mattermost/config/config.json > /mattermost/config/tmp.json && \
    mv /mattermost/config/tmp.json /mattermost/config/config.json


echo "Патчим config.json..."

jq '.ServiceSettings.AllowedUntrustedInternalConnections = "localhost, 127.0.0.1, ::1, bot"' \
    /mattermost/config/config.json > /mattermost/config/tmp.json && \
    mv /mattermost/config/tmp.json /mattermost/config/config.json

chown -R mattermost:mattermost /mattermost/config

echo "Запускаем Mattermost..."
exec /entrypoint.sh mattermost
