FROM python:3.8
ENV PYTHONUNBUFFERED 1
ENV PYTHONIOENCODING UTF-8

WORKDIR /usr/src/app
COPY . .
RUN pip install -r requirements.txt
EXPOSE 8000
CMD ["sh", "scripts/entrypoint.sh" ]

