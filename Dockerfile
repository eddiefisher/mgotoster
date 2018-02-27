FROM scratch

ENV PORT 8000
EXPOSE $PORT

COPY /bin/mgotoster /
CMD ["/mgotoster"]