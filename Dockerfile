FROM scratch

LABEL Author Yoann Le Garff (le-garff-yoann) <pe.weeble@yahoo.fr>

COPY mojochat /
COPY vue/mojochat/index.html /vue/mojochat/
COPY vue/mojochat/dist /vue/mojochat/dist

ENV MOJOCHAT_PORT 8080

EXPOSE 8080

CMD [ "/mojochat" ]
