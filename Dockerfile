FROM golang:1.22.1-bookworm AS base

ENV WORKPATH=/opt/workdir
ENV USERNAME=servo
ARG USER_ID=1000
ARG GROUP_ID=1000
RUN addgroup --gid ${GROUP_ID} $USERNAME
RUN adduser --uid ${USER_ID} --gid ${GROUP_ID} --disabled-password --gecos '' $USERNAME

WORKDIR ${WORKPATH}
RUN chown -R ${USER_ID}:${GROUP_ID} ${WORKPATH}
COPY --chown=${USER_ID}:${GROUP_ID} . $WORKPATH

USER ${USERNAME}

CMD ["go", "run", "main.go"]
