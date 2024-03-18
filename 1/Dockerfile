FROM ubuntu:22.04

ENV TZ = "Europe/Warsaw"
ENV DEBIAN_FRONTEND=noninteractive

# 3.0 obraz ubuntu z Pythonem w wersji 3.8
RUN apt-get update && \
    apt-get install -y software-properties-common && \
    add-apt-repository ppa:deadsnakes/ppa && \
    apt-get update && \
    apt-get install -y python3.8

# 3.5 obraz ubuntu:22.04 z Javą w wersji 8 oraz Kotlinem
RUN apt-get install -y openjdk-8-jdk kotlin

# 4.0 do powyższego należy dodać najnowszego Gradle’a oraz paczkę JDBC SQLite w ramach projektu na Gradle (build.gradle)
RUN apt-get update && \
    apt-get install -y wget && \
    wget -q https://services.gradle.org/distributions/gradle-8.6-bin.zip -P /tmp && \
    unzip -q /tmp/gradle-8.6-bin.zip -d /opt && \
    rm /tmp/gradle-8.6-bin.zip
ENV GRADLE_HOME /opt/gradle-8.6
ENV PATH $GRADLE_HOME/bin:$PATH

# 4.5 stworzyć przykład typu HelloWorld w Javie oraz uruchomienie aplikacji przez CMD oraz gradle
COPY Gradle /Gradle
WORKDIR /Gradle
CMD gradle build && gradle run