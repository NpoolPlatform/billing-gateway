FROM uhub.service.ucloud.cn/entropypool/alpine:3.16

WORKDIR /app

COPY billing-gateway /app

ENTRYPOINT ["/app/billing-gateway"]

EXPOSE 50900 50901 50902 50903

CMD ["run"]
