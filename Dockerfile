FROM scratch

LABEL Author Yoann Le Garff (le-garff-yoann) <pe.weeble@yahoo.fr>

COPY mojochat /
ADD dist /dist

ENV MOJOCHAT_PORT 8080

EXPOSE 8080

CMD [ "/mojochat" ]
