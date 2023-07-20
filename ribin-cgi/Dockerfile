FROM alpine:3.7


RUN mkdir /root/bud-engine-cgi

COPY ./bud-engine-cgi /root/bud-engine-cgi/
# COPY ./conf.yaml /root/engine-data-cgi/

#设置工作目录
WORKDIR /root/bud-engine-cgi/
# 权限
RUN chmod +x bud-engine-cgi
#暴露端口
EXPOSE 9002
#项目入口

CMD ["./bud-engine-cgi"]


#dev
#docker build . -t 397275977064.dkr.ecr.us-west-1.amazonaws.com/buddy-images/engine-data-cgi:v0.7.9
#docker push 397275977064.dkr.ecr.us-west-1.amazonaws.com/buddy-images/engine-data-cgi:v0.7.9

#prod
#docker build . -t 397275977064.dkr.ecr.us-west-1.amazonaws.com/buddy-images/engine-data-cgi:1.0.2
#docker push 397275977064.dkr.ecr.us-west-1.amazonaws.com/buddy-images/engine-data-cgi:1.0.2

