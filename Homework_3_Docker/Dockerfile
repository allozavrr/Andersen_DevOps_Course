
FROM python:3.8 as base
FROM base as builder-stage1
RUN mkdir /install
WORKDIR /install
RUN pip install --prefix=/install flask emoji
FROM base as builder-stage2
RUN mkdir /install
WORKDIR /install
RUN wget https://raw.githubusercontent.com/allozavrr/Andersen_DevOps_Course/main/Homework_1_zoOapp/zoOapp.zip
RUN unzip ./zoOapp.zip
FROM base
RUN mkdir /app
COPY --from=builder-stage1 /install /usr/local
COPY --from=builder-stage2 /install/zoOapp.py /app
WORKDIR /app
EXPOSE 8080
CMD ["python3", "./zoOapp.py"]


