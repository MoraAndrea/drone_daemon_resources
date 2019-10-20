FROM debian
COPY ./app /app
RUN chmod +x /app
ENTRYPOINT /app