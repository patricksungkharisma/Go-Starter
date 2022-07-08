FROM golang:1.17 AS base

# Set starter file location
ARG SERVICE
ENV CMD=${SERVICE}
RUN echo ${CMD}
ENV STARTER_FILE=./.dev/.${CMD}.air.toml
RUN echo ${STARTER_FILE}

# Setup work directory
WORKDIR /go/src/app

# Prepare builder
RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

# Run builder
CMD air -c $STARTER_FILE
