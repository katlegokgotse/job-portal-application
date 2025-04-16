FROM python:3.12
WORKDIR /auth/

COPY requirements.txt ./auth/
RUN pip install -r requirements.txt

EXPOSE 8080
RUN app
ENTRYPOINT [ "app.py" ]
