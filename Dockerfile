FROM registry.red-soft.ru/ubi8/ubi

COPY . /project/

WORKDIR /project/

RUN dnf install -y git

