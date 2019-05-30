FROM ubuntu:16.04
MAINTAINER wubingheng@qiniu.com

ENV LANG C.UTF-8

RUN mkdir -p /root/miaopai
ADD miaopai-aparser-linux_amd64 /root/miaopai
ADD qufop.conf /root/miaopai

#ENV JAVA_HOME /root/walle/jdk1.8.0_11
#ENV CLASSPATH ${JAVA_HOME}/lib/dt.jar:${JAVA_HOME}/lib/tools.jar
#ENV JRE_HOME ${JAVA_HOME}/jre
#ENV PATH ${JAVA_HOME}/bin:$PATH

EXPOSE 9100

WORKDIR /root/miaopai
ENTRYPOINT ./miaopai-aparser-linux_amd64 qufop.conf

