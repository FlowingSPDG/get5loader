FROM flowingspdg/csgo-get5-docker-dev:latest

ENV GOTV_ENABLE 1
ENV PORT 27015
ENV GOTV_PORT 2720
ENV GOTV_DELAY 5
ENV PASSWORD local
ENV RCON_PASSWORD localpassword
ENV GOTV_PASSWORD gotvpass

EXPOSE 27015-27020/udp

CMD ["bash", "server-launch.sh"]