# Use a slim Java base image
FROM openjdk:24-slim

# Install curl and bash
RUN apt-get update && apt-get install -y curl bash

# Define other build arguments
ARG ANTLR_VERSION=4.13.2
ARG APP_HOME=/app
ARG APP_USER=cloud37
ARG HOME_DIR=/home/${APP_USER}

# Set environment variables based on the build arguments
ENV ANTLR_VERSION=${ANTLR_VERSION}
ENV APP_USER=${APP_USER}
ENV HOME=${HOME_DIR}

WORKDIR ${APP_HOME}

# Create a non-root user with a home directory
RUN addgroup --system ${APP_USER} && adduser --system --ingroup ${APP_USER} --home ${HOME_DIR} ${APP_USER}

# Ensure the home directory exists and has the correct permissions
RUN mkdir -p ${HOME_DIR} && chown -R ${APP_USER}:${APP_USER} ${HOME_DIR}

# Download ANTLR and place it in /usr/local/lib
RUN curl -O https://www.antlr.org/download/antlr-${ANTLR_VERSION}-complete.jar && \
    mv antlr-${ANTLR_VERSION}-complete.jar /usr/local/lib/antlr-complete.jar

# Create a shell script for antlr4 and ensure variable expansion
RUN echo "#!/bin/bash\njava -jar /usr/local/lib/antlr-complete.jar \"\$@\"" > /usr/local/bin/antlr4 && \
    chmod +x /usr/local/bin/antlr4

# Set correct ownership and permissions
RUN chown ${APP_USER}:${APP_USER} ${APP_HOME}

# Copy the entrypoint script
COPY --chown=${APP_USER}:${APP_USER} ./docker/antlr/entrypoint /entrypoint
RUN sed -i 's/\r$//g' /entrypoint && chmod +x /entrypoint

# Switch to the non-root user
USER ${APP_USER}

# Set the entrypoint
ENTRYPOINT ["/bin/bash", "/entrypoint"]
