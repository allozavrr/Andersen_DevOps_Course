FROM alpine:3.5
RUN apk add --update python py-pip
COPY requirements.txt //https://raw.githubusercontent.com/allozavrr/Andersen_DevOps_Course/main/exam/Python_app/requirements.txt
RUN pip install -r //https://raw.githubusercontent.com/allozavrr/Andersen_DevOps_Course/main/exam/Python_app/requirements.txt
COPY app.py /tmp
COPY Python_app /tmp/Python_app
CMD python /tmp/app.py
