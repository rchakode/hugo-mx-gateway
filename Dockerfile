FROM alpine:3.9.6

ARG RUNTIME_USER="mxgateway"
ARG RUNTIME_USER_UID=4583

RUN addgroup -g $RUNTIME_USER_UID $RUNTIME_USER && \
    adduser --disabled-password --no-create-home  --gecos "" \
    --home /app --ingroup $RUNTIME_USER --uid $RUNTIME_USER_UID  $RUNTIME_USER

COPY entrypoint.sh \
    bin/hugo-mx-gateway \
    templates \
    LICENSE \
    /app/

RUN chown -R $RUNTIME_USER:$RUNTIME_USER /app

WORKDIR /app
ENTRYPOINT ["sh", "./entrypoint.sh"]
