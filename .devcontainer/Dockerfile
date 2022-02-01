# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.187.0/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version: 1, 1.16, 1.15
ARG VARIANT="1.17"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# [Option] Install Node.js
ARG INSTALL_NODE="true"
ARG NODE_VERSION="lts/*"
RUN if [ "${INSTALL_NODE}" = "true" ]; then su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

RUN curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add -
# [Optional] Uncomment this section to install additional OS packages.
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
  && apt-get -y install --no-install-recommends zip g++ unzip

RUN echo "deb [arch=amd64] http://storage.googleapis.com/bazel-apt stable jdk1.8" | sudo tee /etc/apt/sources.list.d/bazel.list && \
  curl https://bazel.build/bazel-release.pub.gpg | sudo apt-key add - && \
  apt update && \
  apt -y install bazel && \
  apt -y install patch

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# [Optional] Uncomment the next line to use go get to install anything else you need
# RUN go get -x <your-dependency-or-tool>

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1
ENV GO111MODULE=off

RUN go get -u github.com/cosmtrek/air
RUN go get -u github.com/google/wire/cmd/wire
RUN go get -u go.uber.org/zap
RUN go get -u google.golang.org/api/option
RUN go get -u firebase.google.com/go
RUN go get -u cloud.google.com/go/firestore
RUN go build -o /go/bin/wire github.com/google/wire/cmd/wire
RUN go get -u github.com/blendle/zapdriver
RUN go get -u github.com/99designs/gqlgen
RUN go get -u github.com/rs/cors
# go get gorm.io/gorm && \
# go get github.com/99designs/gqlgen/internal/imports@v0.13.0 && \
# go get github.com/99designs/gqlgen/cmd@v0.13.0
# go get github.com/vektah/gqlparser/v2
ENV GO111MODULE on

ENV GO_ENV="development"

RUN export FIRESTORE_EMULATOR_HOST=host.docker.internal:8812
RUN export PUBSUB_PROJECT_ID=dummy
ENV FIRESTORE_EMULATOR_HOST=host.docker.internal:8812
ENV PUBSUB_PROJECT_ID=dummy

COPY ./entrypoint.sh /entrypoint.sh
RUN chmod 0755 /entrypoint.sh
# コンテナ起動時の処理
ENTRYPOINT ["/entrypoint.sh"]
