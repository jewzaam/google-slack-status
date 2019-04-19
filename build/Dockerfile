FROM alpine:3.8

ENV OPERATOR=/usr/local/bin/google-slack-status \
    USER_UID=1001 \
    USER_NAME=google-slack-status

# install operator binary
COPY build/_output/bin/google-slack-status ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
