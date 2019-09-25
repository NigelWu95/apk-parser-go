FROM ubuntu:16.04
MAINTAINER wubingheng@qiniu.com

ENV LANG C.UTF-8

RUN mkdir -p /root/ufop
ADD apk-parser-linux_amd64 /root/ufop
ADD qufop.conf /root/ufop

#ENV JAVA_HOME /root/walle/jdk1.8.0_11
#ENV CLASSPATH ${JAVA_HOME}/lib/dt.jar:${JAVA_HOME}/lib/tools.jar
#ENV JRE_HOME ${JAVA_HOME}/jre
#ENV PATH ${JAVA_HOME}/bin:$PATH

EXPOSE 9100

WORKDIR /root/ufop
ENTRYPOINT ./apk-parser-linux_amd64 qufop.conf

